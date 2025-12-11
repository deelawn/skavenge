import React, { useState, useEffect } from 'react';
import { loadConfig, CHAIN_ID } from '../config.js';
import { getBrowserRpcUrl, getBrowserGatewayUrl } from '../utils.js';
import { Web3 } from 'web3';

/**
 * Check if MetaMask is installed
 */
function isMetaMaskInstalled() {
  return typeof window.ethereum !== 'undefined' && window.ethereum.isMetaMask;
}

/**
 * Convert chain ID to hex format for MetaMask
 */
function toHex(num) {
  return '0x' + num.toString(16);
}

/**
 * Switch MetaMask to the specified network
 * If the network doesn't exist, it will be added automatically
 */
async function switchToNetwork(chainId, rpcUrl) {
  const chainIdHex = toHex(chainId);
  const browserRpcUrl = getBrowserRpcUrl(rpcUrl);

  try {
    // Try to switch to the network
    await window.ethereum.request({
      method: 'wallet_switchEthereumChain',
      params: [{ chainId: chainIdHex }],
    });
  } catch (switchError) {
    // This error code indicates that the chain has not been added to MetaMask
    if (switchError.code === 4902) {
      try {
        // Add the network to MetaMask
        await window.ethereum.request({
          method: 'wallet_addEthereumChain',
          params: [
            {
              chainId: chainIdHex,
              chainName: `Skavenge Network (${chainId})`,
              rpcUrls: [browserRpcUrl],
              nativeCurrency: {
                name: 'Ether',
                symbol: 'ETH',
                decimals: 18,
              },
            },
          ],
        });
      } catch (addError) {
        // Provide more detailed error message
        const errorMsg = addError.message || 'Failed to add network to MetaMask.';
        throw new Error(`Failed to add network: ${errorMsg}`);
      }
    } else if (switchError.code === 4001) {
      throw new Error('Network switch rejected by user.');
    } else {
      throw new Error('Failed to switch network.');
    }
  }
}

/**
 * Connect to MetaMask with automatic network configuration
 */
async function connectMetaMask() {
  if (!isMetaMaskInstalled()) {
    throw new Error('MetaMask is not installed. Please install MetaMask extension.');
  }

  try {
    // Load configuration to get chain ID and RPC URL
    const config = await loadConfig();
    const chainId = config.chainId || CHAIN_ID.value;
    const rpcUrl = config.networkRpcUrl;

    // First, request account access
    const accounts = await window.ethereum.request({
      method: 'eth_requestAccounts'
    });

    if (!accounts || accounts.length === 0) {
      return null;
    }

    // Get current chain ID
    const currentChainId = await window.ethereum.request({
      method: 'eth_chainId'
    });

    // Convert to decimal for comparison
    const currentChainIdDecimal = parseInt(currentChainId, 16);

    // Switch network if needed
    if (currentChainIdDecimal !== chainId) {
      await switchToNetwork(chainId, rpcUrl);
    }

    return accounts[0];
  } catch (error) {
    if (error.code === 4001) {
      throw new Error('MetaMask connection rejected.');
    } else if (error.message) {
      throw error;
    } else {
      throw new Error('Failed to connect to MetaMask.');
    }
  }
}

/**
 * Get MetaMask account
 */
async function getMetaMaskAccount() {
  if (!isMetaMaskInstalled()) {
    return null;
  }

  try {
    const accounts = await window.ethereum.request({
      method: 'eth_accounts'
    });

    if (accounts && accounts.length > 0) {
      return accounts[0];
    }

    return null;
  } catch (error) {
    return null;
  }
}

/**
 * Request user to sign a message linking their MetaMask address to their Skavenge public key
 * @param {string} address - The Ethereum address to link
 * @param {string} skavengerPublicKey - The Skavenge public key to link
 * @returns {Promise<string>} The signature
 */
async function signLinkageMessage(address, skavengerPublicKey) {
  const message = `Link MetaMask address ${address} to Skavenge key ${skavengerPublicKey}`;

  try {
    const signature = await window.ethereum.request({
      method: 'personal_sign',
      params: [message, address],
    });

    return { signature, message };
  } catch (error) {
    if (error.code === 4001) {
      throw new Error('Signature request rejected by user.');
    }
    throw new Error('Failed to sign message: ' + (error.message || 'Unknown error'));
  }
}

/**
 * Verify that the signature was created by the owner of the address
 * @param {string} message - The original message that was signed
 * @param {string} signature - The signature to verify
 * @param {string} address - The address that should have signed the message
 * @returns {boolean} True if the signature is valid
 */
function verifySignature(message, signature, address) {
  try {
    const web3 = new Web3();
    const recoveredAddress = web3.eth.accounts.recover(message, signature);

    // Compare addresses (case-insensitive)
    return recoveredAddress.toLowerCase() === address.toLowerCase();
  } catch (error) {
    console.error('Signature verification failed:', error);
    return false;
  }
}

