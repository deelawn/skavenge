package zkproof

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.org/x/crypto/sha3"
)

// GenerateProof creates a complete zero-knowledge proof for NFT transfer
// This is the full proof generation (not partial) that includes the seller's signature
func (ps *ProofSystem) GenerateProof(
	originalMsg []byte,
	sellerKey *ecdsa.PrivateKey,
	buyerPubKey *ecdsa.PublicKey,
	sellerCipherText []byte,
) (*ProofResult, error) {
	// Generate the buyer's ciphertext
	buyerCipherText, err := ps.EncryptMessage(originalMsg, buyerPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt message for buyer: %v", err)
	}

	// Generate hash of buyer's ciphertext
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(buyerCipherText)
	buyerCipherHash := hasher.Sum(nil)

	// Get curve parameters
	curveN := ps.Curve.Params().N
	if curveN == nil || curveN.Sign() <= 0 {
		return nil, fmt.Errorf("invalid curve order")
	}

	// Generate random scalar k for the proof
	k, err := rand.Int(rand.Reader, curveN)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random k: %v", err)
	}
	k.Mod(k, curveN)

	// Calculate R1 = g^k
	r1x, r1y := ps.Curve.ScalarBaseMult(k.Bytes())
	if !ps.Curve.IsOnCurve(r1x, r1y) {
		return nil, fmt.Errorf("R1 point not on curve")
	}
	r1Bytes := elliptic.Marshal(ps.Curve, r1x, r1y)

	// Calculate R2 = buyer_pub^k
	r2x, r2y := ps.Curve.ScalarMult(buyerPubKey.X, buyerPubKey.Y, k.Bytes())
	if !ps.Curve.IsOnCurve(r2x, r2y) {
		return nil, fmt.Errorf("R2 point not on curve")
	}
	r2Bytes := elliptic.Marshal(ps.Curve, r2x, r2y)

	// Generate challenge using seller ciphertext and buyer cipher hash
	h := sha3.NewLegacyKeccak256()
	h.Write(sellerCipherText)
	h.Write(buyerCipherHash)
	h.Write(elliptic.Marshal(ps.Curve, buyerPubKey.X, buyerPubKey.Y))
	h.Write(elliptic.Marshal(ps.Curve, sellerKey.PublicKey.X, sellerKey.PublicKey.Y))
	h.Write(r1Bytes)
	h.Write(r2Bytes)

	c := new(big.Int).SetBytes(h.Sum(nil))
	c.Mod(c, curveN)

	// Calculate s = k - c*x mod n
	s := new(big.Int).Mul(c, sellerKey.D)
	s.Sub(k, s)
	s.Mod(s, curveN)

	proof := &Proof{
		C:               c,
		S:               s,
		R1:              r1Bytes,
		R2:              r2Bytes,
		BuyerPubKey:     elliptic.Marshal(ps.Curve, buyerPubKey.X, buyerPubKey.Y),
		SellerPubKey:    elliptic.Marshal(ps.Curve, sellerKey.PublicKey.X, sellerKey.PublicKey.Y),
		BuyerCipherHash: buyerCipherHash,
	}

	return &ProofResult{
		Proof:            proof,
		SellerCipherText: sellerCipherText,
		BuyerCipherText:  buyerCipherText,
		BuyerCipherHash:  buyerCipherHash,
	}, nil
}

// FinalizeProof creates a complete proof from a partial proof and the seller's signature
func (ps *ProofSystem) FinalizeProof(partialProof *PartialProof, s *big.Int) *Proof {
	return &Proof{
		C:               partialProof.C,
		S:               s,
		R1:              partialProof.R1,
		R2:              partialProof.R2,
		BuyerPubKey:     partialProof.BuyerPubKey,
		SellerPubKey:    partialProof.SellerPubKey,
		BuyerCipherHash: partialProof.BuyerCipherHash,
	}
}

