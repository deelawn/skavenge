# Skavenge Minting Package

A reusable Go package for minting Skavenge clues with support for encryption, sale listing, and direct transfers to recipients.

## Overview

The `pkg/minting` package provides core functionality for minting Skavenge clues. It is designed to be used by other Go applications and tools, including the `mint-clue` CLI tool.

## Features

- **ElGamal Encryption**: Automatically encrypts clue content using the recipient's Skavenge public key
- **Gateway Integration**: Retrieves recipient public keys from the linked accounts gateway
- **Flexible Minting**: Mint to yourself, list for sale, or send directly to recipients
- **Batch Operations**: Support for minting multiple clues in a single operation
- **Transaction Management**: Handles gas estimation, nonce management, and transaction confirmation

## Package Structure

```
pkg/minting/
├── types.go      # Data structures and types
├── gateway.go    # Gateway client for public key retrieval
├── minter.go     # Core minting functionality
└── README.md     # This file
```

## Quick Start

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
        MinterPrivateKey:   "your-ethereum-private-key",
        SkavengePrivateKey: "your-skavenge-private-key",
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
        Content:     "What is the answer to life, the universe, and everything?",
        Solution:    "42",
        PointValue:  3,
        SolveReward: big.NewInt(1000000000000000000), // 1 ETH
    }

    // Mint the clue
    result, err := minter.MintClue(context.Background(), clueData, nil)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Minted token ID: %s\n", result.TokenID.String())
}
```

## Types

### Config

Configuration for the minter:

```go
type Config struct {
    RPCURL             string // Blockchain RPC URL
    ContractAddress    string // Skavenge contract address
    MinterPrivateKey   string // Ethereum private key (authorized minter)
    SkavengePrivateKey string // Skavenge private key for encryption
    GatewayURL         string // URL of the linked accounts gateway
}
```

### ClueData

Data for a clue to be minted:

```go
type ClueData struct {
    Content     string   // The clue content (plaintext)
    Solution    string   // The solution to the clue
    PointValue  uint8    // Point value (1-5)
    SolveReward *big.Int // ETH reward for solving (in wei), optional
}
```

### MintOptions

Optional parameters for minting:

```go
type MintOptions struct {
    // For listing the clue for sale immediately after minting
    SalePrice *big.Int // Price in wei, 0 or nil means not for sale
    Timeout   uint64   // Transfer timeout in seconds (required if SalePrice > 0)

    // For minting and sending to another address
    RecipientAddress string // Ethereum address of recipient
}
```

### MintResult

Result of a minting operation:

```go
type MintResult struct {
    TokenID     *big.Int
    TxHash      string
    ClueContent string
    Solution    string
    PointValue  uint8
    Error       error
}
```

## API Reference

### NewMinter

Creates a new Minter instance.

```go
func NewMinter(config *Config) (*Minter, error)
```

**Parameters:**
- `config`: Configuration for the minter

**Returns:**
- `*Minter`: Minter instance
- `error`: Error if creation fails

### MintClue

Mints a single clue with the given data and options.

```go
func (m *Minter) MintClue(ctx context.Context, clueData *ClueData, options *MintOptions) (*MintResult, error)
```

**Parameters:**
- `ctx`: Context for cancellation and timeouts
- `clueData`: Clue data (content, solution, point value, etc.)
- `options`: Optional parameters (sale listing, recipient, etc.)

**Returns:**
- `*MintResult`: Result of the minting operation
- `error`: Error if minting fails

**Example - Mint to Self:**

```go
result, err := minter.MintClue(ctx, &minting.ClueData{
    Content:     "Clue content",
    Solution:    "Answer",
    PointValue:  3,
    SolveReward: big.NewInt(0),
}, nil)
```

**Example - Mint with Sale Listing:**

```go
result, err := minter.MintClue(ctx, clueData, &minting.MintOptions{
    SalePrice: big.NewInt(1000000000000000000), // 1 ETH
    Timeout:   3600,                             // 1 hour
})
```

**Example - Mint to Recipient:**

```go
result, err := minter.MintClue(ctx, clueData, &minting.MintOptions{
    RecipientAddress: "0x1234567890123456789012345678901234567890",
})
```

### MintClues

Mints multiple clues in batch.

```go
func (m *Minter) MintClues(ctx context.Context, clues []ClueData, options []MintOptions) ([]*MintResult, error)
```

**Parameters:**
- `ctx`: Context for cancellation and timeouts
- `clues`: Array of clue data
- `options`: Array of options (must match length of clues, or be empty)

**Returns:**
- `[]*MintResult`: Array of minting results
- `error`: Error if batch operation fails

### GetMinterAddress

Returns the minter's Ethereum address.

```go
func (m *Minter) GetMinterAddress() common.Address
```

**Returns:**
- `common.Address`: Minter's Ethereum address

### Close

Closes the blockchain connection.

```go
func (m *Minter) Close()
```

## Gateway Client

The gateway client retrieves Skavenge public keys for Ethereum addresses.

### NewGatewayClient

Creates a new gateway client.

```go
func NewGatewayClient(gatewayURL string) *GatewayClient
```

### GetPublicKey

Retrieves the Skavenge public key for an Ethereum address.

```go
func (g *GatewayClient) GetPublicKey(ethereumAddress string) (*ecdsa.PublicKey, error)
```

**Parameters:**
- `ethereumAddress`: Ethereum address (with or without 0x prefix)

**Returns:**
- `*ecdsa.PublicKey`: Skavenge public key
- `error`: Error if retrieval fails

## Encryption

The package uses ElGamal encryption on the secp256k1 curve to encrypt clue content. The encryption process:

1. Generates a random `r` value
2. Encrypts the clue content using the recipient's Skavenge public key and `r`
3. Stores the encrypted content and `r` value on-chain
4. The recipient can decrypt using their Skavenge private key and the `r` value

## Error Handling

All functions return errors that should be checked. Common errors include:

- **Network errors**: Connection to blockchain or gateway failed
- **Authentication errors**: Invalid private key or unauthorized minter
- **Validation errors**: Invalid point value, missing required fields, etc.
- **Transaction errors**: Transaction failed or reverted

Example:

```go
result, err := minter.MintClue(ctx, clueData, options)
if err != nil {
    // Handle error
    return err
}

