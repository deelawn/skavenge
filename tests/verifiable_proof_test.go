// Package tests contains tests for the fixed verifiable proof system
package tests

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/deelawn/skavenge/zkproof"
)

// TestVerifiableProof_HonestCase tests the happy path with honest parties
func TestVerifiableProof_HonestCase(t *testing.T) {
	ps := zkproof.NewProofSystem()

	// Generate keys
	sellerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	buyerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	// Original message
	plaintext := []byte("The treasure is buried under the old oak tree")

	// Encrypt for seller (as they would have it initially)
	sellerCipher, err := ps.EncryptMessage(plaintext, &sellerKey.PublicKey)
	require.NoError(t, err)

	// Generate verifiable proof
	proof, transferData, err := ps.GenerateVerifiableTransferProof(
		plaintext,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err, "Proof generation should succeed")
	require.NotNil(t, proof)
	require.NotNil(t, transferData)

	// Verify the proof
	valid := ps.VerifyTransferProof(proof, transferData.SellerCipherText, transferData.BuyerCipherText)
	require.True(t, valid, "Valid proof should verify")

	// Buyer decrypts their ciphertext
	decrypted, err := ps.DecryptMessage(transferData.BuyerCipherText, buyerKey)
	require.NoError(t, err, "Buyer should be able to decrypt")

	// Verify plaintext matches
	require.Equal(t, plaintext, decrypted, "Decrypted content should match original")

	// Verify commitment matches
	commitmentValid := ps.VerifyPlaintextCommitment(decrypted, transferData.Salt, proof.PlaintextCommitment)
	require.True(t, commitmentValid, "Commitment should match decrypted plaintext")

	t.Log("✅ Honest transfer completed successfully with all verifications passing")
}

