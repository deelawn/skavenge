package indexer

import (
	"encoding/json"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Config holds the indexer configuration
type Config struct {
	// RPC endpoint for the blockchain
	RPCURL string `json:"rpc_url"`

	// Contract address to index
	ContractAddress common.Address `json:"contract_address"`

	// Starting block number (0 means start from latest)
	StartBlock uint64 `json:"start_block"`

	// Polling interval for new blocks
	PollInterval time.Duration `json:"poll_interval"`

	// Number of blocks to wait before considering a block confirmed (reorg protection)
	ConfirmationBlocks uint64 `json:"confirmation_blocks"`

	// Maximum number of blocks to process in a single batch
	BatchSize uint64 `json:"batch_size"`

	// Number of retries for failed RPC calls
	MaxRetries int `json:"max_retries"`

	// Retry delay
	RetryDelay time.Duration `json:"retry_delay"`
}

// DefaultConfig returns a config with sensible defaults
func DefaultConfig() *Config {
	return &Config{
		RPCURL:             "http://localhost:8545",
		ContractAddress:    common.Address{},
		StartBlock:         0,
		PollInterval:       12 * time.Second, // ~1 block on Ethereum
		ConfirmationBlocks: 6,                // Wait for 6 confirmations
		BatchSize:          100,
		MaxRetries:         3,
		RetryDelay:         2 * time.Second,
	}
}

// LoadConfig loads configuration from a JSON file
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var rawConfig struct {
		RPCURL             string `json:"rpc_url"`
		ContractAddress    string `json:"contract_address"`
		StartBlock         uint64 `json:"start_block"`
		PollIntervalSec    int    `json:"poll_interval_seconds"`
		ConfirmationBlocks uint64 `json:"confirmation_blocks"`
		BatchSize          uint64 `json:"batch_size"`
		MaxRetries         int    `json:"max_retries"`
		RetryDelaySec      int    `json:"retry_delay_seconds"`
	}

	if err := json.Unmarshal(data, &rawConfig); err != nil {
		return nil, err
	}

	config := &Config{
		RPCURL:             rawConfig.RPCURL,
		ContractAddress:    common.HexToAddress(rawConfig.ContractAddress),
		StartBlock:         rawConfig.StartBlock,
		PollInterval:       time.Duration(rawConfig.PollIntervalSec) * time.Second,
		ConfirmationBlocks: rawConfig.ConfirmationBlocks,
		BatchSize:          rawConfig.BatchSize,
		MaxRetries:         rawConfig.MaxRetries,
		RetryDelay:         time.Duration(rawConfig.RetryDelaySec) * time.Second,
	}

	// Apply defaults for zero values
	if config.PollInterval == 0 {
		config.PollInterval = 12 * time.Second
	}
	if config.ConfirmationBlocks == 0 {
		config.ConfirmationBlocks = 6
	}
	if config.BatchSize == 0 {
		config.BatchSize = 100
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}
	if config.RetryDelay == 0 {
		config.RetryDelay = 2 * time.Second
	}

	return config, nil
}
