# Skavenge Blockchain Indexer

A high-performance blockchain indexer for the Skavenge NFT smart contract. This indexer continuously monitors the blockchain, extracts all events related to the Skavenge contract, and stores them for efficient querying. The indexed data powers the frontend dashboard for viewing NFT status and application events.

## Features

- **Real-time Event Indexing**: Continuously polls the blockchain for new blocks and extracts all Skavenge contract events
- **Chain Reorganization Handling**: Automatically detects and handles blockchain reorgs to maintain data consistency
- **Comprehensive Event Coverage**: Indexes all 14 event types from the Skavenge contract:
  - `ClueMinted` - New NFT minting events
  - `ClueAttempted` - Solution attempt events
  - `ClueSolved` - Successful solution events
  - `SalePriceSet` / `SalePriceRemoved` - NFT marketplace events
  - `TransferInitiated` / `ProofProvided` / `ProofVerified` / `TransferCompleted` / `TransferCancelled` - Secure transfer protocol events
  - `Transfer` / `Approval` / `ApprovalForAll` - Standard ERC721 events
  - `AuthorizedMinterUpdated` - Minter permission events
- **NFT State Tracking**: Maintains current state of all NFTs including ownership, sale status, and solved status
- **Flexible Storage Interface**: Pluggable storage backend (in-memory implementation included, easily extensible to PostgreSQL, MongoDB, etc.)
- **Configurable**: Extensive configuration options for network, polling, and reorg protection
- **Resilient**: Automatic retry logic for RPC failures with exponential backoff
- **Timestamp Association**: All events are timestamped with their block time
- **Block Confirmation**: Configurable confirmation depth to avoid indexing temporary forks

## Architecture

```
┌─────────────────┐
│   Blockchain    │
│   (Ethereum)    │
└────────┬────────┘
         │
         │ RPC (eth_getLogs, eth_getBlockByNumber)
         │
┌────────▼────────┐
│  Block Poller   │ ◄── Continuously polls for new blocks
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Event Parser   │ ◄── Decodes contract events using ABI
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Reorg Handler  │ ◄── Detects and handles chain reorganizations
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│Storage Interface│ ◄── Stores events, blocks, and NFT state
└─────────────────┘
```

## Installation

### Prerequisites

- Go 1.21 or higher
- Access to an Ethereum RPC endpoint (local node, Infura, Alchemy, etc.)

### Build

```bash
cd indexer
go mod download
go build -o skavenge-indexer ./cmd
```

## Configuration

Create a `config.json` file based on the example:

```bash
cp config.example.json config.json
```

### Configuration Options

| Field | Type | Description | Default |
|-------|------|-------------|---------|
| `rpc_url` | string | Ethereum RPC endpoint URL | `http://localhost:8545` |
| `contract_address` | string | Skavenge contract address (hex) | Required |
| `start_block` | uint64 | Block number to start indexing from (0 = latest) | `0` |
| `poll_interval_seconds` | int | Seconds between block polls | `12` |
| `confirmation_blocks` | uint64 | Blocks to wait for confirmation (reorg protection) | `6` |
| `batch_size` | uint64 | Maximum blocks to process in one batch | `100` |
| `max_retries` | int | RPC call retry attempts | `3` |
| `retry_delay_seconds` | int | Delay between retries | `2` |

### Example Configuration

```json
{
  "rpc_url": "https://mainnet.infura.io/v3/YOUR-PROJECT-ID",
  "contract_address": "0x1234567890123456789012345678901234567890",
  "start_block": 18000000,
  "poll_interval_seconds": 12,
  "confirmation_blocks": 6,
  "batch_size": 100,
  "max_retries": 3,
  "retry_delay_seconds": 2
}
```

## Usage

### Run the Indexer

```bash
./skavenge-indexer -config config.json
```

The indexer will:
1. Connect to the RPC endpoint
2. Determine the starting block (from config or last indexed block)
3. Begin polling for new blocks
4. Extract and store all Skavenge contract events
5. Monitor for chain reorganizations
6. Continue running until stopped (Ctrl+C)

### Graceful Shutdown

Press `Ctrl+C` to gracefully shutdown the indexer. It will finish processing the current block and save its state before exiting.

## Storage Interface

The indexer uses a pluggable storage interface defined in `storage.go`. The interface provides methods for:

### Event Operations
- `SaveEvent` - Store an event
- `GetEvent` - Retrieve event by ID
- `GetEventsByTokenID` - Get all events for a specific NFT
- `GetEventsByBlock` - Get all events in a block
- `GetEventsByType` - Get events by type (e.g., all ClueMinted events)
- `GetEventsByTimeRange` - Query events within a time range
- `MarkEventsAsRemoved` - Mark events as removed (reorg handling)

### Block Operations
- `SaveBlock` - Store block metadata
- `GetBlock` - Retrieve block by number
- `GetBlockByHash` - Retrieve block by hash
- `GetLatestBlock` - Get the most recently indexed block
- `DeleteBlocksFrom` - Delete blocks from a certain height (reorg handling)

