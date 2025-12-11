// Configuration loader for Skavenge webapp
// Reads configuration from mounted config file

let configCache = null;

/**
 * Chain ID constant - read from config.json file
 * This constant reads from the loaded configuration in config.json
 * Call loadConfig() first to ensure the config is loaded before accessing
 * 
 * Access the value using: CHAIN_ID.valueOf() or Number(CHAIN_ID)
 * Or use the getter: CHAIN_ID.value
 */
export const CHAIN_ID = {
  /**
   * Gets the chain ID value from the loaded configuration
   * @returns {number} The chain ID from config.json, or 1337 as fallback
   */
  get value() {
    if (configCache && configCache.chainId !== undefined) {
      return configCache.chainId;
    }
    // Fallback to 1337 if config not loaded yet (should not happen in normal usage)
    return 1337;
  },

  /**
   * Allows CHAIN_ID to be used in numeric contexts
   */
  valueOf() {
    return this.value;
  },

  /**
   * String representation
   */
  toString() {
    return String(this.value);
  }
};

/**
 * Loads configuration from the config.json file
 * This file should be mounted in the container at /usr/share/nginx/html/config.json
 *
 * @returns {Promise<{contractAddress: string, networkRpcUrl: string, chainId: number, gatewayUrl: string}>}
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

    // Validate chainId is present in config
    if (config.chainId === undefined) {
      throw new Error('Config file must contain chainId');
    }

    // Validate gatewayUrl is present in config
    if (!config.gatewayUrl) {
      throw new Error('Config file must contain gatewayUrl');
    }

    configCache = {
      contractAddress: config.contractAddress,
      networkRpcUrl: config.networkRpcUrl,
      chainId: config.chainId,
      gatewayUrl: config.gatewayUrl
    };

    return configCache;
  } catch (error) {
    console.error('Error loading configuration:', error);

    // Return default values for development (with chainId fallback)
    return {
      contractAddress: '0x0000000000000000000000000000000000000000',
      networkRpcUrl: 'http://localhost:8545',
      chainId: 1337,
      gatewayUrl: 'http://localhost:4591'
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
