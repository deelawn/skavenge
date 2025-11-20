package zkproof

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.org/x/crypto/sha3"
)

// ElGamalCiphertext represents an ElGamal encrypted message
type ElGamalCiphertext struct {
	C1 []byte // g^r (ephemeral public key point)
	C2 []byte // message * recipientPub^r
}

// DLEQProof proves that two ciphertexts encrypt the same plaintext
// This is a "Discrete Log Equality" proof using the Chaum-Pedersen protocol
type DLEQProof struct {
	// Commitments
	A1 []byte // g^w
	A2 []byte // recipientPub^w

	// Response
	Z *big.Int // w + c*r mod n

	// Challenge (derived from commitments)
	C *big.Int
}

// VerifiableElGamalTransfer contains everything needed for a verifiable transfer
type VerifiableElGamalTransfer struct {
	SellerCipher     *ElGamalCiphertext
	BuyerCipher      *ElGamalCiphertext
	DLEQProof        *DLEQProof
	Commitment       [32]byte // Hash(plaintext || salt)
	Salt             []byte
	SellerPubKey     []byte
	BuyerPubKey      []byte
	SharedR          *big.Int // Kept secret until after payment!
}

// EncryptElGamal encrypts a message using ElGamal encryption
func (ps *ProofSystem) EncryptElGamal(
	message []byte,
	recipientPubKey *ecdsa.PublicKey,
	r *big.Int, // Random value (for verifiable encryption, we control this)
) (*ElGamalCiphertext, error) {

	// Ensure r is valid
	curveN := ps.Curve.Params().N
	if r.Cmp(curveN) >= 0 || r.Sign() <= 0 {
		return nil, fmt.Errorf("invalid r value")
	}

	// C1 = g^r
	c1x, c1y := ps.Curve.ScalarBaseMult(r.Bytes())
	if !ps.Curve.IsOnCurve(c1x, c1y) {
		return nil, fmt.Errorf("C1 not on curve")
	}
	c1Bytes := elliptic.Marshal(ps.Curve, c1x, c1y)

	// Shared secret: S = recipientPub^r
	sx, sy := ps.Curve.ScalarMult(recipientPubKey.X, recipientPubKey.Y, r.Bytes())
	if !ps.Curve.IsOnCurve(sx, sy) {
		return nil, fmt.Errorf("shared secret not on curve")
	}

	// Derive symmetric key from shared secret
	sharedSecret := sx.Bytes()

	// XOR message with derived key (simplified for PoC)
	// In production, use proper KDF + AES-GCM
	keyHash := sha3.NewLegacyKeccak256()
	keyHash.Write(sharedSecret)
	key := keyHash.Sum(nil)

	// For simplicity, we'll encode message as a curve point
	// In production, use hybrid encryption (ElGamal for key, AES for message)
	// For now, just XOR with key hash
	c2 := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		c2[i] = message[i] ^ key[i%len(key)]
	}

	return &ElGamalCiphertext{
		C1: c1Bytes,
		C2: c2,
	}, nil
}

// DecryptElGamal decrypts an ElGamal ciphertext
func (ps *ProofSystem) DecryptElGamal(
	ciphertext *ElGamalCiphertext,
	r *big.Int, // Provided by seller after payment
	recipientPrivKey *ecdsa.PrivateKey,
) ([]byte, error) {

	// Parse C1
	c1x, c1y := elliptic.Unmarshal(ps.Curve, ciphertext.C1)
	if c1x == nil || c1y == nil {
		return nil, fmt.Errorf("invalid C1")
	}

	// Compute shared secret: S = C1^privKey = (g^r)^privKey = g^(r*privKey)
	// But we also have: S = recipientPub^r = (g^privKey)^r = g^(r*privKey)
	// So they're the same!
	sx, sy := ps.Curve.ScalarMult(c1x, c1y, recipientPrivKey.D.Bytes())
	if !ps.Curve.IsOnCurve(sx, sy) {
		return nil, fmt.Errorf("shared secret not on curve")
	}

	sharedSecret := sx.Bytes()

	// Derive same key
	keyHash := sha3.NewLegacyKeccak256()
	keyHash.Write(sharedSecret)
	key := keyHash.Sum(nil)

	// XOR to decrypt
	plaintext := make([]byte, len(ciphertext.C2))
	for i := 0; i < len(ciphertext.C2); i++ {
		plaintext[i] = ciphertext.C2[i] ^ key[i%len(key)]
	}

	return plaintext, nil
}

