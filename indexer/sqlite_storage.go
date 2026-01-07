package indexer

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteStorage implements the Storage interface using SQLite
type SQLiteStorage struct {
	db *sql.DB
}

// NewSQLiteStorage creates a new SQLite storage instance
func NewSQLiteStorage(dbPath string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	storage := &SQLiteStorage{db: db}
	if err := storage.initialize(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return storage, nil
}

// initialize creates the database schema
func (s *SQLiteStorage) initialize() error {
	schema := `
	CREATE TABLE IF NOT EXISTS events (
		clue_id INTEGER NOT NULL,
		block_number INTEGER NOT NULL,
		transaction_index INTEGER NOT NULL,
		event_index INTEGER NOT NULL,
		transaction_hash TEXT NOT NULL,
		initiated_by TEXT NOT NULL,
		event_type TEXT NOT NULL,
		metadata BLOB NOT NULL,
		timestamp INTEGER NOT NULL,
		block_hash TEXT NOT NULL,
		removed INTEGER NOT NULL DEFAULT 0,
		PRIMARY KEY (clue_id, block_number, transaction_index, event_index)
	);

	CREATE INDEX IF NOT EXISTS idx_event_type_timestamp ON events(event_type, timestamp);
	CREATE INDEX IF NOT EXISTS idx_initiated_by_timestamp ON events(initiated_by, timestamp);
	CREATE INDEX IF NOT EXISTS idx_block_number ON events(block_number);
	CREATE INDEX IF NOT EXISTS idx_timestamp ON events(timestamp);

	CREATE TABLE IF NOT EXISTS blocks (
		number INTEGER PRIMARY KEY,
		hash TEXT NOT NULL UNIQUE,
		timestamp INTEGER NOT NULL,
		event_count INTEGER NOT NULL DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS nft_state (
		clue_id INTEGER PRIMARY KEY,
		owner TEXT NOT NULL,
		is_solved INTEGER NOT NULL DEFAULT 0,
		sale_price TEXT,
		for_sale INTEGER NOT NULL DEFAULT 0
	);
	`

	_, err := s.db.Exec(schema)
	return err
}

// SaveEvent saves an event to the database
func (s *SQLiteStorage) SaveEvent(ctx context.Context, event *Event) error {
	query := `
		INSERT INTO events (
			clue_id, block_number, transaction_index, event_index,
			transaction_hash, initiated_by, event_type, metadata,
			timestamp, block_hash, removed
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := s.db.ExecContext(ctx, query,
		event.ClueID,
		event.BlockNumber,
		event.TransactionIndex,
		event.EventIndex,
		event.TransactionHash,
		event.InitiatedBy,
		event.EventType,
		event.Metadata,
		event.Timestamp,
		event.BlockHash.Hex(),
		boolToInt(event.Removed),
	)

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return ErrAlreadyExists
		}
		return fmt.Errorf("failed to save event: %w", err)
	}

	// Update NFT state based on event type
	if event.ClueID > 0 {
		return s.updateNFTState(ctx, event)
	}

	return nil
}

// updateNFTState updates the NFT state table based on an event
func (s *SQLiteStorage) updateNFTState(ctx context.Context, event *Event) error {
	switch event.EventType {
	case string(EventTypeClueMinted), string(EventTypeTransfer):
		// Parse metadata to get owner
		owner := event.InitiatedBy // Default to initiator
		// For Transfer events, we'd need to parse metadata for the "to" address
		// For now, we'll handle this in the indexer when creating events

		query := `
			INSERT INTO nft_state (clue_id, owner, is_solved, for_sale)
			VALUES (?, ?, 0, 0)
			ON CONFLICT(clue_id) DO UPDATE SET owner = ?
		`
		_, err := s.db.ExecContext(ctx, query, event.ClueID, owner, owner)
		return err

	case string(EventTypeClueSolved):
		query := `
			UPDATE nft_state
			SET is_solved = 1, for_sale = 0, sale_price = NULL
			WHERE clue_id = ?
		`
		_, err := s.db.ExecContext(ctx, query, event.ClueID)
		return err

	case string(EventTypeSalePriceSet):
		// Parse metadata to get price
		query := `
			INSERT INTO nft_state (clue_id, owner, is_solved, for_sale, sale_price)
			VALUES (?, ?, 0, 1, ?)
			ON CONFLICT(clue_id) DO UPDATE SET for_sale = 1, sale_price = ?
		`
		// We'll need to extract price from metadata - for now use empty string
		price := "" // TODO: parse from metadata
		_, err := s.db.ExecContext(ctx, query, event.ClueID, event.InitiatedBy, price, price)
		return err

	case string(EventTypeSalePriceRemoved):
		query := `
			UPDATE nft_state
			SET for_sale = 0, sale_price = NULL
			WHERE clue_id = ?
		`
		_, err := s.db.ExecContext(ctx, query, event.ClueID)
		return err
	}

	return nil
}

// GetEventByUniqueKey retrieves an event by its unique key
func (s *SQLiteStorage) GetEventByUniqueKey(ctx context.Context, clueID uint64, blockNumber uint64, txIndex uint, eventIndex uint) (*Event, error) {
	query := `
		SELECT clue_id, block_number, transaction_index, event_index,
		       transaction_hash, initiated_by, event_type, metadata,
		       timestamp, block_hash, removed
		FROM events
		WHERE clue_id = ? AND block_number = ? AND transaction_index = ? AND event_index = ?
	`

	event := &Event{}
	var blockHash string
	var removed int

	err := s.db.QueryRowContext(ctx, query, clueID, blockNumber, txIndex, eventIndex).Scan(
		&event.ClueID,
		&event.BlockNumber,
		&event.TransactionIndex,
		&event.EventIndex,
		&event.TransactionHash,
		&event.InitiatedBy,
		&event.EventType,
		&event.Metadata,
		&event.Timestamp,
		&blockHash,
		&removed,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get event: %w", err)
	}

	event.BlockHash = common.HexToHash(blockHash)
	event.Removed = intToBool(removed)

	return event, nil
}

// GetEventsByClueID retrieves all events for a specific clue
func (s *SQLiteStorage) GetEventsByClueID(ctx context.Context, clueID uint64, opts *QueryOptions) ([]*Event, error) {
	query := `
		SELECT clue_id, block_number, transaction_index, event_index,
		       transaction_hash, initiated_by, event_type, metadata,
		       timestamp, block_hash, removed
		FROM events
		WHERE clue_id = ?
	`
	args := []interface{}{clueID}

	query, args = s.applyQueryOptions(query, args, opts)

	return s.queryEvents(ctx, query, args...)
}

// GetEventsByBlock retrieves all events in a specific block
func (s *SQLiteStorage) GetEventsByBlock(ctx context.Context, blockNumber uint64) ([]*Event, error) {
	query := `
		SELECT clue_id, block_number, transaction_index, event_index,
		       transaction_hash, initiated_by, event_type, metadata,
		       timestamp, block_hash, removed
		FROM events
		WHERE block_number = ?
		ORDER BY transaction_index, event_index
	`

	return s.queryEvents(ctx, query, blockNumber)
}

// GetEventsByType retrieves events by type with query options
func (s *SQLiteStorage) GetEventsByType(ctx context.Context, eventType string, opts *QueryOptions) ([]*Event, error) {
	query := `
		SELECT clue_id, block_number, transaction_index, event_index,
		       transaction_hash, initiated_by, event_type, metadata,
		       timestamp, block_hash, removed
		FROM events
		WHERE event_type = ?
	`
	args := []interface{}{eventType}

	query, args = s.applyQueryOptions(query, args, opts)

	return s.queryEvents(ctx, query, args...)
}

// GetEventsByInitiator retrieves events by initiator with query options
func (s *SQLiteStorage) GetEventsByInitiator(ctx context.Context, initiator string, opts *QueryOptions) ([]*Event, error) {
	query := `
		SELECT clue_id, block_number, transaction_index, event_index,
		       transaction_hash, initiated_by, event_type, metadata,
		       timestamp, block_hash, removed
		FROM events
		WHERE initiated_by = ?
	`
	args := []interface{}{initiator}

	query, args = s.applyQueryOptions(query, args, opts)

	return s.queryEvents(ctx, query, args...)
}

// GetAllEvents retrieves all events with query options
func (s *SQLiteStorage) GetAllEvents(ctx context.Context, opts *QueryOptions) ([]*Event, error) {
	query := `
		SELECT clue_id, block_number, transaction_index, event_index,
		       transaction_hash, initiated_by, event_type, metadata,
		       timestamp, block_hash, removed
		FROM events
		WHERE 1=1
	`
	args := []interface{}{}

	query, args = s.applyQueryOptions(query, args, opts)

	return s.queryEvents(ctx, query, args...)
}

// applyQueryOptions applies query options to a SQL query
func (s *SQLiteStorage) applyQueryOptions(query string, args []interface{}, opts *QueryOptions) (string, []interface{}) {
	if opts == nil {
		opts = &QueryOptions{SortOrder: SortOrderAsc}
	}

	// Apply time range filter
	if opts.StartTime > 0 {
		query += " AND timestamp >= ?"
		args = append(args, opts.StartTime)
	}
	if opts.EndTime > 0 {
		query += " AND timestamp <= ?"
		args = append(args, opts.EndTime)
	}

	// Apply ordering
	if opts.SortOrder == SortOrderDesc {
		query += " ORDER BY timestamp DESC"
	} else {
		query += " ORDER BY timestamp ASC"
	}

	// Apply pagination
	if opts.Limit > 0 {
		query += " LIMIT ?"
		args = append(args, opts.Limit)
	}
	if opts.Offset > 0 {
		query += " OFFSET ?"
		args = append(args, opts.Offset)
	}

	return query, args
}

// queryEvents executes a query and returns events
func (s *SQLiteStorage) queryEvents(ctx context.Context, query string, args ...interface{}) ([]*Event, error) {
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query events: %w", err)
	}
	defer rows.Close()

	var events []*Event
	for rows.Next() {
		event := &Event{}
		var blockHash string
		var removed int

		err := rows.Scan(
			&event.ClueID,
			&event.BlockNumber,
			&event.TransactionIndex,
			&event.EventIndex,
			&event.TransactionHash,
			&event.InitiatedBy,
			&event.EventType,
			&event.Metadata,
			&event.Timestamp,
			&blockHash,
			&removed,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan event: %w", err)
		}

		event.BlockHash = common.HexToHash(blockHash)
		event.Removed = intToBool(removed)

		events = append(events, event)
	}

	return events, rows.Err()
}

// MarkEventsAsRemoved marks events in a block as removed
func (s *SQLiteStorage) MarkEventsAsRemoved(ctx context.Context, blockNumber uint64) error {
	query := "UPDATE events SET removed = 1 WHERE block_number = ?"
	_, err := s.db.ExecContext(ctx, query, blockNumber)
	return err
}

// SaveBlock saves a block to the database
func (s *SQLiteStorage) SaveBlock(ctx context.Context, block *Block) error {
	query := `
		INSERT INTO blocks (number, hash, timestamp, event_count)
		VALUES (?, ?, ?, ?)
		ON CONFLICT(number) DO UPDATE SET
			hash = excluded.hash,
			timestamp = excluded.timestamp,
			event_count = excluded.event_count
	`

	_, err := s.db.ExecContext(ctx, query,
		block.Number,
		block.Hash.Hex(),
		block.Timestamp.Unix(),
		block.EventCount,
	)

	return err
}

// GetBlock retrieves a block by number
func (s *SQLiteStorage) GetBlock(ctx context.Context, blockNumber uint64) (*Block, error) {
	query := "SELECT number, hash, timestamp, event_count FROM blocks WHERE number = ?"

	block := &Block{}
	var hashStr string
	var timestamp int64

	err := s.db.QueryRowContext(ctx, query, blockNumber).Scan(
		&block.Number,
		&hashStr,
		&timestamp,
		&block.EventCount,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get block: %w", err)
	}

	block.Hash = common.HexToHash(hashStr)
	block.Timestamp = timeFromUnix(timestamp)

	return block, nil
}

// GetBlockByHash retrieves a block by hash
func (s *SQLiteStorage) GetBlockByHash(ctx context.Context, hash [32]byte) (*Block, error) {
	query := "SELECT number, hash, timestamp, event_count FROM blocks WHERE hash = ?"

	block := &Block{}
	var hashStr string
	var timestamp int64

	// Convert hash to hex string for query
	hashHex := common.BytesToHash(hash[:]).Hex()

	err := s.db.QueryRowContext(ctx, query, hashHex).Scan(
		&block.Number,
		&hashStr,
		&timestamp,
		&block.EventCount,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get block: %w", err)
	}

	block.Hash = common.HexToHash(hashStr)
	block.Timestamp = timeFromUnix(timestamp)

	return block, nil
}

// GetLatestBlock retrieves the latest indexed block
func (s *SQLiteStorage) GetLatestBlock(ctx context.Context) (*Block, error) {
	query := "SELECT number, hash, timestamp, event_count FROM blocks ORDER BY number DESC LIMIT 1"

	block := &Block{}
	var hashStr string
	var timestamp int64

	err := s.db.QueryRowContext(ctx, query).Scan(
		&block.Number,
		&hashStr,
		&timestamp,
		&block.EventCount,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}

	block.Hash = common.HexToHash(hashStr)
	block.Timestamp = timeFromUnix(timestamp)

	return block, nil
}

// DeleteBlocksFrom deletes blocks and events from a certain block number
func (s *SQLiteStorage) DeleteBlocksFrom(ctx context.Context, blockNumber uint64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Delete events
	_, err = tx.ExecContext(ctx, "DELETE FROM events WHERE block_number >= ?", blockNumber)
	if err != nil {
		return fmt.Errorf("failed to delete events: %w", err)
	}

	// Delete blocks
	_, err = tx.ExecContext(ctx, "DELETE FROM blocks WHERE number >= ?", blockNumber)
	if err != nil {
		return fmt.Errorf("failed to delete blocks: %w", err)
	}

	// Rebuild NFT state from remaining events
	if err := s.rebuildNFTState(ctx, tx); err != nil {
		return fmt.Errorf("failed to rebuild NFT state: %w", err)
	}

	return tx.Commit()
}

// rebuildNFTState rebuilds the NFT state table from events
func (s *SQLiteStorage) rebuildNFTState(ctx context.Context, tx *sql.Tx) error {
	// Clear current state
	_, err := tx.ExecContext(ctx, "DELETE FROM nft_state")
	if err != nil {
		return err
	}

	// Rebuild from events (this is a simplified version)
	// In a real implementation, you'd process events in order to rebuild accurate state
	// For now, we'll just mark this as TODO
	// TODO: Implement proper NFT state rebuilding from events

	return nil
}

// GetNFTCurrentOwner retrieves the current owner of an NFT
func (s *SQLiteStorage) GetNFTCurrentOwner(ctx context.Context, clueID uint64) (string, error) {
	query := "SELECT owner FROM nft_state WHERE clue_id = ?"

	var owner string
	err := s.db.QueryRowContext(ctx, query, clueID).Scan(&owner)

	if err == sql.ErrNoRows {
		return "", ErrNotFound
	}
	if err != nil {
		return "", fmt.Errorf("failed to get NFT owner: %w", err)
	}

	return owner, nil
}

// GetNFTsByOwner retrieves all NFTs owned by an address
func (s *SQLiteStorage) GetNFTsByOwner(ctx context.Context, owner string) ([]uint64, error) {
	query := "SELECT clue_id FROM nft_state WHERE owner = ? ORDER BY clue_id"

	rows, err := s.db.QueryContext(ctx, query, owner)
	if err != nil {
		return nil, fmt.Errorf("failed to query NFTs: %w", err)
	}
	defer rows.Close()

	var nfts []uint64
	for rows.Next() {
		var clueID uint64
		if err := rows.Scan(&clueID); err != nil {
			return nil, fmt.Errorf("failed to scan NFT: %w", err)
		}
		nfts = append(nfts, clueID)
	}

	return nfts, rows.Err()
}

// GetNFTsForSale retrieves NFTs currently for sale
func (s *SQLiteStorage) GetNFTsForSale(ctx context.Context, limit, offset int) ([]*NFTSaleInfo, error) {
	query := `
		SELECT clue_id, owner, sale_price, is_solved
		FROM nft_state
		WHERE for_sale = 1
		ORDER BY clue_id
		LIMIT ? OFFSET ?
	`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query NFTs for sale: %w", err)
	}
	defer rows.Close()

	var nfts []*NFTSaleInfo
	for rows.Next() {
		info := &NFTSaleInfo{}
		var isSolved int
		var price sql.NullString

		err := rows.Scan(&info.ClueID, &info.Owner, &price, &isSolved)
		if err != nil {
			return nil, fmt.Errorf("failed to scan NFT: %w", err)
		}

		if price.Valid {
			info.Price = price.String
		}
		info.IsSolved = intToBool(isSolved)

		nfts = append(nfts, info)
	}

	return nfts, rows.Err()
}

// GetSolvedNFTs retrieves NFTs that have been solved
func (s *SQLiteStorage) GetSolvedNFTs(ctx context.Context, limit, offset int) ([]uint64, error) {
	query := `
		SELECT clue_id FROM nft_state
		WHERE is_solved = 1
		ORDER BY clue_id
		LIMIT ? OFFSET ?
	`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query solved NFTs: %w", err)
	}
	defer rows.Close()

	var nfts []uint64
	for rows.Next() {
		var clueID uint64
		if err := rows.Scan(&clueID); err != nil {
			return nil, fmt.Errorf("failed to scan NFT: %w", err)
		}
		nfts = append(nfts, clueID)
	}

	return nfts, rows.Err()
}

// GetTotalEvents returns the total number of events
func (s *SQLiteStorage) GetTotalEvents(ctx context.Context) (int64, error) {
	query := "SELECT COUNT(*) FROM events"

	var count int64
	err := s.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count events: %w", err)
	}

	return count, nil
}

// GetTotalNFTs returns the total number of NFTs
func (s *SQLiteStorage) GetTotalNFTs(ctx context.Context) (int64, error) {
	query := "SELECT COUNT(*) FROM nft_state"

	var count int64
	err := s.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count NFTs: %w", err)
	}

	return count, nil
}

// Close closes the database connection
func (s *SQLiteStorage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

// Helper functions
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func intToBool(i int) bool {
	return i != 0
}

func timeFromUnix(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
