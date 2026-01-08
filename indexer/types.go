package indexer

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// EventType represents the type of event
type EventType string

const (
	EventTypeClueMinted              EventType = "ClueMinted"
	EventTypeClueAttemptFailed       EventType = "ClueAttemptFailed"
	EventTypeClueSolved              EventType = "ClueSolved"
	EventTypeSalePriceSet            EventType = "SalePriceSet"
	EventTypeSalePriceRemoved        EventType = "SalePriceRemoved"
	EventTypeTransferInitiated       EventType = "TransferInitiated"
	EventTypeProofProvided           EventType = "ProofProvided"
	EventTypeProofVerified           EventType = "ProofVerified"
	EventTypeTransferCompleted       EventType = "TransferCompleted"
	EventTypeTransferCancelled       EventType = "TransferCancelled"
	EventTypeAuthorizedMinterUpdated EventType = "AuthorizedMinterUpdated"
	EventTypeTransfer                EventType = "Transfer"
	EventTypeApproval                EventType = "Approval"
	EventTypeApprovalForAll          EventType = "ApprovalForAll"
)

// Event represents a generic blockchain event
type Event struct {
	// Primary fields matching schema requirements
	ClueID           uint64 `json:"clue_id"`           // NFT Token ID (0 if not applicable)
	BlockNumber      uint64 `json:"block_number"`      // Block number
	TransactionIndex uint   `json:"transaction_index"` // Transaction index within block
	EventIndex       uint   `json:"event_index"`       // Event/Log index within transaction
	TransactionHash  string `json:"transaction_hash"`  // Transaction hash
	InitiatedBy      string `json:"initiated_by"`      // Transaction sender address
	EventType        string `json:"event_type"`        // Event type
	Metadata         []byte `json:"metadata"`          // Event-specific data as JSON
	Timestamp        int64  `json:"timestamp"`         // Unix timestamp

	// Additional fields for internal use
	BlockHash common.Hash `json:"block_hash"` // Block hash
	Removed   bool        `json:"removed"`    // True if log was removed due to reorg
}

// EventData contains event-specific data
type EventData interface {
	EventType() EventType
}

// ClueMintedData represents data for ClueMinted event
type ClueMintedData struct {
	TokenID   *big.Int       `json:"token_id"`
	Minter    common.Address `json:"minter"`
	Recipient common.Address `json:"recipient"`
}

func (d ClueMintedData) EventType() EventType { return EventTypeClueMinted }

// ClueAttemptFailedData represents data for ClueAttemptFailed event
type ClueAttemptFailedData struct {
	TokenID           *big.Int `json:"token_id"`
	AttemptedSolution string   `json:"attempted_solution"`
}

func (d ClueAttemptFailedData) EventType() EventType { return EventTypeClueAttemptFailed }

// ClueSolvedData represents data for ClueSolved event
type ClueSolvedData struct {
	TokenID  *big.Int `json:"token_id"`
	Solution string   `json:"solution"`
}

func (d ClueSolvedData) EventType() EventType { return EventTypeClueSolved }

// SalePriceSetData represents data for SalePriceSet event
type SalePriceSetData struct {
	TokenID *big.Int `json:"token_id"`
	Price   *big.Int `json:"price"`
}

func (d SalePriceSetData) EventType() EventType { return EventTypeSalePriceSet }

// SalePriceRemovedData represents data for SalePriceRemoved event
type SalePriceRemovedData struct {
	TokenID *big.Int `json:"token_id"`
}

func (d SalePriceRemovedData) EventType() EventType { return EventTypeSalePriceRemoved }

// TransferInitiatedData represents data for TransferInitiated event
type TransferInitiatedData struct {
	TransferID [32]byte       `json:"transfer_id"`
	Buyer      common.Address `json:"buyer"`
	TokenID    *big.Int       `json:"token_id"`
}

func (d TransferInitiatedData) EventType() EventType { return EventTypeTransferInitiated }

// ProofProvidedData represents data for ProofProvided event
type ProofProvidedData struct {
	TransferID  [32]byte `json:"transfer_id"`
	Proof       []byte   `json:"proof"`
	NewClueHash [32]byte `json:"new_clue_hash"`
	RValueHash  [32]byte `json:"r_value_hash"`
}

func (d ProofProvidedData) EventType() EventType { return EventTypeProofProvided }

// ProofVerifiedData represents data for ProofVerified event
type ProofVerifiedData struct {
	TransferID [32]byte `json:"transfer_id"`
}

func (d ProofVerifiedData) EventType() EventType { return EventTypeProofVerified }

// TransferCompletedData represents data for TransferCompleted event
type TransferCompletedData struct {
	TransferID [32]byte `json:"transfer_id"`
	RValue     *big.Int `json:"r_value"`
}

func (d TransferCompletedData) EventType() EventType { return EventTypeTransferCompleted }

// TransferCancelledData represents data for TransferCancelled event
type TransferCancelledData struct {
	TransferID  [32]byte       `json:"transfer_id"`
	CancelledBy common.Address `json:"cancelled_by"`
}

func (d TransferCancelledData) EventType() EventType { return EventTypeTransferCancelled }

// AuthorizedMinterUpdatedData represents data for AuthorizedMinterUpdated event
type AuthorizedMinterUpdatedData struct {
	OldMinter common.Address `json:"old_minter"`
	NewMinter common.Address `json:"new_minter"`
}

func (d AuthorizedMinterUpdatedData) EventType() EventType { return EventTypeAuthorizedMinterUpdated }

// TransferData represents data for ERC721 Transfer event
type TransferData struct {
	From    common.Address `json:"from"`
	To      common.Address `json:"to"`
	TokenID *big.Int       `json:"token_id"`
}

func (d TransferData) EventType() EventType { return EventTypeTransfer }

// ApprovalData represents data for ERC721 Approval event
type ApprovalData struct {
	Owner    common.Address `json:"owner"`
	Approved common.Address `json:"approved"`
	TokenID  *big.Int       `json:"token_id"`
}

func (d ApprovalData) EventType() EventType { return EventTypeApproval }

// ApprovalForAllData represents data for ERC721 ApprovalForAll event
type ApprovalForAllData struct {
	Owner    common.Address `json:"owner"`
	Operator common.Address `json:"operator"`
	Approved bool           `json:"approved"`
}

func (d ApprovalForAllData) EventType() EventType { return EventTypeApprovalForAll }

// Block represents a processed block
type Block struct {
	Number    uint64      `json:"number"`
	Hash      common.Hash `json:"hash"`
	Timestamp time.Time   `json:"timestamp"`
	EventCount int        `json:"event_count"`
}

// Clue represents NFT clue metadata
type Clue struct {
	ClueID       uint64 `json:"clue_id"`
	Contents     string `json:"contents"`
	SolutionHash string `json:"solution_hash"`
	PointValue   uint64 `json:"point_value"`
	SolveReward  uint64 `json:"solve_reward"`
}

// ClueOwnership represents ownership information for a clue
type ClueOwnership struct {
	OwnerAddress                     string `json:"owner_address"`
	ClueID                          uint64 `json:"clue_id"`
	OwnershipGrantedBlockNumber      uint64 `json:"ownership_granted_block_number"`
	OwnershipGrantedTransactionIndex uint   `json:"ownership_granted_transaction_index"`
	OwnershipGrantedEventIndex       uint   `json:"ownership_granted_event_index"`
	OwnershipGrantedEventType        string `json:"ownership_granted_event_type"`
}

// ClueWithOwner represents a clue with its ownership information
type ClueWithOwner struct {
	Clue
	Ownership *ClueOwnership `json:"ownership,omitempty"`
}
