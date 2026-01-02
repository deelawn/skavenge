package indexer

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

// Storage defines the interface for storing and retrieving indexed data
type Storage interface {
	// Event operations
	SaveEvent(ctx context.Context, event *Event) error
	GetEvent(ctx context.Context, id string) (*Event, error)
	GetEventsByTokenID(ctx context.Context, tokenID *big.Int, limit, offset int) ([]*Event, error)
	GetEventsByBlock(ctx context.Context, blockNumber uint64) ([]*Event, error)
	GetEventsByType(ctx context.Context, eventType EventType, limit, offset int) ([]*Event, error)
	GetEventsByTimeRange(ctx context.Context, start, end int64, limit, offset int) ([]*Event, error)
	MarkEventsAsRemoved(ctx context.Context, blockNumber uint64) error

	// Block operations
	SaveBlock(ctx context.Context, block *Block) error
	GetBlock(ctx context.Context, blockNumber uint64) (*Block, error)
	GetBlockByHash(ctx context.Context, hash common.Hash) (*Block, error)
	GetLatestBlock(ctx context.Context) (*Block, error)
	DeleteBlocksFrom(ctx context.Context, blockNumber uint64) error

	// NFT state queries (aggregated from events)
	GetNFTCurrentOwner(ctx context.Context, tokenID *big.Int) (common.Address, error)
	GetNFTsByOwner(ctx context.Context, owner common.Address) ([]*big.Int, error)
	GetNFTsForSale(ctx context.Context, limit, offset int) ([]*NFTSaleInfo, error)
	GetSolvedNFTs(ctx context.Context, limit, offset int) ([]*big.Int, error)

	// Stats
	GetTotalEvents(ctx context.Context) (int64, error)
	GetTotalNFTs(ctx context.Context) (int64, error)
}

// NFTSaleInfo contains information about an NFT for sale
type NFTSaleInfo struct {
	TokenID *big.Int       `json:"tokenId"`
	Owner   common.Address `json:"owner"`
	Price   *big.Int       `json:"price"`
	IsSolved bool          `json:"isSolved"`
}
