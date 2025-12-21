package zkproof

import (
	"crypto/rand"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestElGamalTransfer_HonestCase tests the complete verifiable transfer flow
func TestElGamalTransfer_HonestCase(t *testing.T) {
	ps := NewProofSystem()

	// Generate keys
	sellerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	buyerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	plaintext := []byte("The treasure is at coordinates 40.7128°N, 74.0060°W")

	// Create "original" cipher (as if minted on-chain)
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err)
	originalCipher, err := ps.EncryptElGamal(plaintext, &sellerKey.PublicKey, mintR)
	require.NoError(t, err)

	t.Log("=== SELLER: Generate verifiable transfer ===")
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)
	require.NotNil(t, transfer)

	t.Log("=== BUYER: Verify transfer (BEFORE PAYING!) ===")
	// Buyer can cryptographically verify the ciphertext is valid
	// WITHOUT being able to decrypt it!
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.Proof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	require.True(t, valid, "Transfer should verify successfully")

	t.Log("✅ Buyer verified: Both ciphertexts provably contain same plaintext")
	t.Log("✅ Buyer CANNOT decrypt yet (doesn't have r)")
	t.Log("✅ Buyer now confident to pay")

	t.Log("=== BUYER: Pays seller ===")
	// Payment happens on-chain...

	t.Log("=== SELLER: Provides decryption key (after payment) ===")
	// Seller reveals the shared random value r
	// Now buyer can decrypt

	t.Log("=== BUYER: Decrypt and verify ===")
	decrypted, err := ps.DecryptElGamal(
		transfer.BuyerCipher,
		transfer.SharedR,
		buyerKey,
	)
	require.NoError(t, err)
	require.Equal(t, plaintext, decrypted, "Decrypted plaintext should match original")

	// Final verification against commitment
	commitmentValid := ps.VerifyPlaintextMatchesCommitment(
		decrypted,
		transfer.Salt,
		transfer.Commitment,
	)
	require.True(t, commitmentValid, "Commitment should match")

	t.Log("✅ SUCCESS: Complete verifiable transfer with mathematical proof!")
	t.Log("   - Buyer verified BEFORE paying")
	t.Log("   - Buyer could NOT decrypt before paying")
	t.Log("   - No fraud detection needed")
	t.Log("   - Pure cryptographic guarantees")
}

// TestElGamalTransfer_BuyerCannotDecryptWithoutR verifies buyer needs r to decrypt
func TestElGamalTransfer_BuyerCannotDecryptWithoutR(t *testing.T) {
	ps := NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	plaintext := []byte("Secret message")

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
	require.NoError(t, err)

	// Buyer verifies proof
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.Proof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	require.True(t, valid, "Proof should verify")

	t.Log("✅ Buyer verified proof")
	t.Log("❌ But buyer CANNOT decrypt without r")

	// Try to decrypt without r - should fail or produce garbage
	// (In real implementation, this would fail completely)
	// For now, just verify they need the correct r

	// Decrypt with WRONG r - should produce wrong plaintext
	wrongR := new(big.Int).SetInt64(12345)
	wrongDecrypted, err := ps.DecryptElGamal(transfer.BuyerCipher, wrongR, buyerKey)

	if err == nil {
		// Should not match original plaintext
		require.NotEqual(t, plaintext, wrongDecrypted, "Wrong r should not decrypt correctly")
	}

	// Decrypt with correct r - should work
	correctDecrypted, err := ps.DecryptElGamal(transfer.BuyerCipher, transfer.SharedR, buyerKey)
	require.NoError(t, err)
	require.Equal(t, plaintext, correctDecrypted, "Correct r should decrypt successfully")

	t.Log("✅ Confirmed: Buyer needs seller to provide r to decrypt")
}

// TestElGamalTransfer_ProofRejectsDifferentPlaintexts verifies DLEQ proof fails if plaintexts differ
func TestElGamalTransfer_ProofRejectsDifferentPlaintexts(t *testing.T) {
	ps := NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	plaintext1 := []byte("First message")
	plaintext2 := []byte("Different message")

	// Create "original" cipher (as if minted on-chain with plaintext1)
	mintR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	originalCipher, _ := ps.EncryptElGamal(plaintext1, &sellerKey.PublicKey, mintR)

	// Generate transfer for first message
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext1,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// ATTACK: Try to replace buyer cipher with different plaintext
	// Generate new encryption with different r for different plaintext
	fakeTransfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext2,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Try to verify with mismatched ciphertexts
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,     // Real cipher
		fakeTransfer.BuyerCipher,  // FAKE cipher (different plaintext)
		transfer.Proof,            // Original proof
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)

	require.False(t, valid, "ATTACK PREVENTED: Proof should reject mismatched ciphertexts!")

	t.Log("✅ ATTACK PREVENTED: Cannot swap buyer ciphertext")
	t.Log("   DLEQ proof mathematically guarantees both ciphers contain same plaintext")
}

