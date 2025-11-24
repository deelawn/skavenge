// Configuration loader for Skavenge webapp
// Reads configuration from mounted config file

let configCache = null;

/**
 * Loads configuration from the config.json file
 * This file should be mounted in the container at /usr/share/nginx/html/config.json
 *
 * @returns {Promise<{contractAddress: string, networkRpcUrl: string}>}
 */
export async function loadConfig() {
  // Return cached config if already loaded
  if (configCache) {
    return configCache;
  }

  try {
    const response = await fetch('/config.json');

    if (!response.ok) {
      throw new Error(`Failed to load config: ${response.status} ${response.statusText}`);
    }

    const config = await response.json();

    // Validate required fields
    if (!config.contractAddress || !config.networkRpcUrl) {
      throw new Error('Config file must contain contractAddress and networkRpcUrl');
    }

    configCache = {
      contractAddress: config.contractAddress,
      networkRpcUrl: config.networkRpcUrl
    };

    return configCache;
  } catch (error) {
    console.error('Error loading configuration:', error);

    // Return default values for development
    return {
      contractAddress: '0x0000000000000000000000000000000000000000',
      networkRpcUrl: 'http://localhost:8545'
    };
  }
}

/**
 * Gets the cached configuration
 * Call loadConfig() first to ensure config is loaded
 *
 * @returns {object|null}
 */
export function getConfig() {
  return configCache;
}

/**
 * Clears the configuration cache
 * Useful for testing or forcing a reload
 */
export function clearConfigCache() {
  configCache = null;
}
