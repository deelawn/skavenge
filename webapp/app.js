// Skavenger Web App
// Connects to Skavenger extension and MetaMask

// Extension ID - hardcoded to match manifest.json
// This must match the extension_id field in skavenger-extension/manifest.json
// For unpacked extensions: find the ID in chrome://extensions and update both files
// For production: this will be the published Chrome Web Store extension ID
const SKAVENGER_EXTENSION_ID = "replace-with-your-extension-id"; // Must match manifest.json extension_id

let extensionId = SKAVENGER_EXTENSION_ID;

// Send message directly with extension ID (for internal use)
async function sendToExtensionDirect(id, message) {
    return new Promise((resolve, reject) => {
        chrome.runtime.sendMessage(id, message, (response) => {
            if (chrome.runtime.lastError) {
                reject(new Error(chrome.runtime.lastError.message));
            } else if (response && response.success !== false) {
                resolve(response);
            } else {
                reject(new Error(response?.error || 'Unknown error'));
            }
        });
    });
}

// Verify extension ID matches the hardcoded value
async function verifyExtensionId() {
    if (!SKAVENGER_EXTENSION_ID || SKAVENGER_EXTENSION_ID === "replace-with-your-extension-id") {
        return null;
    }
    
    try {
        const testResponse = await sendToExtensionDirect(SKAVENGER_EXTENSION_ID, { action: 'getExtensionId' });
        if (testResponse && testResponse.success && testResponse.extensionId) {
            // Verify the ID matches
            if (testResponse.extensionId === SKAVENGER_EXTENSION_ID) {
                return SKAVENGER_EXTENSION_ID;
            }
        }
    } catch (e) {
        // Extension not found or ID doesn't match
        console.error('Extension ID verification failed:', e);
    }
    
    return null;
}

// Send message to Skavenger extension
async function sendToExtension(message) {
    // Verify extension ID if not already verified
    if (!extensionId || extensionId === "replace-with-your-extension-id") {
        extensionId = await verifyExtensionId();
    }
    
    if (!extensionId || extensionId === "replace-with-your-extension-id") {
        throw new Error('Skavenger extension not found. Please make sure the extension ID is configured in both manifest.json and app.js');
    }
    
    return new Promise((resolve, reject) => {
        chrome.runtime.sendMessage(extensionId, message, (response) => {
            if (chrome.runtime.lastError) {
                // If "Could not establish connection", extension might not be installed or ID is wrong
                if (chrome.runtime.lastError.message.includes('connection')) {
                    // Reset extension ID and try to rediscover
                    extensionId = null;
                    reject(new Error('Could not connect to Skavenger extension. Please make sure the extension is installed.'));
                } else {
                    reject(new Error(chrome.runtime.lastError.message));
                }
            } else if (response && response.success !== false) {
                resolve(response);
            } else {
                reject(new Error(response?.error || 'Unknown error'));
            }
        });
    });
}

// State
let skavengerPublicKey = null;
let metamaskAddress = null;

// DOM Elements
const onboardingScreen = document.getElementById('onboarding');
const dashboardScreen = document.getElementById('dashboard');
const skavengerIndicator = document.getElementById('skavenger-indicator');
const skavengerStatusText = document.getElementById('skavenger-status-text');
const linkSkavengerBtn = document.getElementById('link-skavenger-btn');
const linkInstructions = document.getElementById('link-instructions');
const metamaskIndicator = document.getElementById('metamask-indicator');
const metamaskStatusText = document.getElementById('metamask-status-text');
const connectMetamaskBtn = document.getElementById('connect-metamask-btn');
const skavengerPublicKeyDisplay = document.getElementById('skavenger-public-key-display');
const metamaskAddressDisplay = document.getElementById('metamask-address-display');
const copySkavengerBtn = document.getElementById('copy-skavenger-btn');
const copyMetamaskBtn = document.getElementById('copy-metamask-btn');
const disconnectBtn = document.getElementById('disconnect-btn');
const toast = document.getElementById('toast');

// Utility functions
function showToast(message, type = 'info') {
    toast.textContent = message;
    toast.className = `toast ${type}`;
    toast.classList.remove('hidden');
    setTimeout(() => {
        toast.classList.add('hidden');
    }, 3000);
}

function showScreen(screen) {
    onboardingScreen.classList.add('hidden');
    dashboardScreen.classList.add('hidden');
    screen.classList.remove('hidden');
}

// Check if MetaMask is installed
function isMetaMaskInstalled() {
    return typeof window.ethereum !== 'undefined' && window.ethereum.isMetaMask;
}

