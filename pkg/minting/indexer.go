package minting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// IndexerClient handles communication with the indexer API
type IndexerClient struct {
	baseURL string
	client  *http.Client
}

// IndexerClueRequest represents the request body for creating a clue in the indexer
type IndexerClueRequest struct {
	ClueID       uint64 `json:"clueId"`
	Contents     string `json:"contents"`
	SolutionHash string `json:"solutionHash"`
	PointValue   uint64 `json:"pointValue"`
	SolveReward  uint64 `json:"solveReward"`
	Force        bool   `json:"force"`
}

// IndexerErrorResponse represents an error response from the indexer API
type IndexerErrorResponse struct {
	Error string `json:"error"`
}

// NewIndexerClient creates a new indexer client
func NewIndexerClient(indexerURL string) *IndexerClient {
	return &IndexerClient{
		baseURL: strings.TrimSuffix(indexerURL, "/"),
		client:  &http.Client{},
	}
}

// SaveClue sends clue data to the indexer API
func (ic *IndexerClient) SaveClue(clueData *ClueData, tokenID uint64) error {
	// Hash the solution to match what's stored on-chain
	solutionHash := crypto.Keccak256Hash([]byte(clueData.Solution))

	// Convert solve reward to uint64 (wei amount)
	var solveReward uint64
	if clueData.SolveReward != nil {
		solveReward = clueData.SolveReward.Uint64()
	}

	// Prepare the request body
	reqBody := IndexerClueRequest{
		ClueID:       tokenID,
		Contents:     clueData.Content,
		SolutionHash: solutionHash.Hex(),
		PointValue:   uint64(clueData.PointValue),
		SolveReward:  solveReward,
		Force:        false, // Don't overwrite existing clues by default
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal clue data: %w", err)
	}

	// Make the POST request
	url := fmt.Sprintf("%s/api/clues", ic.baseURL)
	resp, err := ic.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send clue to indexer: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read indexer response: %w", err)
	}

	// Check for errors
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		var errResp IndexerErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return fmt.Errorf("indexer returned status %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("indexer error: %s", errResp.Error)
	}

	return nil
}
