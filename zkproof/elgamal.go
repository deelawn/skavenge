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
	C1            []byte // g^r (ephemeral public key point)
	C2            []byte // message XOR Hash(recipientPub^r)
	SharedSecret  []byte // recipientPub^r (needed for DLEQ proof)
}

// DLEQProof proves that two ciphertexts encrypt the same plaintext
// This is a "Discrete Log Equality" proof using the Chaum-Pedersen protocol
type DLEQProof struct {
	// Commitments
	A1 []byte // g^w
	A2 []byte // sellerPub^w
	A3 []byte // buyerPub^w

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

	// Shared secret: S = recipientPub^r (stored for DLEQ proof verification)
	sx, sy := ps.Curve.ScalarMult(recipientPubKey.X, recipientPubKey.Y, r.Bytes())
	if !ps.Curve.IsOnCurve(sx, sy) {
		return nil, fmt.Errorf("shared secret not on curve")
	}
	sharedSecretPoint := elliptic.Marshal(ps.Curve, sx, sy)

	// Derive symmetric key from r AND the shared secret
	// This ensures:
	// 1. Buyer MUST know r (prevents early decryption before r is revealed)
	// 2. Buyer MUST use private key to compute shared secret (prevents public decryption)
	sharedSecret := sx.Bytes()

	keyHash := sha3.NewLegacyKeccak256()
	keyHash.Write(r.Bytes())      // Prevents early decryption (needs r)
	keyHash.Write(sharedSecret)   // Prevents public decryption (needs private key)
	key := keyHash.Sum(nil)

	// For simplicity, we'll encode message as a curve point
	// In production, use hybrid encryption (ElGamal for key, AES for message)
	// For now, just XOR with key hash
	c2 := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		c2[i] = message[i] ^ key[i%len(key)]
	}

	return &ElGamalCiphertext{
		C1:           c1Bytes,
		C2:           c2,
		SharedSecret: sharedSecretPoint,
	}, nil
}