// Connect to MetaMask
async function connectMetaMask() {
    if (!isMetaMaskInstalled()) {
        showToast('MetaMask is not installed. Please install MetaMask extension.', 'error');
        return null;
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
            showToast('MetaMask connection rejected.', 'error');
        } else {
            showToast('Failed to connect to MetaMask.', 'error');
        }
        return null;
    }
}

// Get MetaMask account
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

// Check Skavenger extension status
async function checkSkavengerKeys() {
    try {
        const response = await sendToExtension({ action: 'hasKeys' });
        return response.hasKeys;
    } catch (error) {
        console.error('Error checking Skavenger keys:', error);
        return false;
    }
}

// Get Skavenger public key
async function getSkavengerPublicKey() {
    try {
        // First verify session or password
        const verifyResponse = await sendToExtension({ action: 'verifyPassword' });
        
        if (!verifyResponse.success) {
            // Need password
            return null;
        }
        
        const response = await sendToExtension({ action: 'exportPublicKey' });
        if (response.success) {
            return response.publicKey;
        }
        
        return null;
    } catch (error) {
        console.error('Error getting Skavenger public key:', error);
        return null;
    }
}

// Request link to Skavenger extension
async function requestSkavengerLink() {
    try {
        // Send message to extension to notify it needs attention
        await sendToExtension({ action: 'requestLink' });
        return true;
    } catch (error) {
        // If extension is not found, that's okay - user will need to install it
        console.error('Error requesting link:', error);
        return false;
    }
}

// Poll for keys after user clicks link button
let keyPollInterval = null;

function startKeyPolling() {
    // Clear any existing interval
    if (keyPollInterval) {
        clearInterval(keyPollInterval);
    }
    
    let pollCount = 0;
    const maxPolls = 60; // Poll for up to 60 seconds (60 * 1 second)
    
    keyPollInterval = setInterval(async () => {
        pollCount++;
        
        try {
            const hasKeys = await checkSkavengerKeys();
            if (hasKeys) {
                const publicKey = await getSkavengerPublicKey();
                if (publicKey) {
                    // Keys found!
                    clearInterval(keyPollInterval);
                    keyPollInterval = null;
                    skavengerPublicKey = publicKey;
                    skavengerIndicator.classList.add('connected');
                    skavengerStatusText.textContent = 'Keys found';
                    linkSkavengerBtn.classList.add('hidden');
                    linkInstructions.classList.add('hidden');
                    showToast('Skavenger account linked successfully', 'success');
                    
                    // Check if MetaMask is also connected
                    if (metamaskAddress) {
                        showDashboard();
                    }
                }
            }
            
            // Stop polling after max attempts
            if (pollCount >= maxPolls) {
                clearInterval(keyPollInterval);
                keyPollInterval = null;
            }
        } catch (error) {
            // Continue polling on error
        }
    }, 1000); // Poll every second
}

// Initialize onboarding
async function initOnboarding() {
    // Check Skavenger extension
    try {
        const hasKeys = await checkSkavengerKeys();
        
        if (hasKeys) {
            // Try to get public key (may need password if session expired)
            const publicKey = await getSkavengerPublicKey();
            
            if (publicKey) {
                skavengerPublicKey = publicKey;
                skavengerIndicator.classList.add('connected');
                skavengerStatusText.textContent = 'Keys found';
                linkSkavengerBtn.classList.add('hidden');
                linkInstructions.classList.add('hidden');
            } else {
                // Session expired, need to link again
                skavengerIndicator.classList.remove('connected');
                skavengerStatusText.textContent = 'Session expired - please link again';
                linkSkavengerBtn.classList.remove('hidden');
                linkInstructions.classList.add('hidden');
            }
        } else {
            skavengerIndicator.classList.remove('connected');
            skavengerStatusText.textContent = 'No keys found';
            linkSkavengerBtn.classList.remove('hidden');
            linkInstructions.classList.add('hidden');
        }
    } catch (error) {
        skavengerIndicator.classList.remove('connected');
        skavengerStatusText.textContent = error.message || 'Extension not found - please install the Skavenger extension';
        linkSkavengerBtn.classList.add('hidden');
        linkInstructions.classList.add('hidden');
    }
    
    // Check MetaMask
    const account = await getMetaMaskAccount();
    if (account) {
        metamaskAddress = account;
        metamaskIndicator.classList.add('connected');
        metamaskStatusText.textContent = `Connected: ${account.substring(0, 6)}...${account.substring(38)}`;
        connectMetamaskBtn.textContent = 'Reconnect MetaMask';
    } else {
        metamaskIndicator.classList.remove('connected');
        metamaskStatusText.textContent = 'Not connected';
        connectMetamaskBtn.textContent = 'Connect MetaMask';
    }
    
    // Check if both are connected
    if (skavengerPublicKey && metamaskAddress) {
        showDashboard();
    }
}

