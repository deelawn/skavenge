# Skavenger Chrome Extension

Secure key management for Skavenge proof generation and verification.

## Features

- **Generate Keys**: Create ECDSA P-256 key pairs
- **Export Keys**: Export keys in PEM format
- **Import Keys**: Import existing PEM-formatted keys
- **Secure Storage**: Keys are encrypted with AES-256-GCM using a password-derived key (PBKDF2)

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

Keys are exported/imported in encrypted JSON keystore format (similar to Ethereum's keystore):

```json
{
  "version": 1,
  "id": "uuid-v4",
  "crypto": {
    "cipher": "aes-256-gcm",
    "ciphertext": "hex-encoded",
    "cipherparams": { "iv": "hex-encoded" },
    "kdf": "pbkdf2",
    "kdfparams": {
      "dklen": 32,
      "salt": "hex-encoded",
      "c": 100000,
      "prf": "hmac-sha256"
    },
    "mac": "hex-encoded"
  }
}
```

The keys use ECDSA with the P-256 curve (equivalent to Go's `ecdsa.GenerateKey` with `elliptic.P256()`).
