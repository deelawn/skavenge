# Skavenge NFT Scavenger Hunt

Skavenge is an NFT-based scavenger hunt game implemented on Ethereum. Players can mint, solve, and trade clues as NFTs using zero-knowledge proofs for secure transfers.
![skavenge diagram](skavenge-flow.svg)
## Project Structure

```
/skavenge
├── cmd/
│   ├── generate-query-cli/  # CLI generator for contract queries
│   ├── skavenge-query/      # Generated CLI tool (auto-generated)
│   └── setup/               # Contract deployment tool
├── eth/
│   ├── bindings/            # Go bindings for Ethereum contracts
│   ├── build/               # Compiled contract artifacts
│   └── skavenge.sol         # Solidity smart contract
├── linked-accounts-gateway/ # Gateway for linked accounts
├── skavenger-extension/     # Browser extension
├── tests/
│   ├── util/                # Test utilities
│   ├── mint_test.go         # Tests for minting clues
│   ├── solve_test.go        # Tests for solving clues
│   └── transfer_test.go     # Tests for transferring clues
├── webapp/                  # Web application
├── zkproof/                 # Zero-knowledge proof implementation
└── Makefile                 # Build and test automation
```

## Features

- **Minting Clues**: Create new clues with encrypted content and solution hashes
- **Solving Clues**: Attempt to solve clues with a limited number of tries
- **Trading Clues**: Securely transfer clues to other players using ZK proofs
- **Limited Attempts**: Each clue has a maximum number of solve attempts
- **Transfer Protection**: Prevents transfer of solved clues
- **CLI Query Tool**: Auto-generated command-line interface for contract queries (see [CLI_TOOLS.md](CLI_TOOLS.md))

## Technology Stack

- **Smart Contracts**: Solidity with OpenZeppelin ERC721
- **Zero-Knowledge Proofs**: Custom ZK proof API for secure transfers
- **Blockchain**: Ethereum
- **Testing**: Go with eth bindings, Hardhat for local blockchain

## Setup Instructions

### Prerequisites
- Go 1.17+
- Node.js 14+
- Hardhat
- Access to a ZK proof API server

### Development Setup

To be completed in Phase 3.

## Testing

The project includes a comprehensive test suite for validating the functionality of the smart contract. Tests cover minting, solving, and transferring clues.

To run the tests:

```bash
# Start a local Hardhat node
cd eth
npx hardhat node

# Start the ZK proof API server
# (Instructions TBD)

# Run the tests
go test ./tests/...
```

## License

TBD