// Show dashboard - always fetch fresh data from extension and MetaMask
async function showDashboard() {
    // Always fetch fresh data from extension
    try {
        const publicKey = await getSkavengerPublicKey();
        if (publicKey) {
            skavengerPublicKey = publicKey;
            skavengerPublicKeyDisplay.textContent = publicKey;
        } else {
            skavengerPublicKeyDisplay.textContent = 'Not available';
            // If we can't get the key, go back to onboarding
            showScreen(onboardingScreen);
            initOnboarding();
            return;
        }
    } catch (error) {
        skavengerPublicKeyDisplay.textContent = 'Not available';
        // If we can't get the key, go back to onboarding
        showScreen(onboardingScreen);
        initOnboarding();
        return;
    }
    
    // Always fetch fresh data from MetaMask
    try {
        const account = await getMetaMaskAccount();
        if (account) {
            metamaskAddress = account;
            metamaskAddressDisplay.textContent = account;
        } else {
            metamaskAddressDisplay.textContent = 'Not connected';
            // If MetaMask is not connected, go back to onboarding
            showScreen(onboardingScreen);
            initOnboarding();
            return;
        }
    } catch (error) {
        metamaskAddressDisplay.textContent = 'Not connected';
        // If MetaMask is not connected, go back to onboarding
        showScreen(onboardingScreen);
        initOnboarding();
        return;
    }
    
    // Only show dashboard if both are successfully loaded
    showScreen(dashboardScreen);
}

// Event handlers
linkSkavengerBtn.addEventListener('click', async () => {
    // Show instructions
    linkInstructions.classList.remove('hidden');
    skavengerStatusText.textContent = 'Waiting for you to set up Skavenger extension...';
    skavengerIndicator.classList.add('pending');
    
    // Request link from extension (this may fail if extension not found, but that's okay)
    try {
        await requestSkavengerLink();
        showToast('Please open the Skavenger extension to set up or unlock your account', 'info');
    } catch (error) {
        // Extension might not be found, but show instructions anyway
        showToast('Please open the Skavenger extension icon in your browser toolbar', 'info');
    }
    
    // Start polling for keys
    startKeyPolling();
});

connectMetamaskBtn.addEventListener('click', async () => {
    try {
        const account = await connectMetaMask();
        
        if (account) {
            metamaskAddress = account;
            metamaskIndicator.classList.add('connected');
            metamaskStatusText.textContent = `Connected: ${account.substring(0, 6)}...${account.substring(38)}`;
            showToast('MetaMask connected successfully', 'success');
            
            // Check if Skavenger is also connected
            if (skavengerPublicKey) {
                showDashboard();
            }
        }
    } catch (error) {
        showToast(error.message, 'error');
    }
});

copySkavengerBtn.addEventListener('click', () => {
    if (skavengerPublicKey) {
        navigator.clipboard.writeText(skavengerPublicKey);
        showToast('Skavenger public key copied', 'success');
    }
});

copyMetamaskBtn.addEventListener('click', () => {
    if (metamaskAddress) {
        navigator.clipboard.writeText(metamaskAddress);
        showToast('MetaMask address copied', 'success');
    }
});

disconnectBtn.addEventListener('click', () => {
    // Stop any polling
    if (keyPollInterval) {
        clearInterval(keyPollInterval);
        keyPollInterval = null;
    }
    
    // Clear in-memory state (but not localStorage - we're not storing keys there)
    metamaskAddress = null;
    skavengerPublicKey = null;
    showScreen(onboardingScreen);
    initOnboarding();
});

// Handle MetaMask account changes
if (isMetaMaskInstalled()) {
    window.ethereum.on('accountsChanged', (accounts) => {
        if (accounts.length > 0) {
            metamaskAddress = accounts[0];
            if (dashboardScreen.classList.contains('hidden')) {
                // Update status in onboarding
                metamaskIndicator.classList.add('connected');
                metamaskStatusText.textContent = `Connected: ${accounts[0].substring(0, 6)}...${accounts[0].substring(38)}`;
                if (skavengerPublicKey) {
                    showDashboard();
                }
            } else {
                // Update dashboard
                metamaskAddressDisplay.textContent = accounts[0];
            }
        } else {
            // Disconnected
            metamaskAddress = null;
            if (!dashboardScreen.classList.contains('hidden')) {
                showScreen(onboardingScreen);
                initOnboarding();
            }
        }
    });
}

// Initialize - verify extension ID and fetch fresh data from extension and MetaMask
(async () => {
    // Verify extension ID matches hardcoded value
    extensionId = await verifyExtensionId();
    // Then initialize onboarding
    initOnboarding();
})();

