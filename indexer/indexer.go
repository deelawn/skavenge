package indexer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Indexer is the main blockchain indexer
type Indexer struct {
	config  *Config
	client  *ethclient.Client
	parser  *EventParser
	storage Storage
	logger  *log.Logger
}

// NewIndexer creates a new indexer instance
func NewIndexer(config *Config, storage Storage, logger *log.Logger) (*Indexer, error) {
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC: %w", err)
	}

	parser, err := NewEventParser()
	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %w", err)
	}

	if logger == nil {
		logger = log.Default()
	}

	return &Indexer{
		config:  config,
		client:  client,
		parser:  parser,
		storage: storage,
		logger:  logger,
	}, nil
}

// Start begins the indexing process
func (idx *Indexer) Start(ctx context.Context) error {
	idx.logger.Printf("Starting indexer for contract %s", idx.config.ContractAddress.Hex())

	// Determine starting block
	startBlock, err := idx.determineStartBlock(ctx)
	if err != nil {
		return fmt.Errorf("failed to determine start block: %w", err)
	}

	idx.logger.Printf("Starting from block %d", startBlock)

	// Start the main indexing loop
	return idx.indexLoop(ctx, startBlock)
}

// determineStartBlock determines which block to start indexing from
func (idx *Indexer) determineStartBlock(ctx context.Context) (uint64, error) {
	// If config specifies a start block, use it
	if idx.config.StartBlock > 0 {
		return idx.config.StartBlock, nil
	}

	// Otherwise, check if we have indexed data
	latestIndexed, err := idx.storage.GetLatestBlock(ctx)
	if err == nil && latestIndexed != nil {
		// Resume from last indexed block
		return latestIndexed.Number + 1, nil
	}

	// Start from current block
	header, err := idx.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to get latest block: %w", err)
	}

	return header.Number.Uint64(), nil
}

// indexLoop is the main indexing loop
func (idx *Indexer) indexLoop(ctx context.Context, startBlock uint64) error {
	currentBlock := startBlock
	ticker := time.NewTicker(idx.config.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			idx.logger.Println("Indexer stopped")
			return ctx.Err()

		case <-ticker.C:
			// Get latest block number
			latestBlock, err := idx.getLatestBlockWithRetry(ctx)
			if err != nil {
				idx.logger.Printf("Error getting latest block: %v", err)
				continue
			}

			// Calculate the safe block (accounting for confirmations)
			safeBlock := latestBlock
			if safeBlock > idx.config.ConfirmationBlocks {
				safeBlock -= idx.config.ConfirmationBlocks
			} else {
				safeBlock = 0
			}

			// Check for reorgs
			if currentBlock > 0 {
				reorgDepth, err := idx.detectReorg(ctx, currentBlock)
				if err != nil {
					idx.logger.Printf("Error detecting reorg: %v", err)
					continue
				}
				if reorgDepth > 0 {
					idx.logger.Printf("Detected reorg at depth %d, rolling back from block %d", reorgDepth, currentBlock)
					if err := idx.handleReorg(ctx, currentBlock-reorgDepth); err != nil {
						idx.logger.Printf("Error handling reorg: %v", err)
						continue
					}
					currentBlock = currentBlock - reorgDepth
				}
			}

			// Index new blocks
			if currentBlock <= safeBlock {
				endBlock := currentBlock + idx.config.BatchSize - 1
				if endBlock > safeBlock {
					endBlock = safeBlock
				}

				idx.logger.Printf("Indexing blocks %d to %d", currentBlock, endBlock)

				if err := idx.indexBlockRange(ctx, currentBlock, endBlock); err != nil {
					idx.logger.Printf("Error indexing blocks: %v", err)
					continue
				}

				currentBlock = endBlock + 1
			}
		}
	}
}

// detectReorg checks if a chain reorganization has occurred
func (idx *Indexer) detectReorg(ctx context.Context, fromBlock uint64) (uint64, error) {
	// Check up to confirmation depth
	checkDepth := idx.config.ConfirmationBlocks
	if fromBlock < checkDepth {
		checkDepth = fromBlock
	}

	for i := uint64(0); i < checkDepth; i++ {
		blockNum := fromBlock - i - 1
		if blockNum == 0 {
			break
		}

		// Get stored block
		storedBlock, err := idx.storage.GetBlock(ctx, blockNum)
		if err != nil {
			if err == ErrNotFound {
				// Block not in storage, no reorg
				return 0, nil
			}
			return 0, err
		}

		// Get current block from chain
		chainBlock, err := idx.getBlockWithRetry(ctx, blockNum)
		if err != nil {
			return 0, err
		}

		// Compare hashes
		if storedBlock.Hash != chainBlock.Hash() {
			// Reorg detected
			return i + 1, nil
		}
	}

	return 0, nil
}

// handleReorg handles a chain reorganization by rolling back indexed data
func (idx *Indexer) handleReorg(ctx context.Context, fromBlock uint64) error {
	idx.logger.Printf("Handling reorg: deleting blocks from %d", fromBlock)

	// Delete blocks and events from the reorg point
	if err := idx.storage.DeleteBlocksFrom(ctx, fromBlock); err != nil {
		return fmt.Errorf("failed to delete blocks: %w", err)
	}

	return nil
}

