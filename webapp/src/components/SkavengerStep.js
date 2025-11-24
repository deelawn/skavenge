import React, { useState, useEffect } from 'react';

// Extension ID - hardcoded to match manifest.json
const SKAVENGER_EXTENSION_ID = "hnbligdjmpihmhhgajlfjmckcnmnofbn";

/**
 * Send message to Skavenger extension
 */
async function sendToExtension(message) {
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
 * Check if Skavenger extension has keys
 */
async function checkSkavengerKeys() {
  try {
    const response = await sendToExtension({ action: 'hasKeys' });
    return response.hasKeys;
  } catch (error) {
    console.error('Error checking Skavenger keys:', error);
    return false;
  }
}

/**
 * Get Skavenger public key
 */
async function getSkavengerPublicKey() {
  try {
    const response = await sendToExtension({ action: 'getPublicKey' });
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
 * Request link to Skavenger extension
 */
async function requestSkavengerLink() {
  try {
    await sendToExtension({ action: 'requestLink' });
    return true;
  } catch (error) {
    console.error('Error requesting link:', error);
    return false;
  }
}

/**
 * Skavenger Step Component
 * Handles linking to the Skavenger browser extension
 *
 * @param {object} props
 * @param {function} props.onKeysFound - Callback when keys are found
 * @param {function} props.onToast - Callback to show toast message
 */
function SkavengerStep({ onKeysFound, onToast }) {
  const [status, setStatus] = useState('checking'); // checking, no-keys, has-keys, pending, error
  const [statusText, setStatusText] = useState('Checking for keys...');
  const [showButton, setShowButton] = useState(false);
  const [showInstructions, setShowInstructions] = useState(false);
  const [publicKey, setPublicKey] = useState(null);

  // Check for keys on mount
  useEffect(() => {
    checkInitialStatus();
  }, []);

  const checkInitialStatus = async () => {
    try {
      const hasKeys = await checkSkavengerKeys();

      if (hasKeys) {
        const key = await getSkavengerPublicKey();
        if (key) {
          setPublicKey(key);
          setStatus('has-keys');
          setStatusText('Keys found');
          setShowButton(false);
          onKeysFound(key);
        } else {
          setStatus('no-keys');
          setStatusText('Please generate or import keys');
          setShowButton(true);
        }
      } else {
        setStatus('no-keys');
        setStatusText('No keys found');
        setShowButton(true);
      }
    } catch (error) {
      setStatus('error');
      setStatusText(error.message || 'Extension not found - please install the Skavenger extension');
      setShowButton(false);
    }
  };

  const handleLinkClick = async () => {
    setStatusText('Opening extension...');
    setStatus('pending');

    try {
      const response = await requestSkavengerLink();

      if (response) {
        onToast('Extension opened. Please generate keys or enter your password.', 'info');
        setStatusText('Waiting for you to complete setup in the extension...');
        startPolling();
      }
    } catch (error) {
      setShowInstructions(true);
      setStatusText('Extension not found');
      setStatus('error');
      onToast('Please install the Skavenger extension and try again', 'error');
    }
  };

  const startPolling = () => {
    let pollCount = 0;
    const maxPolls = 60;

    const interval = setInterval(async () => {
      pollCount++;

      try {
        const hasKeys = await checkSkavengerKeys();
        if (hasKeys) {
          const key = await getSkavengerPublicKey();
          if (key) {
            clearInterval(interval);
            setPublicKey(key);
            setStatus('has-keys');
            setStatusText('Keys found');
            setShowButton(false);
            setShowInstructions(false);
            onToast('Skavenger account linked successfully', 'success');
            onKeysFound(key);
          }
        }

        if (pollCount >= maxPolls) {
          clearInterval(interval);
          setStatus('no-keys');
          setStatusText('Link timed out. Please try again.');
          onToast('Could not detect keys. Please try again.', 'error');
        }
      } catch (error) {
        // Continue polling on error
      }
    }, 1000);
  };

  return (
    <div className="step">
      <h2>Step 1: Link Skavenger Account</h2>
      <p className="info">Connect your Skavenger extension to generate or access your cryptographic key pair.</p>

      <div className="status">
        <span className={`status-indicator ${status === 'has-keys' ? 'connected' : ''} ${status === 'pending' ? 'pending' : ''}`}></span>
        <span>{statusText}</span>
      </div>

      {publicKey && (
        <div className="account-value" style={{ marginTop: '12px' }}>
          {publicKey}
        </div>
      )}

      {showButton && (
        <button className="primary" onClick={handleLinkClick}>
          Link Skavenger Account
        </button>
      )}

      {showInstructions && (
        <div className="link-instructions">
          <p>
            <strong>Instructions:</strong> Click the Skavenger extension icon in your browser toolbar.
            If you haven't set up an account, create a password and generate keys.
            If you already have an account, enter your password to unlock.
          </p>
        </div>
      )}
    </div>
  );
}

export default SkavengerStep;
