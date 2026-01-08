/**
 * Safely stringify JSON data that may contain BigInt values
 * @param {any} data - The data to stringify
 * @param {number} space - Number of spaces for indentation (default: 2)
 * @returns {string} - The stringified JSON
 */
export const safeStringify = (data, space = 2) => {
  return JSON.stringify(data, (key, value) => {
    // Convert BigInt to string
    if (typeof value === 'bigint') {
      return value.toString();
    }
    return value;
  }, space);
};

/**
 * Safely parse and stringify event metadata
 * @param {string} metadata - The metadata string to parse and stringify
 * @returns {string} - The formatted metadata string
 */
export const formatEventMetadata = (metadata) => {
  try {
    const parsed = JSON.parse(metadata);
    return safeStringify(parsed);
  } catch (error) {
    console.error('Failed to parse metadata:', error);
    return metadata; // Return original if parsing fails
  }
};