/**
 * Save verified linkage to localStorage
 * @param {string} address - The Ethereum address
 * @param {string} skavengerPublicKey - The Skavenge public key
 */
function saveVerifiedLinkage(address, skavengerPublicKey) {
  try {
    const linkageData = {
      address: address.toLowerCase(),
      skavengerPublicKey,
      timestamp: Date.now(),
    };
    localStorage.setItem('skavenge_verified_linkage', JSON.stringify(linkageData));
  } catch (error) {
    console.error('Failed to save verified linkage:', error);
  }
}

/**
 * Check if the current address/key pair has been previously verified
 * @param {string} address - The Ethereum address
 * @param {string} skavengerPublicKey - The Skavenge public key
 * @returns {boolean} True if this exact pairing was previously verified
 */
function isLinkageVerified(address, skavengerPublicKey) {
  try {
    const stored = localStorage.getItem('skavenge_verified_linkage');
    if (!stored) {
      return false;
    }

    const linkageData = JSON.parse(stored);
    return (
      linkageData.address === address.toLowerCase() &&
      linkageData.skavengerPublicKey === skavengerPublicKey
    );
  } catch (error) {
    console.error('Failed to check verified linkage:', error);
    return false;
  }
}

/**
 * Clear verified linkage from localStorage
 */
function clearVerifiedLinkage() {
  try {
    localStorage.removeItem('skavenge_verified_linkage');
  } catch (error) {
    console.error('Failed to clear verified linkage:', error);
  }
}

/**
 * Send linkage to the gateway server for verification and storage
 * @param {string} ethereumAddress - The Ethereum address
 * @param {string} skavengePublicKey - The Skavenge public key
 * @param {string} message - The signed message
 * @param {string} signature - The signature
 * @returns {Promise<{success: boolean, error?: string}>}
 */
async function sendLinkageToGateway(ethereumAddress, skavengePublicKey, message, signature) {
  try {
    // Load configuration to get gateway URL
    const config = await loadConfig();
    const gatewayUrl = getBrowserGatewayUrl(config.gatewayUrl);

    // Send POST request to /link endpoint
    const response = await fetch(`${gatewayUrl}/link`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        ethereum_address: ethereumAddress,
        skavenge_public_key: skavengePublicKey,
        message: message,
        signature: signature,
      }),
    });

    const data = await response.json();

    if (!response.ok) {
      // Handle different error cases
      if (response.status === 409) {
        // Address already linked to a different key
        return {
          success: false,
          error: 'This Ethereum address is already linked to a different Skavenge key. Each address can only be linked once.',
        };
      } else if (response.status === 401) {
        return {
          success: false,
          error: 'Signature verification failed on the server.',
        };
      } else {
        return {
          success: false,
          error: data.error || 'Failed to link account on the server.',
        };
      }
    }

    return { success: true };
  } catch (error) {
    console.error('Error sending linkage to gateway:', error);
    return {
      success: false,
      error: 'Failed to connect to the gateway server. Please try again later.',
    };
  }
}

/**
 * MetaMask Step Component
 * Handles connecting to MetaMask wallet and verifying ownership
 *
 * @param {object} props
 * @param {string} props.skavengerPublicKey - The Skavenge public key to link
 * @param {function} props.onAccountConnected - Callback when account is connected and verified
 * @param {function} props.onToast - Callback to show toast message
 */
