# Linked Accounts Gateway

A Go web server that provides a gateway for storing and retrieving linked account information, mapping Ethereum addresses to Skavenge public keys, and managing transfer ciphertext for secure NFT transfers.

## Overview

This service provides two main functionalities:
1. **Account Linking**: Links Ethereum addresses (from MetaMask) to Skavenge public keys through cryptographic signature verification
2. **Transfer Ciphertext Management**: Securely stores and retrieves encrypted clue data during NFT transfers

## Features

- **Dual Signature Verification**: 
  - Ethereum signatures using `personal_sign` for account linking
  - Skavenge ECDSA secp256k1 signatures for transfer authentication (same curve as Ethereum)
- **Immutable Storage**: Once an Ethereum address is linked to a Skavenge key, it cannot be modified
- **Transfer Security**: Verifies seller and buyer ownership through on-chain data and signature verification
- **In-Memory Storage**: Current implementation uses in-memory maps (can be replaced with persistent storage)
- **RESTful API**: Simple HTTP endpoints for all operations
- **Blockchain Integration**: Connects to Ethereum RPC to verify transfer data

## Installation

1. Install dependencies:
```bash
go mod tidy
```

2. Run the server:
```bash
# Run with default settings
go run ./linked-accounts-gateway -contract 0xYourContractAddress

# Run with custom port and RPC URL
go run ./linked-accounts-gateway \
  -port 8080 \
  -rpc http://localhost:8545 \
  -contract 0xYourContractAddress
```

**Required flags:**
- `-contract`: Skavenge contract address (required)

