package zkproof

import (
	"bytes"
	"math/big"
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

// TestDecryptionRequiresCorrectR verifies buyer cannot decrypt without correct r
func TestDecryptionRequiresCorrectR(t *testing.T) {
	ps := NewProofSystem()

	buyerKey, _ := ps.GenerateKeyPair()
	plaintext := []byte("Secret message that requires r to decrypt")

	// Generate transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		buyerKey, // Seller is same as buyer for this test
		&buyerKey.PublicKey,
	)
	if err != nil {
		t.Fatalf("Failed to generate transfer: %v", err)
	}

	// Try to decrypt with WRONG r - should fail or produce garbage
	wrongR := new(big.Int).SetInt64(99999)
	wrongDecrypted, err := ps.DecryptElGamal(
		transfer.BuyerCipher,
		wrongR,
		buyerKey,
	)

	if err == nil {
		// Should NOT match original plaintext
		if bytes.Equal(plaintext, wrongDecrypted) {
			t.Fatal("❌ SECURITY FAILURE: Wrong r decrypted correctly!")
		}
		t.Logf("✅ Wrong r produced garbage: %x (expected)", wrongDecrypted[:min(16, len(wrongDecrypted))])
	}

	// Try to decrypt with nil r - should error
	_, err = ps.DecryptElGamal(
		transfer.BuyerCipher,
		nil,
		buyerKey,
	)
	if err == nil {
		t.Fatal("❌ Should reject nil r value")
	}
	t.Log("✅ Nil r correctly rejected")

	// Decrypt with correct r - should succeed
	correctDecrypted, err := ps.DecryptElGamal(
		transfer.BuyerCipher,
		transfer.SharedR,
		buyerKey,
	)
	if err != nil {
		t.Fatalf("Failed to decrypt with correct r: %v", err)
	}

	if !bytes.Equal(plaintext, correctDecrypted) {
		t.Fatal("❌ Correct r should decrypt successfully")
	}

	t.Log("✅ Correct r decrypted successfully")
	t.Log("✅ SECURITY VERIFIED: Decryption requires correct r value")
}

// TestDecryptionRequiresPrivateKey verifies that even with r, you need the private key
func TestDecryptionRequiresPrivateKey(t *testing.T) {
	ps := NewProofSystem()

	buyerKey, _ := ps.GenerateKeyPair()
	attackerKey, _ := ps.GenerateKeyPair() // Different private key
	plaintext := []byte("Secret message only buyer should decrypt")

	// Generate transfer for buyer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		buyerKey, // Seller is same as buyer for this test
		&buyerKey.PublicKey,
	)
	if err != nil {
		t.Fatalf("Failed to generate transfer: %v", err)
	}

	// Attacker has access to:
	// 1. The buyer's ciphertext (public)
	// 2. The revealed r value (revealed on-chain)
	// 3. The buyer's public key (public)
	// But attacker does NOT have buyer's private key

	// Try to decrypt with correct r but WRONG private key
	attackerDecrypted, err := ps.DecryptElGamal(
		transfer.BuyerCipher,
		transfer.SharedR, // Correct r (revealed on-chain)
		attackerKey,      // Wrong private key!
	)

	if err == nil {
		// Should NOT match original plaintext
		if bytes.Equal(plaintext, attackerDecrypted) {
			t.Fatal("❌ PRIVACY FAILURE: Attacker decrypted with wrong private key!")
		}
		t.Logf("✅ Wrong private key produced garbage: %x (expected)", attackerDecrypted[:min(16, len(attackerDecrypted))])
	} else {
		t.Logf("✅ Decryption with wrong private key failed: %v", err)
	}

	// Verify buyer CAN decrypt with correct r and correct private key
	buyerDecrypted, err := ps.DecryptElGamal(
		transfer.BuyerCipher,
		transfer.SharedR,
		buyerKey, // Correct private key
	)
	if err != nil {
		t.Fatalf("Buyer should be able to decrypt: %v", err)
	}

	if !bytes.Equal(plaintext, buyerDecrypted) {
		t.Fatal("❌ Buyer with correct private key should decrypt successfully")
	}

	t.Log("✅ Buyer with correct private key decrypted successfully")
	t.Log("✅ PRIVACY VERIFIED: Decryption requires buyer's private key")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
