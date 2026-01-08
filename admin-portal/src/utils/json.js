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
 * Decode base64 string to UTF-8
 * @param {string} base64String - The base64 encoded string
 * @returns {string} - The decoded string
 */
const decodeBase64 = (base64String) => {
  try {
    // Check if this looks like base64 (only alphanumeric, +, /, and =)
    if (!/^[A-Za-z0-9+/=]+$/.test(base64String)) {
      return null;
    }
    // Use atob for browser-compatible base64 decoding
    return atob(base64String);
  } catch (error) {
    return null;
  }
};

/**
 * Safely parse and stringify event metadata
 * Attempts to decode base64 if the metadata appears to be encoded
 * @param {string} metadata - The metadata string to parse and stringify
 * @returns {string} - The formatted metadata string
 */
export const formatEventMetadata = (metadata) => {
  try {
    // First try to parse as JSON directly
    let parsed;
    try {
      parsed = JSON.parse(metadata);
      return safeStringify(parsed);
    } catch (jsonError) {
      // If direct JSON parsing fails, try to decode as base64 first
      const decoded = decodeBase64(metadata);
      if (decoded) {
        try {
          parsed = JSON.parse(decoded);
          return safeStringify(parsed);
        } catch (decodedJsonError) {
          // If decoded content isn't JSON, return the decoded string
          return decoded;
        }
      }
      // If not base64 either, return original
      return metadata;
    }
  } catch (error) {
    console.error('Failed to parse metadata:', error);
    return metadata; // Return original if all parsing fails
  }
};
