# Quick Start Guide

Get the Skavenge blockchain indexer up and running in 5 minutes.

## Step 1: Configure

Create your configuration file:

```bash
cd indexer
cp config.example.json config.json
```

Edit `config.json` with your settings:

```json
{
  "rpc_url": "https://mainnet.infura.io/v3/YOUR-PROJECT-ID",
  "contract_address": "0xYOUR_CONTRACT_ADDRESS_HERE",
  "start_block": 0,
  "poll_interval_seconds": 12,
  "confirmation_blocks": 6,
  "batch_size": 100,
  "max_retries": 3,
  "retry_delay_seconds": 2
}
```

**Required:**
- `rpc_url`: Your Ethereum RPC endpoint
- `contract_address`: The deployed Skavenge contract address

**Optional:**
- `start_block`: Set to contract deployment block to index from the beginning, or 0 for latest
- Other fields have sensible defaults

## Step 2: Install Dependencies

```bash
go mod download
```

## Step 3: Build

```bash
go build -o skavenge-indexer ./cmd
```

## Step 4: Run

```bash
./skavenge-indexer -config config.json
```

You should see output like:

```
[INDEXER] 2024/01/01 12:00:00 Starting indexer for contract 0x1234...
[INDEXER] 2024/01/01 12:00:00 Starting from block 18000000
[INDEXER] 2024/01/01 12:00:00 Starting blockchain indexer...
[INDEXER] 2024/01/01 12:00:01 Indexing blocks 18000000 to 18000099
[INDEXER] 2024/01/01 12:00:02 Block 18000042: indexed 3 events
[INDEXER] 2024/01/01 12:00:02 Block 18000087: indexed 1 events
...
```

## Step 5: Stop

Press `Ctrl+C` to gracefully shutdown the indexer.

## Next Steps

### Use a Persistent Database

The default in-memory storage will lose data on restart. For production, implement a persistent storage backend:

1. Create `postgres_storage.go` (or your preferred database)
2. Implement the `Storage` interface
3. Update `cmd/main.go` to use your storage

### Build an API

Create a REST API or GraphQL server to expose indexed data to your frontend:

```go
// Example REST API endpoints
GET /api/events?tokenId=1           // Events for NFT #1
GET /api/nfts?owner=0x123...        // NFTs by owner
GET /api/nfts/for-sale              // Marketplace
GET /api/events/recent?limit=50     // Activity feed
GET /api/stats                      // Statistics
```

### Build a Dashboard

Use the indexed data to power a frontend dashboard:
- NFT overview and statistics
- Individual NFT details and history
- Real-time activity feed
- Marketplace of NFTs for sale
- Transfer tracking

## Troubleshooting

### "Failed to connect to RPC"
- Check your `rpc_url` is correct and accessible
- Verify your RPC provider is online
- Check for firewall/network restrictions

### "No events found"
- Verify `contract_address` is correct
- Check `start_block` is at or before contract deployment
- Confirm the contract has activity in the block range

### High memory usage
- Switch to persistent storage instead of in-memory
- Reduce `batch_size`
- Implement data archival for old events

## Configuration Tips

### For Local Development (Hardhat/Ganache)
```json
{
  "rpc_url": "http://localhost:8545",
  "contract_address": "0x...",
  "start_block": 0,
  "poll_interval_seconds": 1,
  "confirmation_blocks": 1
}
```

### For Testnet (Sepolia/Goerli)
```json
{
  "rpc_url": "https://sepolia.infura.io/v3/YOUR-ID",
  "contract_address": "0x...",
  "start_block": 0,
  "poll_interval_seconds": 12,
  "confirmation_blocks": 3
}
```

### For Mainnet
```json
{
  "rpc_url": "https://mainnet.infura.io/v3/YOUR-ID",
  "contract_address": "0x...",
  "start_block": 18000000,
  "poll_interval_seconds": 12,
  "confirmation_blocks": 6
}
```

## Questions?

See the full [README.md](README.md) for detailed documentation.
