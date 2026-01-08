#!/bin/sh
set -e

# Default values from environment or defaults
RPC_URL=${RPC_URL:-http://hardhat:8545}
CONTRACT_ADDRESS=${CONTRACT_ADDRESS:-}

# Try to read configuration from JSON file if it exists
CONFIG_FILE=${CONFIG_FILE:-/app/config.json}
if [ -f "$CONFIG_FILE" ]; then
    echo "Config file found at $CONFIG_FILE, reading configuration..."

    CONFIG_CONTRACT_ADDRESS=$(jq -r '.contractAddress' "$CONFIG_FILE" 2>/dev/null || echo "")
    CONFIG_RPC_URL=$(jq -r '.networkRpcUrl' "$CONFIG_FILE" 2>/dev/null || echo "")

    # Use config file values if available and not null
    if [ -n "$CONFIG_CONTRACT_ADDRESS" ] && [ "$CONFIG_CONTRACT_ADDRESS" != "null" ]; then
        CONTRACT_ADDRESS="$CONFIG_CONTRACT_ADDRESS"
    fi

    if [ -n "$CONFIG_RPC_URL" ] && [ "$CONFIG_RPC_URL" != "null" ]; then
        RPC_URL="$CONFIG_RPC_URL"
    fi
fi

echo "Starting gateway with:"
echo "  Contract Address: ${CONTRACT_ADDRESS:-<not set>}"
echo "  RPC URL: $RPC_URL"

# Start the gateway with the configuration
# Contract address is optional - if not set, only /link endpoint will work
if [ -n "$CONTRACT_ADDRESS" ]; then
    exec ./gateway -contract "$CONTRACT_ADDRESS" -rpc "$RPC_URL"
else
    echo "Warning: Contract address not set. /transfers endpoint will not be available."
    exec ./gateway -rpc "$RPC_URL"
fi

