package main

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// VerifySignature verifies that the signature was created by the owner of the address
// This replicates the Web3 eth.accounts.recover functionality
func VerifySignature(message, signature, address string) (bool, error) {
	// Decode the hex signature
	if !strings.HasPrefix(signature, "0x") {
		signature = "0x" + signature
	}

	sig, err := hexutil.Decode(signature)
	if err != nil {
		return false, fmt.Errorf("invalid signature format: %w", err)
	}

	// Ethereum signatures are 65 bytes: R (32) + S (32) + V (1)
	if len(sig) != 65 {
		return false, fmt.Errorf("invalid signature length: expected 65, got %d", len(sig))
	}

	// The V value in MetaMask's personal_sign is 27 or 28, but go-ethereum expects 0 or 1
	// Convert if necessary
	if sig[64] >= 27 {
		sig[64] -= 27
	}

	// Create the hash of the message with Ethereum's signed message prefix
	hash := textHash([]byte(message))

	// Recover the public key from the signature
	pubKey, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}

	// Derive the address from the public key
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	// Compare addresses (case-insensitive)
	expectedAddr := common.HexToAddress(address)
	return recoveredAddr.Hex() == expectedAddr.Hex(), nil
}

// textHash creates the hash of a message with the Ethereum signed message prefix
// This replicates the behavior of personal_sign in MetaMask
func textHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
