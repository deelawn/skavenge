package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/deelawn/skavenge/pkg/minting"
)

// JSONClue represents a clue in the JSON input file
type JSONClue struct {
	Content         string  `json:"content"`
	Solution        string  `json:"solution"`
	PointValue      uint8   `json:"pointValue"`
	SolveReward     *string `json:"solveReward,omitempty"`     // Wei amount as string
	SalePrice       *string `json:"salePrice,omitempty"`       // Wei amount as string
	Timeout         *uint64 `json:"timeout,omitempty"`         // Seconds
	RecipientAddress *string `json:"recipientAddress,omitempty"` // Ethereum address to send to
}

// JSONConfig represents the configuration in the JSON file
type JSONConfig struct {
	Clues []JSONClue `json:"clues"`
}

func main() {
	// Define command-line flags
	configFile := flag.String("config", "", "Path to configuration file (JSON)")
	rpcURL := flag.String("rpc", "http://localhost:8545", "Blockchain RPC URL")
	contractAddr := flag.String("contract", "", "Skavenge contract address")
	gatewayURL := flag.String("gateway", "http://localhost:4591", "Gateway service URL")
	indexerURL := flag.String("indexer", "http://localhost:4040", "Indexer API URL")

	// Single clue flags
	content := flag.String("content", "", "Clue content (plaintext)")
	solution := flag.String("solution", "", "Clue solution")
	pointValue := flag.Uint("point-value", 1, "Point value (1-5)")
	solveReward := flag.String("solve-reward", "0", "Solve reward in wei")
	salePrice := flag.String("sale-price", "0", "Sale price in wei (0 means not for sale)")
	timeout := flag.Uint64("timeout", 14400, "Transfer timeout in seconds (default: 4 hours)")
	recipientAddr := flag.String("recipient", "", "Recipient Ethereum address (for direct minting to another user)")

	// Batch minting flag
	inputFile := flag.String("input", "", "Path to JSON file containing multiple clues to mint")

	flag.Parse()

	// Read keys from environment variables
	minterKey := os.Getenv("MINTER_PRIVATE_KEY")
	skavengePubKey := os.Getenv("SKAVENGE_PUBLIC_KEY")

	// Validate required flags and environment variables
	if *contractAddr == "" {
		fmt.Fprintf(os.Stderr, "Error: contract address is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if minterKey == "" {
		fmt.Fprintf(os.Stderr, "Error: MINTER_PRIVATE_KEY environment variable is required\n")
		os.Exit(1)
	}

	// Determine mode: single clue vs batch
	if *inputFile != "" {
		// Batch mode
		if err := mintFromFile(*inputFile, *rpcURL, *contractAddr, minterKey, skavengePubKey, *gatewayURL, *indexerURL); err != nil {
			fmt.Fprintf(os.Stderr, "Error minting from file: %v\n", err)
			os.Exit(1)
		}
	} else if *content != "" && *solution != "" {
		// Single clue mode
		if err := mintSingleClue(*content, *solution, uint8(*pointValue), *solveReward, *salePrice, *timeout, *recipientAddr, *rpcURL, *contractAddr, minterKey, skavengePubKey, *gatewayURL, *indexerURL); err != nil {
			fmt.Fprintf(os.Stderr, "Error minting clue: %v\n", err)
			os.Exit(1)
		}
	} else if *configFile != "" {
		// Config file mode (legacy support)
		if err := mintFromConfigFile(*configFile, *rpcURL, *contractAddr, minterKey, skavengePubKey, *gatewayURL, *indexerURL); err != nil {
			fmt.Fprintf(os.Stderr, "Error minting from config file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Error: must provide either --content and --solution for single clue, --input for batch, or --config for config file\n")
		flag.Usage()
		os.Exit(1)
	}
}

func mintSingleClue(content, solution string, pointValue uint8, solveRewardStr, salePriceStr string, timeout uint64, recipientAddr, rpcURL, contractAddr, minterKey, skavengePubKey, gatewayURL, indexerURL string) error {
	// Validate point value
	if pointValue < 1 || pointValue > 5 {
		return fmt.Errorf("point value must be between 1 and 5")
	}

	// Validate skavenge public key requirement
	// Only required when minting to self (no recipient specified)
	if recipientAddr == "" && skavengePubKey == "" {
		return fmt.Errorf("SKAVENGE_PUBLIC_KEY environment variable is required when minting to self")
	}

	// Parse solve reward
	solveReward := new(big.Int)
	if solveRewardStr != "" && solveRewardStr != "0" {
		var ok bool
		solveReward, ok = new(big.Int).SetString(solveRewardStr, 10)
		if !ok {
			return fmt.Errorf("invalid solve reward: %s", solveRewardStr)
		}
	}

	// Parse sale price
	salePrice := new(big.Int)
	if salePriceStr != "" && salePriceStr != "0" {
		var ok bool
		salePrice, ok = new(big.Int).SetString(salePriceStr, 10)
		if !ok {
			return fmt.Errorf("invalid sale price: %s", salePriceStr)
		}
	}

	// Create minter config
	config := &minting.Config{
		RPCURL:            rpcURL,
		ContractAddress:   contractAddr,
		MinterPrivateKey:  minterKey,
		SkavengePublicKey: skavengePubKey,
		GatewayURL:        gatewayURL,
		IndexerURL:        indexerURL,
	}

	// Create minter
	minter, err := minting.NewMinter(config)
	if err != nil {
		return fmt.Errorf("failed to create minter: %w", err)
	}
	defer minter.Close()

	fmt.Printf("Minter address: %s\n", minter.GetMinterAddress().Hex())

	// Prepare clue data
	clueData := &minting.ClueData{
		Content:     content,
		Solution:    solution,
		PointValue:  pointValue,
		SolveReward: solveReward,
	}

	// Prepare options
	options := &minting.MintOptions{}
	if salePrice.Cmp(big.NewInt(0)) > 0 {
		options.SalePrice = salePrice
		options.Timeout = timeout
	}
	if recipientAddr != "" {
		options.RecipientAddress = recipientAddr
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Mint the clue
	fmt.Println("Minting clue...")
	result, err := minter.MintClue(ctx, clueData, options)
	if err != nil {
		return fmt.Errorf("failed to mint clue: %w", err)
	}

	if result.Error != nil {
		return fmt.Errorf("minting error: %w", result.Error)
	}

	// Print result
	fmt.Printf("\n✓ Clue minted successfully!\n")
	fmt.Printf("  Token ID: %s\n", result.TokenID.String())
	fmt.Printf("  Transaction: %s\n", result.TxHash)
	fmt.Printf("  Content: %s\n", result.ClueContent)
	fmt.Printf("  Solution: %s\n", result.Solution)
	fmt.Printf("  Point Value: %d\n", result.PointValue)

	if recipientAddr != "" {
		fmt.Printf("  Recipient: %s\n", recipientAddr)
	}

	if salePrice.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("  Listed for sale at: %s wei\n", salePrice.String())
		fmt.Printf("  Transfer timeout: %d seconds\n", timeout)
	}

	return nil
}

func mintFromFile(inputFile, rpcURL, contractAddr, minterKey, skavengePubKey, gatewayURL, indexerURL string) error {
	// Read the input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Parse JSON
	var jsonConfig JSONConfig
	if err := json.Unmarshal(data, &jsonConfig); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	if len(jsonConfig.Clues) == 0 {
		return fmt.Errorf("no clues found in input file")
	}

	// Check if any clue is being minted to self (requires skavenge public key)
	needsSkavengeKey := false
	for _, clue := range jsonConfig.Clues {
		if clue.RecipientAddress == nil || *clue.RecipientAddress == "" {
			needsSkavengeKey = true
			break
		}
	}

	if needsSkavengeKey && skavengePubKey == "" {
		return fmt.Errorf("SKAVENGE_PUBLIC_KEY environment variable is required when minting clues to self")
	}

	// Create minter config
	config := &minting.Config{
		RPCURL:            rpcURL,
		ContractAddress:   contractAddr,
		MinterPrivateKey:  minterKey,
		SkavengePublicKey: skavengePubKey,
		GatewayURL:        gatewayURL,
		IndexerURL:        indexerURL,
	}

	// Create minter
	minter, err := minting.NewMinter(config)
	if err != nil {
		return fmt.Errorf("failed to create minter: %w", err)
	}
	defer minter.Close()

	fmt.Printf("Minter address: %s\n", minter.GetMinterAddress().Hex())
	fmt.Printf("Minting %d clues...\n\n", len(jsonConfig.Clues))

	// Create context
	ctx := context.Background()

	// Mint each clue
	successCount := 0
	for i, jsonClue := range jsonConfig.Clues {
		fmt.Printf("[%d/%d] Minting clue...\n", i+1, len(jsonConfig.Clues))

		// Validate point value
		if jsonClue.PointValue < 1 || jsonClue.PointValue > 5 {
			fmt.Fprintf(os.Stderr, "  ✗ Error: point value must be between 1 and 5\n\n")
			continue
		}

		// Parse solve reward
		var solveReward *big.Int
		if jsonClue.SolveReward != nil && *jsonClue.SolveReward != "" && *jsonClue.SolveReward != "0" {
			solveReward = new(big.Int)
			var ok bool
			solveReward, ok = solveReward.SetString(*jsonClue.SolveReward, 10)
			if !ok {
				fmt.Fprintf(os.Stderr, "  ✗ Error: invalid solve reward: %s\n\n", *jsonClue.SolveReward)
				continue
			}
		}

		// Prepare clue data
		clueData := &minting.ClueData{
			Content:     jsonClue.Content,
			Solution:    jsonClue.Solution,
			PointValue:  jsonClue.PointValue,
			SolveReward: solveReward,
		}

		// Prepare options
		options := &minting.MintOptions{}

		// Parse sale price if provided
		if jsonClue.SalePrice != nil && *jsonClue.SalePrice != "" && *jsonClue.SalePrice != "0" {
			salePrice := new(big.Int)
			var ok bool
			salePrice, ok = salePrice.SetString(*jsonClue.SalePrice, 10)
			if !ok {
				fmt.Fprintf(os.Stderr, "  ✗ Error: invalid sale price: %s\n\n", *jsonClue.SalePrice)
				continue
			}
			options.SalePrice = salePrice

			// Use timeout from JSON or default
			if jsonClue.Timeout != nil {
				options.Timeout = *jsonClue.Timeout
			} else {
				options.Timeout = 14400 // Default 4 hours
			}
		}

		// Set recipient if provided
		if jsonClue.RecipientAddress != nil && *jsonClue.RecipientAddress != "" {
			options.RecipientAddress = *jsonClue.RecipientAddress
		}

		// Create context with timeout
		mintCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)

		// Mint the clue
		result, err := minter.MintClue(mintCtx, clueData, options)
		cancel()

		if err != nil || result.Error != nil {
			if err != nil {
				fmt.Fprintf(os.Stderr, "  ✗ Error: %v\n\n", err)
			} else {
				fmt.Fprintf(os.Stderr, "  ✗ Error: %v\n\n", result.Error)
			}
			continue
		}

		// Print result
		fmt.Printf("  ✓ Token ID: %s\n", result.TokenID.String())
		fmt.Printf("    Transaction: %s\n", result.TxHash)
		fmt.Printf("    Content: %s\n", result.ClueContent)
		fmt.Printf("    Solution: %s\n", result.Solution)
		fmt.Printf("    Point Value: %d\n", result.PointValue)

		if options.RecipientAddress != "" {
			fmt.Printf("    Recipient: %s\n", options.RecipientAddress)
		}

		if options.SalePrice != nil && options.SalePrice.Cmp(big.NewInt(0)) > 0 {
			fmt.Printf("    Listed for sale at: %s wei\n", options.SalePrice.String())
		}

		fmt.Println()
		successCount++
	}

	fmt.Printf("Successfully minted %d/%d clues\n", successCount, len(jsonConfig.Clues))

	return nil
}

func mintFromConfigFile(configFile, rpcURL, contractAddr, minterKey, skavengePubKey, gatewayURL, indexerURL string) error {
	// This is similar to mintFromFile but kept separate for backward compatibility
	return mintFromFile(configFile, rpcURL, contractAddr, minterKey, skavengePubKey, gatewayURL, indexerURL)
}
