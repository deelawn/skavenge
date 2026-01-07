#!/bin/sh
set -e

# Wait for config file to be available
while [ ! -f /app/webapp-config.json ]; do
  echo "Waiting for webapp config..."
  sleep 2
done

# Extract values from webapp config
CONTRACT_ADDRESS=$(cat /app/webapp-config.json | grep -o '"contractAddress": "[^"]*"' | cut -d'"' -f4)
RPC_URL=$(cat /app/webapp-config.json | grep -o '"networkRpcUrl": "[^"]*"' | cut -d'"' -f4)

# Create indexer config
cat > /app/config.json <<EOF
{
  "rpc_url": "${RPC_URL}",
  "contract_address": "${CONTRACT_ADDRESS}",
  "start_block": 1,
  "poll_interval_seconds": 2,
  "confirmation_blocks": 1,
  "batch_size": 100,
  "max_retries": 3,
  "retry_delay_seconds": 2
}
EOF

echo "Indexer configuration created:"
cat /app/config.json

# Execute the indexer
exec "$@"
