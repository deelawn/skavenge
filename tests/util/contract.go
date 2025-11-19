// Package util provides utility functions for the Skavenge contract tests.
package util

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/deelawn/skavenge/eth/bindings"
)

// GetHardhatURL returns the Hardhat URL from environment variable or default.
func GetHardhatURL() string {
	url := os.Getenv("HARDHAT_URL")
	if url == "" {
		url = "http://localhost:8545"
	}
	return url
}

// DeployContract deploys a fresh instance of the Skavenge contract for testing.
// Returns the contract instance, the contract address, and any error encountered.
func DeployContract(client *ethclient.Client, auth *bind.TransactOpts) (*bindings.Skavenge, common.Address, error) {
	// Use deployer as initial minter
	initialMinter := auth.From
	address, tx, contract, err := bindings.DeploySkavenge(auth, client, initialMinter)
	if err != nil {
		return nil, common.Address{}, err
	}

	// Wait for the transaction to be mined
	if _, err := WaitForTransaction(client, tx); err != nil {
		return nil, common.Address{}, err
	}

	return contract, address, nil
}

// NewTransactOpts creates a new set of transaction options with the given private key.
func NewTransactOpts(client *ethclient.Client, privateKeyHex string) (*bind.TransactOpts, error) {
	// Load the private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	// Get the public address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	// Get the chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	// Create the auth
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)             // in wei
	auth.GasPrice = big.NewInt(1000000000) // in wei (1 gwei)

	return auth, nil
}

// WaitForTransaction waits for a transaction to be mined and returns an error if it failed.
func WaitForTransaction(client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return nil, err
	}

	if receipt.Status == 0 {
		return receipt, ErrTxFailed
	}

	return receipt, nil
}

// ErrTxFailed is returned when a transaction fails.
var ErrTxFailed = errors.New("transaction failed")

// LoadPrivateKey loads a private key from a hexadecimal string.
func LoadPrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(privateKeyHex)
}

// DerivePublicKey derives a public key from a private key.
func DerivePublicKey(privateKey *ecdsa.PrivateKey) string {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return ""
	}

	// Convert the public key to bytes and then to hex
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	return hex.EncodeToString(publicKeyBytes)
}

// GetAddress derives the Ethereum address from a private key.
func GetAddress(privateKey *ecdsa.PrivateKey) common.Address {
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	return crypto.PubkeyToAddress(*publicKeyECDSA)
}
