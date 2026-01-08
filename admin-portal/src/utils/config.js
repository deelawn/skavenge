let config = null;

export const loadConfig = async () => {
  if (config) return config;

  try {
    const response = await fetch('/config.json');
    config = await response.json();
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
