# Skavenger Chrome Extension

Secure key management for Skavenge proof generation and verification.

## Extension ID

```
hnbligdjmpihmhhgajlfjmckcnmnofbn
```

This extension has a stable ID generated from the public key in manifest.json. See [EXTENSION_ID.md](EXTENSION_ID.md) for details on how to use this ID in web applications.

## Features

- **Generate Keys**: Create ECDSA P-256 key pairs
- **Export Keys**: Export keys in PEM format
- **Import Keys**: Import existing PEM-formatted keys
- **Secure Storage**: Keys are encrypted with AES-256-GCM using a password-derived key (PBKDF2)
- **Web App Integration**: Communicate with web applications via externally_connectable API

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

Keys are exported/imported as JSON with hex-encoded keys:

```json
{
  "privateKey": "hex-encoded-pkcs8",
  "publicKey": "hex-encoded-spki"
}
```

The keys use ECDSA with the P-256 curve (equivalent to Go's `ecdsa.GenerateKey` with `elliptic.P256()`).

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
