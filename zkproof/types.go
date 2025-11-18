package zkproof

import (
	"crypto/elliptic"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// ProofResult contains the proof and related ciphertexts
type ProofResult struct {
	Proof            *Proof
	SellerCipherText []byte
	BuyerCipherText  []byte
	BuyerCipherHash  []byte
}

// PartialProof represents a partial proof (before seller signs)
type PartialProof struct {
	C               *big.Int // Challenge value
	K               *big.Int // Random scalar
	R1              []byte   // First commitment (g^k)
	R2              []byte   // Second commitment (buyer_pub^k)
	CurveN          *big.Int // Curve order
	BuyerPubKey     []byte   // Buyer's public key
	SellerPubKey    []byte   // Seller's public key
	BuyerCipherHash []byte   // Hash of buyer's ciphertext
}

// PartialProofWithBuyerCiphertext includes the buyer's ciphertext
type PartialProofWithBuyerCiphertext struct {
	PartialProof
	BuyerCipherText []byte
}

// Proof contains the zero-knowledge proof components
type Proof struct {
	C               *big.Int // Challenge value
	S               *big.Int // Response value
	R1              []byte   // First commitment (g^r)
	R2              []byte   // Second commitment (buyer_pub^r)
	BuyerPubKey     []byte   // Buyer's public key
	SellerPubKey    []byte   // Seller's public key
	BuyerCipherHash []byte   // Hash of buyer's ciphertext
}

// ProofSystem handles the ZKP generation and verification
type ProofSystem struct {
	Curve elliptic.Curve
}

// NewProofSystem creates a new proof system using secp256k1 curve
func NewProofSystem() *ProofSystem {
	return &ProofSystem{
		Curve: secp256k1.S256(),
	}
}

// NewProof creates a new proof with proper initialization
func NewProof() *Proof {
	return &Proof{
		C: new(big.Int),
		S: new(big.Int),
	}
}
