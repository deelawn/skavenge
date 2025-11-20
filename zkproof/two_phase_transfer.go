package zkproof

import (
	"crypto/ecdsa"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// TwoPhaseTransferCommitment contains commitments (hashes only) for phase 1
// This is what the seller provides initially - buyer CAN'T decrypt anything yet
type TwoPhaseTransferCommitment struct {
	ProofHash           [32]byte // Hash of the complete proof
	BuyerCipherHash     [32]byte // Hash of buyer's ciphertext
	PlaintextCommitment [32]byte // Hash(plaintext || salt)
}

// TwoPhaseReveal contains the actual data revealed in phase 3
// Seller only provides this AFTER buyer has locked payment
type TwoPhaseReveal struct {
	Proof           *VerifiableTransferProof
	BuyerCipherText []byte
}

// GenerateTwoPhaseTransfer creates a two-phase transfer with commitments
// Returns both the commitment (to share immediately) and reveal data (to share after payment locks)
func (ps *ProofSystem) GenerateTwoPhaseTransfer(
	plaintext []byte,
	sellerKey *ecdsa.PrivateKey,
	buyerPubKey *ecdsa.PublicKey,
) (*TwoPhaseTransferCommitment, *TwoPhaseReveal, *TransferData, error) {

	// Generate the complete verifiable proof
	proof, transferData, err := ps.GenerateVerifiableTransferProof(
		plaintext,
		sellerKey,
		buyerPubKey,
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to generate proof: %v", err)
	}

	// Phase 1: Create commitments (hashes only)
	// Seller provides these immediately - buyer CAN'T decrypt

	// Hash the marshaled proof
	proofBytes := proof.Marshal()
	proofHash := computeHash(proofBytes)

	// Hash the buyer's ciphertext
	buyerCipherHash := computeHash(transferData.BuyerCipherText)

	commitment := &TwoPhaseTransferCommitment{
		ProofHash:           proofHash,
		BuyerCipherHash:     buyerCipherHash,
		PlaintextCommitment: proof.PlaintextCommitment,
	}

	// Phase 3: Prepare reveal data
	// Seller provides this ONLY AFTER buyer locks payment
	reveal := &TwoPhaseReveal{
		Proof:           proof,
		BuyerCipherText: transferData.BuyerCipherText,
	}

	return commitment, reveal, transferData, nil
}

// VerifyRevealMatchesCommitment verifies that revealed data matches the commitments
// This is called by the smart contract when seller reveals
func (ps *ProofSystem) VerifyRevealMatchesCommitment(
	commitment *TwoPhaseTransferCommitment,
	reveal *TwoPhaseReveal,
) bool {

	// Verify proof hash matches
	proofBytes := reveal.Proof.Marshal()
	proofHash := computeHash(proofBytes)
	if proofHash != commitment.ProofHash {
		return false
	}

	// Verify buyer cipher hash matches
	buyerCipherHash := computeHash(reveal.BuyerCipherText)
	if buyerCipherHash != commitment.BuyerCipherHash {
		return false
	}

	// Verify plaintext commitment matches
	if reveal.Proof.PlaintextCommitment != commitment.PlaintextCommitment {
		return false
	}

	return true
}

// Marshal converts commitment to bytes for on-chain storage
func (c *TwoPhaseTransferCommitment) Marshal() []byte {
	buf := make([]byte, 0, 32*3)
	buf = append(buf, c.ProofHash[:]...)
	buf = append(buf, c.BuyerCipherHash[:]...)
	buf = append(buf, c.PlaintextCommitment[:]...)
	return buf
}

// Unmarshal recreates commitment from bytes
func (c *TwoPhaseTransferCommitment) Unmarshal(data []byte) error {
	if len(data) != 96 {
		return fmt.Errorf("invalid commitment data size: got %d, want 96", len(data))
	}

	copy(c.ProofHash[:], data[0:32])
	copy(c.BuyerCipherHash[:], data[32:64])
	copy(c.PlaintextCommitment[:], data[64:96])

	return nil
}

// Marshal converts proof to bytes (helper for VerifiableTransferProof)
func (p *VerifiableTransferProof) Marshal() []byte {
	// Calculate total size
	totalSize := 32 + // PlaintextCommitment
		len(p.Salt) + 2 + // Salt + length prefix
		32 + 32 + // C, S
		len(p.R1) + len(p.R2) + // R1, R2
		len(p.SellerPubKey) + len(p.BuyerPubKey) + // Public keys
		32 + 32 + // Cipher hashes
		len(p.BindingSignature) + 2 // Signature + length prefix

	buf := make([]byte, 0, totalSize)

	// Add commitment
	buf = append(buf, p.PlaintextCommitment[:]...)

	// Add salt with length prefix
	saltLen := uint16(len(p.Salt))
	buf = append(buf, byte(saltLen>>8), byte(saltLen))
	buf = append(buf, p.Salt...)

	// Add C and S (32 bytes each, padded)
	cBytes := make([]byte, 32)
	sBytes := make([]byte, 32)
	copy(cBytes[32-len(p.C.Bytes()):], p.C.Bytes())
	copy(sBytes[32-len(p.S.Bytes()):], p.S.Bytes())
	buf = append(buf, cBytes...)
	buf = append(buf, sBytes...)

	// Add curve points (fixed size)
	buf = append(buf, p.R1...)
	buf = append(buf, p.R2...)
	buf = append(buf, p.SellerPubKey...)
	buf = append(buf, p.BuyerPubKey...)

	// Add cipher hashes
	buf = append(buf, p.SellerCipherHash[:]...)
	buf = append(buf, p.BuyerCipherHash[:]...)

	// Add binding signature with length prefix
	sigLen := uint16(len(p.BindingSignature))
	buf = append(buf, byte(sigLen>>8), byte(sigLen))
	buf = append(buf, p.BindingSignature...)

	return buf
}

// Example usage demonstrating the two-phase protocol
func ExampleTwoPhaseProtocol() {
	ps := NewProofSystem()

	// Setup keys
	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	clueContent := []byte("The treasure is at coordinates 40.7128°N, 74.0060°W")

	// ========== PHASE 1: SELLER COMMITS ==========
	// Seller generates commitments and reveal data
	commitment, reveal, transferData, _ := ps.GenerateTwoPhaseTransfer(
		clueContent,
		sellerKey,
		&buyerKey.PublicKey,
	)

	// Seller provides commitment to smart contract
	// At this point, buyer sees ONLY hashes - cannot decrypt!
	fmt.Println("Phase 1: Commitments provided")
	fmt.Printf("  ProofHash: %x...\n", commitment.ProofHash[:8])
	fmt.Printf("  BuyerCipherHash: %x...\n", commitment.BuyerCipherHash[:8])
	fmt.Printf("  PlaintextCommitment: %x...\n", commitment.PlaintextCommitment[:8])

	// Buyer reviews commitments (can't see actual data yet)
	// Buyer sees structure is correct
	// Buyer CANNOT decrypt anything at this point!

	// ========== PHASE 2: BUYER LOCKS PAYMENT ==========
	// Buyer calls LockPayment() on smart contract
	// ⚠️ CRITICAL: Payment now NON-REFUNDABLE (except fraud proof)
	fmt.Println("\nPhase 2: Buyer locks payment")
	fmt.Println("  Payment is now committed!")
	fmt.Println("  Buyer CANNOT cancel anymore")

	// ========== PHASE 3: SELLER REVEALS ==========
	// Seller must reveal within timeout (e.g., 1 hour)
	// Seller provides actual proof and buyerCipherText

	// Smart contract verifies reveals match commitments
	valid := ps.VerifyRevealMatchesCommitment(commitment, reveal)
	if !valid {
		fmt.Println("ERROR: Reveal doesn't match commitment!")
		return
	}

	fmt.Println("\nPhase 3: Seller reveals data")
	fmt.Println("  Reveal matches commitment ✓")

	// NOW buyer can see and decrypt the ciphertext
	// But buyer has already locked payment - cannot cancel!
	decrypted, _ := ps.DecryptMessage(reveal.BuyerCipherText, buyerKey)

	fmt.Println("\nPhase 4: Buyer verifies")
	fmt.Printf("  Decrypted: %s\n", string(decrypted))

	// Buyer verifies commitment
	commitmentValid := ps.VerifyPlaintextCommitment(
		decrypted,
		transferData.Salt,
		commitment.PlaintextCommitment,
	)

	if commitmentValid {
		fmt.Println("  Commitment matches ✓")
		fmt.Println("  Transfer will complete after dispute period")
	} else {
		fmt.Println("  FRAUD DETECTED! Submitting fraud proof...")
		// Buyer would call claimFraud() here
	}
}
