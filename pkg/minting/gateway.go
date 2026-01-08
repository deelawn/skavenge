package minting

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// GatewayClient handles communication with the linked accounts gateway
type GatewayClient struct {
	baseURL string
	client  *http.Client
}

// LinkResponse represents the response from the gateway
type LinkResponse struct {
	Success           bool   `json:"success"`
	SkavengePublicKey string `json:"skavenge_public_key,omitempty"`
	Message           string `json:"message,omitempty"`
}

// ErrorResponse represents an error response from the gateway
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// NewGatewayClient creates a new gateway client
func NewGatewayClient(gatewayURL string) *GatewayClient {
	return &GatewayClient{
		baseURL: strings.TrimSuffix(gatewayURL, "/"),
		client:  &http.Client{},
	}
}

// GetPublicKey retrieves the Skavenge public key for an Ethereum address
func (g *GatewayClient) GetPublicKey(ethereumAddress string) (*ecdsa.PublicKey, error) {
	// Normalize the address (remove 0x prefix if present, convert to lowercase)
	normalizedAddress := strings.ToLower(strings.TrimPrefix(ethereumAddress, "0x"))
	if !strings.HasPrefix(normalizedAddress, "0x") {
		normalizedAddress = "0x" + normalizedAddress
	}

	// Build the request URL
	reqURL := fmt.Sprintf("%s/link?ethereumAddress=%s", g.baseURL, url.QueryEscape(normalizedAddress))

	// Make the request
	resp, err := g.client.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to gateway: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, fmt.Errorf("gateway returned status %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("gateway error: %s", errResp.Error)
	}

	// Parse the response
	var linkResp LinkResponse
	if err := json.Unmarshal(body, &linkResp); err != nil {
		return nil, fmt.Errorf("failed to parse gateway response: %w", err)
	}

	if !linkResp.Success {
		return nil, fmt.Errorf("gateway request failed: %s", linkResp.Message)
	}

	if linkResp.SkavengePublicKey == "" {
		return nil, fmt.Errorf("no public key found for address %s", ethereumAddress)
	}

	// Parse the public key (hex format: 0x04 + X + Y, total 65 bytes)
	if !strings.HasPrefix(linkResp.SkavengePublicKey, "0x") {
		linkResp.SkavengePublicKey = "0x" + linkResp.SkavengePublicKey
	}

	pubKeyBytes, err := hexutil.Decode(linkResp.SkavengePublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key: %w", err)
	}

	// Verify format (should be 65 bytes: 0x04 + 32 bytes X + 32 bytes Y)
	if len(pubKeyBytes) != 65 || pubKeyBytes[0] != 0x04 {
		return nil, fmt.Errorf("invalid public key format: expected 65 bytes starting with 0x04, got %d bytes", len(pubKeyBytes))
	}

	// Extract X and Y coordinates
	x := new(big.Int).SetBytes(pubKeyBytes[1:33])
	y := new(big.Int).SetBytes(pubKeyBytes[33:65])

	// Create the public key (using secp256k1 curve)
	pubKey := &ecdsa.PublicKey{
		Curve: crypto.S256(),
		X:     x,
		Y:     y,
	}

	// Verify the point is on the curve
	if !pubKey.Curve.IsOnCurve(x, y) {
		return nil, fmt.Errorf("public key point is not on the curve")
	}

	return pubKey, nil
}
