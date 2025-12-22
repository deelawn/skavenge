package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// LinkRequest represents the JSON payload for POST /link
type LinkRequest struct {
	EthereumAddress   string `json:"ethereum_address"`
	SkavengePublicKey string `json:"skavenge_public_key"`
	Message           string `json:"message"`
	Signature         string `json:"signature"`
}

// LinkResponse represents the JSON response for successful operations
type LinkResponse struct {
	Success           bool   `json:"success"`
	Message           string `json:"message,omitempty"`
	SkavengePublicKey string `json:"skavenge_public_key,omitempty"`
}

// TransferCiphertextRequest represents the JSON payload for POST /transfers
type TransferCiphertextRequest struct {
	TransferID       string `json:"transfer_id"`
	BuyerCiphertext  string `json:"buyer_ciphertext"`
	SellerCiphertext string `json:"seller_ciphertext"`
	Message          string `json:"message"`
	Signature        string `json:"signature"`
}

// TransferCiphertextResponse represents the JSON response for GET /transfers
type TransferCiphertextResponse struct {
	Success          bool   `json:"success"`
	BuyerCiphertext  string `json:"buyer_ciphertext,omitempty"`
	SellerCiphertext string `json:"seller_ciphertext,omitempty"`
}

// ErrorResponse represents the JSON response for errors
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// Server holds the dependencies for the HTTP handlers
type Server struct {
	storage         Storage
	transferStorage TransferCiphertextStorage
	contractClient  ContractClientInterface
}

// NewServer creates a new Server instance
func NewServer(store Storage, transferStore TransferCiphertextStorage, client ContractClientInterface) *Server {
	return &Server{
		storage:         store,
		transferStorage: transferStore,
		contractClient:  client,
	}
}

// corsMiddleware adds CORS headers to all responses
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")

		// Call the next handler
		next(w, r)
	}
}

// HandleLink routes requests to the appropriate handler based on method
func (s *Server) HandleLink(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodPost:
		s.handlePostLink(w, r)
	case http.MethodGet:
		s.handleGetLink(w, r)
	default:
		writeErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// handlePostLink processes POST /link requests
func (s *Server) handlePostLink(w http.ResponseWriter, r *http.Request) {
	var req LinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	// Validate required fields
	if req.EthereumAddress == "" || req.SkavengePublicKey == "" || req.Message == "" || req.Signature == "" {
		writeErrorResponse(w, http.StatusBadRequest, "missing required fields")
		return
	}

	// Verify the signature
	valid, err := VerifySignature(req.Message, req.Signature, req.EthereumAddress)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		writeErrorResponse(w, http.StatusBadRequest, "signature verification failed: "+err.Error())
		return
	}

	if !valid {
		writeErrorResponse(w, http.StatusUnauthorized, "invalid signature")
		return
	}

	// Normalize the ethereum address to lowercase for consistent storage
	normalizedAddress := strings.ToLower(req.EthereumAddress)

	// Check if the address is already linked (immutability check)
	if _, err := s.storage.Get(normalizedAddress); err == nil {
		writeErrorResponse(w, http.StatusConflict, "ethereum address already linked (keys are immutable)")
		return
	}

	// Store the linkage
	s.storage.Set(normalizedAddress, req.SkavengePublicKey)

	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(LinkResponse{
		Success: true,
		Message: "linkage created successfully",
	})
}

// handleGetLink processes GET /link requests
func (s *Server) handleGetLink(w http.ResponseWriter, r *http.Request) {
	// Get the ethereum address from query parameter
	ethereumAddress := r.URL.Query().Get("ethereumAddress")
	if ethereumAddress == "" {
		writeErrorResponse(w, http.StatusBadRequest, "missing ethereumAddress query parameter")
		return
	}

	// Normalize the ethereum address to lowercase
	normalizedAddress := strings.ToLower(ethereumAddress)

	// Retrieve the linkage
	skavengePublicKey, err := s.storage.Get(normalizedAddress)
	if err != nil {
		if errors.Is(err, ErrKeyNotFound) {
			writeErrorResponse(w, http.StatusNotFound, "ethereum address not found")
			return
		}
		log.Printf("Storage error: %v", err)
		writeErrorResponse(w, http.StatusInternalServerError, "failed to retrieve linkage")
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LinkResponse{
		Success:           true,
		SkavengePublicKey: skavengePublicKey,
	})
}

