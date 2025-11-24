import React from 'react';

/**
 * Dashboard Component
 * Displays connected accounts and configuration information
 *
 * @param {object} props
 * @param {string} props.skavengerPublicKey - The Skavenger public key
 * @param {string} props.metamaskAddress - The MetaMask address
 * @param {object} props.config - Configuration object with contractAddress and networkRpcUrl
 * @param {function} props.onToast - Callback to show toast message
 */
function Dashboard({ skavengerPublicKey, metamaskAddress, config, onToast }) {
  const handleCopySkavengerKey = () => {
    if (skavengerPublicKey) {
      navigator.clipboard.writeText(skavengerPublicKey);
      onToast('Skavenger public key copied', 'success');
    }
  };

  const handleCopyMetamaskAddress = () => {
    if (metamaskAddress) {
      navigator.clipboard.writeText(metamaskAddress);
      onToast('MetaMask address copied', 'success');
    }
  };

  return (
    <div className="dashboard">
      <h2>Your Accounts</h2>

      <div className="account-card">
        <div className="account-header">
          <span className="account-label">Skavenger Public Key</span>
          <span className="account-status connected">Connected</span>
        </div>
        <div className="account-value">{skavengerPublicKey || 'Loading...'}</div>
        <button className="copy-btn" onClick={handleCopySkavengerKey}>
          Copy
        </button>
      </div>

      <div className="account-card">
        <div className="account-header">
          <span className="account-label">MetaMask Wallet</span>
          <span className="account-status connected">Connected</span>
        </div>
        <div className="account-value">{metamaskAddress || 'Loading...'}</div>
        <button className="copy-btn" onClick={handleCopyMetamaskAddress}>
          Copy
        </button>
      </div>

      {config && (
        <div className="config-info">
          <h3>Configuration</h3>
          <div className="config-row">
            <span className="config-label">Contract Address</span>
            <span className="config-value">{config.contractAddress}</span>
          </div>
          <div className="config-row">
            <span className="config-label">Network RPC URL</span>
            <span className="config-value">{config.networkRpcUrl}</span>
          </div>
          <div className="config-row">
            <span className="config-label">Chain ID</span>
            <span className="config-value">{config.chainId || 'N/A'}</span>
          </div>
        </div>
      )}
    </div>
  );
}

export default Dashboard;