### NFT State Queries
- `GetNFTCurrentOwner` - Get current owner of an NFT
- `GetNFTsByOwner` - Get all NFTs owned by an address
- `GetNFTsForSale` - Get NFTs currently for sale
- `GetSolvedNFTs` - Get NFTs that have been solved

### Statistics
- `GetTotalEvents` - Total number of indexed events
- `GetTotalNFTs` - Total number of minted NFTs

### Implementing Custom Storage

To implement a custom storage backend (e.g., PostgreSQL):

1. Create a new file (e.g., `postgres_storage.go`)
2. Implement the `Storage` interface
3. Update `cmd/main.go` to use your storage implementation

Example:

```go
type PostgresStorage struct {
    db *sql.DB
}

func (p *PostgresStorage) SaveEvent(ctx context.Context, event *Event) error {
    // Your PostgreSQL implementation
}

// ... implement all other Storage interface methods
```

## Chain Reorganization Handling

The indexer implements robust reorg handling:

1. **Detection**: After each polling interval, checks the last N blocks (where N = `confirmation_blocks`) to see if their hashes still match the chain
2. **Rollback**: If a reorg is detected, deletes all blocks and events from the reorg point onward
3. **Re-indexing**: Automatically re-indexes the affected blocks with the new chain state
4. **State Consistency**: Rebuilds NFT state from remaining events after rollback

### Reorg Protection Configuration

The `confirmation_blocks` setting determines how deep the indexer looks for reorgs:
- **Lower values** (1-3): Faster indexing, less reorg protection
- **Higher values** (6-12): Slower indexing, better reorg protection
- **Recommended**: 6 blocks for most networks (~1 minute on Ethereum)

## Event Data Structures

All events include:
- Unique ID (transaction hash + log index)
- Event type
- Block number and hash
- Transaction hash
- Log index
- Timestamp
- Associated token ID (when applicable)
- Event-specific data

### Example Event JSON

```json
{
  "id": "0xabc...123-0",
  "type": "ClueMinted",
  "blockNumber": 18000000,
  "blockHash": "0xdef...456",
  "transactionHash": "0xabc...123",
  "logIndex": 0,
  "timestamp": "2024-01-01T12:00:00Z",
  "tokenId": "1",
  "data": {
    "tokenId": "1",
    "minter": "0x789...012"
  },
  "removed": false
}
```

## Performance Considerations

- **RPC Rate Limiting**: Adjust `poll_interval_seconds` and `batch_size` to avoid rate limits
- **Memory Usage**: The in-memory storage grows with indexed data; consider a persistent database for production
- **Confirmation Depth**: Higher `confirmation_blocks` increases safety but delays indexing
- **Batch Processing**: Larger `batch_size` improves throughput but increases memory usage

## Frontend Integration

The indexer provides all data needed for a dashboard UI:

### Dashboard Features
- **NFT Overview**: Total minted, solved, and for sale
- **NFT List**: All NFTs with current owner, sale price, solved status
- **NFT Details**: Full event history for each NFT
- **Activity Feed**: Recent events across all NFTs
- **Marketplace**: NFTs currently for sale with prices
- **Transfer Tracking**: Monitor ongoing secure transfers

### API Integration

Build a REST API or GraphQL server on top of the storage interface to serve data to the frontend. Example endpoints:

```
GET /api/events?tokenId=1          # Events for NFT #1
GET /api/nfts?owner=0x123...       # NFTs owned by address
GET /api/nfts/for-sale             # NFTs for sale
GET /api/events/recent?limit=50    # Recent events
GET /api/stats                     # Total NFTs, events, etc.
```

## Troubleshooting

### Indexer stops responding
- Check RPC endpoint connectivity
- Verify `max_retries` and `retry_delay_seconds` are appropriate
- Check for rate limiting from RPC provider

### Events missing
- Verify `contract_address` is correct
- Check `start_block` is before contract deployment
- Ensure RPC endpoint is fully synced

### High memory usage
- Switch from in-memory to persistent storage
- Reduce `batch_size`
- Implement data archival for old events

### Reorg issues
- Increase `confirmation_blocks` for better protection
- Check for network instability
- Verify RPC endpoint is reliable

## Development

### Running Tests

```bash
go test ./...
```

### Project Structure

```
indexer/
├── cmd/
│   └── main.go              # CLI entry point
├── config.go                # Configuration loading
├── config.example.json      # Example configuration
├── indexer.go               # Main indexer orchestrator
├── memory_storage.go        # In-memory storage implementation
├── parser.go                # Event parsing logic
├── storage.go               # Storage interface definition
├── types.go                 # Event and data structures
├── go.mod                   # Go module definition
└── README.md                # This file
```

## License

MIT

## Contributing

Contributions welcome! Please submit issues and pull requests.