// HandleTransfers routes requests to the appropriate handler based on method
func (s *Server) HandleTransfers(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodPost:
		s.handlePostTransfers(w, r)
	case http.MethodGet:
		s.handleGetTransfers(w, r)
	default:
		writeErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// handlePostTransfers processes POST /transfers requests (seller stores ciphertext)
func (s *Server) handlePostTransfers(w http.ResponseWriter, r *http.Request) {
	var req TransferCiphertextRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	// Validate required fields
	if req.TransferID == "" || req.BuyerCiphertext == "" || req.SellerCiphertext == "" || req.Message == "" || req.Signature == "" {
		writeErrorResponse(w, http.StatusBadRequest, "missing required fields")
		return
	}

	// Decode transfer ID from hex to bytes32
	transferIDBytes, err := hex.DecodeString(strings.TrimPrefix(req.TransferID, "0x"))
	if err != nil || len(transferIDBytes) != 32 {
		writeErrorResponse(w, http.StatusBadRequest, "invalid transfer ID format")
		return
	}

	var transferID [32]byte
	copy(transferID[:], transferIDBytes)

	// Retrieve transfer from blockchain
	ctx := context.Background()
	transfer, err := s.contractClient.GetTransferInfo(ctx, transferID)
	if err != nil {
		log.Printf("Failed to retrieve transfer: %v", err)
		writeErrorResponse(w, http.StatusNotFound, "transfer not found: "+err.Error())
		return
	}

	// Get the owner (seller) of the token being transferred
	owner, err := s.contractClient.GetTokenOwner(ctx, transfer.TokenID)
	if err != nil {
		log.Printf("Failed to get token owner: %v", err)
		writeErrorResponse(w, http.StatusInternalServerError, "failed to get token owner")
		return
	}

	// Normalize the owner address
	normalizedOwner := strings.ToLower(owner.Hex())

	// Look up the seller's skavenge public key
	sellerPublicKey, err := s.storage.Get(normalizedOwner)
	if err != nil {
		if errors.Is(err, ErrKeyNotFound) {
			writeErrorResponse(w, http.StatusNotFound, "seller public key not found")
			return
		}
		log.Printf("Storage error: %v", err)
		writeErrorResponse(w, http.StatusInternalServerError, "failed to retrieve seller public key")
		return
	}

	// Verify the signature was signed by the seller's skavenge private key
	valid, err := VerifySkavengeSignature(req.Message, req.Signature, sellerPublicKey)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		writeErrorResponse(w, http.StatusBadRequest, "signature verification failed: "+err.Error())
		return
	}

	if !valid {
		writeErrorResponse(w, http.StatusUnauthorized, "invalid signature")
		return
	}

	// Store the ciphertext
	s.transferStorage.SetTransferCiphertext(req.TransferID, req.BuyerCiphertext, req.SellerCiphertext)

	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(LinkResponse{
		Success: true,
		Message: "transfer ciphertext stored successfully",
	})
}

