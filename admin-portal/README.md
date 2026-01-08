# Skavenge Admin Portal

A modern, single-page React application for administering and viewing Skavenge clues (NFTs), ownership details, and indexed events.

## Features

- **Dashboard**: Overview of clues, owners, events, and system configuration
- **Clues Management**: Browse all clues with pagination, search functionality
- **Clue Details**: View detailed information from both the indexer and smart contract
  - Ownership information
  - Event history (sortable)
  - Contract state data
- **Events Viewer**: Filter and browse all indexed events
  - Filter by event type
  - Sort by timestamp
  - Expandable event details
- **Owner Management**: View all clue owners
  - List of owned clues
  - Activity audit log
  - Linked Skavenge public keys (via gateway API)

## Technology Stack

- **React 18** - UI framework
- **React Router** - Client-side routing
- **Tailwind CSS** - Styling
- **Web3.js** - Ethereum interaction
- **Axios** - HTTP client
- **Vite** - Build tool and dev server

## Development

### Prerequisites

- Node.js 20+
- Docker and Docker Compose (for containerized deployment)

### Local Development

1. Install dependencies:
```bash
cd admin-portal
npm install
```

2. Update configuration in `public/config.json`:
```json
{
  "indexerApiUrl": "http://localhost:4040",
  "gatewayApiUrl": "http://localhost:4591",
  "contractAddress": "0x...",
  "rpcUrl": "http://localhost:8545",
  "chainId": 31337
}
```

3. Start development server:
```bash
npm run dev
```

The application will be available at `http://localhost:3000`

### Building for Production

```bash
npm run build
```

The built files will be in the `dist/` directory.

### Preview Production Build

```bash
npm run serve
```

## Docker Deployment

The admin portal is containerized and part of the Skavenge docker-compose setup.

### Build and Start with Docker Compose

From the project root:

```bash
# Build all services including admin portal
make docker-build

# Start all services
make start-with-setup
```

The admin portal will be available at `http://localhost:3000`

### Rebuild Admin Portal Only

```bash
# Rebuild with cache
make rebuild-admin-portal

# Rebuild without cache
make rebuild-admin-portal-no-cache
```

### Stop Services

```bash
make stop
```

## Configuration

The application reads configuration from `/public/config.json` at runtime. In Docker, this uses service hostnames:

```json
{
  "indexerApiUrl": "http://indexer:4040",
  "gatewayApiUrl": "http://linked-accounts-gateway:4591",
  "contractAddress": "0x5FbDB2315678afecb367f032d93F642f64180aa3",
  "rpcUrl": "http://hardhat-node:8545",
  "chainId": 31337
}
```

## Project Structure

```
admin-portal/
├── public/
│   └── config.json          # Runtime configuration
├── src/
│   ├── components/          # Reusable UI components
│   │   ├── ClueLink.jsx
│   │   ├── ErrorMessage.jsx
│   │   ├── EventBadge.jsx
│   │   ├── Loading.jsx
│   │   ├── Navigation.jsx
│   │   ├── OwnerLink.jsx
│   │   └── Pagination.jsx
│   ├── pages/               # Page components
│   │   ├── ClueDetails.jsx
│   │   ├── CluesList.jsx
│   │   ├── Dashboard.jsx
│   │   ├── EventsList.jsx
│   │   ├── OwnerDetails.jsx
│   │   └── OwnersList.jsx
│   ├── utils/               # Utility functions
│   │   ├── api.js          # API clients (Indexer & Gateway)
│   │   ├── config.js       # Configuration loader
│   │   ├── contractABI.js  # Smart contract ABI
│   │   └── web3.js         # Web3 utilities
│   ├── App.jsx             # Main app component
│   ├── main.jsx            # Entry point
│   └── index.css           # Global styles
├── Dockerfile              # Docker build configuration
├── package.json            # Dependencies
├── vite.config.js          # Vite configuration
└── tailwind.config.js      # Tailwind configuration
```

## API Integration

### Indexer API

The admin portal connects to the indexer API for:
- Fetching all clues with ownership data
- Querying events (with filtering and sorting)
- Getting ownership information

### Gateway API

Used to retrieve linked Skavenge public keys for Ethereum addresses.

### Smart Contract

Direct Web3 calls to the Skavenge contract for:
- Real-time clue state (solved status, sale price, etc.)
- Contract-specific data not indexed

## Features in Detail

### Navigation

All clue IDs and owner addresses are clickable links throughout the application, allowing easy navigation between related data.

### Event Filtering

Events can be filtered by:
- Event type (ClueMinted, ClueSolved, Transfer, etc.)
- Sort order (ascending/descending by timestamp)

### Pagination

Lists are paginated for performance:
- Clues: 20 per page
- Events: 50 per page

### Expandable Details

Events have expandable sections showing:
- Transaction hash and block details
- Initiator address
- Event-specific metadata (JSON)

## Future Enhancements

Potential features for future versions:
- Real-time updates via WebSocket
- Advanced search and filtering
- Export functionality (CSV, JSON)
- Analytics and charts
- User authentication/authorization
