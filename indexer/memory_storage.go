package indexer

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// MemoryStorage is an in-memory implementation of the Storage interface
// Suitable for testing and small-scale deployments
type MemoryStorage struct {
	mu              sync.RWMutex
	events          map[string]*Event           // eventID -> Event
	eventsByBlock   map[uint64][]*Event         // blockNumber -> Events
	eventsByTokenID map[string][]*Event         // tokenID (hex) -> Events
	eventsByType    map[EventType][]*Event      // eventType -> Events
	blocks          map[uint64]*Block           // blockNumber -> Block
	blocksByHash    map[common.Hash]*Block      // blockHash -> Block
	nftOwners       map[string]common.Address   // tokenID (hex) -> current owner
	nftForSale      map[string]*big.Int         // tokenID (hex) -> sale price
	nftSolved       map[string]bool             // tokenID (hex) -> is solved
}

// NewMemoryStorage creates a new in-memory storage instance
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		events:          make(map[string]*Event),
		eventsByBlock:   make(map[uint64][]*Event),
		eventsByTokenID: make(map[string][]*Event),
		eventsByType:    make(map[EventType][]*Event),
		blocks:          make(map[uint64]*Block),
		blocksByHash:    make(map[common.Hash]*Block),
		nftOwners:       make(map[string]common.Address),
		nftForSale:      make(map[string]*big.Int),
		nftSolved:       make(map[string]bool),
	}
}

func (m *MemoryStorage) SaveEvent(ctx context.Context, event *Event) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.events[event.ID]; exists {
		return ErrAlreadyExists
	}

	m.events[event.ID] = event
	m.eventsByBlock[event.BlockNumber] = append(m.eventsByBlock[event.BlockNumber], event)
	m.eventsByType[event.Type] = append(m.eventsByType[event.Type], event)

	if event.TokenID != nil {
		tokenIDKey := event.TokenID.String()
		m.eventsByTokenID[tokenIDKey] = append(m.eventsByTokenID[tokenIDKey], event)
	}

	// Update NFT state based on event type
	m.updateNFTState(event)

	return nil
}

func (m *MemoryStorage) updateNFTState(event *Event) {
	if event.TokenID == nil {
		return
	}
	tokenIDKey := event.TokenID.String()

	switch data := event.Data.(type) {
	case ClueMintedData:
		m.nftOwners[tokenIDKey] = data.Minter
	case TransferData:
		if data.To != (common.Address{}) { // Not a burn
			m.nftOwners[tokenIDKey] = data.To
		}
	case ClueSolvedData:
		m.nftSolved[tokenIDKey] = true
		delete(m.nftForSale, tokenIDKey) // Remove from sale when solved
	case SalePriceSetData:
		m.nftForSale[tokenIDKey] = data.Price
	case SalePriceRemovedData:
		delete(m.nftForSale, tokenIDKey)
	}
}

func (m *MemoryStorage) GetEvent(ctx context.Context, id string) (*Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	event, exists := m.events[id]
	if !exists {
		return nil, ErrNotFound
	}
	return event, nil
}

func (m *MemoryStorage) GetEventsByTokenID(ctx context.Context, tokenID *big.Int, limit, offset int) ([]*Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	tokenIDKey := tokenID.String()
	events := m.eventsByTokenID[tokenIDKey]

	return m.paginateEvents(events, limit, offset), nil
}

func (m *MemoryStorage) GetEventsByBlock(ctx context.Context, blockNumber uint64) ([]*Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	events := m.eventsByBlock[blockNumber]
	return events, nil
}

func (m *MemoryStorage) GetEventsByType(ctx context.Context, eventType EventType, limit, offset int) ([]*Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	events := m.eventsByType[eventType]
	return m.paginateEvents(events, limit, offset), nil
}

func (m *MemoryStorage) GetEventsByTimeRange(ctx context.Context, start, end int64, limit, offset int) ([]*Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var filtered []*Event
	for _, event := range m.events {
		ts := event.Timestamp.Unix()
		if ts >= start && ts <= end {
			filtered = append(filtered, event)
		}
	}

	// Sort by timestamp
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Timestamp.Before(filtered[j].Timestamp)
	})

	return m.paginateEvents(filtered, limit, offset), nil
}

func (m *MemoryStorage) MarkEventsAsRemoved(ctx context.Context, blockNumber uint64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	events := m.eventsByBlock[blockNumber]
	for _, event := range events {
		event.Removed = true
	}
	return nil
}

func (m *MemoryStorage) SaveBlock(ctx context.Context, block *Block) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.blocks[block.Number] = block
	m.blocksByHash[block.Hash] = block
	return nil
}

func (m *MemoryStorage) GetBlock(ctx context.Context, blockNumber uint64) (*Block, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	block, exists := m.blocks[blockNumber]
	if !exists {
		return nil, ErrNotFound
	}
	return block, nil
}

func (m *MemoryStorage) GetBlockByHash(ctx context.Context, hash common.Hash) (*Block, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	block, exists := m.blocksByHash[hash]
	if !exists {
		return nil, ErrNotFound
	}
	return block, nil
}

func (m *MemoryStorage) GetLatestBlock(ctx context.Context) (*Block, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var latest *Block
	for _, block := range m.blocks {
		if latest == nil || block.Number > latest.Number {
			latest = block
		}
	}

	if latest == nil {
		return nil, ErrNotFound
	}
	return latest, nil
}

