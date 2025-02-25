// Package util provides utility functions for the Skavenge contract tests.
package util

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/deelawn/skavenge-zk-proof/zk"
)

const (
	// DefaultAPIBaseURL is the default base URL for the ZK proof API.
	DefaultAPIBaseURL = "http://localhost:8080"

	// API endpoints
	EncryptEndpoint       = "/api/message/encrypt"
	DecryptEndpoint       = "/api/message/decrypt"
	GenerateProofEndpoint = "/api/proof/generate"
	VerifyProofEndpoint   = "/api/proof/verify"
)

// APIClient is a client for the ZK proof API.
type APIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewAPIClient creates a new API client with the default base URL.
func NewAPIClient() *APIClient {
	return &APIClient{
		BaseURL: DefaultAPIBaseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// EncryptMessageRequest represents a request to encrypt a message.
type EncryptMessageRequest struct {
	Message            string `json:"message"`
	RecipientPublicKey string `json:"pubKey"`
}

// EncryptMessageResponse represents a response from the message encryption endpoint.
type EncryptMessageResponse struct {
	Data struct {
		Ciphertext string `json:"ciphertext"`
	} `json:"data"`
	Message string `json:"message"`
}

// DecryptMessageRequest represents a request to decrypt a message.
type DecryptMessageRequest struct {
	EncryptedMessage string `json:"ciphertext"`
	PrivateKey       string `json:"privKey"`
}

// DecryptMessageResponse represents a response from the message decryption endpoint.
type DecryptMessageResponse struct {
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
	Message string `json:"message"`
}

// GenerateProofRequest represents a request to generate a proof.
type GenerateProofRequest struct {
	Message            string `json:"message"`
	RecipientPublicKey string `json:"buyerPubKey"`
	SenderPrivateKey   string `json:"sellerKey"`
	SellerCipherText   string `json:"sellerCipherText"` // TODO
}

// GenerateProofResponse represents a response from the proof generation endpoint.
type GenerateProofResponse struct {
	Data struct {
		Proof            string `json:"proof"`
		SellerCipherText string `json:"sellerCipherText"` // TODO
		BuyerCipherText  string `json:"buyerCipherText"`  // TODO
	} `json:"data"`
	Message string `json:"message"`
}

// VerifyProofRequest represents a request to verify a proof.
type VerifyProofRequest struct {
	Proof            string `json:"proof"`
	SellerCipherText string `json:"sellerCipherText"` // TODO
}

// VerifyProofResponse represents a response from the proof verification endpoint.
type VerifyProofResponse struct {
	Data struct {
		Valid bool `json:"valid"`
	} `json:"data"`
	Message string `json:"message"`
}

// encodePublicKey marshals and base64 encodes an ECDSA public key for API requests
func encodePublicKey(publicKey *ecdsa.PublicKey) (string, error) {
	ps := zk.NewProofSystem()
	keyBytes, err := ps.MarshalPublicKey(publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to marshal public key: %w", err)
	}

	// Base64 encode the DER bytes
	return base64.StdEncoding.EncodeToString(keyBytes), nil
}

// encodePrivateKey marshals and base64 encodes an ECDSA private key for API requests
func encodePrivateKey(privateKey *ecdsa.PrivateKey) (string, error) {
	ps := zk.NewProofSystem()
	keyBytes, err := ps.MarshalPrivateKey(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to marshal private key: %w", err)
	}

	// Base64 encode the DER bytes
	return base64.StdEncoding.EncodeToString(keyBytes), nil
}

// EncryptMessage encrypts a message for a recipient.
func (c *APIClient) EncryptMessage(message string, recipientPublicKey *ecdsa.PublicKey) ([]byte, error) {
	// Base64 encode the message
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))

	// Encode the recipient public key
	encodedPublicKey, err := encodePublicKey(recipientPublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to encode recipient public key: %w", err)
	}

	req := EncryptMessageRequest{
		Message:            encodedMessage,
		RecipientPublicKey: encodedPublicKey,
	}

	var resp EncryptMessageResponse
	err = c.doRequest(EncryptEndpoint, req, &resp)
	if err != nil {
		return nil, fmt.Errorf("encryption failed: %w", err)
	}

	decodedCipherText, err := base64.StdEncoding.DecodeString(resp.Data.Ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	return decodedCipherText, nil
}

// DecryptMessage decrypts a message using a private key.
func (c *APIClient) DecryptMessage(encryptedMessage []byte, privateKey *ecdsa.PrivateKey) (string, error) {
	// Encode the private key
	encodedPrivateKey, err := encodePrivateKey(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to encode private key: %w", err)
	}

	encodedEncryptedMessage := base64.StdEncoding.EncodeToString(encryptedMessage)

	req := DecryptMessageRequest{
		EncryptedMessage: encodedEncryptedMessage,
		PrivateKey:       encodedPrivateKey,
	}

	var resp DecryptMessageResponse
	err = c.doRequest(DecryptEndpoint, req, &resp)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}

	// Decode the base64 encoded message from the response
	decodedMessage, err := base64.StdEncoding.DecodeString(resp.Data.Message)
	if err != nil {
		return "", fmt.Errorf("failed to decode response message: %w", err)
	}

	return string(decodedMessage), nil
}

