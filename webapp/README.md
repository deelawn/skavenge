# Skavenge Web Application

A React-based single page application for onboarding users to the Skavenge NFT scavenger hunt game.

## Features

- **Skavenger Extension Integration**: Connect to the Skavenger browser extension for cryptographic key management
- **MetaMask Integration**: Connect to MetaMask wallet for blockchain transactions
- **Configuration Management**: Read smart contract address and blockchain network RPC URL from a mounted config file
- **Dashboard**: Display connected accounts and configuration information

## Development

### Prerequisites

- Node.js 18 or higher
- npm

### Running Locally

```bash
# Install dependencies
npm install

# Start development server
npm start
```

The app will open at [http://localhost:3000](http://localhost:3000)

### Building for Production

```bash
# Build the app
npm run build
```

The production build will be in the `build/` directory.

## Docker Deployment

### Configuration

The app reads configuration from a `config.json` file that should be mounted at runtime. The configuration file must contain:

```json
{
  "contractAddress": "0x1234567890123456789012345678901234567890",
  "networkRpcUrl": "http://localhost:8545"
}
```

- `contractAddress`: The Ethereum smart contract address for the Skavenge game
- `networkRpcUrl`: The RPC URL for the blockchain network

### Running with Docker Compose

The webapp is part of the docker-compose setup. The `config.json` file in the webapp directory is automatically mounted to the container:

```bash
# Build and start all services
docker-compose up --build

# Start in detached mode
docker-compose up -d

# View logs
docker-compose logs webapp

# Stop all services
docker-compose down
```

The webapp will be accessible at [http://localhost:8080](http://localhost:8080)

### Running Docker Container Directly

```bash
# Build the image
docker build -t skavenge-webapp ./webapp

# Run with mounted config file
docker run -p 8080:80 \
  -v $(pwd)/webapp/config.json:/usr/share/nginx/html/config.json:ro \
  skavenge-webapp
```

## Configuration

### Contract Address

Update the `contractAddress` in `config.json` with your deployed Skavenge contract address.

### Network RPC URL

Update the `networkRpcUrl` in `config.json` with the RPC URL of your target blockchain network:

- Local Hardhat: `http://localhost:8545`
- Ethereum Mainnet: `https://mainnet.infura.io/v3/YOUR_PROJECT_ID`
- Sepolia Testnet: `https://sepolia.infura.io/v3/YOUR_PROJECT_ID`
- Other networks: Use appropriate RPC URL

## Architecture

### Components

- **App.js**: Main application component, manages state and screens
- **SkavengerStep.js**: Handles Skavenger extension connection
- **MetaMaskStep.js**: Handles MetaMask wallet connection
- **Dashboard.js**: Displays connected accounts and configuration
- **Toast.js**: Toast notification component

### Configuration Loading

The `src/config.js` module handles loading configuration from the mounted `config.json` file. It includes:

- Caching to avoid repeated fetches
- Validation of required fields
- Fallback to default values if config loading fails

## Browser Extension Requirements

### Skavenger Extension

The app requires the Skavenger browser extension to be installed. The extension ID is hardcoded in `src/components/SkavengerStep.js`:

```javascript
const SKAVENGER_EXTENSION_ID = "hnbligdjmpihmhhgajlfjmckcnmnofbn";
```

Update this ID if you're using a different extension ID.

### MetaMask Extension

The app requires MetaMask to be installed for wallet connectivity.

## File Structure

```
webapp/
├── public/
│   └── index.html          # HTML template
├── src/
│   ├── components/
│   │   ├── Dashboard.js    # Dashboard component
│   │   ├── MetaMaskStep.js # MetaMask connection component
│   │   ├── SkavengerStep.js # Skavenger connection component
│   │   └── Toast.js        # Toast notification component
│   ├── App.css             # Application styles
│   ├── App.js              # Main application component
│   ├── config.js           # Configuration loader
│   └── index.js            # Application entry point
├── .gitignore              # Git ignore rules
├── config.json             # Configuration file (mounted in Docker)
├── Dockerfile              # Multi-stage Docker build
├── package.json            # npm dependencies and scripts
└── README.md               # This file
```

## Troubleshooting

### Configuration Not Loading

If the configuration is not loading:

1. Check that `config.json` exists in the webapp directory
2. Verify the file is properly mounted in the Docker container
3. Check browser console for errors
4. Ensure the JSON is valid

### Extension Not Connecting

If the Skavenger extension is not connecting:

1. Verify the extension is installed
2. Check that the extension ID in `SkavengerStep.js` matches your extension
3. Ensure the extension has permissions to communicate with the webapp

### MetaMask Not Connecting

If MetaMask is not connecting:

1. Verify MetaMask is installed
2. Check that you're using a browser that supports MetaMask
3. Ensure you've approved the connection request

## Legacy Files

The original vanilla JavaScript implementation has been moved to `old-vanilla-app/` for reference.
