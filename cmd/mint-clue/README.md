# Mint Clue CLI Tool

A command-line tool for minting Skavenge clues. Supports single clue minting, batch minting from JSON files, and advanced features like listing for sale and minting directly to recipients.

## Features

- **Single Clue Minting**: Mint individual clues using command-line arguments
- **Batch Minting**: Mint multiple clues from a JSON configuration file
- **List for Sale**: Optionally list a clue for sale immediately after minting
- **Direct Transfer**: Mint and send clues directly to another user's Ethereum address
- **Encrypted Transfer**: Automatically retrieves recipient's Skavenge public key from gateway and encrypts clue data

## Installation

```bash
go build -o mint-clue
```

## Usage

### Single Clue Minting

Mint a single clue to yourself:

```bash
./mint-clue \
  --contract "0x..." \
  --minter-key "your-private-key" \
  --skavenge-key "your-skavenge-private-key" \
  --content "Find the hidden treasure in the old oak tree" \
  --solution "Oak tree" \
  --point-value 3
```

### Mint with Sale Listing

Mint a clue and list it for sale immediately:

```bash
./mint-clue \
  --contract "0x..." \
  --minter-key "your-private-key" \
  --skavenge-key "your-skavenge-private-key" \
  --content "The answer lies in the stars" \
  --solution "Constellation" \
  --point-value 4 \
  --sale-price "1000000000000000000" \
  --timeout 3600
```

### Mint to Another User

Mint a clue directly to another user's Ethereum address (requires their Skavenge public key to be registered in the gateway):

```bash
./mint-clue \
  --contract "0x..." \
  --minter-key "your-private-key" \
  --skavenge-key "your-skavenge-private-key" \
  --content "Where the river meets the sea" \
  --solution "Delta" \
  --point-value 2 \
  --recipient "0x1234567890123456789012345678901234567890"
```

### Batch Minting from JSON

Create a JSON file with multiple clues (see `example-clues.json`):

```json
{
  "clues": [
    {
      "content": "Find the hidden treasure in the old oak tree",
      "solution": "Oak tree",
      "pointValue": 3,
      "solveReward": "1000000000000000000",
      "salePrice": "500000000000000000",
      "timeout": 3600
    },
    {
      "content": "The answer lies in the stars above",
      "solution": "Constellation",
      "pointValue": 4,
      "solveReward": "2000000000000000000"
    }
  ]
}
```

Then mint all clues:

```bash
./mint-clue \
  --contract "0x..." \
  --minter-key "your-private-key" \
  --skavenge-key "your-skavenge-private-key" \
  --input clues.json
```

## Command-Line Options

### Required Options

- `--contract`: Skavenge contract address
- `--minter-key`: Ethereum private key of authorized minter
- `--skavenge-key`: Skavenge private key for encrypting clue content

### Single Clue Options

- `--content`: Clue content (plaintext)
- `--solution`: Clue solution
- `--point-value`: Point value (1-5, default: 1)
- `--solve-reward`: ETH reward for solving in wei (default: 0)
- `--sale-price`: Sale price in wei, 0 means not for sale (default: 0)
- `--timeout`: Transfer timeout in seconds (default: 3600)
- `--recipient`: Recipient Ethereum address (for direct minting to another user)

### Batch Options

- `--input`: Path to JSON file containing multiple clues

### Connection Options

- `--rpc`: Blockchain RPC URL (default: http://localhost:8545)
- `--gateway`: Gateway service URL (default: http://localhost:4591)

## JSON Clue Format

Each clue in the JSON file can have the following fields:

```json
{
  "content": "Clue text",           // Required
  "solution": "Solution text",      // Required
  "pointValue": 3,                  // Required (1-5)
  "solveReward": "1000000000",      // Optional: Wei amount
  "salePrice": "500000000",         // Optional: Wei amount (lists for sale)
  "timeout": 3600,                  // Optional: Seconds (required if salePrice > 0)
  "recipientAddress": "0x..."       // Optional: Send to this address
}
```

## Examples

### Example 1: Mint a Simple Clue

```bash
./mint-clue \
  --contract "0x5FbDB2315678afecb367f032d93F642f64180aa3" \
  --minter-key "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" \
  --skavenge-key "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef" \
  --content "What has keys but no locks?" \
  --solution "Piano" \
  --point-value 2
```

### Example 2: Mint with Reward and Sale

```bash
./mint-clue \
  --contract "0x5FbDB2315678afecb367f032d93F642f64180aa3" \
  --minter-key "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" \
  --skavenge-key "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef" \
  --content "I speak without a mouth and hear without ears" \
  --solution "Echo" \
  --point-value 4 \
  --solve-reward "1000000000000000000" \
  --sale-price "500000000000000000" \
  --timeout 7200
```

### Example 3: Batch Mint

```bash
./mint-clue \
  --contract "0x5FbDB2315678afecb367f032d93F642f64180aa3" \
  --minter-key "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" \
  --skavenge-key "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef" \
  --input example-clues.json
```

## Using the Minting Package in Your Code

The core minting functionality is available as a reusable Go package at `github.com/deelawn/skavenge/pkg/minting`.

```go
package main

import (
    "context"
    "fmt"
    "math/big"

    "github.com/deelawn/skavenge/pkg/minting"
)

func main() {
    // Create configuration
    config := &minting.Config{
        RPCURL:             "http://localhost:8545",
        ContractAddress:    "0x5FbDB2315678afecb367f032d93F642f64180aa3",
        MinterPrivateKey:   "your-private-key",
        SkavengePrivateKey: "your-skavenge-key",
        GatewayURL:         "http://localhost:4591",
    }

    // Create minter
    minter, err := minting.NewMinter(config)
    if err != nil {
        panic(err)
    }
    defer minter.Close()

    // Prepare clue data
    clueData := &minting.ClueData{
        Content:     "What is the answer?",
        Solution:    "42",
        PointValue:  3,
        SolveReward: big.NewInt(1000000000000000000), // 1 ETH
    }

    // Prepare options (optional)
    options := &minting.MintOptions{
        SalePrice: big.NewInt(500000000000000000), // 0.5 ETH
        Timeout:   3600,                            // 1 hour
    }

    // Mint the clue
    result, err := minter.MintClue(context.Background(), clueData, options)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Minted token ID: %s\n", result.TokenID.String())
    fmt.Printf("Transaction: %s\n", result.TxHash)
}
```

## How Encrypted Minting to Recipients Works

When you mint a clue to a recipient:

1. The CLI tool retrieves the recipient's Skavenge public key from the gateway service using their Ethereum address
2. A new random `r` value is generated for ElGamal encryption
3. The clue content is encrypted using the recipient's public key and the `r` value
4. The encrypted clue data, `r` value, and other clue information are stored on-chain
5. The recipient can decrypt the clue using their Skavenge private key and the `r` value

This ensures that only the recipient can decrypt and read the clue content.

## Troubleshooting

### "ethereum address not found" error

The recipient's Ethereum address must be registered in the gateway service with their Skavenge public key before you can mint a clue to them.

### "UnauthorizedMinter" error

The private key you're using must be authorized as a minter in the Skavenge contract. Use the `updateAuthorizedMinter` function to set the authorized minter.

### Gas limit errors

If you encounter gas limit errors, you may need to increase the gas limit in the `minter.go` file (currently set to 300000).

## License

MIT
