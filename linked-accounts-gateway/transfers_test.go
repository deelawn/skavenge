package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// MockContractClient mocks the ContractClient for testing
type MockContractClient struct {
	transfers map[[32]byte]*TransferInfo
	owners    map[string]common.Address // tokenID -> owner
}

func NewMockContractClient() *MockContractClient {
	return &MockContractClient{
		transfers: make(map[[32]byte]*TransferInfo),
		owners:    make(map[string]common.Address),
	}
}

func (m *MockContractClient) GetTransferInfo(ctx context.Context, transferID [32]byte) (*TransferInfo, error) {
	transfer, exists := m.transfers[transferID]
	if !exists {
		return nil, fmt.Errorf("transfer does not exist")
	}
	return transfer, nil
}

func (m *MockContractClient) GetTokenOwner(ctx context.Context, tokenID *big.Int) (common.Address, error) {
	owner, exists := m.owners[tokenID.String()]
	if !exists {
		return common.Address{}, fmt.Errorf("token not found")
	}
	return owner, nil
}

func (m *MockContractClient) GetAddress() common.Address {
	// Return a mock contract address for testing
	return common.HexToAddress("0x0000000000000000000000000000000000000001")
}

func (m *MockContractClient) Close() {
	// No-op for mock
}

// Helper function to generate secp256k1 key pair and export public key in uncompressed format
func generateSecp256k1KeyPair() (*ecdsa.PrivateKey, []byte, error) {
	// Generate secp256k1 private key (same curve as Ethereum)
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	// Export public key in uncompressed format: 0x04 || X (32 bytes) || Y (32 bytes)
	x := privateKey.PublicKey.X.Bytes()
	y := privateKey.PublicKey.Y.Bytes()

	// Pad to 32 bytes
	xPadded := make([]byte, 32)
	yPadded := make([]byte, 32)
	copy(xPadded[32-len(x):], x)
	copy(yPadded[32-len(y):], y)

	// Create uncompressed point: 0x04 || X || Y
	point := make([]byte, 65)
	point[0] = 0x04
	copy(point[1:33], xPadded)
	copy(point[33:65], yPadded)

	return privateKey, point, nil
}

// Helper function to sign a message with secp256k1 private key
func signMessageSecp256k1(privateKey *ecdsa.PrivateKey, message string) (string, error) {
	hash := sha256.Sum256([]byte(message))

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	// Encode as ASN.1 DER
	sigStruct := struct {
		R, S *big.Int
	}{r, s}

	sigBytes, err := asn1.Marshal(sigStruct)
	if err != nil {
		return "", err
	}

	return "0x" + hex.EncodeToString(sigBytes), nil
}