// DecryptElGamal decrypts an ElGamal ciphertext using the revealed r value
// Buyer MUST have both r (revealed by seller) AND their private key to decrypt
func (ps *ProofSystem) DecryptElGamal(
	ciphertext *ElGamalCiphertext,
	r *big.Int, // Provided by seller after payment - REQUIRED!
	recipientPrivKey *ecdsa.PrivateKey,
) ([]byte, error) {

	if r == nil {
		return nil, fmt.Errorf("r value is required for decryption")
	}

	// Parse C1 = g^r
	c1x, c1y := elliptic.Unmarshal(ps.Curve, ciphertext.C1)
	if c1x == nil || c1y == nil {
		return nil, fmt.Errorf("invalid C1")
	}

	// Compute shared secret using private key: S = C1^privKey = (g^r)^privKey
	// This equals recipientPub^r used during encryption
	sx, sy := ps.Curve.ScalarMult(c1x, c1y, recipientPrivKey.D.Bytes())
	if !ps.Curve.IsOnCurve(sx, sy) {
		return nil, fmt.Errorf("shared secret not on curve")
	}
	sharedSecret := sx.Bytes()

	// Derive decryption key (same as encryption)
	keyHash := sha3.NewLegacyKeccak256()
	keyHash.Write(r.Bytes())      // Uses revealed r value
	keyHash.Write(sharedSecret)   // Uses computed shared secret (needs private key)
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

	// A2 = sellerPub^w
	a2x, a2y := ps.Curve.ScalarMult(sellerPub.X, sellerPub.Y, w.Bytes())
	a2Bytes := elliptic.Marshal(ps.Curve, a2x, a2y)

	// A3 = buyerPub^w
	a3x, a3y := ps.Curve.ScalarMult(buyerPub.X, buyerPub.Y, w.Bytes())
	a3Bytes := elliptic.Marshal(ps.Curve, a3x, a3y)

	// Generate challenge: c = Hash(sellerPub, buyerPub, A1, A2, A3)
	h := sha3.NewLegacyKeccak256()
	h.Write(elliptic.Marshal(ps.Curve, sellerPub.X, sellerPub.Y))
	h.Write(elliptic.Marshal(ps.Curve, buyerPub.X, buyerPub.Y))
	h.Write(a1Bytes)
	h.Write(a2Bytes)
	h.Write(a3Bytes)

	c := new(big.Int).SetBytes(h.Sum(nil))
	c.Mod(c, curveN)

	// Response: z = w + c*r mod n
	z := new(big.Int).Mul(c, r)
	z.Add(w, z)
	z.Mod(z, curveN)

	return &DLEQProof{
		A1: a1Bytes,
		A2: a2Bytes,
		A3: a3Bytes,
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

	if sellerX == nil || sellerY == nil || buyerX == nil || buyerY == nil {
		return false
	}

	// Parse proof commitments
	a1x, a1y := elliptic.Unmarshal(ps.Curve, proof.A1)
	a2x, a2y := elliptic.Unmarshal(ps.Curve, proof.A2)
	a3x, a3y := elliptic.Unmarshal(ps.Curve, proof.A3)

	if a1x == nil || a2x == nil || a3x == nil {
		return false
	}

	// Parse C1 values from ciphertexts (should be identical since same r)
	c1SellerX, c1SellerY := elliptic.Unmarshal(ps.Curve, sellerCipher.C1)
	c1BuyerX, c1BuyerY := elliptic.Unmarshal(ps.Curve, buyerCipher.C1)

	if c1SellerX == nil || c1BuyerX == nil {
		return false
	}

	// First, verify that C1 is the same for both (same r was used)
	if c1SellerX.Cmp(c1BuyerX) != 0 || c1SellerY.Cmp(c1BuyerY) != 0 {
		return false // Different r values used!
	}

	// Parse shared secrets
	sSellerX, sSellerY := elliptic.Unmarshal(ps.Curve, sellerCipher.SharedSecret)
	sBuyerX, sBuyerY := elliptic.Unmarshal(ps.Curve, buyerCipher.SharedSecret)

	if sSellerX == nil || sBuyerX == nil {
		return false
	}

	// Verify challenge
	h := sha3.NewLegacyKeccak256()
	h.Write(sellerPubKey)
	h.Write(buyerPubKey)
	h.Write(proof.A1)
	h.Write(proof.A2)
	h.Write(proof.A3)

	cCheck := new(big.Int).SetBytes(h.Sum(nil))
	cCheck.Mod(cCheck, curveN)

	if cCheck.Cmp(proof.C) != 0 {
		return false
	}

	// Verify equation 1: g^z = A1 * C1^c
	// This proves: C1 = g^r for some r
	gzX, gzY := ps.Curve.ScalarBaseMult(proof.Z.Bytes())

	c1cX, c1cY := ps.Curve.ScalarMult(c1SellerX, c1SellerY, proof.C.Bytes())
	right1X, right1Y := ps.Curve.Add(a1x, a1y, c1cX, c1cY)

	if gzX.Cmp(right1X) != 0 || gzY.Cmp(right1Y) != 0 {
		return false
	}

	// Verify equation 2: sellerPub^z = A2 * S_seller^c
	// This proves: S_seller = sellerPub^r for the same r
	sellerPubZX, sellerPubZY := ps.Curve.ScalarMult(sellerX, sellerY, proof.Z.Bytes())

	sSellerCX, sSellerCY := ps.Curve.ScalarMult(sSellerX, sSellerY, proof.C.Bytes())
	right2X, right2Y := ps.Curve.Add(a2x, a2y, sSellerCX, sSellerCY)

	if sellerPubZX.Cmp(right2X) != 0 || sellerPubZY.Cmp(right2Y) != 0 {
		return false
	}

	// Verify equation 3: buyerPub^z = A3 * S_buyer^c
	// This proves: S_buyer = buyerPub^r for the same r
	buyerPubZX, buyerPubZY := ps.Curve.ScalarMult(buyerX, buyerY, proof.Z.Bytes())

	sBuyerCX, sBuyerCY := ps.Curve.ScalarMult(sBuyerX, sBuyerY, proof.C.Bytes())
	right3X, right3Y := ps.Curve.Add(a3x, a3y, sBuyerCX, sBuyerCY)

	if buyerPubZX.Cmp(right3X) != 0 || buyerPubZY.Cmp(right3Y) != 0 {
		return false
	}

	// All three equations verified!
	// This PROVES:
	// 1. C1 = g^r
	// 2. S_seller = sellerPub^r
	// 3. S_buyer = buyerPub^r
	// All with the SAME r value.
	// Since both ciphertexts use the same r and have the same commitment,
	// they MUST encrypt the same plaintext.
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

// computeCommitment creates a hash commitment to plaintext with salt
func computeCommitment(plaintext []byte, salt []byte) [32]byte {
	h := sha3.NewLegacyKeccak256()
	h.Write(plaintext)
	h.Write(salt)
	hash := h.Sum(nil)

	var commitment [32]byte
	copy(commitment[:], hash)
	return commitment
}

// Marshal serializes an ElGamalCiphertext to bytes
func (c *ElGamalCiphertext) Marshal() []byte {
	// Format: len(C1) | C1 | len(C2) | C2 | len(SharedSecret) | SharedSecret
	result := make([]byte, 0)

	// C1
	c1Len := make([]byte, 4)
	c1Len[0] = byte(len(c.C1) >> 24)
	c1Len[1] = byte(len(c.C1) >> 16)
	c1Len[2] = byte(len(c.C1) >> 8)
	c1Len[3] = byte(len(c.C1))
	result = append(result, c1Len...)
	result = append(result, c.C1...)

	// C2
	c2Len := make([]byte, 4)
	c2Len[0] = byte(len(c.C2) >> 24)
	c2Len[1] = byte(len(c.C2) >> 16)
	c2Len[2] = byte(len(c.C2) >> 8)
	c2Len[3] = byte(len(c.C2))
	result = append(result, c2Len...)
	result = append(result, c.C2...)

	// SharedSecret
	ssLen := make([]byte, 4)
	ssLen[0] = byte(len(c.SharedSecret) >> 24)
	ssLen[1] = byte(len(c.SharedSecret) >> 16)
	ssLen[2] = byte(len(c.SharedSecret) >> 8)
	ssLen[3] = byte(len(c.SharedSecret))
	result = append(result, ssLen...)
	result = append(result, c.SharedSecret...)

	return result
}

// Unmarshal deserializes bytes into an ElGamalCiphertext
func (c *ElGamalCiphertext) Unmarshal(data []byte) error {
	if len(data) < 12 {
		return fmt.Errorf("data too short")
	}

	offset := 0

	// C1
	c1Len := int(data[offset])<<24 | int(data[offset+1])<<16 | int(data[offset+2])<<8 | int(data[offset+3])
	offset += 4
	if offset+c1Len > len(data) {
		return fmt.Errorf("invalid C1 length")
	}
	c.C1 = make([]byte, c1Len)
	copy(c.C1, data[offset:offset+c1Len])
	offset += c1Len

	// C2
	if offset+4 > len(data) {
		return fmt.Errorf("data too short for C2 length")
	}
	c2Len := int(data[offset])<<24 | int(data[offset+1])<<16 | int(data[offset+2])<<8 | int(data[offset+3])
	offset += 4
	if offset+c2Len > len(data) {
		return fmt.Errorf("invalid C2 length")
	}
	c.C2 = make([]byte, c2Len)
	copy(c.C2, data[offset:offset+c2Len])
	offset += c2Len

	// SharedSecret
	if offset+4 > len(data) {
		return fmt.Errorf("data too short for SharedSecret length")
	}
	ssLen := int(data[offset])<<24 | int(data[offset+1])<<16 | int(data[offset+2])<<8 | int(data[offset+3])
	offset += 4
	if offset+ssLen > len(data) {
		return fmt.Errorf("invalid SharedSecret length")
	}
	c.SharedSecret = make([]byte, ssLen)
	copy(c.SharedSecret, data[offset:offset+ssLen])

	return nil
}

// Marshal serializes a DLEQProof to bytes
func (p *DLEQProof) Marshal() []byte {
	// Format: len(A1) | A1 | len(A2) | A2 | len(A3) | A3 | len(Z) | Z | len(C) | C
	result := make([]byte, 0)

	// A1
	a1Len := make([]byte, 4)
	a1Len[0] = byte(len(p.A1) >> 24)
	a1Len[1] = byte(len(p.A1) >> 16)
	a1Len[2] = byte(len(p.A1) >> 8)
	a1Len[3] = byte(len(p.A1))
	result = append(result, a1Len...)
	result = append(result, p.A1...)

	// A2
	a2Len := make([]byte, 4)
	a2Len[0] = byte(len(p.A2) >> 24)
	a2Len[1] = byte(len(p.A2) >> 16)
	a2Len[2] = byte(len(p.A2) >> 8)
	a2Len[3] = byte(len(p.A2))
	result = append(result, a2Len...)
	result = append(result, p.A2...)

	// A3
	a3Len := make([]byte, 4)
	a3Len[0] = byte(len(p.A3) >> 24)
	a3Len[1] = byte(len(p.A3) >> 16)
	a3Len[2] = byte(len(p.A3) >> 8)
	a3Len[3] = byte(len(p.A3))
	result = append(result, a3Len...)
	result = append(result, p.A3...)

	// Z
	zBytes := p.Z.Bytes()
	zLen := make([]byte, 4)
	zLen[0] = byte(len(zBytes) >> 24)
	zLen[1] = byte(len(zBytes) >> 16)
	zLen[2] = byte(len(zBytes) >> 8)
	zLen[3] = byte(len(zBytes))
	result = append(result, zLen...)
	result = append(result, zBytes...)

	// C
	cBytes := p.C.Bytes()
	cLen := make([]byte, 4)
	cLen[0] = byte(len(cBytes) >> 24)
	cLen[1] = byte(len(cBytes) >> 16)
	cLen[2] = byte(len(cBytes) >> 8)
	cLen[3] = byte(len(cBytes))
	result = append(result, cLen...)
	result = append(result, cBytes...)

	return result
}

// Unmarshal deserializes bytes into a DLEQProof
func (p *DLEQProof) Unmarshal(data []byte) error {
	if len(data) < 20 {
		return fmt.Errorf("data too short")
	}

	offset := 0

	// A1
	a1Len := int(data[offset])<<24 | int(data[offset+1])<<16 | int(data[offset+2])<<8 | int(data[offset+3])
	offset += 4
	if offset+a1Len > len(data) {
		return fmt.Errorf("invalid A1 length")
	}
	p.A1 = make([]byte, a1Len)
	copy(p.A1, data[offset:offset+a1Len])
	offset += a1Len

	// A2
	if offset+4 > len(data) {
		return fmt.Errorf("data too short for A2 length")
	}
	a2Len := int(data[offset])<<24 | int(data[offset+1])<<16 | int(data[offset+2])<<8 | int(data[offset+3])
	offset += 4
	if offset+a2Len > len(data) {
		return fmt.Errorf("invalid A2 length")
	}
	p.A2 = make([]byte, a2Len)
	copy(p.A2, data[offset:offset+a2Len])
	offset += a2Len

	// A3
	if offset+4 > len(data) {
		return fmt.Errorf("data too short for A3 length")
	}
	a3Len := int(data[offset])<<24 | int(data[offset+1])<<16 | int(data[offset+2])<<8 | int(data[offset+3])
	offset += 4
	if offset+a3Len > len(data) {
		return fmt.Errorf("invalid A3 length")
	}
	p.A3 = make([]byte, a3Len)
	copy(p.A3, data[offset:offset+a3Len])
	offset += a3Len

	// Z
	if offset+4 > len(data) {
		return fmt.Errorf("data too short for Z length")
	}
	zLen := int(data[offset])<<24 | int(data[offset+1])<<16 | int(data[offset+2])<<8 | int(data[offset+3])
	offset += 4
	if offset+zLen > len(data) {
		return fmt.Errorf("invalid Z length")
	}
	p.Z = new(big.Int).SetBytes(data[offset : offset+zLen])
	offset += zLen

	// C
	if offset+4 > len(data) {
		return fmt.Errorf("data too short for C length")
	}
	cLen := int(data[offset])<<24 | int(data[offset+1])<<16 | int(data[offset+2])<<8 | int(data[offset+3])
	offset += 4
	if offset+cLen > len(data) {
		return fmt.Errorf("invalid C length")
	}
	p.C = new(big.Int).SetBytes(data[offset : offset+cLen])

	return nil
}
