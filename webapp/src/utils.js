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
      skavengePublicKey: data.skavenge_public_key,
    };
  } catch (error) {
    console.error('Error checking linkage on gateway:', error);
    return {
      exists: false,
      error: 'Failed to connect to the gateway server',
    };
  }
}

/**
 * Get Skavenge public key for an Ethereum address from the gateway
 * @param {string} ethereumAddress - The Ethereum address to look up
 * @returns {Promise<{success: boolean, publicKey?: string, error?: string}>}
 */
export async function getPublicKeyByEthereumAddress(ethereumAddress) {
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
      return {
        success: false,
        error: 'Buyer has not linked their Skavenge account',
      };
    }

    if (!response.ok) {
      const data = await response.json();
      return {
        success: false,
        error: data.error || 'Failed to get public key from gateway',
      };
    }

    const data = await response.json();
    return {
      success: true,
      publicKey: data.skavenge_public_key,
    };
  } catch (error) {
    console.error('Error getting public key from gateway:', error);
    return {
      success: false,
      error: 'Failed to connect to the gateway server',
    };
  }
}

/**
 * Store transfer ciphertexts on the gateway after proof submission
 * @param {string} transferId - The transfer ID (hex string with 0x prefix)
 * @param {string} buyerCiphertext - The buyer's ciphertext (hex string)
 * @param {string} sellerCiphertext - The seller's ciphertext (hex string)
 * @param {string} extensionId - The extension ID for signing
 * @returns {Promise<{success: boolean, error?: string}>}
 */
export async function storeTransferCiphertext(transferId, buyerCiphertext, sellerCiphertext, extensionId) {
  try {
    const config = await loadConfig();
    const gatewayUrl = getBrowserGatewayUrl(config.gatewayUrl);

    // Create message to sign
    const message = `Store transfer ciphertext for ${transferId}`;

    // Sign the message with the extension
    const signResponse = await new Promise((resolve) => {
      chrome.runtime.sendMessage(extensionId, {
        action: 'signMessage',
        message: message
      }, resolve);
    });

    if (!signResponse || !signResponse.success) {
      return {
        success: false,
        error: signResponse?.error || 'Failed to sign message with extension',
      };
    }

    // Store the ciphertexts on the gateway
    const response = await fetch(`${gatewayUrl}/transfers`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        transfer_id: transferId,
        buyer_ciphertext: buyerCiphertext,
        seller_ciphertext: sellerCiphertext,
        message: message,
        signature: signResponse.signature,
      }),
    });

    if (!response.ok) {
      const data = await response.json();
      return {
        success: false,
        error: data.error || 'Failed to store ciphertexts on gateway',
      };
    }

    return { success: true };
  } catch (error) {
    console.error('Error storing ciphertexts on gateway:', error);
    return {
      success: false,
      error: 'Failed to connect to the gateway server',
    };
  }
}
