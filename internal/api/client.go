// Package api provides client implementations for interacting with the ZK proof API.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is a client for the ZK proof API.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new ZK proof API client.
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// EncryptRequest represents a request to encrypt a message.
type EncryptRequest struct {
	Message            string `json:"message"`
	RecipientPublicKey string `json:"recipientPublicKey"`
}

// EncryptResponse represents a response from the message encryption endpoint.
type EncryptResponse struct {
	EncryptedMessage string `json:"encryptedMessage"`
}

// Encrypt encrypts a message for a recipient.
func (c *Client) Encrypt(message, recipientPublicKey string) (string, error) {
	req := EncryptRequest{
		Message:            message,
		RecipientPublicKey: recipientPublicKey,
	}

	var resp EncryptResponse
	if err := c.request(http.MethodPost, "/api/message/encrypt", req, &resp); err != nil {
		return "", fmt.Errorf("encrypt message request failed: %w", err)
	}

	return resp.EncryptedMessage, nil
}

// DecryptRequest represents a request to decrypt a message.
type DecryptRequest struct {
	EncryptedMessage string `json:"encryptedMessage"`
	PrivateKey       string `json:"privateKey"`
}

// DecryptResponse represents a response from the message decryption endpoint.
type DecryptResponse struct {
	Message string `json:"message"`
}

// Decrypt decrypts a message using a private key.
func (c *Client) Decrypt(encryptedMessage, privateKey string) (string, error) {
	req := DecryptRequest{
		EncryptedMessage: encryptedMessage,
		PrivateKey:       privateKey,
	}

	var resp DecryptResponse
	if err := c.request(http.MethodPost, "/api/message/decrypt", req, &resp); err != nil {
		return "", fmt.Errorf("decrypt message request failed: %w", err)
	}

	return resp.Message, nil
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

// GenerateProof generates a proof for transferring a clue.
func (c *Client) GenerateProof(message, recipientPublicKey, senderPrivateKey string) (string, string, error) {
	req := GenerateProofRequest{
		Message:            message,
		RecipientPublicKey: recipientPublicKey,
		SenderPrivateKey:   senderPrivateKey,
	}

	var resp GenerateProofResponse
	if err := c.request(http.MethodPost, "/api/proof/generate", req, &resp); err != nil {
		return "", "", fmt.Errorf("generate proof request failed: %w", err)
	}

	return resp.Proof, resp.NewClueHash, nil
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

// VerifyProof verifies a proof for transferring a clue.
func (c *Client) VerifyProof(proof, recipientPrivateKey string) (bool, error) {
	req := VerifyProofRequest{
		Proof:               proof,
		RecipientPrivateKey: recipientPrivateKey,
	}

	var resp VerifyProofResponse
	if err := c.request(http.MethodPost, "/api/proof/verify", req, &resp); err != nil {
		return false, fmt.Errorf("verify proof request failed: %w", err)
	}

	return resp.Valid, nil
}

// ErrorResponse represents an error response from the API.
type ErrorResponse struct {
	Error string `json:"error"`
}

// request performs an HTTP request to the API.
func (c *Client) request(method, endpoint string, body interface{}, target interface{}) error {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	url := c.baseURL + endpoint
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for error response
	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.Unmarshal(respBody, &errResp); err == nil && errResp.Error != "" {
			return fmt.Errorf("API error: %s", errResp.Error)
		}
		return fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(respBody))
	}

	// Unmarshal the response into the target
	if target != nil {
		if err := json.Unmarshal(respBody, target); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}
