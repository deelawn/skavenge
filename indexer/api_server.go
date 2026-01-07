package indexer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// APIServer provides HTTP endpoints for querying indexed data
type APIServer struct {
	storage Storage
	logger  *log.Logger
	server  *http.Server
}

// NewAPIServer creates a new API server instance
func NewAPIServer(storage Storage, logger *log.Logger, port int) *APIServer {
	api := &APIServer{
		storage: storage,
		logger:  logger,
	}

	mux := http.NewServeMux()

	// Clue endpoints
	mux.HandleFunc("/api/clues", api.handleClues)
	mux.HandleFunc("/api/clues/", api.handleClueByID)

	// Event endpoints
	mux.HandleFunc("/api/events", api.handleEvents)

	// Ownership endpoints
	mux.HandleFunc("/api/ownership", api.handleOwnership)

	api.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: corsMiddleware(mux),
	}

	return api
}

// Start starts the API server
func (api *APIServer) Start() error {
	api.logger.Printf("Starting API server on %s", api.server.Addr)
	if err := api.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("API server error: %w", err)
	}
	return nil
}

// Shutdown gracefully shuts down the API server
func (api *APIServer) Shutdown(ctx context.Context) error {
	api.logger.Println("Shutting down API server...")
	return api.server.Shutdown(ctx)
}

// handleClues handles GET /api/clues (list all) and POST /api/clues (create/update)
func (api *APIServer) handleClues(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.handleGetAllClues(w, r)
	case http.MethodPost:
		api.handleCreateClue(w, r)
	default:
		api.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// handleGetAllClues retrieves all clues with their ownership information
func (api *APIServer) handleGetAllClues(w http.ResponseWriter, r *http.Request) {
	clues, err := api.storage.GetAllClues(r.Context())
	if err != nil {
		api.logger.Printf("Failed to get all clues: %v", err)
		api.errorResponse(w, http.StatusInternalServerError, "failed to retrieve clues")
		return
	}

	// Fetch ownership data for all clues
	cluesWithOwnership := make([]*ClueWithOwner, 0, len(clues))
	for _, clue := range clues {
		ownership, err := api.storage.GetClueOwnership(r.Context(), clue.ClueID)
		if err != nil && !errors.Is(err, ErrNotFound) {
			api.logger.Printf("Failed to get ownership for clue %d: %v", clue.ClueID, err)
			api.errorResponse(w, http.StatusInternalServerError, "failed to retrieve ownership data")
			return
		}

		cluesWithOwnership = append(cluesWithOwnership, &ClueWithOwner{
			Clue:      *clue,
			Ownership: ownership,
		})
	}

	api.jsonResponse(w, http.StatusOK, cluesWithOwnership)
}

// handleClueByID handles GET /api/clues/:id
func (api *APIServer) handleClueByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Extract clue ID from path
	path := strings.TrimPrefix(r.URL.Path, "/api/clues/")
	clueID, err := strconv.ParseUint(path, 10, 64)
	if err != nil {
		api.errorResponse(w, http.StatusBadRequest, "invalid clue ID")
		return
	}

	clue, err := api.storage.GetClue(r.Context(), clueID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			api.errorResponse(w, http.StatusNotFound, "clue not found")
			return
		}
		api.logger.Printf("Failed to get clue %d: %v", clueID, err)
		api.errorResponse(w, http.StatusInternalServerError, "failed to retrieve clue")
		return
	}

	// Fetch ownership data
	ownership, err := api.storage.GetClueOwnership(r.Context(), clueID)
	if err != nil && !errors.Is(err, ErrNotFound) {
		api.logger.Printf("Failed to get ownership for clue %d: %v", clueID, err)
		api.errorResponse(w, http.StatusInternalServerError, "failed to retrieve ownership data")
		return
	}

	clueWithOwnership := &ClueWithOwner{
		Clue:      *clue,
		Ownership: ownership,
	}

	api.jsonResponse(w, http.StatusOK, clueWithOwnership)
}

