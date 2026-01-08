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

	CREATE TABLE IF NOT EXISTS clues (
		clue_id INTEGER PRIMARY KEY,
		contents TEXT NOT NULL,
		solution_hash TEXT NOT NULL,
		point_value INTEGER NOT NULL,
		solve_reward INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS clue_owners (
		owner_address TEXT NOT NULL,
		clue_id INTEGER NOT NULL,
		ownership_granted_block_number INTEGER NOT NULL,
		ownership_granted_transaction_index INTEGER NOT NULL,
		ownership_granted_event_index INTEGER NOT NULL,
		ownership_granted_event_type TEXT NOT NULL,
		UNIQUE(clue_id)
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

	return nil
}

// UpdateClueOwnership updates the ownership of a clue with the new owner address
func (s *SQLiteStorage) UpdateClueOwnership(ctx context.Context, clueID uint64, newOwner string, blockNumber uint64, txIndex uint, eventIndex uint, eventType string) error {
	query := `
		INSERT OR REPLACE INTO clue_owners (
			owner_address,
			clue_id,
			ownership_granted_block_number,
			ownership_granted_transaction_index,
			ownership_granted_event_index,
			ownership_granted_event_type
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := s.db.ExecContext(ctx, query,
		newOwner,
		clueID,
		blockNumber,
		txIndex,
		eventIndex,
		eventType,
	)
	return err
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

	// Rebuild clue ownership from remaining events
	if err := s.rebuildClueOwners(ctx, tx); err != nil {
		return fmt.Errorf("failed to rebuild clue owners: %w", err)
	}

	return tx.Commit()
}

// rebuildClueOwners rebuilds the clue_owners table from events
func (s *SQLiteStorage) rebuildClueOwners(ctx context.Context, tx *sql.Tx) error {
	// Clear current ownership state
	_, err := tx.ExecContext(ctx, "DELETE FROM clue_owners")
	if err != nil {
		return err
	}

	// Rebuild ownership from remaining events
	// Find the most recent ownership event (ClueMinted or Transfer) for each clue
	query := `
		INSERT INTO clue_owners (
			owner_address,
			clue_id,
			ownership_granted_block_number,
			ownership_granted_transaction_index,
			ownership_granted_event_index,
			ownership_granted_event_type
		)
		SELECT
			initiated_by,
			clue_id,
			block_number,
			transaction_index,
			event_index,
			event_type
		FROM events
		WHERE event_type IN ('ClueMinted', 'Transfer')
		AND (clue_id, block_number, transaction_index, event_index) IN (
			SELECT clue_id, MAX(block_number), MAX(transaction_index), MAX(event_index)
			FROM events
			WHERE event_type IN ('ClueMinted', 'Transfer')
			GROUP BY clue_id
		)
	`
	_, err = tx.ExecContext(ctx, query)
	return err
}

// GetNFTCurrentOwner retrieves the current owner of a clue
func (s *SQLiteStorage) GetNFTCurrentOwner(ctx context.Context, clueID uint64) (string, error) {
	query := "SELECT owner_address FROM clue_owners WHERE clue_id = ?"

	var owner string
	err := s.db.QueryRowContext(ctx, query, clueID).Scan(&owner)

	if err == sql.ErrNoRows {
		return "", ErrNotFound
	}
	if err != nil {
		return "", fmt.Errorf("failed to get clue owner: %w", err)
	}

	return owner, nil
}

// GetNFTsByOwner retrieves all clues owned by an address
func (s *SQLiteStorage) GetNFTsByOwner(ctx context.Context, owner string) ([]uint64, error) {
	query := "SELECT clue_id FROM clue_owners WHERE owner_address = ? ORDER BY clue_id"

	rows, err := s.db.QueryContext(ctx, query, owner)
	if err != nil {
		return nil, fmt.Errorf("failed to query clues: %w", err)
	}
	defer rows.Close()

	var clues []uint64
	for rows.Next() {
		var clueID uint64
		if err := rows.Scan(&clueID); err != nil {
			return nil, fmt.Errorf("failed to scan clue: %w", err)
		}
		clues = append(clues, clueID)
	}

	return clues, rows.Err()
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

// GetTotalNFTs returns the total number of clues
func (s *SQLiteStorage) GetTotalNFTs(ctx context.Context) (int64, error) {
	query := "SELECT COUNT(*) FROM clues"

	var count int64
	err := s.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count clues: %w", err)
	}

	return count, nil
}

// GetClueOwnership retrieves ownership information for a specific clue
func (s *SQLiteStorage) GetClueOwnership(ctx context.Context, clueID uint64) (*ClueOwnership, error) {
	query := `
		SELECT owner_address, clue_id, ownership_granted_block_number,
		       ownership_granted_transaction_index, ownership_granted_event_index,
		       ownership_granted_event_type
		FROM clue_owners
		WHERE clue_id = ?
	`

	ownership := &ClueOwnership{}
	err := s.db.QueryRowContext(ctx, query, clueID).Scan(
		&ownership.OwnerAddress,
		&ownership.ClueID,
		&ownership.OwnershipGrantedBlockNumber,
		&ownership.OwnershipGrantedTransactionIndex,
		&ownership.OwnershipGrantedEventIndex,
		&ownership.OwnershipGrantedEventType,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get clue ownership: %w", err)
	}

	return ownership, nil
}

// GetAllClueOwnerships retrieves all clue ownership records
func (s *SQLiteStorage) GetAllClueOwnerships(ctx context.Context) ([]*ClueOwnership, error) {
	query := `
		SELECT owner_address, clue_id, ownership_granted_block_number,
		       ownership_granted_transaction_index, ownership_granted_event_index,
		       ownership_granted_event_type
		FROM clue_owners
		ORDER BY clue_id
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query clue ownerships: %w", err)
	}
	defer rows.Close()

	var ownerships []*ClueOwnership
	for rows.Next() {
		ownership := &ClueOwnership{}
		err := rows.Scan(
			&ownership.OwnerAddress,
			&ownership.ClueID,
			&ownership.OwnershipGrantedBlockNumber,
			&ownership.OwnershipGrantedTransactionIndex,
			&ownership.OwnershipGrantedEventIndex,
			&ownership.OwnershipGrantedEventType,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan clue ownership: %w", err)
		}
		ownerships = append(ownerships, ownership)
	}

	return ownerships, rows.Err()
}

