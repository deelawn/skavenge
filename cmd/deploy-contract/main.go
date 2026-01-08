package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/deelawn/skavenge/tests/util"
)

// DeployConfig holds the deployment configuration
type DeployConfig struct {
	DeployerPrivateKey string `json:"deployerPrivateKey"`
	HardhatURL         string `json:"hardhatUrl"`
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
	configData, err := os.ReadFile("deploy-config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading deploy-config.json: %v\n", err)
		os.Exit(1)
	}

	var config DeployConfig
	if err := json.Unmarshal(configData, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing deploy-config.json: %v\n", err)
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

	// Get gateway URL from environment
	gatewayURL := os.Getenv("GATEWAY_URL")
	if gatewayURL == "" {
		gatewayURL = "http://gateway:4591"
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
	fmt.Println("\nContract deployment completed successfully!")
}