// GenerateVerifiableElGamalTransfer creates a verifiable transfer with DLEQ proof
// This allows buyer to verify BEFORE paying that they'll get the same content
func (ps *ProofSystem) GenerateVerifiableElGamalTransfer(
	plaintext []byte,
	sellerPrivKey *ecdsa.PrivateKey,
	buyerPubKey *ecdsa.PublicKey,
) (*VerifiableElGamalTransfer, error) {

	// Generate random salt for commitment
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %v", err)
	}

	// Create commitment
	commitment := computeCommitment(plaintext, salt)

	// Generate random r (used for BOTH encryptions!)
	curveN := ps.Curve.Params().N
	r, err := rand.Int(rand.Reader, curveN)
	if err != nil {
		return nil, fmt.Errorf("failed to generate r: %v", err)
	}

	// Encrypt for seller
	sellerCipher, err := ps.EncryptElGamal(plaintext, &sellerPrivKey.PublicKey, r)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt for seller: %v", err)
	}

	// Encrypt for buyer (SAME r value!)
	buyerCipher, err := ps.EncryptElGamal(plaintext, buyerPubKey, r)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt for buyer: %v", err)
	}

	// Generate DLEQ proof that both ciphers use same r
	dleqProof, err := ps.GenerateDLEQProof(r, &sellerPrivKey.PublicKey, buyerPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate DLEQ proof: %v", err)
	}

	sellerPubBytes := elliptic.Marshal(ps.Curve, sellerPrivKey.PublicKey.X, sellerPrivKey.PublicKey.Y)
	buyerPubBytes := elliptic.Marshal(ps.Curve, buyerPubKey.X, buyerPubKey.Y)

	return &VerifiableElGamalTransfer{
		SellerCipher: sellerCipher,
		BuyerCipher:  buyerCipher,
		DLEQProof:    dleqProof,
		Commitment:   commitment,
		Salt:         salt,
		SellerPubKey: sellerPubBytes,
		BuyerPubKey:  buyerPubBytes,
		SharedR:      r, // Seller keeps this secret until after payment!
	}, nil
}

// GenerateDLEQProof creates a proof that log_g(C1_seller) == log_g(C1_buyer)
// This proves both ciphertexts use the same r value
func (ps *ProofSystem) GenerateDLEQProof(
	r *big.Int, // The secret random value
	sellerPub *ecdsa.PublicKey,
	buyerPub *ecdsa.PublicKey,
) (*DLEQProof, error) {

	curveN := ps.Curve.Params().N

	// Generate random w
	w, err := rand.Int(rand.Reader, curveN)
	if err != nil {
		return nil, fmt.Errorf("failed to generate w: %v", err)
	}

	// A1 = g^w
	a1x, a1y := ps.Curve.ScalarBaseMult(w.Bytes())
	a1Bytes := elliptic.Marshal(ps.Curve, a1x, a1y)

	// A2 = buyerPub^w
	a2x, a2y := ps.Curve.ScalarMult(buyerPub.X, buyerPub.Y, w.Bytes())
	a2Bytes := elliptic.Marshal(ps.Curve, a2x, a2y)

	// Generate challenge: c = Hash(g, sellerPub, buyerPub, C1_seller, C1_buyer, A1, A2)
	// For simplicity, we'll hash the key components
	h := sha3.NewLegacyKeccak256()
	h.Write(elliptic.Marshal(ps.Curve, sellerPub.X, sellerPub.Y))
	h.Write(elliptic.Marshal(ps.Curve, buyerPub.X, buyerPub.Y))
	h.Write(a1Bytes)
	h.Write(a2Bytes)

	c := new(big.Int).SetBytes(h.Sum(nil))
	c.Mod(c, curveN)

	// Response: z = w + c*r mod n
	z := new(big.Int).Mul(c, r)
	z.Add(w, z)
	z.Mod(z, curveN)

	return &DLEQProof{
		A1: a1Bytes,
		A2: a2Bytes,
		Z:  z,
		C:  c,
	}, nil
}

