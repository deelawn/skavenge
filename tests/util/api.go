// Package util provides utility functions for the Skavenge contract tests.
package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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
	RecipientPublicKey string `json:"recipientPublicKey"`
}

// EncryptMessageResponse represents a response from the message encryption endpoint.
type EncryptMessageResponse struct {
	EncryptedMessage string `json:"encryptedMessage"`
}

// DecryptMessageRequest represents a request to decrypt a message.
type DecryptMessageRequest struct {
	EncryptedMessage string `json:"encryptedMessage"`
	PrivateKey       string `json:"privateKey"`
}

// DecryptMessageResponse represents a response from the message decryption endpoint.
type DecryptMessageResponse struct {
	Message string `json:"message"`
}

// GenerateProofRequest represents a request to generate a proof.
type GenerateProofRequest struct {
	Message            string `json:"message"`
	RecipientPublicKey string `json:"recipientPublicKey"`
	SenderPrivateKey   string `json:"senderPrivateKey"`
}

// GenerateProofResponse represents a response from the proof generation endpoint.
type GenerateProofResponse struct {
	Proof       string `json:"proof"`
	NewClueHash string `json:"newClueHash"`
}

// VerifyProofRequest represents a request to verify a proof.
type VerifyProofRequest struct {
	Proof               string `json:"proof"`
	RecipientPrivateKey string `json:"recipientPrivateKey"`
}

// VerifyProofResponse represents a response from the proof verification endpoint.
type VerifyProofResponse struct {
	Valid bool `json:"valid"`
}

// EncryptMessage encrypts a message for a recipient.
func (c *APIClient) EncryptMessage(message, recipientPublicKey string) (string, error) {
	req := EncryptMessageRequest{
		Message:            message,
		RecipientPublicKey: recipientPublicKey,
	}

	var resp EncryptMessageResponse
	err := c.doRequest(EncryptEndpoint, req, &resp)
	if err != nil {
		return "", fmt.Errorf("encryption failed: %w", err)
	}

	return resp.EncryptedMessage, nil
}

// DecryptMessage decrypts a message using a private key.
func (c *APIClient) DecryptMessage(encryptedMessage, privateKey string) (string, error) {
	req := DecryptMessageRequest{
		EncryptedMessage: encryptedMessage,
		PrivateKey:       privateKey,
	}

	var resp DecryptMessageResponse
	err := c.doRequest(DecryptEndpoint, req, &resp)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}

	return resp.Message, nil
}

// GenerateProof generates a proof for transferring a clue.
func (c *APIClient) GenerateProof(message, recipientPublicKey, senderPrivateKey string) (string, string, error) {
	req := GenerateProofRequest{
		Message:            message,
		RecipientPublicKey: recipientPublicKey,
		SenderPrivateKey:   senderPrivateKey,
	}

	var resp GenerateProofResponse
	err := c.doRequest(GenerateProofEndpoint, req, &resp)
	if err != nil {
		return "", "", fmt.Errorf("proof generation failed: %w", err)
	}

	return resp.Proof, resp.NewClueHash, nil
}

// VerifyProof verifies a proof for transferring a clue.
func (c *APIClient) VerifyProof(proof, recipientPrivateKey string) (bool, error) {
	req := VerifyProofRequest{
		Proof:               proof,
		RecipientPrivateKey: recipientPrivateKey,
	}

	var resp VerifyProofResponse
	err := c.doRequest(VerifyProofEndpoint, req, &resp)
	if err != nil {
		return false, fmt.Errorf("proof verification failed: %w", err)
	}

	return resp.Valid, nil
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
