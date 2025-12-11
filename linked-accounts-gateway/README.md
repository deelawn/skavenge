# Linked Accounts Gateway

A Go web server that provides a gateway for storing and retrieving linked account information, mapping Ethereum addresses to Skavenge public keys.

## Overview

This service allows users to link their Ethereum addresses (from MetaMask) to their Skavenge public keys through cryptographic signature verification. Once linked, the associations are immutable and can be retrieved via the API.

## Features

- **Signature Verification**: Uses Ethereum's `personal_sign` method to verify ownership of Ethereum addresses
- **Immutable Storage**: Once an Ethereum address is linked to a Skavenge key, it cannot be modified
- **In-Memory Storage**: Current implementation uses an in-memory map (can be replaced with persistent storage via the Storage interface)
- **RESTful API**: Simple HTTP endpoints for linking and retrieving account information

## Installation

1. Install dependencies:
```bash
go mod tidy
```

2. Run the server:
```bash
go run main.go verify.go
```

The server will start on port 8080 by default.

## API Endpoints

### POST /link

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

## Signature Verification

The server verifies signatures using the same method as Web3's `eth.accounts.recover`:

1. The message is hashed with Ethereum's signed message prefix: `\x19Ethereum Signed Message:\n{length}{message}`
2. The public key is recovered from the signature
3. The Ethereum address is derived from the public key
4. The derived address is compared with the provided address (case-insensitive)

This ensures that the signature was created by the owner of the Ethereum address.

## Architecture

### Storage Layer

The storage layer is defined as an interface with two methods:
- `Set(key, value string) error`: Stores a key-value pair, returns error if key exists
- `Get(key string) (string, error)`: Retrieves value by key, returns error if not found

The current implementation uses an in-memory map with mutex protection for thread safety. This can easily be replaced with a persistent storage implementation (e.g., database, Redis) by implementing the `Storage` interface.

### Project Structure

```
linked-accounts-gateway/
├── main.go              # HTTP server and handlers
├── verify.go            # Ethereum signature verification
├── storage/
│   └── storage.go       # Storage interface and in-memory implementation
├── go.mod               # Go module definition
└── README.md            # This file
```

## Testing Example

### Using curl

1. **Link an account** (you'll need a real signature from MetaMask):
```bash
curl -X POST http://localhost:8080/link \
  -H "Content-Type: application/json" \
  -d '{
    "ethereumAddress": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "skavengePublicKey": "skv1abc123",
    "message": "Link MetaMask address 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb to Skavenge key skv1abc123",
    "signature": "0x..."
  }'
```

2. **Retrieve a linked account**:
```bash
curl "http://localhost:8080/link?ethereumAddress=0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
```

## Future Enhancements

- Add persistent storage implementation (PostgreSQL, Redis, etc.)
- Add authentication/authorization for administrative endpoints
- Add rate limiting
- Add logging middleware
- Add metrics and monitoring
- Add CORS configuration for web client integration
- Add health check endpoint