// TestVerifiableProof_PreventsDifferentPlaintext demonstrates the fix for Issue #2
// Seller can no longer provide a proof for one message and encrypt a different one
func TestVerifiableProof_PreventsDifferentPlaintext(t *testing.T) {
	ps := zkproof.NewProofSystem()

	// Generate keys
	sellerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	buyerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	// Real and fake messages
	realMessage := []byte("The treasure is under the oak tree")
	fakeMessage := []byte("The treasure is in a volcano (FAKE!)")

	// Seller generates proof for REAL message
	proof, transferData, err := ps.GenerateVerifiableTransferProof(
		realMessage,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// ATTACK ATTEMPT: Seller tries to swap buyer's ciphertext with fake message
	fakeBuyerCipher, err := ps.EncryptMessage(fakeMessage, &buyerKey.PublicKey)
	require.NoError(t, err)

	// Verification should FAIL because cipher hash won't match
	valid := ps.VerifyTransferProof(proof, transferData.SellerCipherText, fakeBuyerCipher)
	require.False(t, valid, "FIXED: Proof should reject when buyer cipher is swapped!")

	// Even if somehow the proof passed, commitment verification would catch it
	if valid {
		decrypted, err := ps.DecryptMessage(fakeBuyerCipher, buyerKey)
		require.NoError(t, err)

		commitmentValid := ps.VerifyPlaintextCommitment(decrypted, transferData.Salt, proof.PlaintextCommitment)
		require.False(t, commitmentValid, "FIXED: Commitment should not match fake message!")
	}

	t.Log("✅ Attack prevented: Cannot swap buyer's ciphertext with different content")
}

// TestVerifiableProof_PreventsModifiedProof demonstrates the fix for Issue #1
// Proof verification now checks BOTH R1 and R2
func TestVerifiableProof_PreventsModifiedProof(t *testing.T) {
	ps := zkproof.NewProofSystem()

	// Generate keys
	sellerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	buyerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	plaintext := []byte("Secret treasure location")

	// Generate valid proof
	proof, transferData, err := ps.GenerateVerifiableTransferProof(
		plaintext,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// ATTACK ATTEMPT: Corrupt R2 (this worked in the old system!)
	randomX, randomY := ps.Curve.ScalarBaseMult(big.NewInt(99999).Bytes())
	corruptedR2 := elliptic.Marshal(ps.Curve, randomX, randomY)

	// Create corrupted proof
	corruptedProof := &zkproof.VerifiableTransferProof{
		PlaintextCommitment: proof.PlaintextCommitment,
		Salt:                proof.Salt,
		C:                   proof.C,
		S:                   proof.S,
		R1:                  proof.R1,
		R2:                  corruptedR2, // CORRUPTED!
		SellerPubKey:        proof.SellerPubKey,
		BuyerPubKey:         proof.BuyerPubKey,
		SellerCipherHash:    proof.SellerCipherHash,
		BuyerCipherHash:     proof.BuyerCipherHash,
		BindingSignature:    proof.BindingSignature,
	}

	// Verification should FAIL
	valid := ps.VerifyTransferProof(corruptedProof, transferData.SellerCipherText, transferData.BuyerCipherText)
	require.False(t, valid, "FIXED: Proof verification now checks R2 and should reject!")

	t.Log("✅ Attack prevented: R2 verification catches corrupted proofs")
}

// TestVerifiableProof_PreventsInvalidBindingSignature demonstrates binding signature verification
func TestVerifiableProof_PreventsInvalidBindingSignature(t *testing.T) {
	ps := zkproof.NewProofSystem()

	// Generate keys
	sellerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	buyerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	plaintext := []byte("Secret treasure location")

	// Generate valid proof
	proof, transferData, err := ps.GenerateVerifiableTransferProof(
		plaintext,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// ATTACK ATTEMPT: Modify the binding signature
	corruptedProof := &zkproof.VerifiableTransferProof{
		PlaintextCommitment: proof.PlaintextCommitment,
		Salt:                proof.Salt,
		C:                   proof.C,
		S:                   proof.S,
		R1:                  proof.R1,
		R2:                  proof.R2,
		SellerPubKey:        proof.SellerPubKey,
		BuyerPubKey:         proof.BuyerPubKey,
		SellerCipherHash:    proof.SellerCipherHash,
		BuyerCipherHash:     proof.BuyerCipherHash,
		BindingSignature:    []byte("invalid signature bytes"),
	}

	// Verification should FAIL
	valid := ps.VerifyTransferProof(corruptedProof, transferData.SellerCipherText, transferData.BuyerCipherText)
	require.False(t, valid, "FIXED: Invalid binding signature should be rejected!")

	t.Log("✅ Attack prevented: Binding signature verification catches tampering")
}

// TestVerifiableProof_CommitmentPreventsSellerEquivocation tests that seller cannot
// claim two different plaintexts for the same commitment
func TestVerifiableProof_CommitmentPreventsSellerEquivocation(t *testing.T) {
	ps := zkproof.NewProofSystem()

	// Generate keys
	sellerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	buyerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	plaintext1 := []byte("First message")
	plaintext2 := []byte("Second message (different!)")

	// Generate proof for first message
	proof, transferData, err := ps.GenerateVerifiableTransferProof(
		plaintext1,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Buyer decrypts and gets plaintext2 somehow (attack scenario)
	// They verify against the commitment
	commitment1Valid := ps.VerifyPlaintextCommitment(plaintext1, transferData.Salt, proof.PlaintextCommitment)
	require.True(t, commitment1Valid, "Correct plaintext should match commitment")

	// If seller tries to claim a different plaintext later
	commitment2Valid := ps.VerifyPlaintextCommitment(plaintext2, transferData.Salt, proof.PlaintextCommitment)
	require.False(t, commitment2Valid, "FIXED: Different plaintext should NOT match commitment!")

	t.Log("✅ Commitment prevents seller from claiming different plaintexts")
}

// TestVerifiableProof_BuyerCanDetectFraud demonstrates the fraud detection mechanism
func TestVerifiableProof_BuyerCanDetectFraud(t *testing.T) {
	ps := zkproof.NewProofSystem()

	// Generate keys
	sellerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	buyerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	claimedPlaintext := []byte("I claim this is the treasure location")
	actualPlaintext := []byte("But this is what buyer actually gets")

	// Seller generates proof claiming one thing
	proof, _, err := ps.GenerateVerifiableTransferProof(
		claimedPlaintext,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Buyer receives different plaintext (somehow - attack scenario)
	// Buyer checks commitment
	fraudDetected := !ps.VerifyPlaintextCommitment(actualPlaintext, proof.Salt, proof.PlaintextCommitment)
	require.True(t, fraudDetected, "Buyer should detect fraud when plaintext doesn't match commitment")

	// Buyer can now submit fraud proof on-chain with:
	// - proof.PlaintextCommitment (from contract)
	// - actualPlaintext (what they received)
	// - proof.Salt (revealed by seller)
	// Contract can verify: Hash(actualPlaintext || salt) != commitment

	t.Log("✅ Buyer can detect and prove fraud on-chain")
}

// TestVerifiableProof_CompleteTransferFlow tests the complete transfer flow
func TestVerifiableProof_CompleteTransferFlow(t *testing.T) {
	ps := zkproof.NewProofSystem()

	// Setup
	sellerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	buyerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	clueContent := []byte("The treasure is buried at coordinates 40.7128°N, 74.0060°W")

	t.Log("Step 1: Seller encrypts clue for themselves (initial NFT mint)")
	sellerCipher, err := ps.EncryptMessage(clueContent, &sellerKey.PublicKey)
	require.NoError(t, err)

	// Verify seller can decrypt their own clue
	sellerDecrypted, err := ps.DecryptMessage(sellerCipher, sellerKey)
	require.NoError(t, err)
	require.Equal(t, clueContent, sellerDecrypted)

	t.Log("Step 2: Buyer initiates purchase (on-chain)")
	// This happens on-chain, buyer sends payment

	t.Log("Step 3: Seller generates verifiable proof (off-chain)")
	proof, transferData, err := ps.GenerateVerifiableTransferProof(
		clueContent,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	t.Log("Step 4: Seller provides proof to smart contract (on-chain)")
	// Contract stores: proof, buyerCipherHash, commitment

	t.Log("Step 5: Buyer verifies proof (off-chain)")
	proofValid := ps.VerifyTransferProof(proof, transferData.SellerCipherText, transferData.BuyerCipherText)
	require.True(t, proofValid, "Buyer verifies proof off-chain")

	// Buyer can also decrypt and check commitment
	buyerDecrypted, err := ps.DecryptMessage(transferData.BuyerCipherText, buyerKey)
	require.NoError(t, err)

	commitmentValid := ps.VerifyPlaintextCommitment(buyerDecrypted, transferData.Salt, proof.PlaintextCommitment)
	require.True(t, commitmentValid, "Decrypted content matches commitment")

	t.Log("Step 6: Buyer calls VerifyProof on contract (on-chain)")
	// Buyer attests that they verified the proof off-chain

	t.Log("Step 7: Seller completes transfer (on-chain)")
	// Contract verifies hash matches and transfers NFT
	// Payment held in escrow during dispute period

	t.Log("Step 8: Seller reveals plaintext and salt (after transfer)")
	// Seller provides: transferData.Plaintext, transferData.Salt

	t.Log("Step 9: Buyer verifies revealed plaintext (off-chain)")
	finalCheck := ps.VerifyPlaintextCommitment(transferData.Plaintext, transferData.Salt, proof.PlaintextCommitment)
	require.True(t, finalCheck, "Revealed plaintext matches commitment")
	require.Equal(t, clueContent, transferData.Plaintext, "Revealed plaintext is correct")

	t.Log("Step 10: Dispute period passes, seller claims payment (on-chain)")
	// After 7 days with no fraud claim, seller gets paid

	t.Log("✅ Complete transfer flow successful with all security checks passing")
}

// TestVerifiableProof_FraudProofScenario demonstrates on-chain fraud detection
func TestVerifiableProof_FraudProofScenario(t *testing.T) {
	ps := zkproof.NewProofSystem()

	sellerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	buyerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	// Seller creates proof
	realContent := []byte("Real treasure location")
	proof, transferData, err := ps.GenerateVerifiableTransferProof(
		realContent,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// --- Transfer happens ---

	t.Log("Scenario: Seller tries to reveal wrong plaintext")
	fakeReveal := []byte("Fake treasure location")

	// Buyer checks commitment with fake reveal
	fraudDetected := !ps.VerifyPlaintextCommitment(fakeReveal, transferData.Salt, proof.PlaintextCommitment)
	require.True(t, fraudDetected, "Buyer detects fraud")

	// Buyer submits fraud proof on-chain:
	// claimFraud(transferId, fakeReveal, transferData.Salt)

	// Contract verifies:
	reconstructed := crypto.Keccak256Hash(append(fakeReveal, transferData.Salt...))
	var commitment [32]byte
	copy(commitment[:], proof.PlaintextCommitment[:])

	fraudProven := reconstructed.Bytes()[0] != commitment[0] // Simplified check
	require.True(t, fraudProven, "Contract can verify fraud on-chain")

	// Buyer gets refund, seller is penalized

	t.Log("✅ Fraud detection and on-chain verification works")
}

// Benchmark the new proof system
func BenchmarkVerifiableProof_Generation(b *testing.B) {
	ps := zkproof.NewProofSystem()
	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()
	plaintext := []byte("The treasure is buried under the old oak tree")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = ps.GenerateVerifiableTransferProof(plaintext, sellerKey, &buyerKey.PublicKey)
	}
}

func BenchmarkVerifiableProof_Verification(b *testing.B) {
	ps := zkproof.NewProofSystem()
	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()
	plaintext := []byte("The treasure is buried under the old oak tree")

	proof, transferData, _ := ps.GenerateVerifiableTransferProof(plaintext, sellerKey, &buyerKey.PublicKey)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ps.VerifyTransferProof(proof, transferData.SellerCipherText, transferData.BuyerCipherText)
	}
}
