package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/deelawn/skavenge/linked-accounts-gateway/storage"
)

// LinkRequest represents the JSON payload for POST /link
type LinkRequest struct {
	EthereumAddress    string `json:"ethereumAddress"`
	SkavengePublicKey  string `json:"skavengePublicKey"`
	Message            string `json:"message"`
	Signature          string `json:"signature"`
}

// LinkResponse represents the JSON response for successful operations
type LinkResponse struct {
	Success           bool   `json:"success"`
	Message           string `json:"message,omitempty"`
	SkavengePublicKey string `json:"skavengePublicKey,omitempty"`
}

// ErrorResponse represents the JSON response for errors
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// Server holds the dependencies for the HTTP handlers
type Server struct {
	storage storage.Storage
}

// NewServer creates a new Server instance
func NewServer(store storage.Storage) *Server {
	return &Server{
		storage: store,
	}
}

// HandleLink routes requests to the appropriate handler based on method
func (s *Server) HandleLink(w http.ResponseWriter, r *http.Request) {
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

	// Store the linkage
	if err := s.storage.Set(normalizedAddress, req.SkavengePublicKey); err != nil {
		if errors.Is(err, storage.ErrKeyExists) {
			writeErrorResponse(w, http.StatusConflict, "ethereum address already linked (keys are immutable)")
			return
		}
		log.Printf("Storage error: %v", err)
		writeErrorResponse(w, http.StatusInternalServerError, "failed to store linkage")
		return
	}

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
		if errors.Is(err, storage.ErrKeyNotFound) {
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

// writeErrorResponse is a helper to write error responses
func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Error:   message,
	})
}

func main() {
	// Initialize in-memory storage
	store := storage.NewInMemoryStorage()

	// Create server
	server := NewServer(store)

	// Register handlers
	http.HandleFunc("/link", server.HandleLink)

	// Start server
	port := ":8080"
	log.Printf("Starting linked accounts gateway server on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