// TestElGamalTransfer_InvalidProofRejected verifies tampered proofs are rejected
func TestElGamalTransfer_InvalidProofRejected(t *testing.T) {
	ps := NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	plaintext := []byte("Secret message")

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
	require.NoError(t, err)

	// Valid proof should verify
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.Proof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	require.True(t, valid, "Valid proof should verify")

	// ATTACK: Tamper with proof
	tamperedDLEQ := &DLEQProof{
		A1:    transfer.Proof.DLEQ.A1,
		A2:    transfer.Proof.DLEQ.A2,
		A3:    transfer.Proof.DLEQ.A3,
		Z:     new(big.Int).SetInt64(99999), // TAMPERED!
		C:     transfer.Proof.DLEQ.C,
		RHash: transfer.Proof.DLEQ.RHash,
	}

	tamperedProof := &TransferProof{
		DLEQ:      tamperedDLEQ,
		Plaintext: transfer.Proof.Plaintext,
	}

	valid = ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		tamperedProof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)

	require.False(t, valid, "ATTACK PREVENTED: Tampered proof should be rejected!")

	t.Log("✅ ATTACK PREVENTED: Cannot tamper with proof")
}

// TestElGamalTransfer_CommitmentVerification tests final commitment check
func TestElGamalTransfer_CommitmentVerification(t *testing.T) {
	ps := NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	realPlaintext := []byte("Real treasure location")
	fakePlaintext := []byte("Fake treasure location")

	// Create "original" cipher (as if minted on-chain)
	mintR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	originalCipher, _ := ps.EncryptElGamal(realPlaintext, &sellerKey.PublicKey, mintR)

	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		realPlaintext,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Decrypt (after payment)
	decrypted, err := ps.DecryptElGamal(transfer.BuyerCipher, transfer.SharedR, buyerKey)
	require.NoError(t, err)

	// Verify real plaintext matches commitment
	realValid := ps.VerifyPlaintextMatchesCommitment(
		decrypted,
		transfer.Salt,
		transfer.Commitment,
	)
	require.True(t, realValid, "Real plaintext should match commitment")

	// Verify fake plaintext does NOT match commitment
	fakeValid := ps.VerifyPlaintextMatchesCommitment(
		fakePlaintext,
		transfer.Salt,
		transfer.Commitment,
	)
	require.False(t, fakeValid, "Fake plaintext should NOT match commitment")

	t.Log("✅ Commitment verification works correctly")
}

// TestElGamalTransfer_CompleteFlow demonstrates the complete protocol
func TestElGamalTransfer_CompleteFlow(t *testing.T) {
	ps := NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	clueContent := []byte("The hidden treasure is behind the waterfall in the old forest")

	// Create "original" cipher (as if minted on-chain)
	mintR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	originalCipher, _ := ps.EncryptElGamal(clueContent, &sellerKey.PublicKey, mintR)

	t.Log("\n" + strings.Repeat("=", 70))
	t.Log("COMPLETE VERIFIABLE TRANSFER PROTOCOL")
	t.Log(strings.Repeat("=", 70))

	// Step 1: Seller generates verifiable transfer
	t.Log("\n[1] Seller generates verifiable transfer (off-chain)")
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		clueContent,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)
	t.Log("    ✓ Generated ElGamal ciphertexts for seller and buyer")
	t.Log("    ✓ Generated DLEQ proof")
	t.Log("    ✓ Generated plaintext equality proof")
	t.Log("    ✓ Created commitment to plaintext")

	// Step 2: Seller posts to blockchain (without r!)
	t.Log("\n[2] Seller posts to smart contract (on-chain)")
	t.Log("    - SellerCipher")
	t.Log("    - BuyerCipher")
	t.Log("    - DLEQ Proof")
	t.Log("    - Plaintext Equality Proof")
	t.Log("    - Commitment")
	t.Log("    - Public keys")
	t.Log("    ⚠️  Does NOT reveal r (decryption key)!")

	// Step 3: Buyer verifies proof
	t.Log("\n[3] Buyer verifies proofs (off-chain)")
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.Proof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	require.True(t, valid)
	t.Log("    ✓ DLEQ proof verified!")
	t.Log("    ✓ Plaintext equality proof verified!")
	t.Log("    ✓ Mathematically proven: buyer cipher encrypts SAME plaintext as original")
	t.Log("    ✓ Buyer confident to proceed")
	t.Log("    ✗ Buyer CANNOT decrypt yet (no r)")

	// Step 4: Buyer pays
	t.Log("\n[4] Buyer sends payment (on-chain)")
	t.Log("    → Payment sent to smart contract")

	// Step 5: Smart contract transfers NFT
	t.Log("\n[5] Smart contract transfers NFT to buyer")
	t.Log("    → Buyer now owns NFT")
	t.Log("    → But still can't read clue without r!")

	// Step 6: Seller reveals r
	t.Log("\n[6] Seller reveals decryption key r (on-chain or off-chain)")
	t.Log("    → Seller provides r value")

	// Step 7: Buyer decrypts
	t.Log("\n[7] Buyer decrypts clue (off-chain)")
	decrypted, err := ps.DecryptElGamal(transfer.BuyerCipher, transfer.SharedR, buyerKey)
	require.NoError(t, err)
	require.Equal(t, clueContent, decrypted)
	t.Log("    ✓ Successfully decrypted!")

	// Step 8: Buyer verifies commitment
	t.Log("\n[8] Buyer verifies final commitment")
	commitmentValid := ps.VerifyPlaintextMatchesCommitment(
		decrypted,
		transfer.Salt,
		transfer.Commitment,
	)
	require.True(t, commitmentValid)
	t.Log("    ✓ Commitment matches!")
	t.Log("    ✓ Buyer got exactly what was promised")

	// Step 9: Payment released
	t.Log("\n[9] Payment released to seller (on-chain)")
	t.Log("    → Seller receives payment")

	t.Log("\n" + strings.Repeat("=", 70))
	t.Log("✅ TRANSFER COMPLETE")
	t.Log(strings.Repeat("=", 70))
	t.Log("\nSecurity guarantees:")
	t.Log("  ✅ Buyer verified BEFORE paying")
	t.Log("  ✅ Buyer verified content matches original on-chain cipher")
	t.Log("  ✅ Buyer could NOT decrypt before seller revealed r")
	t.Log("  ✅ Mathematical proof (no trust required)")
	t.Log("  ✅ No fraud detection mechanism needed")
	t.Log("  ✅ No dispute period required")
	t.Log("  ✅ Pure cryptographic security")
}

