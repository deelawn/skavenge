// Background service worker for Skavenger extension
// Handles key generation, storage, import, and export

const STORAGE_KEY = 'skavenger_encrypted_keys';
const KEYSTORE_VERSION = 1;

// Generate UUID v4
function generateUUID() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, (c) => {
    const r = crypto.getRandomValues(new Uint8Array(1))[0] % 16;
    const v = c === 'x' ? r : (r & 0x3) | 0x8;
    return v.toString(16);
  });
}

// Convert ArrayBuffer to hex string
function bufferToHex(buffer) {
  return Array.from(new Uint8Array(buffer))
    .map((b) => b.toString(16).padStart(2, '0'))
    .join('');
}

// Convert hex string to ArrayBuffer
function hexToBuffer(hex) {
  const bytes = new Uint8Array(hex.length / 2);
  for (let i = 0; i < hex.length; i += 2) {
    bytes[i / 2] = parseInt(hex.substr(i, 2), 16);
  }
  return bytes.buffer;
}

// Derive encryption key from password using PBKDF2
async function deriveKey(password, salt, iterations = 100000) {
  const encoder = new TextEncoder();
  const keyMaterial = await crypto.subtle.importKey(
    'raw',
    encoder.encode(password),
    'PBKDF2',
    false,
    ['deriveKey']
  );

  return crypto.subtle.deriveKey(
    {
      name: 'PBKDF2',
      salt: salt,
      iterations: iterations,
      hash: 'SHA-256'
    },
    keyMaterial,
    { name: 'AES-GCM', length: 256 },
    false,
    ['encrypt', 'decrypt']
  );
}

// Derive key bytes for MAC calculation
async function deriveKeyBytes(password, salt, iterations = 100000) {
  const encoder = new TextEncoder();
  const keyMaterial = await crypto.subtle.importKey(
    'raw',
    encoder.encode(password),
    'PBKDF2',
    false,
    ['deriveBits']
  );

  return crypto.subtle.deriveBits(
    {
      name: 'PBKDF2',
      salt: salt,
      iterations: iterations,
      hash: 'SHA-256'
    },
    keyMaterial,
    256
  );
}

// Calculate MAC (SHA-256 of derived key + ciphertext)
async function calculateMAC(derivedKeyBytes, ciphertext) {
  const macInput = new Uint8Array(derivedKeyBytes.byteLength + ciphertext.byteLength);
  macInput.set(new Uint8Array(derivedKeyBytes), 0);
  macInput.set(new Uint8Array(ciphertext), derivedKeyBytes.byteLength);
  return crypto.subtle.digest('SHA-256', macInput);
}

// Encrypt data with AES-GCM
async function encryptData(data, password) {
  const encoder = new TextEncoder();
  const salt = crypto.getRandomValues(new Uint8Array(16));
  const iv = crypto.getRandomValues(new Uint8Array(12));
  const key = await deriveKey(password, salt);

  const encrypted = await crypto.subtle.encrypt(
    { name: 'AES-GCM', iv: iv },
    key,
    encoder.encode(JSON.stringify(data))
  );

  return {
    salt: Array.from(salt),
    iv: Array.from(iv),
    data: Array.from(new Uint8Array(encrypted))
  };
}

// Decrypt data with AES-GCM
async function decryptData(encryptedObj, password) {
  const salt = new Uint8Array(encryptedObj.salt);
  const iv = new Uint8Array(encryptedObj.iv);
  const data = new Uint8Array(encryptedObj.data);
  const key = await deriveKey(password, salt);

  const decrypted = await crypto.subtle.decrypt(
    { name: 'AES-GCM', iv: iv },
    key,
    data
  );

  const decoder = new TextDecoder();
  return JSON.parse(decoder.decode(decrypted));
}

// Generate ECDSA P-256 key pair
async function generateKeyPair() {
  const keyPair = await crypto.subtle.generateKey(
    {
      name: 'ECDSA',
      namedCurve: 'P-256'
    },
    true,
    ['sign', 'verify']
  );

  return keyPair;
}

// Convert ArrayBuffer to base64
function arrayBufferToBase64(buffer) {
  const bytes = new Uint8Array(buffer);
  let binary = '';
  for (let i = 0; i < bytes.byteLength; i++) {
    binary += String.fromCharCode(bytes[i]);
  }
  return btoa(binary);
}

// Convert base64 to ArrayBuffer
function base64ToArrayBuffer(base64) {
  const binary = atob(base64);
  const bytes = new Uint8Array(binary.length);
  for (let i = 0; i < binary.length; i++) {
    bytes[i] = binary.charCodeAt(i);
  }
  return bytes.buffer;
}

// Export keys to JSON keystore format
async function exportToKeystore(privateKeyRaw, publicKeyRaw, password) {
  const salt = crypto.getRandomValues(new Uint8Array(32));
  const iv = crypto.getRandomValues(new Uint8Array(12));
  const iterations = 100000;

  const key = await deriveKey(password, salt, iterations);
  const derivedKeyBytes = await deriveKeyBytes(password, salt, iterations);

  // Combine private and public keys
  const keyData = JSON.stringify({
    privateKey: bufferToHex(privateKeyRaw),
    publicKey: bufferToHex(publicKeyRaw)
  });

  const encoder = new TextEncoder();
  const ciphertext = await crypto.subtle.encrypt(
    { name: 'AES-GCM', iv: iv },
    key,
    encoder.encode(keyData)
  );

  const mac = await calculateMAC(derivedKeyBytes, ciphertext);

  return {
    version: KEYSTORE_VERSION,
    id: generateUUID(),
    crypto: {
      cipher: 'aes-256-gcm',
      ciphertext: bufferToHex(ciphertext),
      cipherparams: {
        iv: bufferToHex(iv)
      },
      kdf: 'pbkdf2',
      kdfparams: {
        dklen: 32,
        salt: bufferToHex(salt),
        c: iterations,
        prf: 'hmac-sha256'
      },
      mac: bufferToHex(mac)
    }
  };
}

