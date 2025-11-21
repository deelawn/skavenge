// Background service worker for Skavenger extension
// Handles key generation, storage, import, and export

const STORAGE_KEY = 'skavenger_encrypted_keys';

// Derive encryption key from password using PBKDF2
async function deriveKey(password, salt) {
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
      iterations: 100000,
      hash: 'SHA-256'
    },
    keyMaterial,
    { name: 'AES-GCM', length: 256 },
    false,
    ['encrypt', 'decrypt']
  );
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

// Export key to PEM format
async function exportKeyToPEM(key, isPrivate) {
  const format = isPrivate ? 'pkcs8' : 'spki';
  const exported = await crypto.subtle.exportKey(format, key);
  const base64 = arrayBufferToBase64(exported);

  const type = isPrivate ? 'PRIVATE KEY' : 'PUBLIC KEY';
  const lines = base64.match(/.{1,64}/g) || [];

  return `-----BEGIN ${type}-----\n${lines.join('\n')}\n-----END ${type}-----`;
}

// Import key from PEM format
async function importKeyFromPEM(pem, isPrivate) {
  const pemHeader = isPrivate ? '-----BEGIN PRIVATE KEY-----' : '-----BEGIN PUBLIC KEY-----';
  const pemFooter = isPrivate ? '-----END PRIVATE KEY-----' : '-----END PUBLIC KEY-----';

  const pemContents = pem
    .replace(pemHeader, '')
    .replace(pemFooter, '')
    .replace(/\s/g, '');

  const binaryDer = base64ToArrayBuffer(pemContents);

  const format = isPrivate ? 'pkcs8' : 'spki';
  const keyUsages = isPrivate ? ['sign'] : ['verify'];

  return crypto.subtle.importKey(
    format,
    binaryDer,
    {
      name: 'ECDSA',
      namedCurve: 'P-256'
    },
    true,
    keyUsages
  );
}

// Store encrypted keys
async function storeKeys(keyData, password) {
  const encrypted = await encryptData(keyData, password);
  await chrome.storage.local.set({ [STORAGE_KEY]: encrypted });
}

// Retrieve and decrypt keys
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
        const privateKeyPEM = await exportKeyToPEM(keyPair.privateKey, true);
        const publicKeyPEM = await exportKeyToPEM(keyPair.publicKey, false);

        await storeKeys({ privateKey: privateKeyPEM, publicKey: publicKeyPEM }, request.password);

        return { success: true, publicKey: publicKeyPEM };
      }

      case 'exportKeys': {
        const keys = await retrieveKeys(request.password);
        if (!keys) {
          return { success: false, error: 'Invalid password or no keys found' };
        }
        return { success: true, privateKey: keys.privateKey, publicKey: keys.publicKey };
      }

      case 'exportPublicKey': {
        const keys = await retrieveKeys(request.password);
        if (!keys) {
          return { success: false, error: 'Invalid password or no keys found' };
        }
        return { success: true, publicKey: keys.publicKey };
      }

      case 'importKeys': {
        // Validate the keys by attempting to import them
        await importKeyFromPEM(request.privateKey, true);
        await importKeyFromPEM(request.publicKey, false);

        await storeKeys({ privateKey: request.privateKey, publicKey: request.publicKey }, request.password);

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
