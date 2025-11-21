package zkproof

import (
	"crypto/ecdsa"
	"crypto/rand"
)

// GenerateKeyPair creates a new ECDSA key pair
func (ps *ProofSystem) GenerateKeyPair() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(ps.Curve, rand.Reader)
}