// TestHandlePostTransfers tests the POST /transfers endpoint
func TestHandlePostTransfers(t *testing.T) {
	// Setup mock contract client
	mockClient := NewMockContractClient()

	// Generate seller key pair
	sellerPrivateKey, sellerPublicKey, err := generateSecp256k1KeyPair()
	if err != nil {
		t.Fatalf("Failed to generate seller key pair: %v", err)
	}

	// Generate Ethereum address for seller
	sellerEthKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate seller Ethereum key: %v", err)
	}
	sellerAddress := crypto.PubkeyToAddress(sellerEthKey.PublicKey)

	// Setup storage
	store := NewInMemoryStorage()
	transferStore := NewInMemoryTransferStorage()

	// Link seller's Ethereum address to Skavenge public key (normalized to lowercase)
	sellerPublicKeyHex := "0x" + hex.EncodeToString(sellerPublicKey)
	store.Set(strings.ToLower(sellerAddress.Hex()), sellerPublicKeyHex)

	// Create a mock transfer
	transferID := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	tokenID := big.NewInt(1)
	buyerAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")

	mockClient.transfers[transferID] = &TransferInfo{
		Buyer:   buyerAddress,
		TokenID: tokenID,
		Value:   big.NewInt(1000000000000000000), // 1 ETH
	}
	mockClient.owners[tokenID.String()] = sellerAddress

	// Create server
	server := NewServer(store, transferStore, mockClient)

	// Test successful POST
	t.Run("successful POST", func(t *testing.T) {
		message := "Transfer ciphertext for transfer " + hex.EncodeToString(transferID[:])
		signature, err := signMessageSecp256k1(sellerPrivateKey, message)
		if err != nil {
			t.Fatalf("Failed to sign message: %v", err)
		}

		reqBody := TransferCiphertextRequest{
			TransferID:       "0x" + hex.EncodeToString(transferID[:]),
			BuyerCiphertext:  "buyer_ciphertext_data",
			SellerCiphertext: "seller_ciphertext_data",
			Message:          message,
			Signature:        signature,
		}

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		server.handlePostTransfers(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status %d, got %d. Body: %s", http.StatusCreated, w.Code, w.Body.String())
		}

		// Verify ciphertext was stored
		buyerCipher, sellerCipher, err := transferStore.GetTransferCiphertext("0x" + hex.EncodeToString(transferID[:]))
		if err != nil {
			t.Errorf("Failed to retrieve stored ciphertext: %v", err)
		}
		if buyerCipher != "buyer_ciphertext_data" || sellerCipher != "seller_ciphertext_data" {
			t.Errorf("Stored ciphertext doesn't match. Got buyer=%s, seller=%s", buyerCipher, sellerCipher)
		}
	})

	// Test missing fields
	t.Run("missing fields", func(t *testing.T) {
		reqBody := TransferCiphertextRequest{
			TransferID: "0x" + hex.EncodeToString(transferID[:]),
			// Missing other fields
		}

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		server.handlePostTransfers(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	// Test invalid transfer ID
	t.Run("invalid transfer ID", func(t *testing.T) {
		reqBody := TransferCiphertextRequest{
			TransferID:       "invalid",
			BuyerCiphertext:  "buyer_ciphertext_data",
			SellerCiphertext: "seller_ciphertext_data",
			Message:          "test",
			Signature:        "0x1234",
		}

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		server.handlePostTransfers(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	// Test transfer not found
	t.Run("transfer not found", func(t *testing.T) {
		nonExistentTransferID := [32]byte{99, 99, 99}
		message := "Transfer ciphertext"
		signature, _ := signMessageSecp256k1(sellerPrivateKey, message)

		reqBody := TransferCiphertextRequest{
			TransferID:       "0x" + hex.EncodeToString(nonExistentTransferID[:]),
			BuyerCiphertext:  "buyer_ciphertext_data",
			SellerCiphertext: "seller_ciphertext_data",
			Message:          message,
			Signature:        signature,
		}

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		server.handlePostTransfers(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
		}
	})

	// Test invalid signature
	t.Run("invalid signature", func(t *testing.T) {
		message := "Transfer ciphertext"
		// Sign with a different key
		wrongKey, _, _ := generateSecp256k1KeyPair()
		signature, _ := signMessageSecp256k1(wrongKey, message)

		reqBody := TransferCiphertextRequest{
			TransferID:       "0x" + hex.EncodeToString(transferID[:]),
			BuyerCiphertext:  "buyer_ciphertext_data",
			SellerCiphertext: "seller_ciphertext_data",
			Message:          message,
			Signature:        signature,
		}

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		server.handlePostTransfers(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
		}
	})
}

// TestHandleGetTransfers tests the GET /transfers endpoint
func TestHandleGetTransfers(t *testing.T) {
	// Setup mock contract client
	mockClient := NewMockContractClient()

	// Generate buyer key pair
	buyerPrivateKey, buyerPublicKey, err := generateSecp256k1KeyPair()
	if err != nil {
		t.Fatalf("Failed to generate buyer key pair: %v", err)
	}

	// Generate Ethereum address for buyer
	buyerEthKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate buyer Ethereum key: %v", err)
	}
	buyerAddress := crypto.PubkeyToAddress(buyerEthKey.PublicKey)

	// Setup storage
	store := NewInMemoryStorage()
	transferStore := NewInMemoryTransferStorage()

	// Link buyer's Ethereum address to Skavenge public key (normalized to lowercase)
	buyerPublicKeyHex := "0x" + hex.EncodeToString(buyerPublicKey)
	store.Set(strings.ToLower(buyerAddress.Hex()), buyerPublicKeyHex)

	// Create a mock transfer
	transferID := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	tokenID := big.NewInt(1)

	mockClient.transfers[transferID] = &TransferInfo{
		Buyer:   buyerAddress,
		TokenID: tokenID,
		Value:   big.NewInt(1000000000000000000), // 1 ETH
	}

	// Store ciphertext
	transferIDHex := "0x" + hex.EncodeToString(transferID[:])
	transferStore.SetTransferCiphertext(transferIDHex, "buyer_ciphertext_data", "seller_ciphertext_data")

	// Create server
	server := NewServer(store, transferStore, mockClient)

	// Test successful GET
	t.Run("successful GET", func(t *testing.T) {
		message := "Retrieve ciphertext for transfer " + hex.EncodeToString(transferID[:])
		signature, err := signMessageSecp256k1(buyerPrivateKey, message)
		if err != nil {
			t.Fatalf("Failed to sign message: %v", err)
		}

		// Properly URL encode the query parameters
		urlStr := fmt.Sprintf("/transfers?transferId=%s&message=%s&signature=%s",
			url.QueryEscape(transferIDHex),
			url.QueryEscape(message),
			url.QueryEscape(signature))
		req := httptest.NewRequest(http.MethodGet, urlStr, nil)
		w := httptest.NewRecorder()

		server.handleGetTransfers(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d. Body: %s", http.StatusOK, w.Code, w.Body.String())
		}

		var resp TransferCiphertextResponse
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if !resp.Success {
			t.Errorf("Expected success=true, got false")
		}
		if resp.BuyerCiphertext != "buyer_ciphertext_data" || resp.SellerCiphertext != "seller_ciphertext_data" {
			t.Errorf("Unexpected ciphertext. Got buyer=%s, seller=%s", resp.BuyerCiphertext, resp.SellerCiphertext)
		}
	})

	// Test missing query parameters
	t.Run("missing query parameters", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/transfers", nil)
		w := httptest.NewRecorder()

		server.handleGetTransfers(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	// Test transfer not found
	t.Run("transfer not found", func(t *testing.T) {
		nonExistentTransferID := [32]byte{99, 99, 99}
		message := "Retrieve ciphertext"
		signature, _ := signMessageSecp256k1(buyerPrivateKey, message)

		urlStr := fmt.Sprintf("/transfers?transferId=%s&message=%s&signature=%s",
			url.QueryEscape("0x"+hex.EncodeToString(nonExistentTransferID[:])),
			url.QueryEscape(message),
			url.QueryEscape(signature))
		req := httptest.NewRequest(http.MethodGet, urlStr, nil)
		w := httptest.NewRecorder()

		server.handleGetTransfers(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
		}
	})

	// Test invalid signature
	t.Run("invalid signature", func(t *testing.T) {
		message := "Retrieve ciphertext"
		// Sign with a different key
		wrongKey, _, _ := generateSecp256k1KeyPair()
		signature, _ := signMessageSecp256k1(wrongKey, message)

		urlStr := fmt.Sprintf("/transfers?transferId=%s&message=%s&signature=%s",
			url.QueryEscape(transferIDHex),
			url.QueryEscape(message),
			url.QueryEscape(signature))
		req := httptest.NewRequest(http.MethodGet, urlStr, nil)
		w := httptest.NewRecorder()

		server.handleGetTransfers(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
		}
	})

	// Test ciphertext not found
	t.Run("ciphertext not found", func(t *testing.T) {
		// Create a new transfer without stored ciphertext
		newTransferID := [32]byte{50, 50, 50}
		mockClient.transfers[newTransferID] = &TransferInfo{
			Buyer:   buyerAddress,
			TokenID: tokenID,
			Value:   big.NewInt(1000000000000000000),
		}

		message := "Retrieve ciphertext"
		signature, _ := signMessageSecp256k1(buyerPrivateKey, message)

		urlStr := fmt.Sprintf("/transfers?transferId=%s&message=%s&signature=%s",
			url.QueryEscape("0x"+hex.EncodeToString(newTransferID[:])),
			url.QueryEscape(message),
			url.QueryEscape(signature))
		req := httptest.NewRequest(http.MethodGet, urlStr, nil)
		w := httptest.NewRecorder()

		server.handleGetTransfers(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

// TestVerifySkavengeSignature tests the Skavenge signature verification
func TestVerifySkavengeSignature(t *testing.T) {
	// Generate key pair
	privateKey, publicKey, err := generateSecp256k1KeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	publicKeyHex := "0x" + hex.EncodeToString(publicKey)

	tests := []struct {
		name      string
		message   string
		wantValid bool
		wantError bool
	}{
		{
			name:      "valid signature",
			message:   "Test message",
			wantValid: true,
			wantError: false,
		},
		{
			name:      "empty message",
			message:   "",
			wantValid: true,
			wantError: false,
		},
		{
			name:      "long message",
			message:   "This is a very long message that tests the signature verification with more data to ensure it works correctly regardless of message length.",
			wantValid: true,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signature, err := signMessageSecp256k1(privateKey, tt.message)
			if err != nil {
				t.Fatalf("Failed to sign message: %v", err)
			}

			valid, err := VerifySkavengeSignature(tt.message, signature, publicKeyHex)

			if tt.wantError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if valid != tt.wantValid {
				t.Errorf("Expected valid=%v, got %v", tt.wantValid, valid)
			}
		})
	}

	// Test with wrong message
	t.Run("wrong message", func(t *testing.T) {
		signature, _ := signMessageSecp256k1(privateKey, "Original message")
		valid, err := VerifySkavengeSignature("Different message", signature, publicKeyHex)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if valid {
			t.Errorf("Expected invalid signature for wrong message")
		}
	})

	// Test with wrong key
	t.Run("wrong key", func(t *testing.T) {
		_, wrongPublicKey, _ := generateSecp256k1KeyPair()
		wrongPublicKeyHex := "0x" + hex.EncodeToString(wrongPublicKey)

		signature, _ := signMessageSecp256k1(privateKey, "Test message")
		valid, err := VerifySkavengeSignature("Test message", signature, wrongPublicKeyHex)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if valid {
			t.Errorf("Expected invalid signature for wrong key")
		}
	})
}