// GeneratePartialProof creates a partial proof (without seller's signature)
func (ps *ProofSystem) GeneratePartialProof(
	originalMsg []byte,
	sellerPubKey *ecdsa.PublicKey,
	buyerPubKey *ecdsa.PublicKey,
	sellerCipherText []byte,
) (*PartialProofWithBuyerCiphertext, error) {
	// Generate the buyer's ciphertext
	buyerCipherText, err := ps.EncryptMessage(originalMsg, buyerPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt message for buyer: %v", err)
	}

	// Generate hash of buyer's ciphertext
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(buyerCipherText)
	buyerCipherHash := hasher.Sum(nil)

	// Get curve parameters
	curveN := ps.Curve.Params().N
	if curveN == nil || curveN.Sign() <= 0 {
		return nil, fmt.Errorf("invalid curve order")
	}

	// Generate random scalar k for the proof
	k, err := rand.Int(rand.Reader, curveN)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random k: %v", err)
	}
	k.Mod(k, curveN)

	// Calculate R1 = g^k
	r1x, r1y := ps.Curve.ScalarBaseMult(k.Bytes())
	if !ps.Curve.IsOnCurve(r1x, r1y) {
		return nil, fmt.Errorf("R1 point not on curve")
	}
	r1Bytes := elliptic.Marshal(ps.Curve, r1x, r1y)

	// Calculate R2 = buyer_pub^k
	r2x, r2y := ps.Curve.ScalarMult(buyerPubKey.X, buyerPubKey.Y, k.Bytes())
	if !ps.Curve.IsOnCurve(r2x, r2y) {
		return nil, fmt.Errorf("R2 point not on curve")
	}
	r2Bytes := elliptic.Marshal(ps.Curve, r2x, r2y)

	// Generate challenge using seller ciphertext and buyer cipher hash
	h := sha3.NewLegacyKeccak256()
	h.Write(sellerCipherText)
	h.Write(buyerCipherHash)
	h.Write(elliptic.Marshal(ps.Curve, buyerPubKey.X, buyerPubKey.Y))
	h.Write(elliptic.Marshal(ps.Curve, sellerPubKey.X, sellerPubKey.Y))
	h.Write(r1Bytes)
	h.Write(r2Bytes)

	c := new(big.Int).SetBytes(h.Sum(nil))
	c.Mod(c, curveN)

	return &PartialProofWithBuyerCiphertext{
		PartialProof: PartialProof{
			C:               c,
			K:               k,
			R1:              r1Bytes,
			R2:              r2Bytes,
			CurveN:          curveN,
			BuyerPubKey:     elliptic.Marshal(ps.Curve, buyerPubKey.X, buyerPubKey.Y),
			SellerPubKey:    elliptic.Marshal(ps.Curve, sellerPubKey.X, sellerPubKey.Y),
			BuyerCipherHash: buyerCipherHash,
		},
		BuyerCipherText: buyerCipherText,
	}, nil
}

// VerifyProof verifies the zero-knowledge proof
func (ps *ProofSystem) VerifyProof(proof *Proof, sellerCipherText []byte) bool {
	// Get curve parameters
	curveN := ps.Curve.Params().N
	if curveN == nil || curveN.Sign() <= 0 {
		return false
	}

	// Unmarshal points
	r1x, r1y := elliptic.Unmarshal(ps.Curve, proof.R1)
	r2x, r2y := elliptic.Unmarshal(ps.Curve, proof.R2)
	buyerX, buyerY := elliptic.Unmarshal(ps.Curve, proof.BuyerPubKey)
	sellerX, sellerY := elliptic.Unmarshal(ps.Curve, proof.SellerPubKey)

	if r1x == nil || r1y == nil || r2x == nil || r2y == nil ||
		buyerX == nil || buyerY == nil || sellerX == nil || sellerY == nil {
		return false
	}

	// Verify points are on curve
	if !ps.Curve.IsOnCurve(r1x, r1y) || !ps.Curve.IsOnCurve(r2x, r2y) ||
		!ps.Curve.IsOnCurve(buyerX, buyerY) || !ps.Curve.IsOnCurve(sellerX, sellerY) {
		return false
	}

	// Regenerate challenge using seller ciphertext and buyer cipher hash
	h := sha3.NewLegacyKeccak256()
	h.Write(sellerCipherText)
	h.Write(proof.BuyerCipherHash)
	h.Write(proof.BuyerPubKey)
	h.Write(proof.SellerPubKey)
	h.Write(proof.R1)
	h.Write(proof.R2)

	c := new(big.Int).SetBytes(h.Sum(nil))
	c.Mod(c, curveN)

	if c.Cmp(proof.C) != 0 {
		return false
	}

	// Verify equation: g^s * Y^c = R1
	gsX, gsY := ps.Curve.ScalarBaseMult(proof.S.Bytes())
	ycX, ycY := ps.Curve.ScalarMult(sellerX, sellerY, c.Bytes())
	rightX, rightY := ps.Curve.Add(gsX, gsY, ycX, ycY)

	return rightX.Cmp(r1x) == 0 && rightY.Cmp(r1y) == 0
}

// ComputeS computes the S value for finalizing a partial proof
func ComputeS(c, k, sellerKeyD, curveN *big.Int) *big.Int {
	s := new(big.Int).Mul(c, sellerKeyD)
	s.Sub(k, s)
	s.Mod(s, curveN)
	return s
}
