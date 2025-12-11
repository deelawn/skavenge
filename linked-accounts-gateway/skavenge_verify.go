package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/asn1"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// VerifySkavengeSignature verifies that the signature was created using a Skavenge private key
// corresponding to the provided public key
// The Skavenge extension uses ECDSA secp256k1 (same curve as Ethereum)
func VerifySkavengeSignature(message, signature, publicKeyHex string) (bool, error) {
	// Decode the hex signature
	sig, err := hexutil.Decode(signature)
	if err != nil {
		return false, fmt.Errorf("invalid signature format: %w", err)
	}

	// Parse the public key from hex (uncompressed secp256k1 format: 0x04 || X || Y)
	pubKeyBytes, err := hexutil.Decode(publicKeyHex)
	if err != nil {
		return false, fmt.Errorf("invalid public key format: %w", err)
	}

	// Parse the uncompressed public key (65 bytes: 0x04 + 32 bytes X + 32 bytes Y)
	if len(pubKeyBytes) != 65 || pubKeyBytes[0] != 0x04 {
		return false, fmt.Errorf("invalid public key format: expected 65 bytes starting with 0x04, got %d bytes", len(pubKeyBytes))
	}

	// Extract X and Y coordinates
	x := new(big.Int).SetBytes(pubKeyBytes[1:33])
	y := new(big.Int).SetBytes(pubKeyBytes[33:65])

	// Create secp256k1 public key
	publicKey := &ecdsa.PublicKey{
		Curve: crypto.S256(), // secp256k1
		X:     x,
		Y:     y,
	}

	// Verify the point is on the curve
	if !publicKey.Curve.IsOnCurve(x, y) {
		return false, fmt.Errorf("public key point is not on the secp256k1 curve")
	}

	// Parse the signature (ASN.1 DER format)
	var sigStruct struct {
		R, S *big.Int
	}
	_, err = asn1.Unmarshal(sig, &sigStruct)
	if err != nil {
		return false, fmt.Errorf("failed to parse signature: %w", err)
	}

	// Hash the message with SHA-256
	hash := sha256.Sum256([]byte(message))

	// Verify the signature using secp256k1
	valid := ecdsa.Verify(publicKey, hash[:], sigStruct.R, sigStruct.S)
	return valid, nil
}