// GenerateProof generates a proof for transferring a clue.
func (c *APIClient) GenerateProof(
	message string,
	b64EncryptedMessage string,
	recipientPublicKey *ecdsa.PublicKey,
	senderPrivateKey *ecdsa.PrivateKey,
) ([]byte, []byte, error) {
	// Base64 encode the message
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))

	// Encode the public key
	encodedPublicKey, err := encodePublicKey(recipientPublicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encode recipient public key: %w", err)
	}

	// Encode the private key
	encodedPrivateKey, err := encodePrivateKey(senderPrivateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encode sender private key: %w", err)
	}

	req := GenerateProofRequest{
		Message:            encodedMessage,
		RecipientPublicKey: encodedPublicKey,
		SenderPrivateKey:   encodedPrivateKey,
		SellerCipherText:   b64EncryptedMessage,
	}

	var resp GenerateProofResponse
	err = c.doRequest(GenerateProofEndpoint, req, &resp)
	if err != nil {
		return nil, nil, fmt.Errorf("proof generation failed: %w", err)
	}

	decodedProof, err := base64.StdEncoding.DecodeString(resp.Data.Proof)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode proof: %w", err)
	}

	decodedBuyerCipherText, err := base64.StdEncoding.DecodeString(resp.Data.BuyerCipherText)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode buyer ciphertext: %w", err)
	}

	return decodedProof, decodedBuyerCipherText, nil
}

// VerifyProof verifies a proof for transferring a clue.
func (c *APIClient) VerifyProof(proof []byte, sellerCipherText []byte) (bool, error) {
	encodedProof := base64.StdEncoding.EncodeToString(proof)
	encodedSellerCipherText := base64.StdEncoding.EncodeToString(sellerCipherText)

	req := VerifyProofRequest{
		Proof:            encodedProof,
		SellerCipherText: encodedSellerCipherText,
	}

	var resp VerifyProofResponse
	if err := c.doRequest(VerifyProofEndpoint, req, &resp); err != nil {
		return false, fmt.Errorf("proof verification failed: %w", err)
	}

	return resp.Data.Valid, nil
}

// doRequest performs an HTTP request to the API and unmarshals the response.
func (c *APIClient) doRequest(endpoint string, reqBody interface{}, respBody interface{}) error {
	// Marshal the request body
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", c.BaseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	// TODO: Use decoder and check for error response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for error response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned non-OK status code %d: %s", resp.StatusCode, string(body))
	}

	// Unmarshal the response body
	err = json.Unmarshal(body, respBody)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return nil
}
