# Skavenge Query CLI Generator

This tool automatically generates a CLI program for making read-only calls to the Skavenge smart contract. When the contract is updated and the Go bindings are regenerated, you can easily regenerate the CLI tool to reflect the new functions.

## Overview

The generator parses the `bindings.go` file created by `abigen`, identifies all read-only functions (methods on `SkavengeCaller`), and generates a complete CLI program with:

- Automatic command creation for each contract function
- Proper argument parsing based on parameter types
- Type conversion (addresses, big integers, bytes)
- Formatted output
- Comprehensive help documentation

## Usage

### Generate the CLI Tool

```bash
# Using Make (recommended)
make generate-cli

# Or run directly
go run ./cmd/generate-query-cli eth/bindings/bindings.go cmd/skavenge-query/main.go
```

### Build the CLI Binary

```bash
# Generate and build in one step
make build-cli

# This creates: bin/skavenge-query
```

### Use the Generated CLI

```bash
# Set the contract address
export SKAVENGE_CONTRACT_ADDRESS=0x1234567890123456789012345678901234567890

# Query contract functions
./bin/skavenge-query name
./bin/skavenge-query total-supply
./bin/skavenge-query balance-of 0xabcdef...
./bin/skavenge-query owner-of 1
./bin/skavenge-query get-clues-for-sale 0 10

# Use custom node address
./bin/skavenge-query -node http://localhost:8545 name
```

## Workflow

When you update the smart contract:

1. **Compile the contract:**
   ```bash
   make compile
   ```

2. **Regenerate the CLI:**
   ```bash
   make generate-cli
   ```

3. **Build the CLI binary:**
   ```bash
   make build-cli
   ```

All done! The CLI now reflects your updated contract.

## How It Works

1. **Parse Bindings:** The generator uses Go's AST parser to read `eth/bindings/bindings.go`
2. **Identify Functions:** It finds all methods on `SkavengeCaller` (read-only contract functions)
3. **Categorize Functions:** Functions are grouped by parameter count and types
4. **Generate Code:** A template generates the complete CLI program with:
   - Command routing
   - Argument parsing
   - Type conversion
   - Error handling
   - Formatted output

## Generated CLI Features

- **Environment Variables:** 
  - `SKAVENGE_CONTRACT_ADDRESS` - Contract address (required)
  
- **Flags:**
  - `-node` - Ethereum node address (default: http://localhost:8545)

- **Command Naming:**
  - camelCase contract functions → kebab-case CLI commands
  - Example: `getCurrentTokenId` → `get-current-token-id`

- **Automatic Type Handling:**
  - Ethereum addresses (validates and converts)
  - Big integers (string to big.Int conversion)
  - Byte arrays (hex string parsing)

- **Output Formatting:**
  - Addresses: hex format
  - Big integers: decimal string
  - Bytes: hex format
  - Booleans: true/false
  - Structs: formatted with field names

## Example Commands

```bash
# No arguments
./bin/skavenge-query name
./bin/skavenge-query symbol
./bin/skavenge-query total-supply
./bin/skavenge-query get-current-token-id
./bin/skavenge-query authorized-minter

# Single argument (address)
./bin/skavenge-query balance-of 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb

# Single argument (token ID)
./bin/skavenge-query owner-of 1
./bin/skavenge-query get-clue-contents 1
./bin/skavenge-query get-r-value 1
./bin/skavenge-query clues 1

# Multiple arguments
./bin/skavenge-query get-clues-for-sale 0 10
./bin/skavenge-query token-of-owner-by-index 0x742d35... 0
./bin/skavenge-query generate-transfer-id 0x742d35... 1
```

## Benefits

- **Always Up-to-Date:** Regenerate whenever the contract changes
- **No Manual Coding:** All functions are automatically included
- **Type Safe:** Uses the same bindings as your Go code
- **Comprehensive:** Every read-only function is exposed
- **Easy to Use:** Clear command names and help text
- **Error Handling:** Proper validation and error messages

## Technical Details

- **Language:** Go
- **Dependencies:** 
  - `go-ethereum` (ethclient, accounts/abi/bind, common)
  - Generated bindings package
- **Parser:** Uses Go's `go/ast` and `go/parser` packages
- **Template Engine:** Go's `text/template` package

