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
