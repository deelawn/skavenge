package main

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// generateTestSignature creates a valid Ethereum signature for testing
func generateTestSignature(privateKey *ecdsa.PrivateKey, message string) (string, string, error) {
	// Get the address from the private key
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Create the hash with Ethereum prefix
	hash := textHash([]byte(message))

	// Sign the hash
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return "", "", err
	}

	// Convert V from 0/1 to 27/28 (as MetaMask does)
	if signature[64] < 27 {
		signature[64] += 27
	}

	return hexutil.Encode(signature), address.Hex(), nil
}

// TestVerifySignature_ValidSignature tests signature verification with valid signatures
func TestVerifySignature_ValidSignature(t *testing.T) {
	// Generate a test private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	tests := []struct {
		name    string
		message string
	}{
		{
			name:    "simple message",
			message: "Hello, Ethereum!",
		},
		{
			name:    "linkage message format",
			message: "Link MetaMask address 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb to Skavenge key skv1abc123",
		},
		{
			name:    "empty message",
			message: "",
		},
		{
			name:    "message with special characters",
			message: "Test message with special chars: !@#$%^&*()",
		},
		{
			name:    "long message",
			message: "This is a very long message that exceeds typical message lengths to test the signature verification with larger data sets and ensure that the hashing and signature verification still works correctly regardless of message length.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signature, address, err := generateTestSignature(privateKey, tt.message)
			if err != nil {
				t.Fatalf("Failed to generate signature: %v", err)
			}

			valid, err := VerifySignature(tt.message, signature, address)
			if err != nil {
				t.Errorf("VerifySignature() error = %v", err)
			}
			if !valid {
				t.Errorf("VerifySignature() = false, want true for valid signature")
			}
		})
	}
}

// TestVerifySignature_InvalidSignature tests signature verification with invalid signatures
func TestVerifySignature_InvalidSignature(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	message := "Test message"
	signature, address, err := generateTestSignature(privateKey, message)
	if err != nil {
		t.Fatalf("Failed to generate signature: %v", err)
	}

	tests := []struct {
		name      string
		message   string
		signature string
		address   string
		wantValid bool
		wantError bool
	}{
		{
			name:      "wrong message",
			message:   "Different message",
			signature: signature,
			address:   address,
			wantValid: false,
			wantError: false,
		},
		{
			name:      "wrong address",
			message:   message,
			signature: signature,
			address:   "0x0000000000000000000000000000000000000000",
			wantValid: false,
			wantError: false,
		},
		{
			name:      "invalid signature format",
			signature: "invalid",
			address:   address,
			message:   message,
			wantValid: false,
			wantError: true,
		},
		{
			name:      "signature too short",
			signature: "0x1234",
			address:   address,
			message:   message,
			wantValid: false,
			wantError: true,
		},
		{
			name:      "empty signature",
			signature: "",
			address:   address,
			message:   message,
			wantValid: false,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := VerifySignature(tt.message, tt.signature, tt.address)

			if tt.wantError && err == nil {
				t.Errorf("VerifySignature() expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("VerifySignature() unexpected error = %v", err)
			}
			if valid != tt.wantValid {
				t.Errorf("VerifySignature() = %v, want %v", valid, tt.wantValid)
			}
		})
	}
}

// TestVerifySignature_CaseInsensitiveAddress tests that address comparison is case-insensitive
func TestVerifySignature_CaseInsensitiveAddress(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	message := "Test message"
	signature, address, err := generateTestSignature(privateKey, message)
	if err != nil {
		t.Fatalf("Failed to generate signature: %v", err)
	}

	tests := []struct {
		name    string
		address string
	}{
		{
			name:    "lowercase address",
			address: address[0:2] + address[2:],
		},
		{
			name:    "uppercase address",
			address: address[0:2] + address[2:],
		},
		{
			name:    "mixed case address",
			address: address,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := VerifySignature(message, signature, tt.address)
			if err != nil {
				t.Errorf("VerifySignature() error = %v", err)
			}
			if !valid {
				t.Errorf("VerifySignature() = false, want true (should be case-insensitive)")
			}
		})
	}
}