**Optional flags:**
- `-port`: Port to listen on (default: 4591)
- `-rpc`: Blockchain RPC URL (default: http://localhost:8545)

### Docker

Run using Docker Compose:
```bash
# Start all services including gateway
make start

# Or start just the gateway
docker compose up -d gateway
```

The gateway will be available at `http://localhost:4591`.

## API Endpoints

### Account Linking Endpoints

#### POST /link

Links an Ethereum address to a Skavenge public key after verifying the signature.

**Request Body:**
```json
{
  "ethereumAddress": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "skavengePublicKey": "skv1abc123...",
  "message": "Link MetaMask address 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb to Skavenge key skv1abc123...",
  "signature": "0x1234567890abcdef..."
}
```

**Success Response (201 Created):**
```json
{
  "success": true,
  "message": "linkage created successfully"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid JSON, missing fields, or signature verification failed
- `401 Unauthorized`: Invalid signature
- `409 Conflict`: Ethereum address already linked (keys are immutable)
- `500 Internal Server Error`: Storage failure

### GET /link

Retrieves the Skavenge public key associated with an Ethereum address.

**Query Parameters:**
- `ethereumAddress` (required): The Ethereum address to look up

**Example:**
```
GET /link?ethereumAddress=0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb
```

**Success Response (200 OK):**
```json
{
  "success": true,
  "skavengePublicKey": "skv1abc123..."
}
```

**Error Responses:**
- `400 Bad Request`: Missing ethereumAddress parameter
- `404 Not Found`: Ethereum address not found
- `500 Internal Server Error`: Storage failure

### Transfer Ciphertext Endpoints

#### POST /transfers

Stores buyer and seller ciphertext for a transfer. Must be called by the seller (token owner).

**Request Body:**
```json
{
  "transfer_id": "0x1234567890abcdef...",
  "buyer_ciphertext": "encrypted_data_for_buyer",
  "seller_ciphertext": "encrypted_data_for_seller",
  "message": "Transfer ciphertext for transfer 0x1234...",
  "signature": "0xabcdef..."
}
```

**Authentication:**
- The signature must be signed by the seller's Skavenge private key
- The server verifies the seller owns the token being transferred
- The seller's Skavenge public key must be registered via `/link`

**Success Response (201 Created):**
```json
{
  "success": true,
  "message": "transfer ciphertext stored successfully"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid JSON, missing fields, or invalid transfer ID
- `401 Unauthorized`: Invalid signature
- `404 Not Found`: Transfer not found or seller public key not registered
- `500 Internal Server Error`: Storage or blockchain query failure

#### GET /transfers

Retrieves buyer and seller ciphertext for a transfer. Must be called by the buyer.

**Query Parameters:**
- `transferId` (required): The transfer ID (hex string with 0x prefix)
- `message` (required): Message that was signed
- `signature` (required): Signature from buyer's Skavenge private key

**Example:**
```
GET /transfers?transferId=0x1234...&message=Retrieve%20ciphertext&signature=0xabcdef...
```

**Authentication:**
- The signature must be signed by the buyer's Skavenge private key
- The server verifies the buyer address matches the transfer's buyer
- The buyer's Skavenge public key must be registered via `/link`

**Success Response (200 OK):**
```json
{
  "success": true,
  "buyer_ciphertext": "encrypted_data_for_buyer",
  "seller_ciphertext": "encrypted_data_for_seller"
}
```

**Error Responses:**
- `400 Bad Request`: Missing parameters or invalid transfer ID
- `401 Unauthorized`: Invalid signature
- `404 Not Found`: Transfer not found, buyer public key not registered, or ciphertext not stored
- `500 Internal Server Error`: Storage or blockchain query failure

## Signature Verification

### Ethereum Signatures (Account Linking)

The server verifies signatures using the same method as Web3's `eth.accounts.recover`:

1. The message is hashed with Ethereum's signed message prefix: `\x19Ethereum Signed Message:\n{length}{message}`
2. The public key is recovered from the signature
3. The Ethereum address is derived from the public key
4. The derived address is compared with the provided address (case-insensitive)

This ensures that the signature was created by the owner of the Ethereum address.

### Skavenge Signatures (Transfer Authentication)

For transfer endpoints, the server verifies ECDSA secp256k1 signatures:

1. The message is hashed with SHA-256
2. The signature is parsed from ASN.1 DER format
3. The public key (in uncompressed format: 0x04 || X || Y) is retrieved from the linked accounts storage
4. The signature is verified using ECDSA with the secp256k1 curve (same as Ethereum)

This ensures that only the legitimate buyer or seller can access transfer ciphertext. Using secp256k1 allows the same key to be used for both ElGamal decryption and message signing.

## Architecture

### Storage Layer

The service uses two storage interfaces:

**Account Linkage Storage:**
- `Set(key, value string)`: Stores Ethereum address to Skavenge public key mappings
- `Get(key string) (string, error)`: Retrieves Skavenge public key by Ethereum address

**Transfer Ciphertext Storage:**
- `SetTransferCiphertext(transferID, buyerCiphertext, sellerCiphertext string)`: Stores encrypted data
- `GetTransferCiphertext(transferID string) (buyerCiphertext, sellerCiphertext string, err error)`: Retrieves encrypted data

Both use in-memory maps with mutex protection for thread safety. These can be replaced with persistent storage by implementing the respective interfaces.

### Business Logic Layer

Immutability enforcement happens at the HTTP handler level, not in storage:
- Before storing a new linkage, the handler checks if the address already exists using `Get()`
- If it exists, a 409 Conflict error is returned
- Only if it doesn't exist does the handler call `Set()` to store the new linkage

This separation of concerns makes the storage layer more flexible and reusable.

### Blockchain Integration

The service connects to an Ethereum RPC endpoint to:
- Retrieve transfer information from the Skavenge smart contract
- Verify token ownership
- Validate that transfers exist before storing/retrieving ciphertext

### Project Structure

```
linked-accounts-gateway/
├── main.go              # HTTP server and handlers (business logic)
├── verify.go            # Ethereum signature verification
├── skavenge_verify.go   # Skavenge ECDSA P-256 signature verification
├── contract.go          # Blockchain contract client
├── storage.go           # Storage interfaces and in-memory implementations
├── storage_test.go      # Storage layer tests
├── verify_test.go       # Ethereum signature verification tests
├── transfers_test.go    # Transfer endpoint tests
└── README.md            # This file
```

## Testing

Run all tests:
```bash
go test -v
```

Run specific test suites:
```bash
# Test account linking
go test -v -run TestVerifySignature

# Test transfer endpoints
go test -v -run TestHandle

# Test Skavenge signature verification
go test -v -run TestVerifySkavengeSignature
```

### Using curl

1. **Link an account** (you'll need a real signature from MetaMask):
```bash
curl -X POST http://localhost:4591/link \
  -H "Content-Type: application/json" \
  -d '{
    "ethereumAddress": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "skavengePublicKey": "0x04a1b2c3d4...",
    "message": "Link MetaMask address 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb to Skavenge key",
    "signature": "0x..."
  }'
```

Note: The `skavengePublicKey` should be in uncompressed secp256k1 format (65 bytes: 0x04 || X || Y).

2. **Retrieve a linked account**:
```bash
curl "http://localhost:4591/link?ethereumAddress=0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
```

3. **Store transfer ciphertext** (as seller):
```bash
curl -X POST http://localhost:4591/transfers \
  -H "Content-Type: application/json" \
  -d '{
    "transfer_id": "0x1234567890abcdef...",
    "buyer_ciphertext": "encrypted_buyer_data",
    "seller_ciphertext": "encrypted_seller_data",
    "message": "Transfer ciphertext for transfer 0x1234...",
    "signature": "0x..."
  }'
```

4. **Retrieve transfer ciphertext** (as buyer):
```bash
curl "http://localhost:4591/transfers?transferId=0x1234...&message=Retrieve%20ciphertext&signature=0x..."
```

## Future Enhancements

- Add persistent storage implementation (PostgreSQL, Redis, etc.)
- Add authentication/authorization for administrative endpoints
- Add rate limiting
- Add logging middleware
- Add metrics and monitoring
- Add CORS configuration for web client integration
- Add health check endpoint
