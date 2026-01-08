package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/deelawn/skavenge/pkg/minting"
	"github.com/deelawn/skavenge/tests/util"
)

// KeyPair represents an Ethereum/Skavenge key pair to register with the gateway
type KeyPair struct {
	EthereumPrivateKey string `json:"ethereumPrivateKey"`
	SkavengePrivateKey string `json:"skavengePrivateKey"`
}

// NFTConfig combines clue data and mint options for configuration
type NFTConfig struct {
	Content          string `json:"content"`
	Solution         string `json:"solution"`
	PointValue       uint8  `json:"pointValue"`
	SolveReward      string `json:"solveReward,omitempty"`      // Wei as string
	SalePrice        string `json:"salePrice,omitempty"`        // Wei as string
	Timeout          uint64 `json:"timeout,omitempty"`          // Transfer timeout in seconds
	RecipientAddress string `json:"recipientAddress,omitempty"` // Optional recipient
}

// SetupConfig holds the complete setup configuration
type SetupConfig struct {
	Minting            minting.Config `json:"minting"`
	DeployerPrivateKey string         `json:"deployerPrivateKey"`
	KeyPairs           []KeyPair      `json:"keyPairs"`
	NFTs               []NFTConfig    `json:"nfts"`
}

// WebappConfig holds the webapp configuration
type WebappConfig struct {
	ContractAddress string `json:"contractAddress"`
	NetworkRpcUrl   string `json:"networkRpcUrl"`
	ChainId         int64  `json:"chainId"`
	GatewayUrl      string `json:"gatewayUrl"`
}

// LinkRequest represents the JSON payload for POST /link
type LinkRequest struct {
	EthereumAddress   string `json:"ethereum_address"`
	SkavengePublicKey string `json:"skavenge_public_key"`
	Message           string `json:"message"`
	Signature         string `json:"signature"`
}

// ContractRequest represents the JSON payload for POST /contract
type ContractRequest struct {
	ContractAddress string `json:"contract_address"`
}