// GetClueOwnershipsByOwner retrieves all clue ownership records for a specific owner
func (s *SQLiteStorage) GetClueOwnershipsByOwner(ctx context.Context, owner string) ([]*ClueOwnership, error) {
	query := `
		SELECT owner_address, clue_id, ownership_granted_block_number,
		       ownership_granted_transaction_index, ownership_granted_event_index,
		       ownership_granted_event_type
		FROM clue_owners
		WHERE owner_address = ?
		ORDER BY clue_id
	`

	rows, err := s.db.QueryContext(ctx, query, owner)
	if err != nil {
		return nil, fmt.Errorf("failed to query clue ownerships by owner: %w", err)
	}
	defer rows.Close()

	var ownerships []*ClueOwnership
	for rows.Next() {
		ownership := &ClueOwnership{}
		err := rows.Scan(
			&ownership.OwnerAddress,
			&ownership.ClueID,
			&ownership.OwnershipGrantedBlockNumber,
			&ownership.OwnershipGrantedTransactionIndex,
			&ownership.OwnershipGrantedEventIndex,
			&ownership.OwnershipGrantedEventType,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan clue ownership: %w", err)
		}
		ownerships = append(ownerships, ownership)
	}

	return ownerships, rows.Err()
}

// SaveClue saves a clue to the database
func (s *SQLiteStorage) SaveClue(ctx context.Context, clue *Clue, force bool) error {
	if force {
		// Use INSERT OR REPLACE to overwrite existing record
		query := `
			INSERT OR REPLACE INTO clues (clue_id, contents, solution_hash, point_value, solve_reward)
			VALUES (?, ?, ?, ?, ?)
		`
		_, err := s.db.ExecContext(ctx, query,
			clue.ClueID,
			clue.Contents,
			clue.SolutionHash,
			clue.PointValue,
			clue.SolveReward,
		)
		if err != nil {
			return fmt.Errorf("failed to save clue: %w", err)
		}
		return nil
	}

	// Regular insert - will fail on uniqueness constraint violation
	query := `
		INSERT INTO clues (clue_id, contents, solution_hash, point_value, solve_reward)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := s.db.ExecContext(ctx, query,
		clue.ClueID,
		clue.Contents,
		clue.SolutionHash,
		clue.PointValue,
		clue.SolveReward,
	)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return ErrAlreadyExists
		}
		return fmt.Errorf("failed to save clue: %w", err)
	}

	return nil
}

// GetClue retrieves a clue by ID
func (s *SQLiteStorage) GetClue(ctx context.Context, clueID uint64) (*Clue, error) {
	query := `
		SELECT clue_id, contents, solution_hash, point_value, solve_reward
		FROM clues
		WHERE clue_id = ?
	`

	clue := &Clue{}
	err := s.db.QueryRowContext(ctx, query, clueID).Scan(
		&clue.ClueID,
		&clue.Contents,
		&clue.SolutionHash,
		&clue.PointValue,
		&clue.SolveReward,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get clue: %w", err)
	}

	return clue, nil
}

// GetAllClues retrieves all clues
func (s *SQLiteStorage) GetAllClues(ctx context.Context) ([]*Clue, error) {
	query := `
		SELECT clue_id, contents, solution_hash, point_value, solve_reward
		FROM clues
		ORDER BY clue_id
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query clues: %w", err)
	}
	defer rows.Close()

	var clues []*Clue
	for rows.Next() {
		clue := &Clue{}
		err := rows.Scan(
			&clue.ClueID,
			&clue.Contents,
			&clue.SolutionHash,
			&clue.PointValue,
			&clue.SolveReward,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan clue: %w", err)
		}
		clues = append(clues, clue)
	}

	return clues, rows.Err()
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
