# Skavenger Chrome Extension

Secure key management for Skavenge proof generation and verification.

## Extension ID

```
hnbligdjmpihmhhgajlfjmckcnmnofbn
```

This extension has a stable ID generated from the public key in manifest.json. See [EXTENSION_ID.md](EXTENSION_ID.md) for details on how to use this ID in web applications.

## Features

- **Generate Keys**: Create ECDSA secp256k1 key pairs
- **Export Keys**: Export keys in hex format
- **Import Keys**: Import existing hex-formatted keys
- **Secure Storage**: Keys are encrypted with AES-256-GCM using a password-derived key (PBKDF2)
- **Web App Integration**: Communicate with web applications via externally_connectable API
- **Signing**: Sign messages using secp256k1 ECDSA with deterministic k-values (RFC 6979)
- **ElGamal Decryption**: Decrypt ElGamal ciphertexts using secp256k1

## Installation

1. Open Chrome and navigate to `chrome://extensions/`
2. Enable "Developer mode" (toggle in top right)
3. Click "Load unpacked"
4. Select the `skavenger-extension` directory
5. Add placeholder icons to `icons/` directory (icon16.png, icon48.png, icon128.png)

## Security

- Private keys are never accessible to web applications
- Keys are encrypted at rest using AES-256-GCM
- Encryption key is derived from user password using PBKDF2 (100,000 iterations)
- Keys remain encrypted in Chrome's local storage

## Key Format

Keys are stored and exported as hex-encoded strings:

- **Private Key**: 64 hex characters (32 bytes)
- **Public Key**: 130 hex characters (65 bytes, uncompressed format: `0x04 || X || Y`)

The keys use ECDSA with the secp256k1 curve (Bitcoin/Ethereum curve), compatible with Go's `ecdsa.GenerateKey` with `secp256k1.S256()`.

Signatures are encoded in ASN.1 DER format to match the Go backend's signature verification.

## Web Application Integration

Web applications can communicate with this extension using the Chrome extension messaging API. The extension is configured to accept messages from:
- `http://localhost:*/*` (local development)
- `http://127.0.0.1:*/*` (local development)
- `https://*.skavenge.io/*` (production)

Example usage in a web app:

```javascript
const SKAVENGER_EXTENSION_ID = 'hnbligdjmpihmhhgajlfjmckcnmnofbn';

chrome.runtime.sendMessage(SKAVENGER_EXTENSION_ID, {
  action: 'hasKeys'
}, (response) => {
  if (chrome.runtime.lastError) {
    console.error('Extension error:', chrome.runtime.lastError);
    return;
  }
  console.log('Has keys:', response.hasKeys);
});
```

See the [webapp.html](../webapp.html) example in the repository root for a complete demo.
