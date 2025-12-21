package zkproof

import (
	"bytes"
	"crypto/rand"
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

	// Create "original" cipher (as if minted on-chain)
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	if err != nil {
		t.Fatalf("Failed to generate mintR: %v", err)
	}
	originalCipher, err := ps.EncryptElGamal(plaintext, &sellerKey.PublicKey, mintR)
	if err != nil {
		t.Fatalf("Failed to encrypt original cipher: %v", err)
	}

	// Generate verifiable transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	if err != nil {
		t.Fatalf("Failed to generate transfer: %v", err)
	}

	// Verify the transfer (BEFORE payment)
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.DLEQProof,
		transfer.PlaintextProof,
		mintR,
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

	// Create "original" cipher (as if minted on-chain)
	mintR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	originalCipher, _ := ps.EncryptElGamal(plaintext, &sellerKey.PublicKey, mintR)

	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	if err != nil {
		t.Fatalf("Failed to generate transfer: %v", err)
	}

	// Valid proof should verify
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.DLEQProof,
		transfer.PlaintextProof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	if !valid {
		t.Fatal("Valid proof should verify")
	}

	// Tamper with A1
	tamperedProof := &DLEQProof{
		A1:    []byte{1, 2, 3, 4}, // Invalid point
		A2:    transfer.DLEQProof.A2,
		A3:    transfer.DLEQProof.A3,
		Z:     transfer.DLEQProof.Z,
		C:     transfer.DLEQProof.C,
		RHash: transfer.DLEQProof.RHash,
	}

	valid = ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		tamperedProof,
		transfer.PlaintextProof,
		mintR,
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

	// Create "original" cipher (as if minted on-chain)
	mintR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	originalCipher, _ := ps.EncryptElGamal(plaintext, &buyerKey.PublicKey, mintR)

	// Generate transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		originalCipher,
		mintR,
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

	// Create "original" cipher (as if minted on-chain)
	mintR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	originalCipher, _ := ps.EncryptElGamal(plaintext, &buyerKey.PublicKey, mintR)

	// Generate transfer for buyer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		originalCipher,
		mintR,
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

// TestPlaintextEqualityProofDetectsAlteredContent verifies that altered content is detected
func TestPlaintextEqualityProofDetectsAlteredContent(t *testing.T) {
	ps := NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	originalContent := []byte("The treasure is buried under the old oak tree")
	alteredContent := []byte("FAKE CONTENT - this is not the real clue!")

	// Create "original" cipher (as if minted on-chain with the correct content)
	mintR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	originalCipher, _ := ps.EncryptElGamal(originalContent, &sellerKey.PublicKey, mintR)

	// ATTACK: Seller generates transfer with ALTERED content
	// The seller provides the wrong content hoping buyer won't notice
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		alteredContent, // ATTACK: different from originalContent!
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	if err != nil {
		t.Fatalf("Failed to generate transfer: %v", err)
	}

	// ATTACK DETECTION: Buyer verifies the proof
	// This should FAIL because the content doesn't match
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.DLEQProof,
		transfer.PlaintextProof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)

	if valid {
		t.Fatal("❌ SECURITY FAILURE: Altered content was NOT detected!")
	}

	t.Log("✅ ATTACK PREVENTED: Altered content was detected")
	t.Log("   Plaintext equality proof correctly rejected fraudulent transfer")
}

// TestPlaintextEqualityProofAcceptsCorrectContent verifies correct content passes verification
func TestPlaintextEqualityProofAcceptsCorrectContent(t *testing.T) {
	ps := NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	content := []byte("The treasure is buried under the old oak tree")

	// Create "original" cipher (as if minted on-chain)
	mintR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	originalCipher, _ := ps.EncryptElGamal(content, &sellerKey.PublicKey, mintR)

	// Honest seller generates transfer with CORRECT content
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		content, // Same as original content
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	if err != nil {
		t.Fatalf("Failed to generate transfer: %v", err)
	}

	// Buyer verifies the proof - should PASS
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.DLEQProof,
		transfer.PlaintextProof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)

	if !valid {
		t.Fatal("❌ Correct content should pass verification!")
	}

	t.Log("✅ Correct content passed verification")

	// Buyer should be able to decrypt and get the original content
	decrypted, err := ps.DecryptElGamal(transfer.BuyerCipher, transfer.SharedR, buyerKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	if !bytes.Equal(content, decrypted) {
		t.Fatal("❌ Decrypted content doesn't match original!")
	}

	t.Log("✅ Buyer received correct content after decryption")
	t.Log("✅ Complete honest transfer flow verified!")
}

// TestPlaintextEqualityProofMarshalUnmarshal tests marshaling/unmarshaling
func TestPlaintextEqualityProofMarshalUnmarshal(t *testing.T) {
	ps := NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	content := []byte("Test marshaling content")

	// Create cipher and transfer
	mintR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	originalCipher, _ := ps.EncryptElGamal(content, &sellerKey.PublicKey, mintR)

	transfer, _ := ps.GenerateVerifiableElGamalTransfer(
		content,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)

	// Marshal the plaintext proof
	proofBytes := transfer.PlaintextProof.Marshal()

	// Unmarshal into new struct
	unmarshaled := &CrossRPlaintextEqualityProof{}
	err := unmarshaled.Unmarshal(proofBytes)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Verify the unmarshaled proof matches original
	if !bytes.Equal(transfer.PlaintextProof.KeyBuyer[:], unmarshaled.KeyBuyer[:]) {
		t.Fatal("Unmarshaled KeyBuyer doesn't match original")
	}

	// Verify unmarshaled proof still works for verification
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.DLEQProof,
		unmarshaled, // Use unmarshaled proof
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)

	if !valid {
		t.Fatal("Unmarshaled proof should verify!")
	}

	t.Log("✅ Marshal/Unmarshal works correctly")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
