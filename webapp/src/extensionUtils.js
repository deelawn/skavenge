// Extension ID - hardcoded to match manifest.json
const SKAVENGER_EXTENSION_ID = "hnbligdjmpihmhhgajlfjmckcnmnofbn";

/**
 * Basic send message to Skavenger extension (without auto-unlock)
 */
async function sendToExtensionBasic(message) {
  return new Promise((resolve, reject) => {
    if (typeof window.chrome === 'undefined' || !window.chrome.runtime) {
      reject(new Error('Chrome extension API not available'));
      return;
    }

    window.chrome.runtime.sendMessage(SKAVENGER_EXTENSION_ID, message, (response) => {
      if (window.chrome.runtime.lastError) {
        reject(new Error(window.chrome.runtime.lastError.message));
      } else if (response && response.success !== false) {
        resolve(response);
      } else {
        reject(new Error(response?.error || 'Unknown error'));
      }
    });
  });
}

/**
 * Request the extension to open (for password unlock)
 */
async function requestExtensionUnlock() {
  try {
    await sendToExtensionBasic({ action: 'requestLink' });
    return true;
  } catch (error) {
    console.error('Error requesting extension unlock:', error);
    return false;
  }
}

/**
 * Check if the extension session is valid
 */
async function checkExtensionSession() {
  try {
    const response = await sendToExtensionBasic({ action: 'verifyPassword' });
    return response.success === true;
  } catch (error) {
    return false;
  }
}

/**
 * Poll for extension unlock (wait for user to enter password)
 * @param {number} maxAttempts - Maximum number of polling attempts
 * @param {number} intervalMs - Interval between polls in milliseconds
 * @returns {Promise<boolean>} - True if unlocked, false if timeout
 */
async function pollForUnlock(maxAttempts = 60, intervalMs = 1000) {
  for (let i = 0; i < maxAttempts; i++) {
    const isUnlocked = await checkExtensionSession();
    if (isUnlocked) {
      return true;
    }
    // Wait before next poll
    await new Promise(resolve => setTimeout(resolve, intervalMs));
  }
  return false;
}

/**
 * Send message to Skavenger extension with automatic password unlock handling
 * If the extension returns "Password required or session expired", this function will:
 * 1. Open the extension popup to prompt for password
 * 2. Wait for the user to unlock
 * 3. Retry the original operation
 * 4. If unlock fails, throw the original error
 * 
 * @param {object} message - The message to send to the extension
 * @param {function} onUnlockPrompt - Optional callback when unlock is prompted
 * @returns {Promise} - Resolves with the extension response
 */
export async function sendToExtension(message, onUnlockPrompt = null) {
  try {
    // First attempt - try the operation directly
    const response = await sendToExtensionBasic(message);
    return response;
  } catch (error) {
    // Check if this is a password/session error
    if (error.message === 'Password required or session expired') {
      // Notify caller that we're prompting for unlock
      if (onUnlockPrompt) {
        onUnlockPrompt();
      }

      // Request extension to open for unlock
      const unlockRequested = await requestExtensionUnlock();
      if (!unlockRequested) {
        throw error; // Could not open extension, throw original error
      }

      // Poll for unlock (wait for user to enter password)
      const unlocked = await pollForUnlock();
      if (!unlocked) {
        // Timeout waiting for unlock
        throw new Error('Extension unlock timed out. Please try again.');
      }

      // Retry the original operation
      try {
        const retryResponse = await sendToExtensionBasic(message);
        return retryResponse;
      } catch (retryError) {
        // If retry also fails, throw the retry error
        throw retryError;
      }
    }

    // Not a password error, throw original error
    throw error;
  }
}

/**
 * Check if Skavenger extension has keys
 */
export async function checkSkavengerKeys() {
  try {
    const response = await sendToExtensionBasic({ action: 'hasKeys' });
    return response.hasKeys;
  } catch (error) {
    console.error('Error checking Skavenger keys:', error);
    return false;
  }
}

/**
 * Get Skavenger public key (with auto-unlock)
 */
export async function getSkavengerPublicKey(onUnlockPrompt = null) {
  try {
    const response = await sendToExtension({ action: 'getPublicKey' }, onUnlockPrompt);
    if (response.success) {
      return response.publicKey;
    }
    return null;
  } catch (error) {
    console.error('Error getting Skavenger public key:', error);
    return null;
  }
}

/**
 * Request link to Skavenger extension (open popup)
 */
export async function requestSkavengerLink() {
  try {
    await sendToExtensionBasic({ action: 'requestLink' });
    return true;
  } catch (error) {
    console.error('Error requesting link:', error);
    return false;
  }
}

export { SKAVENGER_EXTENSION_ID };

