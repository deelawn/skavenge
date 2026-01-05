package minting

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// ClueData represents the data needed to mint a clue
type ClueData struct {
	Content     string   // The clue content (plaintext)
	Solution    string   // The solution to the clue
	PointValue  uint8    // Point value (1-5)
	SolveReward *big.Int // ETH reward for solving (in wei), optional
}

// MintOptions contains optional parameters for minting
type MintOptions struct {
	// For listing the clue for sale immediately after minting
	SalePrice *big.Int // Price in wei, 0 or nil means not for sale
	Timeout   uint64   // Transfer timeout in seconds (required if SalePrice > 0)

	// For minting and sending to another address
	RecipientAddress string // Ethereum address of recipient
}

// MintResult contains the result of a minting operation
type MintResult struct {
	TokenID     *big.Int
	TxHash      string
	ClueContent string
	Solution    string
	PointValue  uint8
	Error       error
}

// Config holds the configuration for the minter
type Config struct {
	// Blockchain connection
	RPCURL          string
	ContractAddress string

	// Minter credentials
	MinterPrivateKey string // Ethereum private key (authorized minter)

	// Skavenge encryption key
	SkavengePublicKey string // Skavenge public key for encryption (hex format, optional - only needed when minting to self)

	// Gateway service
	GatewayURL string // URL of the linked accounts gateway
}

// RecipientInfo contains the recipient's information for encrypted minting
type RecipientInfo struct {
	EthereumAddress   common.Address
	SkavengePublicKey *ecdsa.PublicKey
}
