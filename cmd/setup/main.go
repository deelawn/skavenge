package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/deelawn/skavenge/tests/util"
	"github.com/deelawn/skavenge/zkproof"
)

// Config holds the setup configuration
type Config struct {
	PrivateKey         string `json:"privateKey"`
	HardhatURL         string `json:"hardhatUrl"`
	SkavengePrivateKey string `json:"skavengePrivateKey"`
}

// WebappConfig holds the webapp configuration
type WebappConfig struct {
	ContractAddress string `json:"contractAddress"`
	NetworkRpcUrl   string `json:"networkRpcUrl"`
	ChainId         int64  `json:"chainId"`
	GatewayUrl      string `json:"gatewayUrl"`
}

func main() {
	// Load config
	configData, err := os.ReadFile("test-config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading test-config.json: %v\n", err)
		os.Exit(1)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing test-config.json: %v\n", err)
		os.Exit(1)
	}

	// Use hardhat URL from config or environment
	hardhatURL := config.HardhatURL
	if hardhatURL == "" {
		hardhatURL = os.Getenv("HARDHAT_URL")
		if hardhatURL == "" {
			hardhatURL = "http://localhost:8545"
		}
	}

	// Connect to Hardhat network
	client, err := ethclient.Dial(hardhatURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to Hardhat: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	fmt.Printf("Connected to Hardhat at %s\n", hardhatURL)

	// Setup account (same account for deployer and minter)
	auth, err := util.NewTransactOpts(client, config.PrivateKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating transaction options: %v\n", err)
		os.Exit(1)
	}

	accountAddress := auth.From
	fmt.Printf("Using account: %s\n", accountAddress.Hex())

	// Deploy contract
	fmt.Println("Deploying Skavenge contract...")
	contract, contractAddress, err := util.DeployContract(client, auth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deploying contract: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Contract deployed at: %s\n", contractAddress.Hex())

	// Load private key for encryption
	skavengePrivateKey, err := crypto.HexToECDSA(config.SkavengePrivateKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading private key: %v\n", err)
		os.Exit(1)
	}

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Sample data for the clue
	clueContent := []byte("Welcome to Skavenge! This is your first clue.")
	solution := "test-solution"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating r value: %v\n", err)
		os.Exit(1)
	}

	// Encrypt the clue content using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &skavengePrivateKey.PublicKey, mintR)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encrypting clue content: %v\n", err)
		os.Exit(1)
	}

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()

	// Mint a token to the same account
	fmt.Println("Minting token to account...")
	auth, err = util.NewTransactOpts(client, config.PrivateKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating transaction options for mint: %v\n", err)
		os.Exit(1)
	}

	auth.Value = big.NewInt(1000000000000000000)
	tx, err := contract.MintClue(auth, encryptedClueContent, solutionHash, mintR, 1, common.Address{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error minting token: %v\n", err)
		os.Exit(1)
	}

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for mint transaction: %v\n", err)
		os.Exit(1)
	}

	if receipt.Status == 0 {
		fmt.Fprintf(os.Stderr, "Mint transaction failed\n")
		os.Exit(1)
	}

	// Get the minted token ID
	tokenId, err := contract.GetCurrentTokenId(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current token ID: %v\n", err)
		os.Exit(1)
	}

	// CurrentTokenId is the next token ID to be minted, so we need to subtract 1 to get the actual token ID
	tokenId = new(big.Int).Sub(tokenId, big.NewInt(1))
	fmt.Printf("Token minted successfully! Token ID: %s\n", tokenId.String())

	// Verify ownership
	owner, err := contract.OwnerOf(nil, tokenId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting token owner: %v\n", err)
		os.Exit(1)
	}

	if owner != accountAddress {
		fmt.Fprintf(os.Stderr, "Error: Token owner mismatch. Expected %s, got %s\n", accountAddress.Hex(), owner.Hex())
		os.Exit(1)
	}

	fmt.Printf("Token ownership verified: %s owns token %s\n", owner.Hex(), tokenId.String())

	// Get chain ID
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting chain ID: %v\n", err)
		os.Exit(1)
	}

	// Get gateway URL from environment, default to internal Docker service name
	gatewayURL := os.Getenv("GATEWAY_URL")
	if gatewayURL == "" {
		gatewayURL = "http://gateway:4591"
	}

	// Update webapp config.json
	webappConfig := WebappConfig{
		ContractAddress: contractAddress.Hex(),
		NetworkRpcUrl:   hardhatURL,
		ChainId:         chainID.Int64(),
		GatewayUrl:      gatewayURL,
	}

	configJSON, err := json.MarshalIndent(webappConfig, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling webapp config: %v\n", err)
		os.Exit(1)
	}

	configPath := "webapp/config.json"
	if err := os.WriteFile(configPath, configJSON, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing webapp config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Updated %s with contract address: %s\n", configPath, contractAddress.Hex())
	fmt.Println("Setup completed successfully!")
}
