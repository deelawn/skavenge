import React, { useState } from 'react';
import { getBrowserRpcUrl } from '../utils.js';
import TokenDisplay from './TokenDisplay';
import ClueMarketplace from './ClueMarketplace';
import Transfers from './Transfers';

/**
 * Dashboard Component
 * Displays connected accounts, configuration, and user's NFT tokens
 *
 * @param {object} props
 * @param {string} props.skavengerPublicKey - The Skavenger public key
 * @param {string} props.metamaskAddress - The MetaMask address
 * @param {object} props.config - Configuration object with contractAddress and networkRpcUrl
 * @param {function} props.onToast - Callback to show toast message
 */
function Dashboard({ skavengerPublicKey, metamaskAddress, config, onToast }) {
  const [currentView, setCurrentView] = useState('myTokens'); // 'myTokens', 'marketplace', or 'transfers'
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
    <div className="dashboard-layout">
      {/* Left Sidebar Navigation */}
      <div className="dashboard-sidebar">
        <h2>Account Linkage</h2>

        <div className="account-card">
          <div className="account-header">
            <span className="account-label">Linked Accounts</span>
            <span className="account-status connected">Verified</span>
          </div>

          <div style={{ marginTop: '12px', marginBottom: '12px' }}>
            <div style={{
              fontSize: '11px',
              color: '#718096',
              marginBottom: '6px',
              fontWeight: '600',
              textTransform: 'uppercase',
              letterSpacing: '0.5px'
            }}>
              Ethereum Address
            </div>
            <div className="account-value" style={{ marginBottom: '12px' }}>
              {metamaskAddress || 'Loading...'}
            </div>
            <button
              className="copy-btn"
              onClick={handleCopyMetamaskAddress}
              style={{ marginBottom: '16px' }}
            >
              Copy Address
            </button>

            <div style={{
              textAlign: 'center',
              fontSize: '20px',
              color: '#667eea',
              margin: '12px 0',
              fontWeight: 'bold'
            }}>
              ↓
            </div>

            <div style={{
              fontSize: '11px',
              color: '#718096',
              marginBottom: '6px',
              fontWeight: '600',
              textTransform: 'uppercase',
              letterSpacing: '0.5px'
            }}>
              Skavenge Public Key
            </div>
            <div className="account-value" style={{ marginBottom: '12px' }}>
              {skavengerPublicKey || 'Loading...'}
            </div>
            <button className="copy-btn" onClick={handleCopySkavengerKey}>
              Copy Key
            </button>
          </div>
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
              <span className="config-value">{getBrowserRpcUrl(config.networkRpcUrl)}</span>
            </div>
            <div className="config-row">
              <span className="config-label">Chain ID</span>
              <span className="config-value">{config.chainId || 'N/A'}</span>
            </div>
          </div>
        )}
      </div>

      {/* Main Content Area - Token Display or Marketplace */}
      <div className="dashboard-main">
        {/* View Toggle Navigation */}
        <div className="view-toggle" style={{
          display: 'flex',
          gap: '12px',
          marginBottom: '20px',
          borderBottom: '2px solid #e2e8f0',
          paddingBottom: '0'
        }}>
          <button
            onClick={() => setCurrentView('myTokens')}
            className={`view-toggle-btn ${currentView === 'myTokens' ? 'active' : ''}`}
            style={{
              padding: '12px 24px',
              fontSize: '16px',
              fontWeight: '600',
              backgroundColor: 'transparent',
              color: currentView === 'myTokens' ? '#667eea' : '#718096',
              border: 'none',
              borderBottom: currentView === 'myTokens' ? '3px solid #667eea' : '3px solid transparent',
              cursor: 'pointer',
              transition: 'all 0.2s',
              marginBottom: '-2px'
            }}
          >
            My Tokens
          </button>
          <button
            onClick={() => setCurrentView('marketplace')}
            className={`view-toggle-btn ${currentView === 'marketplace' ? 'active' : ''}`}
            style={{
              padding: '12px 24px',
              fontSize: '16px',
              fontWeight: '600',
              backgroundColor: 'transparent',
              color: currentView === 'marketplace' ? '#667eea' : '#718096',
              border: 'none',
              borderBottom: currentView === 'marketplace' ? '3px solid #667eea' : '3px solid transparent',
              cursor: 'pointer',
              transition: 'all 0.2s',
              marginBottom: '-2px'
            }}
          >
            Marketplace
          </button>
          <button
            onClick={() => setCurrentView('transfers')}
            className={`view-toggle-btn ${currentView === 'transfers' ? 'active' : ''}`}
            style={{
              padding: '12px 24px',
              fontSize: '16px',
              fontWeight: '600',
              backgroundColor: 'transparent',
              color: currentView === 'transfers' ? '#667eea' : '#718096',
              border: 'none',
              borderBottom: currentView === 'transfers' ? '3px solid #667eea' : '3px solid transparent',
              cursor: 'pointer',
              transition: 'all 0.2s',
              marginBottom: '-2px'
            }}
          >
            Transfers
          </button>
        </div>

        {/* Conditional Rendering Based on Current View */}
        {currentView === 'myTokens' ? (
          <TokenDisplay
            metamaskAddress={metamaskAddress}
            config={config}
            onToast={onToast}
          />
        ) : currentView === 'marketplace' ? (
          <ClueMarketplace
            metamaskAddress={metamaskAddress}
            config={config}
            onToast={onToast}
          />
        ) : (
          <Transfers
            metamaskAddress={metamaskAddress}
            config={config}
            onToast={onToast}
          />
        )}
      </div>
    </div>
  );
}

export default Dashboard;
