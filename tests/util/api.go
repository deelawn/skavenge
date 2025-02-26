// Package util provides utility functions for the Skavenge contract tests.
package util

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"time"

	proto "github.com/deelawn/skavenge-zk-proof/api/zkproof"
	"github.com/deelawn/skavenge-zk-proof/zk"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// DefaultGRPCAddr is the default address for the ZK proof gRPC API.
	DefaultGRPCAddr = "localhost:9090"

	// DefaultTimeout for gRPC calls
	DefaultTimeout = 30 * time.Second
)

// GRPCClient is a client for the ZK proof gRPC API.
type GRPCClient struct {
	conn   *grpc.ClientConn
	client proto.ServiceClient
}

// NewGRPCClient creates a new gRPC client with the default address.
func NewGRPCClient() (*GRPCClient, error) {
	return NewGRPCClientWithAddr(DefaultGRPCAddr)
}

// NewGRPCClientWithAddr creates a new gRPC client with the specified address.
func NewGRPCClientWithAddr(addr string) (*GRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	client := proto.NewServiceClient(conn)
	return &GRPCClient{
		conn:   conn,
		client: client,
	}, nil
}

// Close closes the gRPC connection.
func (c *GRPCClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// encodePublicKey marshals an ECDSA public key for gRPC requests
func encodePublicKey(publicKey *ecdsa.PublicKey) ([]byte, error) {
	ps := zk.NewProofSystem()
	keyBytes, err := ps.MarshalPublicKey(publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %w", err)
	}

	return keyBytes, nil
}

// encodePrivateKey marshals an ECDSA private key for gRPC requests
func encodePrivateKey(privateKey *ecdsa.PrivateKey) ([]byte, error) {
	ps := zk.NewProofSystem()
	keyBytes, err := ps.MarshalPrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal private key: %w", err)
	}

	return keyBytes, nil
}

// EncryptMessage encrypts a message for a recipient.
func (c *GRPCClient) EncryptMessage(message string, recipientPublicKey *ecdsa.PublicKey) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	// Marshal the recipient public key
	encodedPublicKey, err := encodePublicKey(recipientPublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to encode recipient public key: %w", err)
	}

	// Create the request
	req := &proto.EncryptMessageRequest{
		Message: []byte(message),
		PubKey:  encodedPublicKey,
	}

	// Call the gRPC service
	resp, err := c.client.EncryptMessage(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("encryption failed: %w", err)
	}

	return resp.Ciphertext, nil
}

// DecryptMessage decrypts a message using a private key.
func (c *GRPCClient) DecryptMessage(encryptedMessage []byte, privateKey *ecdsa.PrivateKey) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	// Marshal the private key
	encodedPrivateKey, err := encodePrivateKey(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to encode private key: %w", err)
	}

	// Create the request
	req := &proto.DecryptMessageRequest{
		Ciphertext: encryptedMessage,
		PrivKey:    encodedPrivateKey,
	}

	// Call the gRPC service
	resp, err := c.client.DecryptMessage(ctx, req)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}

	return string(resp.Message), nil
}

// GeneratePartialProof generates a partial proof for transferring a clue.
func (c *GRPCClient) GeneratePartialProof(
	message string,
	sellerCipherText []byte,
	buyerPublicKey *ecdsa.PublicKey,
	sellerPublicKey *ecdsa.PublicKey,
) (*proto.GeneratePartialProofResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	// Marshal the keys
	encodedBuyerPublicKey, err := encodePublicKey(buyerPublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to encode buyer public key: %w", err)
	}

	encodedSellerPublicKey, err := encodePublicKey(sellerPublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to encode seller public key: %w", err)
	}

	// Create the partial proof request
	partialProofReq := &proto.GeneratePartialProofRequest{
		Message:          []byte(message),
		SellerPubKey:     encodedSellerPublicKey,
		BuyerPubKey:      encodedBuyerPublicKey,
		SellerCipherText: sellerCipherText,
	}

	// Call the gRPC service to generate the partial proof
	partialProofResp, err := c.client.GeneratePartialProof(ctx, partialProofReq)
	if err != nil {
		return nil, fmt.Errorf("partial proof generation failed: %w", err)
	}

	return partialProofResp, nil
}

// MarshalProof marshals a proof for on-chain use.
func (c *GRPCClient) MarshalProof(proof *proto.Proof) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	// Create the marshal request
	marshalReq := &proto.MarshalProofRequest{
		Proof: proof,
	}

	// Call the gRPC service to marshal the proof
	marshalResp, err := c.client.MarshalProof(ctx, marshalReq)
	if err != nil {
		return nil, fmt.Errorf("proof marshaling failed: %w", err)
	}

	return marshalResp.Proof, nil
}

// VerifyProof verifies a proof for a clue transfer.
func (c *GRPCClient) VerifyProof(proof []byte, sellerCipherText []byte) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	req := &proto.VerifyProofRequest{
		Proof:            proof,
		SellerCipherText: sellerCipherText,
	}

	resp, err := c.client.VerifyProof(ctx, req)
	if err != nil {
		return false, fmt.Errorf("proof verification failed: %w", err)
	}

	return resp.IsValid, nil
}
