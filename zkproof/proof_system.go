package zkproof

import (
	"crypto/elliptic"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

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
