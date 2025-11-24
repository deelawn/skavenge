// DOM Elements
const unlockScreen = document.getElementById('unlock-screen');
const setupSection = document.getElementById('setup-section');
const loginSection = document.getElementById('login-section');
const mainScreen = document.getElementById('main-screen');

const newPasswordInput = document.getElementById('new-password');
const confirmPasswordInput = document.getElementById('confirm-password');
const createPasswordBtn = document.getElementById('create-password-btn');

const loginPasswordInput = document.getElementById('login-password');
const unlockBtn = document.getElementById('unlock-btn');
const resetBtn = document.getElementById('reset-btn');

const keyIndicator = document.getElementById('key-indicator');
const keyStatusText = document.getElementById('key-status-text');
const publicKeySection = document.getElementById('public-key-section');
const publicKeyDisplay = document.getElementById('public-key-display');
const copyPublicKeyBtn = document.getElementById('copy-public-key-btn');

const generateBtn = document.getElementById('generate-btn');
const exportBtn = document.getElementById('export-btn');
const importBtn = document.getElementById('import-btn');
const lockBtn = document.getElementById('lock-btn');

const exportModal = document.getElementById('export-modal');
const exportOutput = document.getElementById('export-output');
const copyExportBtn = document.getElementById('copy-export-btn');
const downloadExportBtn = document.getElementById('download-export-btn');
const closeExportBtn = document.getElementById('close-export-btn');

const importModal = document.getElementById('import-modal');
const importKeysInput = document.getElementById('import-keys');
const confirmImportBtn = document.getElementById('confirm-import-btn');
const closeImportBtn = document.getElementById('close-import-btn');

const confirmModal = document.getElementById('confirm-modal');
const confirmTitle = document.getElementById('confirm-title');
const confirmMessage = document.getElementById('confirm-message');
const confirmYesBtn = document.getElementById('confirm-yes-btn');
const confirmNoBtn = document.getElementById('confirm-no-btn');

const toast = document.getElementById('toast');

// State
let currentPassword = null;
let hasKeys = false;

// Utility functions
function showToast(message, isError = false) {
  toast.textContent = message;
  toast.className = isError ? 'toast error' : 'toast';
  setTimeout(() => toast.classList.add('hidden'), 3000);
}

function sendMessage(message) {
  return chrome.runtime.sendMessage(message);
}

function showScreen(screen) {
  unlockScreen.classList.add('hidden');
  mainScreen.classList.add('hidden');
  screen.classList.remove('hidden');
}

async function updateKeyStatus(keysExist, publicKey = null) {
  hasKeys = keysExist;
  if (keysExist) {
    keyIndicator.classList.add('active');
    keyStatusText.textContent = 'Keys loaded';
    exportBtn.disabled = false;

    if (publicKey) {
      publicKeyDisplay.value = publicKey;
      publicKeySection.classList.remove('hidden');
    } else {
      // Fetch public key if not provided (no password required - public keys are public)
      const result = await sendMessage({ action: 'getPublicKey' });
      if (result.success) {
        publicKeyDisplay.value = result.publicKey;
        publicKeySection.classList.remove('hidden');
      }
    }
  } else {
    keyIndicator.classList.remove('active');
    keyStatusText.textContent = 'No keys generated';
    exportBtn.disabled = true;
    publicKeySection.classList.add('hidden');
    publicKeyDisplay.value = '';
  }
}

async function showConfirmDialog(title, message) {
  return new Promise((resolve) => {
    confirmTitle.textContent = title;
    confirmMessage.textContent = message;
    confirmModal.classList.remove('hidden');

    const handleYes = () => {
      confirmModal.classList.add('hidden');
      confirmYesBtn.removeEventListener('click', handleYes);
      confirmNoBtn.removeEventListener('click', handleNo);
      resolve(true);
    };

    const handleNo = () => {
      confirmModal.classList.add('hidden');
      confirmYesBtn.removeEventListener('click', handleYes);
      confirmNoBtn.removeEventListener('click', handleNo);
      resolve(false);
    };

    confirmYesBtn.addEventListener('click', handleYes);
    confirmNoBtn.addEventListener('click', handleNo);
  });
}

// Initialize
async function init() {
  // Clear badge when popup opens
  try {
    await chrome.action.setBadgeText({ text: '' });
  } catch (error) {
    // Ignore if badge API is not available
  }

  const result = await sendMessage({ action: 'hasKeys' });

  if (result.hasKeys) {
    // Check if there's a valid session
    const sessionCheck = await sendMessage({ action: 'verifyPassword' });

    if (sessionCheck.success) {
      // Valid session exists - auto-unlock and go to main screen
      currentPassword = true; // Mark as unlocked (we don't need the actual password in popup)
      showScreen(mainScreen);
      await updateKeyStatus(true);
    } else {
      // No valid session - show login screen
      loginSection.classList.remove('hidden');
      setupSection.classList.add('hidden');
      showScreen(unlockScreen);
    }
  } else {
    // No keys - show setup screen
    setupSection.classList.remove('hidden');
    loginSection.classList.add('hidden');
    showScreen(unlockScreen);
  }
}