// handleGetTransfers processes GET /transfers requests (buyer retrieves ciphertext)
func (s *Server) handleGetTransfers(w http.ResponseWriter, r *http.Request) {
	// Get the transfer ID from query parameter
	transferID := r.URL.Query().Get("transferId")
	if transferID == "" {
		writeErrorResponse(w, http.StatusBadRequest, "missing transferId query parameter")
		return
	}

	// Get the signature and message from query parameters (or headers)
	signature := r.URL.Query().Get("signature")
	message := r.URL.Query().Get("message")

	if signature == "" || message == "" {
		writeErrorResponse(w, http.StatusBadRequest, "missing signature or message query parameters")
		return
	}

	// Decode transfer ID from hex to bytes32
	transferIDBytes, err := hex.DecodeString(strings.TrimPrefix(transferID, "0x"))
	if err != nil || len(transferIDBytes) != 32 {
		writeErrorResponse(w, http.StatusBadRequest, "invalid transfer ID format")
		return
	}

	var transferIDArray [32]byte
	copy(transferIDArray[:], transferIDBytes)

	// Retrieve transfer from blockchain
	ctx := context.Background()
	transfer, err := s.contractClient.GetTransferInfo(ctx, transferIDArray)
	if err != nil {
		log.Printf("Failed to retrieve transfer: %v", err)
		writeErrorResponse(w, http.StatusNotFound, "transfer not found: "+err.Error())
		return
	}

	// Normalize the buyer address
	normalizedBuyer := strings.ToLower(transfer.Buyer.Hex())

	// Look up the buyer's skavenge public key
	buyerPublicKey, err := s.storage.Get(normalizedBuyer)
	if err != nil {
		if errors.Is(err, ErrKeyNotFound) {
			writeErrorResponse(w, http.StatusNotFound, "buyer public key not found")
			return
		}
		log.Printf("Storage error: %v", err)
		writeErrorResponse(w, http.StatusInternalServerError, "failed to retrieve buyer public key")
		return
	}

	// Verify the signature was signed by the buyer's skavenge private key
	valid, err := VerifySkavengeSignature(message, signature, buyerPublicKey)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		writeErrorResponse(w, http.StatusBadRequest, "signature verification failed: "+err.Error())
		return
	}

	if !valid {
		writeErrorResponse(w, http.StatusUnauthorized, "invalid signature")
		return
	}

	// Retrieve the ciphertext
	buyerCiphertext, sellerCiphertext, err := s.transferStorage.GetTransferCiphertext(transferID)
	if err != nil {
		if errors.Is(err, ErrKeyNotFound) {
			writeErrorResponse(w, http.StatusNotFound, "transfer ciphertext not found")
			return
		}
		log.Printf("Storage error: %v", err)
		writeErrorResponse(w, http.StatusInternalServerError, "failed to retrieve transfer ciphertext")
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TransferCiphertextResponse{
		Success:          true,
		BuyerCiphertext:  buyerCiphertext,
		SellerCiphertext: sellerCiphertext,
	})
}

// writeErrorResponse is a helper to write error responses
func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Error:   message,
	})
}

// HandleHealth processes health check requests
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}

func main() {
	// Parse command-line flags
	port := flag.Int("port", 4591, "Port to listen on")
	rpcURL := flag.String("rpc", "http://localhost:8545", "Blockchain RPC URL")
	contractAddress := flag.String("contract", "", "Skavenge contract address")
	flag.Parse()

	if *contractAddress == "" {
		log.Fatal("Contract address is required. Use -contract flag.")
	}

	// Initialize in-memory storage
	store := NewInMemoryStorage()
	transferStore := NewInMemoryTransferStorage()

	// Initialize contract client
	contractClient, err := NewContractClient(*rpcURL, *contractAddress)
	if err != nil {
		log.Fatalf("Failed to create contract client: %v", err)
	}
	defer contractClient.Close()

	// Create server
	server := NewServer(store, transferStore, contractClient)

	// Register handlers with CORS middleware
	http.HandleFunc("/link", corsMiddleware(server.HandleLink))
	http.HandleFunc("/transfers", corsMiddleware(server.HandleTransfers))
	http.HandleFunc("/health", HandleHealth)

	// Start server
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting linked accounts gateway server on %s", addr)
	log.Printf("Connected to blockchain at %s", *rpcURL)
	log.Printf("Using contract at %s", *contractAddress)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