if result.Error != nil {
    // Handle minting-specific error
    return result.Error
}
```

## Best Practices

1. **Always close the minter**: Use `defer minter.Close()` after creation
2. **Use contexts with timeouts**: Prevent hanging operations
3. **Check both function errors and result errors**: Some errors are returned in the result
4. **Validate input**: Point values must be 1-5, addresses must be valid
5. **Handle network failures gracefully**: Retry with exponential backoff
6. **Secure private keys**: Never hardcode private keys, use environment variables or secure vaults

## Integration Examples

### Using in a Web Service

```go
type ClueService struct {
    minter *minting.Minter
}

func NewClueService(config *minting.Config) (*ClueService, error) {
    minter, err := minting.NewMinter(config)
    if err != nil {
        return nil, err
    }
    return &ClueService{minter: minter}, nil
}

func (s *ClueService) CreateClue(ctx context.Context, content, solution string, pointValue uint8) (string, error) {
    result, err := s.minter.MintClue(ctx, &minting.ClueData{
        Content:    content,
        Solution:   solution,
        PointValue: pointValue,
    }, nil)

    if err != nil {
        return "", err
    }

    return result.TokenID.String(), nil
}
```

### Batch Minting with Progress Tracking

```go
func batchMintWithProgress(minter *minting.Minter, clues []minting.ClueData) error {
    total := len(clues)

    for i, clue := range clues {
        ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)

        result, err := minter.MintClue(ctx, &clue, nil)
        cancel()

        if err != nil || result.Error != nil {
            fmt.Printf("[%d/%d] Failed: %v\n", i+1, total, err)
            continue
        }

        fmt.Printf("[%d/%d] Minted token %s\n", i+1, total, result.TokenID.String())
    }

    return nil
}
```

## Testing

The package can be tested against a local Hardhat network:

```go
import "testing"

func TestMinting(t *testing.T) {
    config := &minting.Config{
        RPCURL:             "http://localhost:8545",
        ContractAddress:    "0x5FbDB2315678afecb367f032d93F642f64180aa3",
        MinterPrivateKey:   "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
        SkavengePrivateKey: "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
        GatewayURL:         "http://localhost:4591",
    }

    minter, err := minting.NewMinter(config)
    if err != nil {
        t.Fatalf("Failed to create minter: %v", err)
    }
    defer minter.Close()

    // Test minting
    result, err := minter.MintClue(context.Background(), &minting.ClueData{
        Content:    "Test clue",
        Solution:   "Test solution",
        PointValue: 1,
    }, nil)

    if err != nil || result.Error != nil {
        t.Fatalf("Failed to mint clue: %v %v", err, result.Error)
    }

    t.Logf("Minted token ID: %s", result.TokenID.String())
}
```

## License

MIT
