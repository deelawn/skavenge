/**
 * Utility functions for Skavenge webapp
 */

/**
 * Convert RPC URL for browser access
 * Replaces Docker internal hostnames with localhost
 *
 * @param {string} rpcUrl - The RPC URL from config
 * @returns {string} Browser-accessible RPC URL
 */
export function getBrowserRpcUrl(rpcUrl) {
  if (!rpcUrl) {
    return 'http://localhost:8545';
  }

  // Replace hardhat (Docker service name) with localhost for browser access
  return rpcUrl.replace(/http:\/\/hardhat:/, 'http://localhost:');
}

/**
 * Convert Gateway URL for browser access
 * Replaces Docker internal hostnames with localhost
 *
 * @param {string} gatewayUrl - The Gateway URL from config
 * @returns {string} Browser-accessible Gateway URL
 */
export function getBrowserGatewayUrl(gatewayUrl) {
  if (!gatewayUrl) {
    return 'http://localhost:4591';
  }

  // Replace gateway (Docker service name) with localhost for browser access
  return gatewayUrl.replace(/http:\/\/gateway:/, 'http://localhost:');
}