func (m *MemoryStorage) DeleteBlocksFrom(ctx context.Context, blockNumber uint64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var blocksToDelete []uint64
	for num := range m.blocks {
		if num >= blockNumber {
			blocksToDelete = append(blocksToDelete, num)
		}
	}

	for _, num := range blocksToDelete {
		block := m.blocks[num]
		delete(m.blocksByHash, block.Hash)
		delete(m.blocks, num)

		// Delete events from this block
		events := m.eventsByBlock[num]
		for _, event := range events {
			delete(m.events, event.ID)

			// Remove from type index
			typeEvents := m.eventsByType[event.Type]
			for i, e := range typeEvents {
				if e.ID == event.ID {
					m.eventsByType[event.Type] = append(typeEvents[:i], typeEvents[i+1:]...)
					break
				}
			}

			// Remove from token index
			if event.TokenID != nil {
				tokenIDKey := event.TokenID.String()
				tokenEvents := m.eventsByTokenID[tokenIDKey]
				for i, e := range tokenEvents {
					if e.ID == event.ID {
						m.eventsByTokenID[tokenIDKey] = append(tokenEvents[:i], tokenEvents[i+1:]...)
						break
					}
				}
			}
		}
		delete(m.eventsByBlock, num)
	}

	// Rebuild NFT state from remaining events
	m.rebuildNFTState()

	return nil
}

func (m *MemoryStorage) rebuildNFTState() {
	m.nftOwners = make(map[string]common.Address)
	m.nftForSale = make(map[string]*big.Int)
	m.nftSolved = make(map[string]bool)

	// Collect all events and sort by block number and log index
	var allEvents []*Event
	for _, event := range m.events {
		if !event.Removed {
			allEvents = append(allEvents, event)
		}
	}
	sort.Slice(allEvents, func(i, j int) bool {
		if allEvents[i].BlockNumber != allEvents[j].BlockNumber {
			return allEvents[i].BlockNumber < allEvents[j].BlockNumber
		}
		return allEvents[i].LogIndex < allEvents[j].LogIndex
	})

	// Replay events to rebuild state
	for _, event := range allEvents {
		m.updateNFTState(event)
	}
}

func (m *MemoryStorage) GetNFTCurrentOwner(ctx context.Context, tokenID *big.Int) (common.Address, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	tokenIDKey := tokenID.String()
	owner, exists := m.nftOwners[tokenIDKey]
	if !exists {
		return common.Address{}, ErrNotFound
	}
	return owner, nil
}

func (m *MemoryStorage) GetNFTsByOwner(ctx context.Context, owner common.Address) ([]*big.Int, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var nfts []*big.Int
	for tokenIDStr, nftOwner := range m.nftOwners {
		if nftOwner == owner {
			tokenID := new(big.Int)
			tokenID.SetString(tokenIDStr, 10)
			nfts = append(nfts, tokenID)
		}
	}
	return nfts, nil
}

func (m *MemoryStorage) GetNFTsForSale(ctx context.Context, limit, offset int) ([]*NFTSaleInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var sales []*NFTSaleInfo
	for tokenIDStr, price := range m.nftForSale {
		tokenID := new(big.Int)
		tokenID.SetString(tokenIDStr, 10)

		sales = append(sales, &NFTSaleInfo{
			TokenID:  tokenID,
			Owner:    m.nftOwners[tokenIDStr],
			Price:    price,
			IsSolved: m.nftSolved[tokenIDStr],
		})
	}

	// Sort by token ID for consistent ordering
	sort.Slice(sales, func(i, j int) bool {
		return sales[i].TokenID.Cmp(sales[j].TokenID) < 0
	})

	return m.paginateSales(sales, limit, offset), nil
}

func (m *MemoryStorage) GetSolvedNFTs(ctx context.Context, limit, offset int) ([]*big.Int, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var solved []*big.Int
	for tokenIDStr, isSolved := range m.nftSolved {
		if isSolved {
			tokenID := new(big.Int)
			tokenID.SetString(tokenIDStr, 10)
			solved = append(solved, tokenID)
		}
	}

	// Sort for consistent ordering
	sort.Slice(solved, func(i, j int) bool {
		return solved[i].Cmp(solved[j]) < 0
	})

	return m.paginateBigInts(solved, limit, offset), nil
}

func (m *MemoryStorage) GetTotalEvents(ctx context.Context) (int64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return int64(len(m.events)), nil
}

func (m *MemoryStorage) GetTotalNFTs(ctx context.Context) (int64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return int64(len(m.nftOwners)), nil
}

func (m *MemoryStorage) paginateEvents(events []*Event, limit, offset int) []*Event {
	if offset >= len(events) {
		return []*Event{}
	}
	end := offset + limit
	if end > len(events) {
		end = len(events)
	}
	return events[offset:end]
}

func (m *MemoryStorage) paginateSales(sales []*NFTSaleInfo, limit, offset int) []*NFTSaleInfo {
	if offset >= len(sales) {
		return []*NFTSaleInfo{}
	}
	end := offset + limit
	if end > len(sales) {
		end = len(sales)
	}
	return sales[offset:end]
}

func (m *MemoryStorage) paginateBigInts(items []*big.Int, limit, offset int) []*big.Int {
	if offset >= len(items) {
		return []*big.Int{}
	}
	end := offset + limit
	if end > len(items) {
		end = len(items)
	}
	return items[offset:end]
}

var _ Storage = (*MemoryStorage)(nil) // Ensure MemoryStorage implements Storage
