package zkproof

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/asn1"
	"fmt"
	"math/big"
)

// GenerateKeyPair creates a new ECDSA key pair
func (ps *ProofSystem) GenerateKeyPair() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(ps.Curve, rand.Reader)
}

// ECDSAPrivateKey represents an ASN.1 encoded secp256k1 private key
type ECDSAPrivateKey struct {
	Version       int
	PrivateKey    []byte
	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

// MarshalPrivateKey encodes a secp256k1 private key to ASN.1 DER format
func (ps *ProofSystem) MarshalPrivateKey(key *ecdsa.PrivateKey) ([]byte, error) {
	if key == nil {
		return nil, fmt.Errorf("nil private key")
	}

	// Verify it's a secp256k1 key
	if key.Curve != ps.Curve {
		return nil, fmt.Errorf("private key is not on secp256k1 curve")
	}

	// Prepare private key bytes in big-endian format
	privateKeyBytes := key.D.Bytes()
	paddedPrivateKey := make([]byte, 32)
	copy(paddedPrivateKey[32-len(privateKeyBytes):], privateKeyBytes)

	// Encode public key bytes
	publicKeyBytes := elliptic.Marshal(key.Curve, key.X, key.Y)
	publicKey := asn1.BitString{
		Bytes:     publicKeyBytes,
		BitLength: 8 * len(publicKeyBytes),
	}

	// Create ASN.1 structure
	privateKey := ECDSAPrivateKey{
		Version:    1,
		PrivateKey: paddedPrivateKey,
		// We use a custom OID for secp256k1 (Bitcoin curve)
		// This is the OID commonly used for secp256k1
		NamedCurveOID: asn1.ObjectIdentifier{1, 3, 132, 0, 10},
		PublicKey:     publicKey,
	}

	return asn1.Marshal(privateKey)
}

// UnmarshalPrivateKey decodes a private key from ASN.1 DER format
func (ps *ProofSystem) UnmarshalPrivateKey(derBytes []byte) (*ecdsa.PrivateKey, error) {
	var privateKey ECDSAPrivateKey
	_, err := asn1.Unmarshal(derBytes, &privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// Check OID for secp256k1
	if !privateKey.NamedCurveOID.Equal(asn1.ObjectIdentifier{1, 3, 132, 0, 10}) {
		return nil, fmt.Errorf("private key is not on secp256k1 curve")
	}

	// Create private key
	key := new(ecdsa.PrivateKey)
	key.Curve = ps.Curve
	key.D = new(big.Int).SetBytes(privateKey.PrivateKey)

	// Derive public key if not present in ASN.1 structure
	if len(privateKey.PublicKey.Bytes) == 0 {
		key.X, key.Y = key.Curve.ScalarBaseMult(key.D.Bytes())
	} else {
		key.X, key.Y = elliptic.Unmarshal(key.Curve, privateKey.PublicKey.Bytes)
	}

	if key.X == nil || key.Y == nil {
		return nil, fmt.Errorf("invalid public key")
	}

	return key, nil
}

// MarshalPublicKey encodes a secp256k1 public key to compressed format
func (ps *ProofSystem) MarshalPublicKey(key *ecdsa.PublicKey) ([]byte, error) {
	if key == nil {
		return nil, fmt.Errorf("nil public key")
	}

	// Verify it's a secp256k1 key
	if key.Curve != ps.Curve {
		return nil, fmt.Errorf("public key is not on secp256k1 curve")
	}

	return elliptic.Marshal(key.Curve, key.X, key.Y), nil
}

// UnmarshalPublicKey decodes a public key from compressed format
func (ps *ProofSystem) UnmarshalPublicKey(data []byte) (*ecdsa.PublicKey, error) {
	x, y := elliptic.Unmarshal(ps.Curve, data)
	if x == nil || y == nil {
		return nil, fmt.Errorf("invalid public key")
	}

	return &ecdsa.PublicKey{
		Curve: ps.Curve,
		X:     x,
		Y:     y,
	}, nil
}