function MetaMaskStep({ skavengerPublicKey, onAccountConnected, onToast }) {
  const [status, setStatus] = useState('disconnected'); // disconnected, connected
  const [statusText, setStatusText] = useState('Not connected');
  const [address, setAddress] = useState(null);

  // Check for existing connection on mount and when skavengerPublicKey loads
  useEffect(() => {
    checkInitialConnection();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [skavengerPublicKey]);

  // Listen for account changes
  useEffect(() => {
    if (!isMetaMaskInstalled()) {
      return;
    }

    const handleAccountsChanged = async (accounts) => {
      if (accounts.length > 0) {
        const account = accounts[0];

        // If this is the same account already connected, do nothing
        // This prevents double verification during initial connection
        if (address && address.toLowerCase() === account.toLowerCase()) {
          return;
        }

        // Require signature verification when account changes
        if (!skavengerPublicKey) {
          onToast('Cannot verify new account - Skavenge key not found', 'error');
          return;
        }

        // If no address is set yet, this is the initial connection
        // Let handleConnectClick handle the verification to avoid double prompts
        if (!address) {
          return;
        }

        // Account switched - check if this account/key pair is already verified
        if (isLinkageVerified(account, skavengerPublicKey)) {
          // Already verified, no need to re-verify
          setAddress(account);
          setStatus('connected');
          setStatusText(`Connected: ${account.substring(0, 6)}...${account.substring(38)}`);
          onAccountConnected(account);
          return;
        }

        try {
          onToast('Account changed - please sign to verify ownership', 'info');
          const { signature, message } = await signLinkageMessage(account, skavengerPublicKey);
          const isValid = verifySignature(message, signature, account);

          if (!isValid) {
            onToast('Signature verification failed for new account', 'error');
            return;
          }

          // Send linkage to gateway server for verification and storage
          onToast('Sending linkage to server for verification...', 'info');
          const result = await sendLinkageToGateway(account, skavengerPublicKey, message, signature);

          if (!result.success) {
            onToast(result.error, 'error');
            return;
          }

          // Save verified linkage
          saveVerifiedLinkage(account, skavengerPublicKey);

          setAddress(account);
          setStatus('connected');
          setStatusText(`Connected: ${account.substring(0, 6)}...${account.substring(38)}`);
          onToast('New account verified successfully', 'success');
          onAccountConnected(account);
        } catch (error) {
          onToast(error.message, 'error');
        }
      } else {
        // Disconnected - clear verified linkage
        clearVerifiedLinkage();
        setAddress(null);
        setStatus('disconnected');
        setStatusText('Not connected');
        onAccountConnected(null);
      }
    };

    const handleChainChanged = async (chainId) => {
      // Get expected chain ID from config
      const config = await loadConfig();
      const expectedChainId = config.chainId || CHAIN_ID.value;
      const currentChainIdDecimal = parseInt(chainId, 16);

      if (currentChainIdDecimal !== expectedChainId) {
        onToast(`Warning: Wrong network detected. Expected chain ID ${expectedChainId}, but connected to ${currentChainIdDecimal}. Please reconnect.`, 'error');
      } else {
        onToast('Network switched successfully', 'success');
      }
    };

    window.ethereum.on('accountsChanged', handleAccountsChanged);
    window.ethereum.on('chainChanged', handleChainChanged);

    return () => {
      window.ethereum.removeListener('accountsChanged', handleAccountsChanged);
      window.ethereum.removeListener('chainChanged', handleChainChanged);
    };
  }, [onAccountConnected, onToast, skavengerPublicKey, address]);

  const checkInitialConnection = async () => {
    const account = await getMetaMaskAccount();
    if (account && skavengerPublicKey) {
      // Check if this address/key pair has already been verified
      if (isLinkageVerified(account, skavengerPublicKey)) {
        // Already verified, no need to re-verify
        setAddress(account);
        setStatus('connected');
        setStatusText(`Connected: ${account.substring(0, 6)}...${account.substring(38)}`);
        onAccountConnected(account);
        return;
      }

      // Not yet verified, clear any old verification and skip on initial load
      // User will need to click "Connect MetaMask" to verify
      clearVerifiedLinkage();
    }
  };

  const handleConnectClick = async () => {
    try {
      // Check if Skavenge key is available
      if (!skavengerPublicKey) {
        onToast('Please complete Step 1 first (Link Skavenger Account)', 'error');
        return;
      }

      // Connect to MetaMask
      const account = await connectMetaMask();

      if (!account) {
        return;
      }

      // Request user to sign a message to prove ownership
      onToast('Please sign the message in MetaMask to verify ownership', 'info');
      const { signature, message } = await signLinkageMessage(account, skavengerPublicKey);

      // Verify the signature locally
      const isValid = verifySignature(message, signature, account);

      if (!isValid) {
        onToast('Signature verification failed. Unable to confirm address ownership.', 'error');
        return;
      }

      // Send linkage to gateway server for verification and storage
      onToast('Sending linkage to server for verification...', 'info');
      const result = await sendLinkageToGateway(account, skavengerPublicKey, message, signature);

      if (!result.success) {
        onToast(result.error, 'error');
        return;
      }

      // Signature verified successfully - save to localStorage
      saveVerifiedLinkage(account, skavengerPublicKey);

      setAddress(account);
      setStatus('connected');
      setStatusText(`Connected: ${account.substring(0, 6)}...${account.substring(38)}`);
      onToast('MetaMask connected and verified successfully', 'success');
      onAccountConnected(account);
    } catch (error) {
      onToast(error.message, 'error');
    }
  };

  return (
    <div className="step">
      <h2>Step 2: Connect MetaMask</h2>
      <p className="info">Connect your MetaMask wallet and sign a message to verify address ownership.</p>

      <div className="status">
        <span className={`status-indicator ${status === 'connected' ? 'connected' : ''}`}></span>
        <span>{statusText}</span>
      </div>

      <button className="primary" onClick={handleConnectClick}>
        {status === 'connected' ? 'Reconnect MetaMask' : 'Connect MetaMask'}
      </button>
    </div>
  );
}

export default MetaMaskStep;
