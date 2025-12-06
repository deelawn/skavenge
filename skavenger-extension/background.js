// Background service worker for Skavenger extension
// Handles key generation, storage, import, and export

const STORAGE_KEY = 'skavenger_encrypted_keys';
const PUBLIC_KEY_STORAGE = 'skavenger_public_key';
const SESSION_KEY = 'skavenger_session';
const KEYSTORE_VERSION = 1;
const SESSION_TIMEOUT_MS = 15 * 60 * 1000; // 15 minutes

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

// Keccak-256 implementation (legacy Keccak, as used by Go's sha3 package)
// Uses 32-bit array representation for simplicity and correctness
function keccak256(msgBytes) {
  const RC = [
    1, 0, 32898, 0, 32906, 2147483648, 2147516416, 2147483648, 32907, 0, 2147483649,
    0, 2147516545, 2147483648, 32777, 2147483648, 138, 0, 136, 0, 2147516425, 0,
    2147483658, 0, 2147516555, 0, 139, 2147483648, 32905, 2147483648, 32771,
    2147483648, 32770, 2147483648, 128, 2147483648, 32778, 0, 2147483658, 2147483648,
    2147516545, 2147483648, 32896, 2147483648, 2147483649, 0, 2147516424, 2147483648
  ];

  const r = [0, 1, 62, 28, 27, 36, 44, 6, 55, 20, 3, 10, 43, 25, 39, 41, 45, 15, 21, 8, 18, 2, 61, 56, 14];

  const rotate = function(v, n) {
    const h = v[0], l = v[1];
    if (n === 0) return v;
    if (n === 32) return [l, h];
    if (n < 32) {
      return [(h << n) | (l >>> (32 - n)), (l << n) | (h >>> (32 - n))];
    } else {
      n -= 32;
      return [(l << n) | (h >>> (32 - n)), (h << n) | (l >>> (32 - n))];
    }
  };

  // Initialize state (25 lanes x 2 words each = 50 words)
  const s = new Array(50);
  for (let i = 0; i < 50; ++i) s[i] = 0;

  const blocks = [];
  const blockCount = 17; // (1600 - 512) / 32 for Keccak-256
  const byteCount = 136; // rate for Keccak-256

  // Absorb phase
  let blockIndex = 0, lastByteIndex = 0;

  for (let i = 0; i < msgBytes.length; ++i) {
    const index = i % byteCount;
    const wIndex = index >> 2;
    const shift = (index & 3) << 3;

    if (!blocks[wIndex]) blocks[wIndex] = 0;
    blocks[wIndex] |= msgBytes[i] << shift;

    if ((i + 1) % byteCount === 0 || i === msgBytes.length - 1) {
      lastByteIndex = index + 1;

      if (lastByteIndex === byteCount || i === msgBytes.length - 1) {
        // Padding
        if (lastByteIndex < byteCount) {
          blocks[lastByteIndex >> 2] |= 0x01 << ((lastByteIndex & 3) << 3);
          if (lastByteIndex === byteCount - 1) {
            for (let j = 0; j < blockCount; ++j) s[j] ^= blocks[j] || 0;
            f(s);
            blocks.fill(0);
          }
        }
        blocks[(byteCount - 1) >> 2] |= 0x80 << (((byteCount - 1) & 3) << 3);

        for (let j = 0; j < blockCount; ++j) s[j] ^= blocks[j] || 0;
        f(s);
        break;
      }

      for (let j = 0; j < blockCount; ++j) s[j] ^= blocks[j] || 0;
      f(s);
      blocks.fill(0);
    }
  }

  // Keccak-f permutation
  function f(s) {
    for (let n = 0; n < 48; n += 2) {
      const c0 = s[0] ^ s[10] ^ s[20] ^ s[30] ^ s[40];
      const c1 = s[1] ^ s[11] ^ s[21] ^ s[31] ^ s[41];
      const c2 = s[2] ^ s[12] ^ s[22] ^ s[32] ^ s[42];
      const c3 = s[3] ^ s[13] ^ s[23] ^ s[33] ^ s[43];
      const c4 = s[4] ^ s[14] ^ s[24] ^ s[34] ^ s[44];
      const c5 = s[5] ^ s[15] ^ s[25] ^ s[35] ^ s[45];
      const c6 = s[6] ^ s[16] ^ s[26] ^ s[36] ^ s[46];
      const c7 = s[7] ^ s[17] ^ s[27] ^ s[37] ^ s[47];
      const c8 = s[8] ^ s[18] ^ s[28] ^ s[38] ^ s[48];
      const c9 = s[9] ^ s[19] ^ s[29] ^ s[39] ^ s[49];

      let h = c8 ^ ((c2 << 1) | (c3 >>> 31));
      let l = c9 ^ ((c3 << 1) | (c2 >>> 31));
      s[0] ^= h; s[1] ^= l;
      s[10] ^= h; s[11] ^= l;
      s[20] ^= h; s[21] ^= l;
      s[30] ^= h; s[31] ^= l;
      s[40] ^= h; s[41] ^= l;
      h = c0 ^ ((c4 << 1) | (c5 >>> 31));
      l = c1 ^ ((c5 << 1) | (c4 >>> 31));
      s[2] ^= h; s[3] ^= l;
      s[12] ^= h; s[13] ^= l;
      s[22] ^= h; s[23] ^= l;
      s[32] ^= h; s[33] ^= l;
      s[42] ^= h; s[43] ^= l;
      h = c2 ^ ((c6 << 1) | (c7 >>> 31));
      l = c3 ^ ((c7 << 1) | (c6 >>> 31));
      s[4] ^= h; s[5] ^= l;
      s[14] ^= h; s[15] ^= l;
      s[24] ^= h; s[25] ^= l;
      s[34] ^= h; s[35] ^= l;
      s[44] ^= h; s[45] ^= l;
      h = c4 ^ ((c8 << 1) | (c9 >>> 31));
      l = c5 ^ ((c9 << 1) | (c8 >>> 31));
      s[6] ^= h; s[7] ^= l;
      s[16] ^= h; s[17] ^= l;
      s[26] ^= h; s[27] ^= l;
      s[36] ^= h; s[37] ^= l;
      s[46] ^= h; s[47] ^= l;
      h = c6 ^ ((c0 << 1) | (c1 >>> 31));
      l = c7 ^ ((c1 << 1) | (c0 >>> 31));
      s[8] ^= h; s[9] ^= l;
      s[18] ^= h; s[19] ^= l;
      s[28] ^= h; s[29] ^= l;
      s[38] ^= h; s[39] ^= l;
      s[48] ^= h; s[49] ^= l;

      // Rho Pi
      const b0 = s[0], b1 = s[1];
      h = s[2]; l = s[3];
      s[2] = (l << 12) | (h >>> 20);
      s[3] = (h << 12) | (l >>> 20);
      h = s[14]; l = s[15];
      s[14] = (l << 20) | (h >>> 12);
      s[15] = (h << 20) | (l >>> 12);
      h = s[26]; l = s[27];
      s[26] = (l << 29) | (h >>> 3);
      s[27] = (h << 29) | (l >>> 3);
      h = s[38]; l = s[39];
      s[38] = (h << 7) | (l >>> 25);
      s[39] = (l << 7) | (h >>> 25);
      h = s[48]; l = s[49];
      s[48] = (h << 14) | (l >>> 18);
      s[49] = (l << 14) | (h >>> 18);

      h = s[6]; l = s[7];
      s[6] = (h << 3) | (l >>> 29);
      s[7] = (l << 3) | (h >>> 29);
      h = s[18]; l = s[19];
      s[18] = (l << 10) | (h >>> 22);
      s[19] = (h << 10) | (l >>> 22);
      h = s[30]; l = s[31];
      s[30] = (l << 18) | (h >>> 14);
      s[31] = (h << 18) | (l >>> 14);
      h = s[42]; l = s[43];
      s[42] = (h << 27) | (l >>> 5);
      s[43] = (l << 27) | (h >>> 5);
      h = s[4]; l = s[5];
      s[4] = (l << 1) | (h >>> 31);
      s[5] = (h << 1) | (l >>> 31);

      h = s[16]; l = s[17];
      s[16] = (l << 6) | (h >>> 26);
      s[17] = (h << 6) | (l >>> 26);
      h = s[28]; l = s[29];
      s[28] = (h << 25) | (l >>> 7);
      s[29] = (l << 25) | (h >>> 7);
      h = s[40]; l = s[41];
      s[40] = (l << 8) | (h >>> 24);
      s[41] = (h << 8) | (l >>> 24);
      h = s[2]; l = s[3];
      s[2] = (l << 23) | (h >>> 9);
      s[3] = (h << 23) | (l >>> 9);

      h = s[24]; l = s[25];
      s[24] = (h << 2) | (l >>> 30);
      s[25] = (l << 2) | (h >>> 30);
      h = s[36]; l = s[37];
      s[36] = (h << 13) | (l >>> 19);
      s[37] = (l << 13) | (h >>> 19);
      h = s[48]; l = s[49];
      s[48] = (h << 21) | (l >>> 11);
      s[49] = (l << 21) | (h >>> 11);
      h = s[10]; l = s[11];
      s[10] = (l << 30) | (h >>> 2);
      s[11] = (h << 30) | (l >>> 2);
      h = s[22]; l = s[23];
      s[22] = (h << 9) | (l >>> 23);
      s[23] = (l << 9) | (h >>> 23);

      h = s[34]; l = s[35];
      s[34] = (l << 15) | (h >>> 17);
      s[35] = (h << 15) | (l >>> 17);
      h = s[46]; l = s[47];
      s[46] = (l << 24) | (h >>> 8);
      s[47] = (h << 24) | (l >>> 8);
      h = s[8]; l = s[9];
      s[8] = (h << 28) | (l >>> 4);
      s[9] = (l << 28) | (h >>> 4);
      h = s[20]; l = s[21];
      s[20] = (l << 4) | (h >>> 28);
      s[21] = (h << 4) | (l >>> 28);
      h = s[32]; l = s[33];
      s[32] = (h << 11) | (l >>> 21);
      s[33] = (l << 11) | (h >>> 21);

      h = s[44]; l = s[45];
      s[44] = (l << 27) | (h >>> 5);
      s[45] = (h << 27) | (l >>> 5);
      h = s[6]; l = s[7];
      s[6] = (l << 19) | (h >>> 13);
      s[7] = (h << 19) | (l >>> 13);
      h = s[18]; l = s[19];
      s[18] = (h << 20) | (l >>> 12);
      s[19] = (l << 20) | (h >>> 12);
      h = s[30]; l = s[31];
      s[30] = (h << 10) | (l >>> 22);
      s[31] = (l << 10) | (h >>> 22);

      h = s[42]; l = s[43];
      s[42] = (l << 26) | (h >>> 6);
      s[43] = (h << 26) | (l >>> 6);
      h = s[4]; l = s[5];
      s[4] = (h << 18) | (l >>> 14);
      s[5] = (l << 18) | (h >>> 14);
      h = s[16]; l = s[17];
      s[16] = (h << 15) | (l >>> 17);
      s[17] = (l << 15) | (h >>> 17);

      h = s[28]; l = s[29];
      s[28] = (l << 14) | (h >>> 18);
      s[29] = (h << 14) | (l >>> 18);
      h = s[40]; l = s[41];
      s[40] = (l << 5) | (h >>> 27);
      s[41] = (h << 5) | (l >>> 27);
      h = s[2]; l = s[3];
      s[2] = (h << 3) | (l >>> 29);
      s[3] = (l << 3) | (h >>> 29);
      h = s[12]; l = s[13];
      s[12] = (l << 31) | (h >>> 1);
      s[13] = (h << 31) | (l >>> 1);

      h = s[24]; l = s[25];
      s[24] = (l << 25) | (h >>> 7);
      s[25] = (h << 25) | (l >>> 7);
      h = s[36]; l = s[37];
      s[36] = (l << 16) | (h >>> 16);
      s[37] = (h << 16) | (l >>> 16);
      h = s[48]; l = s[49];
      s[48] = (l << 7) | (h >>> 25);
      s[49] = (h << 7) | (l >>> 25);
      h = s[10]; l = s[11];
      s[10] = (h << 29) | (l >>> 3);
      s[11] = (l << 29) | (h >>> 3);

      h = s[22]; l = s[23];
      s[22] = (l << 24) | (h >>> 8);
      s[23] = (h << 24) | (l >>> 8);
      h = s[34]; l = s[35];
      s[34] = (h << 1) | (l >>> 31);
      s[35] = (l << 1) | (h >>> 31);
      h = s[46]; l = s[47];
      s[46] = (h << 6) | (l >>> 26);
      s[47] = (l << 6) | (h >>> 26);
      h = s[8]; l = s[9];
      s[8] = (l << 23) | (h >>> 9);
      s[9] = (h << 23) | (l >>> 9);

      h = s[20]; l = s[21];
      s[20] = (h << 12) | (l >>> 20);
      s[21] = (l << 12) | (h >>> 20);
      h = s[32]; l = s[33];
      s[32] = (l << 21) | (h >>> 11);
      s[33] = (h << 21) | (l >>> 11);
      h = s[44]; l = s[45];
      s[44] = (l << 2) | (h >>> 30);
      s[45] = (h << 2) | (l >>> 30);
      h = s[6]; l = s[7];
      s[6] = (h << 8) | (l >>> 24);
      s[7] = (l << 8) | (h >>> 24);

      h = s[18]; l = s[19];
      s[18] = (l << 22) | (h >>> 10);
      s[19] = (h << 22) | (l >>> 10);
      h = s[30]; l = s[31];
      s[30] = (l << 11) | (h >>> 21);
      s[31] = (h << 11) | (l >>> 21);
      h = s[42]; l = s[43];
      s[42] = (h << 30) | (l >>> 2);
      s[43] = (l << 30) | (h >>> 2);
      h = s[4]; l = s[5];
      s[4] = (h << 21) | (l >>> 11);
      s[5] = (l << 21) | (h >>> 11);

      h = s[16]; l = s[17];
      s[16] = (l << 17) | (h >>> 15);
      s[17] = (h << 17) | (l >>> 15);
      h = s[28]; l = s[29];
      s[28] = (h << 28) | (l >>> 4);
      s[29] = (l << 28) | (h >>> 4);
      h = s[40]; l = s[41];
      s[40] = (h << 19) | (l >>> 13);
      s[41] = (l << 19) | (h >>> 13);
      s[0] = b0;
      s[1] = b1;

      // Chi
      for (let j = 0; j < 50; j += 10) {
        const c0 = s[j], c1 = s[j + 1];
        h = s[j + 2]; l = s[j + 3];
        const c2 = h, c3 = l;
        h = s[j + 4]; l = s[j + 5];
        const c4 = h, c5 = l;
        h = s[j + 6]; l = s[j + 7];
        const c6 = h, c7 = l;
        h = s[j + 8]; l = s[j + 9];
        const c8 = h, c9 = l;

        s[j] ^= (~c2) & c4;
        s[j + 1] ^= (~c3) & c5;
        s[j + 2] ^= (~c4) & c6;
        s[j + 3] ^= (~c5) & c7;
        s[j + 4] ^= (~c6) & c8;
        s[j + 5] ^= (~c7) & c9;
        s[j + 6] ^= (~c8) & c0;
        s[j + 7] ^= (~c9) & c1;
        s[j + 8] ^= (~c0) & c2;
        s[j + 9] ^= (~c1) & c3;
      }

      // Iota
      s[0] ^= RC[n];
      s[1] ^= RC[n + 1];
    }
  }

  // Extract output
  const output = new Uint8Array(32);
  for (let i = 0, j = 0; i < 32; i += 4) {
    const n = s[j++];
    output[i] = n;
    output[i + 1] = n >>> 8;
    output[i + 2] = n >>> 16;
    output[i + 3] = n >>> 24;
  }

  return output;
}

