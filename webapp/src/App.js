import React, { useState, useEffect } from 'react';
import './App.css';
import { loadConfig } from './config';
import Toast from './components/Toast';
import SkavengerStep from './components/SkavengerStep';
import MetaMaskStep from './components/MetaMaskStep';
import Dashboard from './components/Dashboard';

/**
 * Main App Component
 * Manages the onboarding flow and dashboard
 */
function App() {
  const [screen, setScreen] = useState('onboarding'); // onboarding, dashboard
  const [skavengerPublicKey, setSkavengerPublicKey] = useState(null);
  const [metamaskAddress, setMetamaskAddress] = useState(null);
  const [config, setConfig] = useState(null);

  // Toast state
  const [toastMessage, setToastMessage] = useState('');
  const [toastType, setToastType] = useState('info');
  const [toastVisible, setToastVisible] = useState(false);

  // Load configuration on mount
  useEffect(() => {
    const initConfig = async () => {
      const cfg = await loadConfig();
      setConfig(cfg);
    };

    initConfig();
  }, []);

  // Show toast message
  const showToast = (message, type = 'info') => {
    setToastMessage(message);
    setToastType(type);
    setToastVisible(true);

    setTimeout(() => {
      setToastVisible(false);
    }, 3000);
  };

  // Handle Skavenger keys found
  const handleSkavengerKeysFound = (publicKey) => {
    setSkavengerPublicKey(publicKey);

    // If MetaMask is also connected, show dashboard
    if (metamaskAddress) {
      setScreen('dashboard');
    }
  };

  // Handle MetaMask account connected
  const handleMetaMaskConnected = (address) => {
    if (address) {
      setMetamaskAddress(address);

      // If Skavenger is also connected, show dashboard
      if (skavengerPublicKey) {
        setScreen('dashboard');
      }
    } else {
      // Disconnected - go back to onboarding
      setMetamaskAddress(null);
      if (screen === 'dashboard') {
        setScreen('onboarding');
      }
    }
  };

  return (
    <div className="container">
      <header>
        <h1>Skavenger</h1>
        <p className="subtitle">Secure key management for Skavenge</p>
      </header>

      {screen === 'onboarding' && (
        <div className="screen">
          <SkavengerStep
            onKeysFound={handleSkavengerKeysFound}
            onToast={showToast}
          />
          <MetaMaskStep
            onAccountConnected={handleMetaMaskConnected}
            onToast={showToast}
          />
        </div>
      )}

      {screen === 'dashboard' && (
        <div className="screen">
          <Dashboard
            skavengerPublicKey={skavengerPublicKey}
            metamaskAddress={metamaskAddress}
            config={config}
            onToast={showToast}
          />
        </div>
      )}

      <Toast message={toastMessage} type={toastType} visible={toastVisible} />
    </div>
  );
}

export default App;
