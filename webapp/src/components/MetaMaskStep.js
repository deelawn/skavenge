import React, { useState, useEffect } from 'react';
import { loadConfig, CHAIN_ID } from '../config.js';

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
              rpcUrls: [rpcUrl],
              nativeCurrency: {
                name: 'Ether',
                symbol: 'ETH',
                decimals: 18,
              },
            },
          ],
        });
      } catch (addError) {
        throw new Error('Failed to add network to MetaMask.');
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
 * MetaMask Step Component
 * Handles connecting to MetaMask wallet
 *
 * @param {object} props
 * @param {function} props.onAccountConnected - Callback when account is connected
 * @param {function} props.onToast - Callback to show toast message
 */
function MetaMaskStep({ onAccountConnected, onToast }) {
  const [status, setStatus] = useState('disconnected'); // disconnected, connected
  const [statusText, setStatusText] = useState('Not connected');
  const [address, setAddress] = useState(null);

  // Check for existing connection on mount
  useEffect(() => {
    checkInitialConnection();
  }, []);

  // Listen for account changes
  useEffect(() => {
    if (!isMetaMaskInstalled()) {
      return;
    }

    const handleAccountsChanged = (accounts) => {
      if (accounts.length > 0) {
        const account = accounts[0];
        setAddress(account);
        setStatus('connected');
        setStatusText(`Connected: ${account.substring(0, 6)}...${account.substring(38)}`);
        onAccountConnected(account);
      } else {
        // Disconnected
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
  }, [onAccountConnected, onToast]);

  const checkInitialConnection = async () => {
    const account = await getMetaMaskAccount();
    if (account) {
      setAddress(account);
      setStatus('connected');
      setStatusText(`Connected: ${account.substring(0, 6)}...${account.substring(38)}`);
      onAccountConnected(account);
    }
  };

  const handleConnectClick = async () => {
    try {
      const account = await connectMetaMask();

      if (account) {
        setAddress(account);
        setStatus('connected');
        setStatusText(`Connected: ${account.substring(0, 6)}...${account.substring(38)}`);
        onToast('MetaMask connected successfully', 'success');
        onAccountConnected(account);
      }
    } catch (error) {
      onToast(error.message, 'error');
    }
  };

  return (
    <div className="step">
      <h2>Step 2: Connect MetaMask</h2>
      <p className="info">Connect your MetaMask wallet for transaction signing.</p>

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
