/**
 * Utility functions for Skavenge webapp
 */

import { loadConfig } from './config.js';

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

/**
 * Check if a linkage exists on the gateway server
 * @param {string} ethereumAddress - The Ethereum address to check
 * @returns {Promise<{exists: boolean, skavengePublicKey?: string, error?: string}>}
 */
export async function checkLinkageOnGateway(ethereumAddress) {
  try {
    const config = await loadConfig();
    const gatewayUrl = getBrowserGatewayUrl(config.gatewayUrl);

    const response = await fetch(`${gatewayUrl}/link?ethereumAddress=${encodeURIComponent(ethereumAddress)}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (response.status === 404) {
      // Linkage not found
      return { exists: false };
    }

    if (!response.ok) {
      const data = await response.json();
      return {
        exists: false,
        error: data.error || 'Failed to check linkage on server',
      };
    }

    const data = await response.json();
    return {
      exists: true,
      skavengePublicKey: data.skavengePublicKey,
    };
  } catch (error) {
    console.error('Error checking linkage on gateway:', error);
    return {
      exists: false,
      error: 'Failed to connect to the gateway server',
    };
  }
}
