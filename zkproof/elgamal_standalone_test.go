package zkproof

import (
	"bytes"
	"testing"
)

// TestDLEQProofStandalone tests DLEQ proof without external dependencies
func TestDLEQProofStandalone(t *testing.T) {
	ps := NewProofSystem()

	// Generate keys
	sellerKey, err := ps.GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate seller key: %v", err)
	}

	buyerKey, err := ps.GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate buyer key: %v", err)
	}

	plaintext := []byte("Test message for DLEQ proof")

	// Generate verifiable transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		sellerKey,
		&buyerKey.PublicKey,
	)
	if err != nil {
		t.Fatalf("Failed to generate transfer: %v", err)
	}

	// Verify the transfer (BEFORE payment)
	valid := ps.VerifyElGamalTransfer(
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.DLEQProof,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	if !valid {
		t.Fatal("DLEQ proof verification failed!")
	}

	t.Log("✅ DLEQ proof verified successfully")

	// Verify buyer can decrypt with r
	decrypted, err := ps.DecryptElGamal(
		transfer.BuyerCipher,
		transfer.SharedR,
		buyerKey,
	)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Fatalf("Decrypted plaintext doesn't match!\nExpected: %s\nGot: %s", plaintext, decrypted)
	}

	t.Log("✅ Decryption successful")

	// Verify commitment
	commitmentValid := ps.VerifyPlaintextMatchesCommitment(
		decrypted,
		transfer.Salt,
		transfer.Commitment,
	)
	if !commitmentValid {
		t.Fatal("Commitment verification failed!")
	}

	t.Log("✅ Commitment verified")
	t.Log("✅ Complete flow successful!")
}

// TestDLEQProofRejectsTampering verifies tampered proofs are rejected
func TestDLEQProofRejectsTampering(t *testing.T) {
	ps := NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	plaintext := []byte("Test message")

	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		sellerKey,
		&buyerKey.PublicKey,
	)
	if err != nil {
		t.Fatalf("Failed to generate transfer: %v", err)
	}

	// Valid proof should verify
	valid := ps.VerifyElGamalTransfer(
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.DLEQProof,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	if !valid {
		t.Fatal("Valid proof should verify")
	}

	// Tamper with A1
	tamperedProof := &DLEQProof{
		A1: []byte{1, 2, 3, 4}, // Invalid point
		A2: transfer.DLEQProof.A2,
		A3: transfer.DLEQProof.A3,
		Z:  transfer.DLEQProof.Z,
		C:  transfer.DLEQProof.C,
	}

	valid = ps.VerifyElGamalTransfer(
		transfer.SellerCipher,
		transfer.BuyerCipher,
		tamperedProof,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	if valid {
		t.Fatal("Tampered proof should be rejected!")
	}

	t.Log("✅ Tampered proof correctly rejected")
}
