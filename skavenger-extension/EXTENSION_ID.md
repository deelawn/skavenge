# Skavenger Extension ID

## Extension ID
```
hnbligdjmpihmhhgajlfjmckcnmnofbn
```

This is the stable extension ID for the Skavenger Chrome Extension. This ID is generated from the public key stored in the manifest.json file.

## Using the Extension ID

### In Web Applications
Web applications need to use this extension ID to communicate with the Skavenger extension. The extension must be installed and the web page must be served from an allowed origin.

```javascript
const SKAVENGER_EXTENSION_ID = 'hnbligdjmpihmhhgajlfjmckcnmnofbn';

// Send a message to the extension
chrome.runtime.sendMessage(SKAVENGER_EXTENSION_ID, {
  action: 'hasKeys'
}, (response) => {
  if (chrome.runtime.lastError) {
    console.error('Extension not found:', chrome.runtime.lastError);
    return;
  }
  console.log('Extension response:', response);
});
```

### Allowed Origins
The extension can only be accessed from the following origins (configured in manifest.json):
- `http://localhost:*/*` - Local development
- `http://127.0.0.1:*/*` - Local development
- `https://*.skavenge.io/*` - Production domain

## How the Extension ID is Generated

The extension ID is deterministic and derived from the public key in the manifest.json:

1. The public key is decoded from base64
2. SHA-256 hash is calculated
3. First 16 bytes (128 bits) of the hash are taken
4. Each byte is converted to two characters using the alphabet 'a-p' (instead of hex '0-9a-f')

This ensures the extension ID remains constant across installations as long as the same key is used in the manifest.

## Security Note

The private key used to generate the extension ID is stored in `private_key.pem`. This file should be kept secure and backed up, as it ensures the extension maintains the same ID across versions and installations.

**Do not commit `private_key.pem` to public repositories.**

## Regenerating the Extension ID

If you need to regenerate the extension ID:

```bash
# Generate new private key
openssl genrsa 2048 | openssl pkcs8 -topk8 -nocrypt -out private_key.pem

# Extract public key in base64 format
openssl rsa -in private_key.pem -pubout -outform DER | base64 -w 0 > public_key_base64.txt

# Calculate the extension ID
python3 << 'EOF'
import hashlib
import subprocess

result = subprocess.run(['openssl', 'rsa', '-in', 'private_key.pem', '-pubout', '-outform', 'DER'],
                       capture_output=True)
public_key_der = result.stdout
sha256_hash = hashlib.sha256(public_key_der).digest()
first_16_bytes = sha256_hash[:16]

extension_id = ''
for byte in first_16_bytes:
    extension_id += chr(ord('a') + (byte >> 4))
    extension_id += chr(ord('a') + (byte & 0x0f))

print(f"Extension ID: {extension_id}")
EOF

# Update manifest.json with the new public key from public_key_base64.txt
```

After regenerating, update the extension ID in all web applications that communicate with the extension.