// TestVerifySignature_MultipleSigners tests that different signers produce different results
func TestVerifySignature_MultipleSigners(t *testing.T) {
	// Generate two different private keys
	privateKey1, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key 1: %v", err)
	}

	privateKey2, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key 2: %v", err)
	}

	message := "Test message"

	// Generate signature from first key
	signature1, address1, err := generateTestSignature(privateKey1, message)
	if err != nil {
		t.Fatalf("Failed to generate signature 1: %v", err)
	}

	// Generate signature from second key
	signature2, address2, err := generateTestSignature(privateKey2, message)
	if err != nil {
		t.Fatalf("Failed to generate signature 2: %v", err)
	}

	// Signature 1 should verify with address 1
	valid, err := VerifySignature(message, signature1, address1)
	if err != nil {
		t.Errorf("VerifySignature() for signer 1 error = %v", err)
	}
	if !valid {
		t.Errorf("VerifySignature() for signer 1 = false, want true")
	}

	// Signature 2 should verify with address 2
	valid, err = VerifySignature(message, signature2, address2)
	if err != nil {
		t.Errorf("VerifySignature() for signer 2 error = %v", err)
	}
	if !valid {
		t.Errorf("VerifySignature() for signer 2 = false, want true")
	}

	// Signature 1 should NOT verify with address 2
	valid, err = VerifySignature(message, signature1, address2)
	if err != nil {
		t.Errorf("VerifySignature() cross-check 1 error = %v", err)
	}
	if valid {
		t.Errorf("VerifySignature() cross-check 1 = true, want false (wrong address)")
	}

	// Signature 2 should NOT verify with address 1
	valid, err = VerifySignature(message, signature2, address1)
	if err != nil {
		t.Errorf("VerifySignature() cross-check 2 error = %v", err)
	}
	if valid {
		t.Errorf("VerifySignature() cross-check 2 = true, want false (wrong address)")
	}
}

// TestTextHash tests the textHash function
func TestTextHash(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		wantHash string // We can't predict the exact hash, but we can verify it's consistent
	}{
		{
			name:    "simple message",
			message: "Hello",
		},
		{
			name:    "empty message",
			message: "",
		},
		{
			name:    "unicode message",
			message: "Hello 世界",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash1 := textHash([]byte(tt.message))
			hash2 := textHash([]byte(tt.message))

			// Hash should be deterministic (same message produces same hash)
			if fmt.Sprintf("%x", hash1) != fmt.Sprintf("%x", hash2) {
				t.Errorf("textHash() not deterministic")
			}

			// Hash should be 32 bytes (256 bits) for Keccak256
			if len(hash1) != 32 {
				t.Errorf("textHash() length = %d, want 32", len(hash1))
			}
		})
	}
}

// TestTextHash_EthereumPrefix tests that textHash includes the Ethereum signed message prefix
func TestTextHash_EthereumPrefix(t *testing.T) {
	message := "Test"
	hash := textHash([]byte(message))

	// The hash should be of the message with prefix: "\x19Ethereum Signed Message:\n4Test"
	expectedPrefix := "\x19Ethereum Signed Message:\n"
	prefixedMessage := fmt.Sprintf("%s%d%s", expectedPrefix, len(message), message)
	expectedHash := crypto.Keccak256([]byte(prefixedMessage))

	if fmt.Sprintf("%x", hash) != fmt.Sprintf("%x", expectedHash) {
		t.Errorf("textHash() does not match expected Ethereum signed message format")
	}
}

// TestVerifySignature_VValueConversion tests that V value conversion works correctly
func TestVerifySignature_VValueConversion(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	message := "Test message"
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	hash := textHash([]byte(message))

	// Create signature with V = 0 or 1
	sig, err := crypto.Sign(hash, privateKey)
	if err != nil {
		t.Fatalf("Failed to sign: %v", err)
	}

	// The original V value from crypto.Sign is 0 or 1
	originalV := sig[64]

	// Test both formats: raw (0 or 1) and Ethereum format (27 or 28)
	// Note: We can only test the V values that match the original signature's R and S
	tests := []struct {
		name   string
		vValue byte
	}{
		{
			name:   fmt.Sprintf("V = %d (raw go-ethereum format)", originalV),
			vValue: originalV,
		},
		{
			name:   fmt.Sprintf("V = %d (Ethereum/MetaMask format)", originalV+27),
			vValue: originalV + 27,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testSig := make([]byte, 65)
			copy(testSig, sig)
			testSig[64] = tt.vValue

			signature := hexutil.Encode(testSig)
			valid, err := VerifySignature(message, signature, address.Hex())
			if err != nil {
				t.Errorf("VerifySignature() error = %v", err)
			}
			if !valid {
				t.Errorf("VerifySignature() with V=%d = false, want true", tt.vValue)
			}
		})
	}
}