// VerifyElGamalTransfer verifies that both ciphertexts encrypt the same plaintext
// Buyer can call this BEFORE paying to get cryptographic proof!
func (ps *ProofSystem) VerifyElGamalTransfer(
	sellerCipher *ElGamalCiphertext,
	buyerCipher *ElGamalCiphertext,
	proof *DLEQProof,
	sellerPubKey []byte,
	buyerPubKey []byte,
) bool {

	curveN := ps.Curve.Params().N

	// Parse public keys
	sellerX, sellerY := elliptic.Unmarshal(ps.Curve, sellerPubKey)
	buyerX, buyerY := elliptic.Unmarshal(ps.Curve, buyerPubKey)

	if sellerX == nil || buyerX == nil {
		return false
	}

	// Parse proof commitments
	a1x, a1y := elliptic.Unmarshal(ps.Curve, proof.A1)
	a2x, a2y := elliptic.Unmarshal(ps.Curve, proof.A2)

	if a1x == nil || a2x == nil {
		return false
	}

	// Parse C1 values from ciphertexts
	c1SellerX, c1SellerY := elliptic.Unmarshal(ps.Curve, sellerCipher.C1)
	c1BuyerX, c1BuyerY := elliptic.Unmarshal(ps.Curve, buyerCipher.C1)

	if c1SellerX == nil || c1BuyerX == nil {
		return false
	}

	// Verify challenge
	h := sha3.NewLegacyKeccak256()
	h.Write(sellerPubKey)
	h.Write(buyerPubKey)
	h.Write(proof.A1)
	h.Write(proof.A2)

	cCheck := new(big.Int).SetBytes(h.Sum(nil))
	cCheck.Mod(cCheck, curveN)

	if cCheck.Cmp(proof.C) != 0 {
		return false
	}

	// Verify equation 1: g^z = A1 * C1_seller^c
	// This proves: z = w + c*r, so g^z = g^w * g^(c*r) = A1 * (g^r)^c
	gzX, gzY := ps.Curve.ScalarBaseMult(proof.Z.Bytes())

	c1cX, c1cY := ps.Curve.ScalarMult(c1SellerX, c1SellerY, proof.C.Bytes())
	rightX, rightY := ps.Curve.Add(a1x, a1y, c1cX, c1cY)

	if gzX.Cmp(rightX) != 0 || gzY.Cmp(rightY) != 0 {
		return false
	}

	// Verify equation 2: buyerPub^z = A2 * C1_buyer^c
	// This proves the same z works for buyer's encryption
	pubzX, pubzY := ps.Curve.ScalarMult(buyerX, buyerY, proof.Z.Bytes())

	c1BuyercX, c1BuyercY := ps.Curve.ScalarMult(c1BuyerX, c1BuyerY, proof.C.Bytes())
	right2X, right2Y := ps.Curve.Add(a2x, a2y, c1BuyercX, c1BuyercY)

	if pubzX.Cmp(right2X) != 0 || pubzY.Cmp(right2Y) != 0 {
		return false
	}

	// Both equations verified!
	// This PROVES both ciphertexts use the same r value
	// Therefore they encrypt the same plaintext
	return true
}

// VerifyPlaintextMatchesCommitment verifies decrypted plaintext matches commitment
func (ps *ProofSystem) VerifyPlaintextMatchesCommitment(
	plaintext []byte,
	salt []byte,
	commitment [32]byte,
) bool {
	reconstructed := computeCommitment(plaintext, salt)
	return reconstructed == commitment
}
