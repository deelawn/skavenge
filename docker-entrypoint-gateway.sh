#!/bin/sh
set -e

# Default config file path
CONFIG_FILE=${CONFIG_FILE:-/app/config.json}

# Check if config file exists
if [ ! -f "$CONFIG_FILE" ]; then
    echo "Error: Config file not found at $CONFIG_FILE"
    exit 1
fi

# Read configuration from JSON file
CONTRACT_ADDRESS=$(jq -r '.contractAddress' "$CONFIG_FILE")
RPC_URL=$(jq -r '.networkRpcUrl' "$CONFIG_FILE")

# Validate required configuration
if [ -z "$CONTRACT_ADDRESS" ] || [ "$CONTRACT_ADDRESS" = "null" ]; then
    echo "Error: contractAddress not found in config file"
    exit 1
fi

if [ -z "$RPC_URL" ] || [ "$RPC_URL" = "null" ]; then
    echo "Error: networkRpcUrl not found in config file"
    exit 1
fi

echo "Starting gateway with:"
echo "  Contract Address: $CONTRACT_ADDRESS"
echo "  RPC URL: $RPC_URL"

# Start the gateway with the configuration
exec ./gateway -contract "$CONTRACT_ADDRESS" -rpc "$RPC_URL"

