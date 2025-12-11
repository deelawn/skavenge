import React, { useState, useEffect } from 'react';
import './App.css';
import { loadConfig } from './config';
import { checkLinkageOnGateway } from './utils';
import Toast from './components/Toast';
import SkavengerStep from './components/SkavengerStep';
import MetaMaskStep from './components/MetaMaskStep';
import Dashboard from './components/Dashboard';
import RelinkPrompt from './components/RelinkPrompt';

/**
 * Main App Component
 * Manages the onboarding flow and dashboard
 */
function App() {
  const [screen, setScreen] = useState('onboarding'); // onboarding, dashboard, relink
  const [skavengerPublicKey, setSkavengerPublicKey] = useState(null);
  const [metamaskAddress, setMetamaskAddress] = useState(null);
  const [config, setConfig] = useState(null);

  // Linkage verification state
  const [linkageVerified, setLinkageVerified] = useState(false);
  const [isCheckingLinkage, setIsCheckingLinkage] = useState(false);

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

  // Check linkage on gateway when both accounts are connected
  useEffect(() => {
    const validateLinkage = async () => {
      if (!skavengerPublicKey || !metamaskAddress) {
        return;
      }

      setIsCheckingLinkage(true);

      try {
        const result = await checkLinkageOnGateway(metamaskAddress);

        if (result.exists) {
          // Verify that the server-side linkage matches the local one
          if (result.skavengePublicKey === skavengerPublicKey) {
            setLinkageVerified(true);
            setScreen('dashboard');
          } else {
            // Server has a different key linked - show error
            showToast(
              'This Ethereum address is already linked to a different Skavenge key on the server',
              'error'
            );
            setLinkageVerified(false);
            setScreen('relink');
          }
        } else {
          // Linkage doesn't exist on server - need to re-link
          setLinkageVerified(false);
          setScreen('relink');
        }
      } catch (error) {
        console.error('Error validating linkage:', error);
        showToast('Failed to validate linkage with server', 'error');
        setLinkageVerified(false);
        setScreen('relink');
      } finally {
        setIsCheckingLinkage(false);
      }
    };

    validateLinkage();
  }, [skavengerPublicKey, metamaskAddress]);

  // Handle MetaMask disconnection
  useEffect(() => {
    if (!metamaskAddress && screen === 'dashboard') {
      // Switch back to onboarding if MetaMask disconnects
      setScreen('onboarding');
      setLinkageVerified(false);
    }
  }, [metamaskAddress, screen]);

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
    // Dashboard switch is handled by useEffect
  };

  // Handle MetaMask account connected
  const handleMetaMaskConnected = (address) => {
    setMetamaskAddress(address);
    // Dashboard switch is handled by useEffect
  };

  // Handle successful re-linking
  const handleRelinkSuccess = () => {
    setLinkageVerified(true);
    setScreen('dashboard');
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
            skavengerPublicKey={skavengerPublicKey}
            onAccountConnected={handleMetaMaskConnected}
            onToast={showToast}
          />
        </div>
      )}

      {screen === 'relink' && (
        <RelinkPrompt
          metamaskAddress={metamaskAddress}
          skavengerPublicKey={skavengerPublicKey}
          onSuccess={handleRelinkSuccess}
          onToast={showToast}
        />
      )}

      {screen === 'dashboard' && linkageVerified && (
        <div className="screen">
          <Dashboard
            skavengerPublicKey={skavengerPublicKey}
            metamaskAddress={metamaskAddress}
            config={config}
            onToast={showToast}
          />
        </div>
      )}

      {isCheckingLinkage && (
        <div className="screen">
          <div className="step">
            <h2>Verifying Account Linkage...</h2>
            <p className="info">Checking your account linkage with the server. Please wait.</p>
          </div>
        </div>
      )}

      <Toast message={toastMessage} type={toastType} visible={toastVisible} />
    </div>
  );
}

export default App;