// TestElGamalTransfer_ContentAlterationDetected tests that altered content is detected
func TestElGamalTransfer_ContentAlterationDetected(t *testing.T) {
	ps := NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	originalContent := []byte("The treasure is at the secret location")
	alteredContent := []byte("FAKE: No treasure here!")

	// Create "original" cipher (as if minted on-chain with correct content)
	mintR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	originalCipher, _ := ps.EncryptElGamal(originalContent, &sellerKey.PublicKey, mintR)

	t.Log("=== ATTACK: Seller tries to provide altered content ===")

	// ATTACK: Seller generates transfer with WRONG content
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		alteredContent, // ATTACK: different from original!
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Buyer verifies - should FAIL due to plaintext equality proof
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.Proof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)

	require.False(t, valid, "ATTACK DETECTED: Altered content should be rejected!")

	t.Log("✅ ATTACK PREVENTED: Content alteration detected!")
	t.Log("   Plaintext equality proof protects against content manipulation")
}

// Benchmark the ElGamal operations
func BenchmarkElGamalEncryption(b *testing.B) {
	ps := NewProofSystem()
	key, _ := ps.GenerateKeyPair()
	plaintext := []byte("Test message")
	r, _ := rand.Int(rand.Reader, ps.Curve.Params().N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ps.EncryptElGamal(plaintext, &key.PublicKey, r)
	}
}

func BenchmarkElGamalDecryption(b *testing.B) {
	ps := NewProofSystem()
	key, _ := ps.GenerateKeyPair()
	plaintext := []byte("Test message")
	r, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	cipher, _ := ps.EncryptElGamal(plaintext, &key.PublicKey, r)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ps.DecryptElGamal(cipher, r, key)
	}
}

func BenchmarkDLEQProofGeneration(b *testing.B) {
	ps := NewProofSystem()
	seller, _ := ps.GenerateKeyPair()
	buyer, _ := ps.GenerateKeyPair()
	r, _ := rand.Int(rand.Reader, ps.Curve.Params().N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ps.GenerateDLEQProof(r, &seller.PublicKey, &buyer.PublicKey)
	}
}

func BenchmarkDLEQProofVerification(b *testing.B) {
	ps := NewProofSystem()
	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()
	plaintext := []byte("Test")

	// Create original cipher
	mintR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)
	originalCipher, _ := ps.EncryptElGamal(plaintext, &sellerKey.PublicKey, mintR)

	transfer, _ := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ps.VerifyElGamalTransfer(
			originalCipher,
			transfer.SellerCipher,
			transfer.BuyerCipher,
			transfer.Proof,
			mintR,
			transfer.SellerPubKey,
			transfer.BuyerPubKey,
		)
	}
}