// Convert base64url to hex
function base64urlToHex(base64url) {
  // Add padding if needed
  let base64 = base64url.replace(/-/g, '+').replace(/_/g, '/');
  while (base64.length % 4) {
    base64 += '=';
  }
  const binary = atob(base64);
  const bytes = new Uint8Array(binary.length);
  for (let i = 0; i < binary.length; i++) {
    bytes[i] = binary.charCodeAt(i);
  }
  return bufferToHex(bytes.buffer);
}

// Convert hex to base64url
function hexToBase64url(hex) {
  const bytes = new Uint8Array(hexToBuffer(hex));
  let binary = '';
  for (let i = 0; i < bytes.length; i++) {
    binary += String.fromCharCode(bytes[i]);
  }
  const base64 = btoa(binary);
  return base64.replace(/\+/g, '-').replace(/\//g, '_').replace(/=/g, '');
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

// Extract raw private key from ECDSA key (32 bytes for P-256)
async function extractRawPrivateKey(privateKey) {
  const jwk = await crypto.subtle.exportKey('jwk', privateKey);
  // The 'd' parameter contains the base64url-encoded raw private key
  return base64urlToHex(jwk.d);
}

// P-256 curve parameters
const P256_P = BigInt('0xffffffff00000001000000000000000000000000ffffffffffffffffffffffff');
const P256_A = BigInt('0xffffffff00000001000000000000000000000000fffffffffffffffffffffffc');
const P256_B = BigInt('0x5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b');
const P256_GX = BigInt('0x6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c296');
const P256_GY = BigInt('0x4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5');
const P256_N = BigInt('0xffffffff00000000ffffffffffffffffbce6faada7179e84f3b9cac2fc632551');

// Modular inverse using Extended Euclidean Algorithm
function modInverse(a, m) {
  a = ((a % m) + m) % m;
  let [oldR, r] = [a, m];
  let [oldS, s] = [1n, 0n];

  while (r !== 0n) {
    const quotient = oldR / r;
    [oldR, r] = [r, oldR - quotient * r];
    [oldS, s] = [s, oldS - quotient * s];
  }

  return ((oldS % m) + m) % m;
}

// EC point addition on P-256
function pointAdd(p1, p2) {
  if (p1 === null) return p2;
  if (p2 === null) return p1;

  const [x1, y1] = p1;
  const [x2, y2] = p2;

  if (x1 === x2 && y1 === y2) {
    // Point doubling
    const s = (3n * x1 * x1 + P256_A) * modInverse(2n * y1, P256_P) % P256_P;
    const x3 = (s * s - 2n * x1) % P256_P;
    const y3 = (s * (x1 - x3) - y1) % P256_P;
    return [(x3 + P256_P) % P256_P, (y3 + P256_P) % P256_P];
  }

  const s = (y2 - y1) * modInverse((x2 - x1 + P256_P) % P256_P, P256_P) % P256_P;
  const x3 = (s * s - x1 - x2) % P256_P;
  const y3 = (s * (x1 - x3) - y1) % P256_P;
  return [(x3 + P256_P) % P256_P, (y3 + P256_P) % P256_P];
}

// Scalar multiplication on P-256
function scalarMult(k, point) {
  let result = null;
  let addend = point;

  while (k > 0n) {
    if (k & 1n) {
      result = pointAdd(result, addend);
    }
    addend = pointAdd(addend, addend);
    k >>= 1n;
  }

  return result;
}

// Compute public key from private key
function derivePublicKey(privateKeyHex) {
  const d = BigInt('0x' + privateKeyHex);
  const [x, y] = scalarMult(d, [P256_GX, P256_GY]);

  // Convert to base64url for JWK
  const xHex = x.toString(16).padStart(64, '0');
  const yHex = y.toString(16).padStart(64, '0');

  return {
    x: hexToBase64url(xHex),
    y: hexToBase64url(yHex)
  };
}

// Import raw private key (64 hex chars) back to CryptoKey with derived public key
async function importRawPrivateKey(rawPrivateKeyHex) {
  const publicKeyCoords = derivePublicKey(rawPrivateKeyHex);

  // Convert raw key to JWK format with public key coordinates
  const jwk = {
    kty: 'EC',
    crv: 'P-256',
    d: hexToBase64url(rawPrivateKeyHex),
    x: publicKeyCoords.x,
    y: publicKeyCoords.y,
    ext: true
  };

  return await crypto.subtle.importKey(
    'jwk',
    jwk,
    { name: 'ECDSA', namedCurve: 'P-256' },
    true,
    ['sign']
  );
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
  // Also store public key separately (unencrypted) for webapp access
  await chrome.storage.local.set({
    [STORAGE_KEY]: encrypted,
    [PUBLIC_KEY_STORAGE]: keyData.publicKey
  });
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

// Get public key (no password required - public keys are meant to be public)
async function getPublicKey() {
  const result = await chrome.storage.local.get(PUBLIC_KEY_STORAGE);
  return result[PUBLIC_KEY_STORAGE] || null;
}

// Clear stored keys
async function clearKeys() {
  await chrome.storage.local.remove(STORAGE_KEY);
  await chrome.storage.local.remove(PUBLIC_KEY_STORAGE);
  await chrome.storage.local.remove(SESSION_KEY);
}

// Session management
async function createSession(password) {
  const session = {
    password: password,
    timestamp: Date.now()
  };
  await chrome.storage.local.set({ [SESSION_KEY]: session });
  return session;
}

async function getSession() {
  const result = await chrome.storage.local.get(SESSION_KEY);
  if (!result[SESSION_KEY]) {
    return null;
  }
  const session = result[SESSION_KEY];
  const now = Date.now();

  // Check if session expired
  if (now - session.timestamp > SESSION_TIMEOUT_MS) {
    await chrome.storage.local.remove(SESSION_KEY);
    return null;
  }

  return session;
}

async function clearSession() {
  await chrome.storage.local.remove(SESSION_KEY);
}

async function getPasswordFromSession(request) {
  // If password is provided in request, use it (for popup)
  if (request.password) {
    // Verify password is correct and create/update session
    const keys = await retrieveKeys(request.password);
    if (keys) {
      await createSession(request.password);
      return request.password;
    }
    return null;
  }

  // Try to get password from session
  const session = await getSession();
  if (session) {
    return session.password;
  }

  return null;
}

// Unmarshal ElGamal ciphertext from bytes
// Format: len(C1) | C1 | len(C2) | C2 | len(SharedSecret) | SharedSecret
function unmarshalElGamalCiphertext(data) {
  const bytes = new Uint8Array(data);
  let offset = 0;

  if (bytes.length < 12) {
    throw new Error('Ciphertext data too short');
  }

  // Read C1 length (4 bytes, big endian)
  const c1Len = (bytes[offset] << 24) | (bytes[offset + 1] << 16) | (bytes[offset + 2] << 8) | bytes[offset + 3];
  offset += 4;

  if (offset + c1Len > bytes.length) {
    throw new Error('Invalid C1 length');
  }

  // Read C1
  const c1 = bytes.slice(offset, offset + c1Len);
  offset += c1Len;

  if (offset + 4 > bytes.length) {
    throw new Error('Data too short for C2 length');
  }

  // Read C2 length (4 bytes, big endian)
  const c2Len = (bytes[offset] << 24) | (bytes[offset + 1] << 16) | (bytes[offset + 2] << 8) | bytes[offset + 3];
  offset += 4;

  if (offset + c2Len > bytes.length) {
    throw new Error('Invalid C2 length');
  }

  // Read C2
  const c2 = bytes.slice(offset, offset + c2Len);
  offset += c2Len;

  if (offset + 4 > bytes.length) {
    throw new Error('Data too short for SharedSecret length');
  }

  // Read SharedSecret length (4 bytes, big endian)
  const ssLen = (bytes[offset] << 24) | (bytes[offset + 1] << 16) | (bytes[offset + 2] << 8) | bytes[offset + 3];
  offset += 4;

  if (offset + ssLen > bytes.length) {
    throw new Error('Invalid SharedSecret length');
  }

  // Read SharedSecret
  const sharedSecret = bytes.slice(offset, offset + ssLen);

  return { c1, c2, sharedSecret };
}

// Parse elliptic curve point from bytes (uncompressed format)
// Format: 0x04 | X (32 bytes) | Y (32 bytes)
function parseECPoint(bytes) {
  if (bytes.length !== 65 || bytes[0] !== 0x04) {
    throw new Error('Invalid EC point format');
  }

  // Extract X and Y coordinates (skip first byte which is 0x04)
  const xBytes = bytes.slice(1, 33);
  const yBytes = bytes.slice(33, 65);

  // Convert to BigInt
  let x = 0n;
  for (let i = 0; i < xBytes.length; i++) {
    x = (x << 8n) | BigInt(xBytes[i]);
  }

  let y = 0n;
  for (let i = 0; i < yBytes.length; i++) {
    y = (y << 8n) | BigInt(yBytes[i]);
  }

  return [x, y];
}

// Convert BigInt to byte array (big endian)
function bigIntToBytes(value) {
  const hex = value.toString(16);
  const paddedHex = hex.length % 2 === 0 ? hex : '0' + hex;
  const bytes = new Uint8Array(paddedHex.length / 2);
  for (let i = 0; i < paddedHex.length; i += 2) {
    bytes[i / 2] = parseInt(paddedHex.substr(i, 2), 16);
  }
  return bytes;
}

// Decrypt ElGamal ciphertext
// This mirrors the DecryptElGamal function in elgamal.go
async function decryptElGamal(encryptedHex, rValueHex, privateKeyHex) {
  try {
    console.log('=== DECRYPTION DEBUG ===');
    console.log('encryptedHex length:', encryptedHex.length);
    console.log('rValueHex:', rValueHex);
    console.log('privateKeyHex:', privateKeyHex);

    // Parse inputs
    const ciphertextBytes = hexToBuffer(encryptedHex);
    const ciphertext = unmarshalElGamalCiphertext(ciphertextBytes);

    console.log('C1 length:', ciphertext.c1.length);
    console.log('C2 length:', ciphertext.c2.length);
    console.log('SharedSecret length:', ciphertext.sharedSecret.length);
    console.log('C2 hex:', bufferToHex(ciphertext.c2));

    // Parse r value
    const rValue = BigInt('0x' + rValueHex);

    // Parse private key
    const privateKeyD = BigInt('0x' + privateKeyHex);

    // Parse C1 as EC point (g^r)
    const [c1x, c1y] = parseECPoint(ciphertext.c1);

    console.log('C1 parsed successfully');

    // Compute shared secret: S = C1^privKey = (g^r)^privKey
    const [sx, sy] = scalarMult(privateKeyD, [c1x, c1y]);

    console.log('sx:', sx.toString(16));

    // Convert sx to bytes (for hashing)
    const sxBytes = bigIntToBytes(sx);

    // Derive decryption key: Hash(r || sharedSecret)
    // This matches the Go implementation: keyHash.Write(r.Bytes()); keyHash.Write(sharedSecret)
    const rBytes = bigIntToBytes(rValue);

    console.log('rBytes length:', rBytes.length, 'hex:', bufferToHex(rBytes));
    console.log('sxBytes length:', sxBytes.length, 'hex:', bufferToHex(sxBytes));

    const keyInput = new Uint8Array(rBytes.length + sxBytes.length);
    keyInput.set(rBytes, 0);
    keyInput.set(sxBytes, rBytes.length);

    console.log('keyInput length:', keyInput.length, 'hex:', bufferToHex(keyInput));

    const key = keccak256(keyInput);

    console.log('key hex:', bufferToHex(key));

    // XOR C2 with key to decrypt
    const plaintext = new Uint8Array(ciphertext.c2.length);
    for (let i = 0; i < ciphertext.c2.length; i++) {
      plaintext[i] = ciphertext.c2[i] ^ key[i % key.length];
    }

    // Convert plaintext bytes to string
    const decoder = new TextDecoder('utf-8');
    const result = decoder.decode(plaintext);
    console.log('plaintext:', result);
    console.log('=== END DEBUG ===');
    return result;
  } catch (error) {
    console.error('Decryption error:', error);
    throw new Error('Decryption failed: ' + error.message);
  }
}

// Internal message handler (from popup)
chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
  handleMessage(request, true).then(sendResponse);
  return true; // Keep channel open for async response
});

// External message handler (from web pages)
chrome.runtime.onMessageExternal.addListener((request, sender, sendResponse) => {
  handleMessage(request, false).then(sendResponse);
  return true; // Keep channel open for async response
});

async function handleMessage(request, isInternal = true) {
  try {
    // Get password - from request for internal (popup), from session for external
    let password = null;
    if (isInternal) {
      // Internal messages (from popup) can provide password directly
      password = request.password || (await getSession())?.password;
    } else {
      // External messages (from web pages) must use session
      password = await getPasswordFromSession(request);

      // For some actions, we need password from request to establish session
      if (!password && (request.action === 'verifyPassword' || request.action === 'generateKeys')) {
        if (request.password) {
          const keys = await retrieveKeys(request.password);
          if (keys) {
            await createSession(request.password);
            password = request.password;
          }
        }
      }
    }

    switch (request.action) {
      case 'generateKeys': {
        if (!password) {
          return { success: false, error: 'Password required or session expired' };
        }

        const keyPair = await generateKeyPair();
        // Extract raw 32-byte private key (64 hex chars)
        const rawPrivateKeyHex = await extractRawPrivateKey(keyPair.privateKey);
        const publicKeyRaw = await crypto.subtle.exportKey('spki', keyPair.publicKey);

        await storeKeys({
          privateKey: rawPrivateKeyHex,
          publicKey: bufferToHex(publicKeyRaw)
        }, password);

        // Always create session so webapp can access the public key
        await createSession(password);

        return { success: true, publicKey: bufferToHex(publicKeyRaw) };
      }

      case 'exportKeys': {
        if (!password) {
          return { success: false, error: 'Password required or session expired' };
        }

        const keys = await retrieveKeys(password);
        if (!keys) {
          return { success: false, error: 'Invalid password or no keys found' };
        }

        // Return only the raw private key (64 hex chars)
        return {
          success: true,
          privateKey: keys.privateKey
        };
      }

      case 'exportPublicKey': {
        if (!password) {
          return { success: false, error: 'Password required or session expired' };
        }

        const keys = await retrieveKeys(password);
        if (!keys) {
          return { success: false, error: 'Invalid password or no keys found' };
        }

        return { success: true, publicKey: keys.publicKey };
      }

      case 'importKeys': {
        if (!password) {
          return { success: false, error: 'Password required or session expired' };
        }

        const rawPrivateKeyHex = request.privateKey.trim();

        // Validate it's exactly 64 hex characters
        if (!/^[0-9a-fA-F]{64}$/.test(rawPrivateKeyHex)) {
          return { success: false, error: 'Invalid private key format. Expected 64 hex characters.' };
        }

        try {
          // Import private key and derive public key
          const privateKey = await importRawPrivateKey(rawPrivateKeyHex);

          // Export the full JWK to get a complete keypair that we can use
          const privateJwk = await crypto.subtle.exportKey('jwk', privateKey);

          // Create public key from the coordinates
          const publicJwk = {
            kty: privateJwk.kty,
            crv: privateJwk.crv,
            x: privateJwk.x,
            y: privateJwk.y,
            ext: true
          };

          const publicKey = await crypto.subtle.importKey(
            'jwk',
            publicJwk,
            { name: 'ECDSA', namedCurve: 'P-256' },
            true,
            ['verify']
          );

          // Export public key in SPKI format for storage
          const publicKeyRaw = await crypto.subtle.exportKey('spki', publicKey);

          // Store both keys
          await storeKeys({
            privateKey: rawPrivateKeyHex,
            publicKey: bufferToHex(publicKeyRaw)
          }, password);

          // Always create session so webapp can access the public key
          await createSession(password);

          return { success: true };
        } catch (error) {
          return { success: false, error: 'Failed to import private key: ' + error.message };
        }
      }

      case 'hasKeys': {
        const hasKeys = await hasStoredKeys();
        return { success: true, hasKeys };
      }

      case 'getPublicKey': {
        // Public key is always accessible (no password/session required)
        // because public keys are meant to be public
        const publicKey = await getPublicKey();
        if (publicKey) {
          return { success: true, publicKey };
        }
        return { success: false, error: 'No public key found' };
      }

      case 'clearKeys': {
        await clearKeys();
        return { success: true };
      }

      case 'verifyPassword': {
        if (!request.password) {
          // Check if session is valid
          const session = await getSession();
          return { success: !!session };
        }

        const keys = await retrieveKeys(request.password);
        if (keys) {
          // Always create/update session
          await createSession(request.password);
        }
        return { success: !!keys };
      }

      case 'clearSession': {
        await clearSession();
        return { success: true };
      }

      case 'requestLink': {
        // Open the extension popup
        try {
          await chrome.action.openPopup();
          return { success: true, message: 'Extension popup opened' };
        } catch (error) {
          // If openPopup fails (e.g., popup already open or user interaction required),
          // fall back to setting a badge to notify the user
          console.log('Could not open popup, setting badge instead:', error);
          try {
            await chrome.action.setBadgeText({ text: '!' });
            await chrome.action.setBadgeBackgroundColor({ color: '#667eea' });
            // Clear badge after 30 seconds
            setTimeout(async () => {
              await chrome.action.setBadgeText({ text: '' });
            }, 30000);
          } catch (badgeError) {
            console.log('Could not set badge:', badgeError);
          }
          return { success: true, message: 'Please click the Skavenger extension icon to set up or unlock your account' };
        }
      }

      case 'getExtensionId': {
        // Return extension ID for webapp discovery
        return { success: true, extensionId: chrome.runtime.id };
      }

      case 'decryptElGamal': {
        if (!password) {
          return { success: false, error: 'Password required or session expired' };
        }

        if (!request.encryptedHex || !request.rValueHex) {
          return { success: false, error: 'Missing required parameters: encryptedHex and rValueHex' };
        }

        const keys = await retrieveKeys(password);
        if (!keys) {
          return { success: false, error: 'Invalid password or no keys found' };
        }

        try {
          const plaintext = await decryptElGamal(
            request.encryptedHex,
            request.rValueHex,
            keys.privateKey
          );
          return { success: true, plaintext };
        } catch (error) {
          return { success: false, error: error.message };
        }
      }

      default:
        return { success: false, error: 'Unknown action' };
    }
  } catch (error) {
    return { success: false, error: error.message };
  }
}