// Import keys from JSON keystore format
async function importFromKeystore(keystore, password) {
  if (keystore.version !== KEYSTORE_VERSION) {
    throw new Error('Unsupported keystore version');
  }

  const cryptoData = keystore.crypto;
  const salt = new Uint8Array(hexToBuffer(cryptoData.kdfparams.salt));
  const iv = new Uint8Array(hexToBuffer(cryptoData.cipherparams.iv));
  const ciphertext = hexToBuffer(cryptoData.ciphertext);
  const iterations = cryptoData.kdfparams.c;

  // Verify MAC
  const derivedKeyBytes = await deriveKeyBytes(password, salt, iterations);
  const calculatedMAC = await calculateMAC(derivedKeyBytes, ciphertext);

  if (bufferToHex(calculatedMAC) !== cryptoData.mac) {
    throw new Error('Invalid password or corrupted keystore');
  }

  // Decrypt
  const key = await deriveKey(password, salt, iterations);
  const decrypted = await crypto.subtle.decrypt(
    { name: 'AES-GCM', iv: iv },
    key,
    ciphertext
  );

  const decoder = new TextDecoder();
  const keyData = JSON.parse(decoder.decode(decrypted));

  return {
    privateKey: hexToBuffer(keyData.privateKey),
    publicKey: hexToBuffer(keyData.publicKey)
  };
}

// Store encrypted keys (internal storage format)
async function storeKeys(keyData, password) {
  const encrypted = await encryptData(keyData, password);
  await chrome.storage.local.set({ [STORAGE_KEY]: encrypted });
}

// Retrieve and decrypt keys (internal storage format)
async function retrieveKeys(password) {
  const result = await chrome.storage.local.get(STORAGE_KEY);
  if (!result[STORAGE_KEY]) {
    return null;
  }
  return decryptData(result[STORAGE_KEY], password);
}

// Check if keys exist
async function hasStoredKeys() {
  const result = await chrome.storage.local.get(STORAGE_KEY);
  return !!result[STORAGE_KEY];
}

// Clear stored keys
async function clearKeys() {
  await chrome.storage.local.remove(STORAGE_KEY);
}

// Message handler
chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
  handleMessage(request).then(sendResponse);
  return true; // Keep channel open for async response
});

async function handleMessage(request) {
  try {
    switch (request.action) {
      case 'generateKeys': {
        const keyPair = await generateKeyPair();
        const privateKeyRaw = await crypto.subtle.exportKey('pkcs8', keyPair.privateKey);
        const publicKeyRaw = await crypto.subtle.exportKey('spki', keyPair.publicKey);

        await storeKeys({
          privateKey: bufferToHex(privateKeyRaw),
          publicKey: bufferToHex(publicKeyRaw)
        }, request.password);

        return { success: true, publicKey: bufferToHex(publicKeyRaw) };
      }

      case 'exportKeys': {
        const keys = await retrieveKeys(request.password);
        if (!keys) {
          return { success: false, error: 'Invalid password or no keys found' };
        }

        // Export as JSON keystore
        const keystore = await exportToKeystore(
          hexToBuffer(keys.privateKey),
          hexToBuffer(keys.publicKey),
          request.exportPassword || request.password
        );

        return { success: true, keystore: JSON.stringify(keystore, null, 2) };
      }

      case 'exportPublicKey': {
        const keys = await retrieveKeys(request.password);
        if (!keys) {
          return { success: false, error: 'Invalid password or no keys found' };
        }
        return { success: true, publicKey: keys.publicKey };
      }

      case 'importKeys': {
        // Parse and import from JSON keystore
        const keystore = JSON.parse(request.keystore);
        const importedKeys = await importFromKeystore(keystore, request.keystorePassword);

        // Validate keys by importing them as CryptoKey objects
        await crypto.subtle.importKey(
          'pkcs8',
          importedKeys.privateKey,
          { name: 'ECDSA', namedCurve: 'P-256' },
          true,
          ['sign']
        );
        await crypto.subtle.importKey(
          'spki',
          importedKeys.publicKey,
          { name: 'ECDSA', namedCurve: 'P-256' },
          true,
          ['verify']
        );

        await storeKeys({
          privateKey: bufferToHex(importedKeys.privateKey),
          publicKey: bufferToHex(importedKeys.publicKey)
        }, request.password);

        return { success: true };
      }

      case 'hasKeys': {
        const hasKeys = await hasStoredKeys();
        return { success: true, hasKeys };
      }

      case 'clearKeys': {
        await clearKeys();
        return { success: true };
      }

      case 'verifyPassword': {
        const keys = await retrieveKeys(request.password);
        return { success: !!keys };
      }

      default:
        return { success: false, error: 'Unknown action' };
    }
  } catch (error) {
    return { success: false, error: error.message };
  }
}
