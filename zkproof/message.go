package zkproof

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// EncryptMessage encrypts a message using ECIES-like encryption
func (ps *ProofSystem) EncryptMessage(message []byte, pubKey *ecdsa.PublicKey) ([]byte, error) {
	// Generate ephemeral key pair
	ephemeral, err := ecdsa.GenerateKey(ps.Curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ephemeral key: %v", err)
	}

	// Perform ECDH to get shared secret
	sx, _ := ps.Curve.ScalarMult(pubKey.X, pubKey.Y, ephemeral.D.Bytes())
	if sx == nil {
		return nil, fmt.Errorf("ECDH key generation failed")
	}

	// Generate symmetric key using shared secret
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(sx.Bytes())
	sharedSecret := hasher.Sum(nil)

	// Create AES cipher
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %v", err)
	}

	// Create GCM mode
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM mode: %v", err)
	}

	// Generate nonce
	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %v", err)
	}

	// Marshal ephemeral public key
	ephPubBytes := elliptic.Marshal(ps.Curve, ephemeral.PublicKey.X, ephemeral.PublicKey.Y)

	// Encrypt the content
	ciphertext := aead.Seal(nil, nonce, message, nil)

	// Format final output: ephPubKey || nonce || ciphertext
	output := make([]byte, 0, len(ephPubBytes)+len(nonce)+len(ciphertext))
	output = append(output, ephPubBytes...)
	output = append(output, nonce...)
	output = append(output, ciphertext...)

	return output, nil
}

// DecryptMessage decrypts a message using the recipient's private key
func (ps *ProofSystem) DecryptMessage(ciphertext []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	if len(ciphertext) < 65 {
		return nil, fmt.Errorf("ciphertext too short")
	}

	// Extract ephemeral public key (first 65 bytes)
	ephX, ephY := elliptic.Unmarshal(ps.Curve, ciphertext[:65])
	if ephX == nil {
		return nil, fmt.Errorf("failed to unmarshal ephemeral public key")
	}

	// Perform ECDH
	sx, _ := ps.Curve.ScalarMult(ephX, ephY, privKey.D.Bytes())
	if sx == nil {
		return nil, fmt.Errorf("ECDH key generation failed")
	}

	// Generate symmetric key
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(sx.Bytes())
	sharedSecret := hasher.Sum(nil)

	// Create AES cipher
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %v", err)
	}

	// Create GCM mode
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM mode: %v", err)
	}

	nonceSize := aead.NonceSize()
	if len(ciphertext) < 65+nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	// Extract nonce and encrypted data
	nonce := ciphertext[65 : 65+nonceSize]
	encrypted := ciphertext[65+nonceSize:]

	// Decrypt the content
	plaintext, err := aead.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %v", err)
	}

	return plaintext, nil
}

// VerifyBuyerCipherTextHash verifies the hash of the buyer's ciphertext matches the one in the proof
func (ps *ProofSystem) VerifyBuyerCipherTextHash(proof *Proof, buyerCipherText []byte) bool {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(buyerCipherText)
	computedHash := hasher.Sum(nil)
	return bytes.Equal(computedHash, proof.BuyerCipherHash)
}