// ContractResponse represents the JSON response for contract operations
type ContractResponse struct {
	Success         bool   `json:"success"`
	Message         string `json:"message,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
}

func main() {
	// Load config
	configData, err := os.ReadFile("test-config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading test-config.json: %v\n", err)
		os.Exit(1)
	}

	var config SetupConfig
	if err := json.Unmarshal(configData, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing test-config.json: %v\n", err)
		os.Exit(1)
	}

	// Use RPCURL from config or environment
	hardhatURL := config.Minting.RPCURL
	if hardhatURL == "" {
		hardhatURL = os.Getenv("HARDHAT_URL")
		if hardhatURL == "" {
			hardhatURL = "http://localhost:8545"
		}
	}

	// Use gateway URL from config or environment
	gatewayURL := config.Minting.GatewayURL
	if gatewayURL == "" {
		gatewayURL = os.Getenv("GATEWAY_URL")
		if gatewayURL == "" {
			gatewayURL = "http://gateway:4591"
		}
	}

	// Use indexer URL from config or environment
	indexerURL := config.Minting.IndexerURL
	if indexerURL == "" {
		indexerURL = os.Getenv("INDEXER_URL")
		if indexerURL == "" {
			indexerURL = "http://indexer:4040"
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

	// Setup deployer account
	auth, err := util.NewTransactOpts(client, config.DeployerPrivateKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating transaction options: %v\n", err)
		os.Exit(1)
	}

	deployerAddress := auth.From
	fmt.Printf("Using deployer account: %s\n", deployerAddress.Hex())

	// Deploy contract
	fmt.Println("Deploying Skavenge contract...")
	_, contractAddress, err := util.DeployContract(client, auth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deploying contract: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Contract deployed at: %s\n", contractAddress.Hex())

	// Get chain ID
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting chain ID: %v\n", err)
		os.Exit(1)
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

	// Register contract address with gateway
	fmt.Printf("\nRegistering contract address with gateway...\n")
	if err := registerContract(gatewayURL, contractAddress.Hex()); err != nil {
		fmt.Fprintf(os.Stderr, "Error registering contract with gateway: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully registered contract address: %s\n", contractAddress.Hex())

	// Register key pairs with the gateway
	fmt.Printf("\nRegistering %d key pair(s) with the gateway...\n", len(config.KeyPairs))
	for i, keyPair := range config.KeyPairs {
		if err := registerKeyPair(gatewayURL, keyPair); err != nil {
			fmt.Fprintf(os.Stderr, "Error registering key pair %d: %v\n", i+1, err)
			os.Exit(1)
		}
		fmt.Printf("Successfully registered key pair %d\n", i+1)
	}

	// Update minting config with deployed contract address
	config.Minting.RPCURL = hardhatURL
	config.Minting.ContractAddress = contractAddress.Hex()
	config.Minting.GatewayURL = gatewayURL
	config.Minting.IndexerURL = indexerURL

	// Create minter
	fmt.Printf("\nInitializing minter...\n")
	minter, err := minting.NewMinter(&config.Minting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating minter: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Minter address: %s\n", minter.GetMinterAddress().Hex())

	// Mint NFTs
	fmt.Printf("\nMinting %d NFT(s)...\n", len(config.NFTs))
	for i, nftConfig := range config.NFTs {
		if err := mintNFT(minter, nftConfig, i+1); err != nil {
			fmt.Fprintf(os.Stderr, "Error minting NFT %d: %v\n", i+1, err)
			os.Exit(1)
		}
	}

	fmt.Println("\nSetup completed successfully!")
}

// registerContract registers the contract address with the gateway
func registerContract(gatewayURL string, contractAddress string) error {
	// Create request
	contractReq := ContractRequest{
		ContractAddress: contractAddress,
	}

	reqBody, err := json.Marshal(contractReq)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	// Send request to gateway
	resp, err := http.Post(gatewayURL+"/contract", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("gateway returned status %d: %s", resp.StatusCode, string(respBody))
	}

	return nil
}

// registerKeyPair registers an Ethereum/Skavenge key pair with the gateway
func registerKeyPair(gatewayURL string, keyPair KeyPair) error {
	// Load Ethereum private key
	ethPrivateKey, err := crypto.HexToECDSA(keyPair.EthereumPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to parse ethereum private key: %w", err)
	}

	// Derive Ethereum address
	ethAddress := crypto.PubkeyToAddress(ethPrivateKey.PublicKey)

	// Load Skavenge private key
	skavengePrivateKey, err := crypto.HexToECDSA(keyPair.SkavengePrivateKey)
	if err != nil {
		return fmt.Errorf("failed to parse skavenge private key: %w", err)
	}

	// Derive Skavenge public key (uncompressed format)
	skavengePublicKeyBytes := crypto.FromECDSAPub(&skavengePrivateKey.PublicKey)
	skavengePublicKeyHex := hex.EncodeToString(skavengePublicKeyBytes)

	// Create message to sign
	message := fmt.Sprintf("Link Ethereum address %s to Skavenge public key %s", ethAddress.Hex(), skavengePublicKeyHex)

	// Sign message with Ethereum private key (EIP-191 personal_sign)
	messageHash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
	signature, err := crypto.Sign(messageHash.Bytes(), ethPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign message: %w", err)
	}

	// Adjust v value for Ethereum signature format
	if signature[64] < 27 {
		signature[64] += 27
	}

	signatureHex := hex.EncodeToString(signature)

	// Create request
	linkReq := LinkRequest{
		EthereumAddress:   ethAddress.Hex(),
		SkavengePublicKey: skavengePublicKeyHex,
		Message:           message,
		Signature:         signatureHex,
	}

	reqBody, err := json.Marshal(linkReq)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	// Send request to gateway
	resp, err := http.Post(gatewayURL+"/link", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("gateway returned status %d: %s", resp.StatusCode, string(respBody))
	}

	return nil
}

// mintNFT mints a single NFT using the minting package
func mintNFT(minter *minting.Minter, nftConfig NFTConfig, index int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Convert NFTConfig to ClueData
	clueData := minting.ClueData{
		Content:    nftConfig.Content,
		Solution:   nftConfig.Solution,
		PointValue: nftConfig.PointValue,
	}

	// Parse solve reward if provided
	if nftConfig.SolveReward != "" {
		solveReward, ok := new(big.Int).SetString(nftConfig.SolveReward, 10)
		if !ok {
			return fmt.Errorf("invalid solve reward value: %s", nftConfig.SolveReward)
		}
		clueData.SolveReward = solveReward
	}

	// Convert NFTConfig to MintOptions
	mintOptions := minting.MintOptions{
		RecipientAddress: nftConfig.RecipientAddress,
		Timeout:          nftConfig.Timeout,
	}

	// Parse sale price if provided
	if nftConfig.SalePrice != "" {
		salePrice, ok := new(big.Int).SetString(nftConfig.SalePrice, 10)
		if !ok {
			return fmt.Errorf("invalid sale price value: %s", nftConfig.SalePrice)
		}
		mintOptions.SalePrice = salePrice
	}

	fmt.Printf("  [%d] Minting clue: %s (point value: %d)\n", index, truncateString(nftConfig.Content, 50), nftConfig.PointValue)

	// Mint the NFT
	result, err := minter.MintClue(ctx, &clueData, &mintOptions)
	if err != nil {
		return fmt.Errorf("minting failed: %w", err)
	}

	if result.Error != nil {
		return fmt.Errorf("minting error: %w", result.Error)
	}

	fmt.Printf("  [%d] ✓ Minted token ID %s (tx: %s)\n", index, result.TokenID.String(), result.TxHash)

	return nil
}

// truncateString truncates a string to maxLen characters
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// derivePublicKeyFromPrivate derives public key from private key string
func derivePublicKeyFromPrivate(privateKeyHex string) (*ecdsa.PublicKey, error) {
	// Remove 0x prefix if present
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	return &privateKey.PublicKey, nil
}
