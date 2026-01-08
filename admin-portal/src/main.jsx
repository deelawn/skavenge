import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import './index.css';
import { loadConfig } from './utils/config';
import { initWeb3 } from './utils/web3';

const init = async () => {
  try {
    await loadConfig();
    initWeb3();

    ReactDOM.createRoot(document.getElementById('root')).render(
      <React.StrictMode>
        <App />
      </React.StrictMode>
    );
  } catch (error) {
    console.error('Failed to initialize application:', error);
    document.getElementById('root').innerHTML = `
      <div style="display: flex; align-items: center; justify-content: center; height: 100vh; font-family: sans-serif;">
        <div style="text-align: center; padding: 2rem; background: #fee; border: 1px solid #fcc; border-radius: 8px; max-width: 500px;">
          <h1 style="color: #c00; margin-bottom: 1rem;">Failed to Load Application</h1>
          <p style="color: #600;">${error.message}</p>
          <p style="color: #600; margin-top: 1rem;">Please check that the config.json file is available and properly formatted.</p>
        </div>
      </div>
    `;
  }
};

init();