// handleCreateClue handles POST /api/clues
func (api *APIServer) handleCreateClue(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ClueID       uint64 `json:"clueId"`
		Contents     string `json:"contents"`
		SolutionHash string `json:"solutionHash"`
		PointValue   uint64 `json:"pointValue"`
		SolveReward  uint64 `json:"solveReward"`
		Force        bool   `json:"force"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.errorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	clue := &Clue{
		ClueID:       req.ClueID,
		Contents:     req.Contents,
		SolutionHash: req.SolutionHash,
		PointValue:   req.PointValue,
		SolveReward:  req.SolveReward,
	}

	if err := api.storage.SaveClue(r.Context(), clue, req.Force); err != nil {
		if errors.Is(err, ErrAlreadyExists) {
			api.errorResponse(w, http.StatusConflict, "clue already exists, use force=true to overwrite")
			return
		}
		api.logger.Printf("Failed to save clue: %v", err)
		api.errorResponse(w, http.StatusInternalServerError, "failed to save clue")
		return
	}

	api.jsonResponse(w, http.StatusCreated, clue)
}

// handleEvents handles GET /api/events with query parameters
func (api *APIServer) handleEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	query := r.URL.Query()

	// Parse query options
	opts := &QueryOptions{
		SortOrder: SortOrderAsc,
	}

	if limit := query.Get("limit"); limit != "" {
		if val, err := strconv.Atoi(limit); err == nil {
			opts.Limit = val
		}
	}

	if offset := query.Get("offset"); offset != "" {
		if val, err := strconv.Atoi(offset); err == nil {
			opts.Offset = val
		}
	}

	if order := query.Get("order"); order != "" {
		if strings.ToLower(order) == "desc" {
			opts.SortOrder = SortOrderDesc
		}
	}

	if startTime := query.Get("startTime"); startTime != "" {
		if val, err := strconv.ParseInt(startTime, 10, 64); err == nil {
			opts.StartTime = val
		}
	}

	if endTime := query.Get("endTime"); endTime != "" {
		if val, err := strconv.ParseInt(endTime, 10, 64); err == nil {
			opts.EndTime = val
		}
	}

	var events []*Event
	var err error

	// Filter by clueId
	if clueIDStr := query.Get("clueId"); clueIDStr != "" {
		clueID, parseErr := strconv.ParseUint(clueIDStr, 10, 64)
		if parseErr != nil {
			api.errorResponse(w, http.StatusBadRequest, "invalid clueId")
			return
		}
		events, err = api.storage.GetEventsByClueID(r.Context(), clueID, opts)
	} else if eventType := query.Get("type"); eventType != "" {
		// Filter by event type
		events, err = api.storage.GetEventsByType(r.Context(), eventType, opts)
	} else if initiator := query.Get("initiator"); initiator != "" {
		// Filter by initiator
		events, err = api.storage.GetEventsByInitiator(r.Context(), initiator, opts)
	} else {
		// Get all events
		events, err = api.storage.GetAllEvents(r.Context(), opts)
	}

	if err != nil {
		api.logger.Printf("Failed to get events: %v", err)
		api.errorResponse(w, http.StatusInternalServerError, "failed to retrieve events")
		return
	}

	api.jsonResponse(w, http.StatusOK, events)
}

// handleOwnership handles GET /api/ownership with query parameters
func (api *APIServer) handleOwnership(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	query := r.URL.Query()

	var ownerships []*ClueOwnership
	var err error

	// Filter by owner address
	if owner := query.Get("owner"); owner != "" {
		ownerships, err = api.storage.GetClueOwnershipsByOwner(r.Context(), owner)
	} else if clueIDStr := query.Get("clueId"); clueIDStr != "" {
		// Filter by clue ID
		clueID, parseErr := strconv.ParseUint(clueIDStr, 10, 64)
		if parseErr != nil {
			api.errorResponse(w, http.StatusBadRequest, "invalid clueId")
			return
		}
		ownership, err := api.storage.GetClueOwnership(r.Context(), clueID)
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				api.errorResponse(w, http.StatusNotFound, "ownership not found")
				return
			}
			api.logger.Printf("Failed to get ownership: %v", err)
			api.errorResponse(w, http.StatusInternalServerError, "failed to retrieve ownership")
			return
		}
		api.jsonResponse(w, http.StatusOK, ownership)
		return
	} else {
		// Get all ownerships
		ownerships, err = api.storage.GetAllClueOwnerships(r.Context())
	}

	if err != nil {
		api.logger.Printf("Failed to get ownerships: %v", err)
		api.errorResponse(w, http.StatusInternalServerError, "failed to retrieve ownerships")
		return
	}

	api.jsonResponse(w, http.StatusOK, ownerships)
}

// jsonResponse sends a JSON response
func (api *APIServer) jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		api.logger.Printf("Failed to encode JSON response: %v", err)
	}
}

// errorResponse sends an error JSON response
func (api *APIServer) errorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := map[string]string{"error": message}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		api.logger.Printf("Failed to encode error response: %v", err)
	}
}

// corsMiddleware adds CORS headers to allow cross-origin requests
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
