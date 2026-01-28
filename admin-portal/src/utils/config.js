let config = null;

/**
 * Transform localhost URLs to use the current browser host.
 * This allows the same config to work in both local development
 * and remote deployments.
 *
 * @param {string} url - The URL to transform
 * @returns {string} - The transformed URL
 */
const transformUrl = (url) => {
  if (!url) return url;

  // Only transform if we're not already on localhost
  const currentHost = window.location.hostname;
  if (currentHost === 'localhost' || currentHost === '127.0.0.1') {
    return url;
  }

  // Replace localhost with the current host
  return url.replace(/http:\/\/localhost:/, `http://${currentHost}:`);
};

export const loadConfig = async () => {
  if (config) return config;

  try {
    const response = await fetch('/config.json');
    const rawConfig = await response.json();

    // Transform localhost URLs to use current host for remote deployments
    config = {
      ...rawConfig,
      indexerApiUrl: transformUrl(rawConfig.indexerApiUrl),
      gatewayApiUrl: transformUrl(rawConfig.gatewayApiUrl),
      rpcUrl: transformUrl(rawConfig.rpcUrl),
    };

    return config;
  } catch (error) {
    console.error('Failed to load config:', error);
    throw error;
  }
};

export const getConfig = () => {
  if (!config) {
    throw new Error('Config not loaded. Call loadConfig() first.');
  }
  return config;
};