// indexBlockRange indexes a range of blocks
func (idx *Indexer) indexBlockRange(ctx context.Context, fromBlock, toBlock uint64) error {
	for blockNum := fromBlock; blockNum <= toBlock; blockNum++ {
		if err := idx.indexBlock(ctx, blockNum); err != nil {
			return fmt.Errorf("failed to index block %d: %w", blockNum, err)
		}
	}
	return nil
}

// indexBlock indexes a single block
func (idx *Indexer) indexBlock(ctx context.Context, blockNumber uint64) error {
	// Get block
	block, err := idx.getBlockWithRetry(ctx, blockNumber)
	if err != nil {
		return err
	}

	// Build transaction index map
	txIndexMap := make(map[common.Hash]uint)
	txSenderMap := make(map[common.Hash]common.Address)
	for i, tx := range block.Transactions() {
		txIndexMap[tx.Hash()] = uint(i)
		// Get transaction sender
		sender, err := idx.client.TransactionSender(ctx, tx, block.Hash(), uint(i))
		if err != nil {
			idx.logger.Printf("Warning: failed to get sender for tx %s: %v", tx.Hash().Hex(), err)
			continue
		}
		txSenderMap[tx.Hash()] = sender
	}

	// Query logs for this block
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(blockNumber),
		ToBlock:   new(big.Int).SetUint64(blockNumber),
		Addresses: []common.Address{idx.config.ContractAddress},
	}

	logs, err := idx.getLogsWithRetry(ctx, query)
	if err != nil {
		return err
	}

	// Get block timestamp
	timestamp := time.Unix(int64(block.Time()), 0)

	// Parse and save events
	var events []*Event
	for _, vLog := range logs {
		// Get transaction info
		txIndex, ok := txIndexMap[vLog.TxHash]
		if !ok {
			idx.logger.Printf("Warning: transaction index not found for log %s-%d", vLog.TxHash.Hex(), vLog.Index)
			continue
		}

		txSender, ok := txSenderMap[vLog.TxHash]
		if !ok {
			idx.logger.Printf("Warning: transaction sender not found for log %s-%d", vLog.TxHash.Hex(), vLog.Index)
			continue
		}

		event, err := idx.parser.ParseLog(vLog, txSender, txIndex, timestamp)
		if err != nil {
			idx.logger.Printf("Warning: failed to parse log %s-%d: %v", vLog.TxHash.Hex(), vLog.Index, err)
			continue
		}

		idx.logger.Printf("Saving event: %+v\n", event)

		if err := idx.storage.SaveEvent(ctx, event); err != nil {
			// If the event already exists, log it and continue instead of failing
			if errors.Is(err, ErrAlreadyExists) {
				idx.logger.Printf("Warning: event already exists (clueID=%d, block=%d, txIndex=%d, eventIndex=%d), skipping",
					event.ClueID, event.BlockNumber, event.TransactionIndex, event.EventIndex)
			} else {
				return fmt.Errorf("failed to save event: %w", err)
			}
		}

		events = append(events, event)
	}

	// Save block
	blockInfo := &Block{
		Number:     blockNumber,
		Hash:       block.Hash(),
		Timestamp:  timestamp,
		EventCount: len(events),
	}

	if err := idx.storage.SaveBlock(ctx, blockInfo); err != nil {
		return fmt.Errorf("failed to save block: %w", err)
	}

	if len(events) > 0 {
		idx.logger.Printf("Block %d: indexed %d events", blockNumber, len(events))
	}

	return nil
}

// getLatestBlockWithRetry gets the latest block number with retry logic
func (idx *Indexer) getLatestBlockWithRetry(ctx context.Context) (uint64, error) {
	var lastErr error
	for i := 0; i < idx.config.MaxRetries; i++ {
		header, err := idx.client.HeaderByNumber(ctx, nil)
		if err == nil {
			return header.Number.Uint64(), nil
		}
		lastErr = err
		time.Sleep(idx.config.RetryDelay)
	}
	return 0, fmt.Errorf("failed after %d retries: %w", idx.config.MaxRetries, lastErr)
}

// getBlockWithRetry gets a block with retry logic
func (idx *Indexer) getBlockWithRetry(ctx context.Context, blockNumber uint64) (*types.Block, error) {
	var lastErr error
	for i := 0; i < idx.config.MaxRetries; i++ {
		block, err := idx.client.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
		if err == nil {
			return block, nil
		}
		lastErr = err
		time.Sleep(idx.config.RetryDelay)
	}
	return nil, fmt.Errorf("failed after %d retries: %w", idx.config.MaxRetries, lastErr)
}

// getLogsWithRetry gets logs with retry logic
func (idx *Indexer) getLogsWithRetry(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	var lastErr error
	for i := 0; i < idx.config.MaxRetries; i++ {
		logs, err := idx.client.FilterLogs(ctx, query)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		time.Sleep(idx.config.RetryDelay)
	}
	return nil, fmt.Errorf("failed after %d retries: %w", idx.config.MaxRetries, lastErr)
}

// Close closes the indexer and releases resources
func (idx *Indexer) Close() {
	if idx.client != nil {
		idx.client.Close()
	}
}
