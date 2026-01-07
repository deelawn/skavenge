package indexer

import (
	"context"
	"errors"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

// SortOrder represents the sort order for queries
type SortOrder string

const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

// QueryOptions contains options for querying events
type QueryOptions struct {
	Limit     int       // Maximum number of results (0 = no limit)
	Offset    int       // Number of results to skip
	SortOrder SortOrder // Sort order by timestamp
	StartTime int64     // Start of time range (Unix timestamp, 0 = no limit)
	EndTime   int64     // End of time range (Unix timestamp, 0 = no limit)
}

// Storage defines the interface for storing and retrieving indexed data
type Storage interface {
	// Event operations
	SaveEvent(ctx context.Context, event *Event) error

	// GetEventByUniqueKey retrieves an event by its unique index (ClueID, BlockNumber, TransactionIndex, EventIndex)
	GetEventByUniqueKey(ctx context.Context, clueID uint64, blockNumber uint64, txIndex uint, eventIndex uint) (*Event, error)

	// GetEventsByClueID retrieves all events for a specific clue/NFT with optional query options
	GetEventsByClueID(ctx context.Context, clueID uint64, opts *QueryOptions) ([]*Event, error)

	// GetEventsByBlock retrieves all events in a specific block
	GetEventsByBlock(ctx context.Context, blockNumber uint64) ([]*Event, error)

	// GetEventsByType retrieves events by type with query options (supports time range and ordering)
	GetEventsByType(ctx context.Context, eventType string, opts *QueryOptions) ([]*Event, error)

	// GetEventsByInitiator retrieves events by initiator address with query options (supports time range and ordering)
	GetEventsByInitiator(ctx context.Context, initiator string, opts *QueryOptions) ([]*Event, error)

	// GetAllEvents retrieves all events with query options
	GetAllEvents(ctx context.Context, opts *QueryOptions) ([]*Event, error)

	// MarkEventsAsRemoved marks events in a block as removed (for reorg handling)
	MarkEventsAsRemoved(ctx context.Context, blockNumber uint64) error

	// Block operations
	SaveBlock(ctx context.Context, block *Block) error
	GetBlock(ctx context.Context, blockNumber uint64) (*Block, error)
	GetBlockByHash(ctx context.Context, hash [32]byte) (*Block, error)
	GetLatestBlock(ctx context.Context) (*Block, error)
	DeleteBlocksFrom(ctx context.Context, blockNumber uint64) error

	// NFT state queries (aggregated from events)
	GetNFTCurrentOwner(ctx context.Context, clueID uint64) (string, error)
	GetNFTsByOwner(ctx context.Context, owner string) ([]uint64, error)
	GetNFTsForSale(ctx context.Context, limit, offset int) ([]*NFTSaleInfo, error)
	GetSolvedNFTs(ctx context.Context, limit, offset int) ([]uint64, error)

	// Stats
	GetTotalEvents(ctx context.Context) (int64, error)
	GetTotalNFTs(ctx context.Context) (int64, error)

	// Close closes the storage and releases resources
	Close() error
}

// NFTSaleInfo contains information about an NFT for sale
type NFTSaleInfo struct {
	ClueID   uint64 `json:"clueId"`
	Owner    string `json:"owner"`
	Price    string `json:"price"` // Price in wei as string
	IsSolved bool   `json:"isSolved"`
}