// Event handlers
createPasswordBtn.addEventListener('click', async () => {
  const password = newPasswordInput.value;
  const confirm = confirmPasswordInput.value;

  if (password.length < 8) {
    showToast('Password must be at least 8 characters', true);
    return;
  }

  if (password !== confirm) {
    showToast('Passwords do not match', true);
    return;
  }

  currentPassword = password;
  showScreen(mainScreen);
  updateKeyStatus(false);
  showToast('Password created');
});

unlockBtn.addEventListener('click', async () => {
  const password = loginPasswordInput.value;

  const result = await sendMessage({ action: 'verifyPassword', password });

  if (result.success) {
    currentPassword = password;
    showScreen(mainScreen);
    await updateKeyStatus(true);
    loginPasswordInput.value = '';
  } else {
    showToast('Invalid password', true);
  }
});

resetBtn.addEventListener('click', async () => {
  const confirmed = await showConfirmDialog(
    'Reset Extension',
    'This will delete all stored keys. This action cannot be undone.'
  );

  if (confirmed) {
    await sendMessage({ action: 'clearKeys' });
    currentPassword = null;
    hasKeys = false;
    setupSection.classList.remove('hidden');
    loginSection.classList.add('hidden');
    showToast('Extension reset');
  }
});

generateBtn.addEventListener('click', async () => {
  if (hasKeys) {
    const confirmed = await showConfirmDialog(
      'Generate New Keys',
      'This will replace your existing keys. Make sure you have exported them first.'
    );
    if (!confirmed) return;
  }

  const result = await sendMessage({ action: 'generateKeys', password: currentPassword });

  if (result.success) {
    await updateKeyStatus(true, result.publicKey);
    showToast('Keys generated');
  } else {
    showToast(result.error || 'Failed to generate keys', true);
  }
});

exportBtn.addEventListener('click', async () => {
  if (!hasKeys) return;

  const result = await sendMessage({
    action: 'exportKeys',
    password: currentPassword
  });

  if (result.success) {
    exportOutput.value = result.keyData;
    exportModal.classList.remove('hidden');
  } else {
    showToast(result.error || 'Failed to export keys', true);
  }
});

copyExportBtn.addEventListener('click', () => {
  if (!exportOutput.value) {
    showToast('Generate keystore first', true);
    return;
  }
  navigator.clipboard.writeText(exportOutput.value);
  showToast('Copied to clipboard');
});

downloadExportBtn.addEventListener('click', () => {
  if (!exportOutput.value) {
    showToast('Generate keystore first', true);
    return;
  }
  const blob = new Blob([exportOutput.value], { type: 'application/json' });
  const url = URL.createObjectURL(blob);
  const a = document.createElement('a');
  a.href = url;
  a.download = 'skavenger-keys.json';
  a.click();
  URL.revokeObjectURL(url);
  showToast('Downloaded');
});

closeExportBtn.addEventListener('click', () => {
  exportModal.classList.add('hidden');
  exportOutput.value = '';
});

importBtn.addEventListener('click', () => {
  importKeysInput.value = '';
  importModal.classList.remove('hidden');
});

confirmImportBtn.addEventListener('click', async () => {
  const keyData = importKeysInput.value.trim();

  if (!keyData) {
    showToast('Keys JSON is required', true);
    return;
  }

  if (hasKeys) {
    const confirmed = await showConfirmDialog(
      'Import Keys',
      'This will replace your existing keys. Make sure you have exported them first.'
    );
    if (!confirmed) return;
  }

  const result = await sendMessage({
    action: 'importKeys',
    keyData,
    password: currentPassword
  });

  if (result.success) {
    await updateKeyStatus(true);
    importModal.classList.add('hidden');
    showToast('Keys imported');
  } else {
    showToast(result.error || 'Invalid keys format', true);
  }
});

copyPublicKeyBtn.addEventListener('click', () => {
  navigator.clipboard.writeText(publicKeyDisplay.value);
  showToast('Public key copied');
});

closeImportBtn.addEventListener('click', () => {
  importModal.classList.add('hidden');
});

lockBtn.addEventListener('click', async () => {
  await sendMessage({ action: 'clearSession' });
  currentPassword = null;
  hasKeys = false;
  loginPasswordInput.value = '';
  loginSection.classList.remove('hidden');
  setupSection.classList.add('hidden');
  showScreen(unlockScreen);
});

// Handle Enter key
newPasswordInput.addEventListener('keypress', (e) => {
  if (e.key === 'Enter') confirmPasswordInput.focus();
});

confirmPasswordInput.addEventListener('keypress', (e) => {
  if (e.key === 'Enter') createPasswordBtn.click();
});

loginPasswordInput.addEventListener('keypress', (e) => {
  if (e.key === 'Enter') unlockBtn.click();
});

// Initialize on load
init();
