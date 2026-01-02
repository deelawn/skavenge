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
	EventTypeClueAttempted           EventType = "ClueAttempted"
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
	ID              string    `json:"id"`              // Unique identifier (txHash + logIndex)
	Type            EventType `json:"type"`            // Event type
	BlockNumber     uint64    `json:"blockNumber"`     // Block number
	BlockHash       common.Hash `json:"blockHash"`     // Block hash
	TransactionHash common.Hash `json:"transactionHash"` // Transaction hash
	LogIndex        uint       `json:"logIndex"`       // Log index in transaction
	Timestamp       time.Time `json:"timestamp"`      // Block timestamp
	TokenID         *big.Int  `json:"tokenId,omitempty"` // Associated NFT token ID (if applicable)
	Data            EventData `json:"data"`           // Event-specific data
	Removed         bool      `json:"removed"`        // True if log was removed due to reorg
}

// EventData contains event-specific data
type EventData interface {
	EventType() EventType
}

// ClueMintedData represents data for ClueMinted event
type ClueMintedData struct {
	TokenID *big.Int       `json:"tokenId"`
	Minter  common.Address `json:"minter"`
}

func (d ClueMintedData) EventType() EventType { return EventTypeClueMinted }

// ClueAttemptedData represents data for ClueAttempted event
type ClueAttemptedData struct {
	TokenID            *big.Int `json:"tokenId"`
	RemainingAttempts  *big.Int `json:"remainingAttempts"`
}

func (d ClueAttemptedData) EventType() EventType { return EventTypeClueAttempted }

// ClueSolvedData represents data for ClueSolved event
type ClueSolvedData struct {
	TokenID  *big.Int `json:"tokenId"`
	Solution string   `json:"solution"`
}

func (d ClueSolvedData) EventType() EventType { return EventTypeClueSolved }

// SalePriceSetData represents data for SalePriceSet event
type SalePriceSetData struct {
	TokenID *big.Int `json:"tokenId"`
	Price   *big.Int `json:"price"`
}

func (d SalePriceSetData) EventType() EventType { return EventTypeSalePriceSet }

// SalePriceRemovedData represents data for SalePriceRemoved event
type SalePriceRemovedData struct {
	TokenID *big.Int `json:"tokenId"`
}

func (d SalePriceRemovedData) EventType() EventType { return EventTypeSalePriceRemoved }

// TransferInitiatedData represents data for TransferInitiated event
type TransferInitiatedData struct {
	TransferID [32]byte       `json:"transferId"`
	Buyer      common.Address `json:"buyer"`
	TokenID    *big.Int       `json:"tokenId"`
}

func (d TransferInitiatedData) EventType() EventType { return EventTypeTransferInitiated }

// ProofProvidedData represents data for ProofProvided event
type ProofProvidedData struct {
	TransferID  [32]byte `json:"transferId"`
	Proof       []byte   `json:"proof"`
	NewClueHash [32]byte `json:"newClueHash"`
	RValueHash  [32]byte `json:"rValueHash"`
}

func (d ProofProvidedData) EventType() EventType { return EventTypeProofProvided }

// ProofVerifiedData represents data for ProofVerified event
type ProofVerifiedData struct {
	TransferID [32]byte `json:"transferId"`
}

func (d ProofVerifiedData) EventType() EventType { return EventTypeProofVerified }

// TransferCompletedData represents data for TransferCompleted event
type TransferCompletedData struct {
	TransferID [32]byte `json:"transferId"`
	RValue     *big.Int `json:"rValue"`
}

func (d TransferCompletedData) EventType() EventType { return EventTypeTransferCompleted }

// TransferCancelledData represents data for TransferCancelled event
type TransferCancelledData struct {
	TransferID [32]byte `json:"transferId"`
}

func (d TransferCancelledData) EventType() EventType { return EventTypeTransferCancelled }

// AuthorizedMinterUpdatedData represents data for AuthorizedMinterUpdated event
type AuthorizedMinterUpdatedData struct {
	OldMinter common.Address `json:"oldMinter"`
	NewMinter common.Address `json:"newMinter"`
}

func (d AuthorizedMinterUpdatedData) EventType() EventType { return EventTypeAuthorizedMinterUpdated }

// TransferData represents data for ERC721 Transfer event
type TransferData struct {
	From    common.Address `json:"from"`
	To      common.Address `json:"to"`
	TokenID *big.Int       `json:"tokenId"`
}

func (d TransferData) EventType() EventType { return EventTypeTransfer }

// ApprovalData represents data for ERC721 Approval event
type ApprovalData struct {
	Owner    common.Address `json:"owner"`
	Approved common.Address `json:"approved"`
	TokenID  *big.Int       `json:"tokenId"`
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
	EventCount int        `json:"eventCount"`
}
