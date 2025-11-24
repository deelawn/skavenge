import React, { useState, useEffect } from 'react';

/**
 * Check if MetaMask is installed
 */
function isMetaMaskInstalled() {
  return typeof window.ethereum !== 'undefined' && window.ethereum.isMetaMask;
}

/**
 * Connect to MetaMask
 */
async function connectMetaMask() {
  if (!isMetaMaskInstalled()) {
    throw new Error('MetaMask is not installed. Please install MetaMask extension.');
  }

  try {
    const accounts = await window.ethereum.request({
      method: 'eth_requestAccounts'
    });

    if (accounts && accounts.length > 0) {
      return accounts[0];
    }

    return null;
  } catch (error) {
    if (error.code === 4001) {
      throw new Error('MetaMask connection rejected.');
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

    window.ethereum.on('accountsChanged', handleAccountsChanged);

    return () => {
      window.ethereum.removeListener('accountsChanged', handleAccountsChanged);
    };
  }, [onAccountConnected]);

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
