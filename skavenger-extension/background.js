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

// Convert hex string to Uint8Array
function hexToBytes(hex) {
  const bytes = new Uint8Array(hex.length / 2);
  for (let i = 0; i < hex.length; i += 2) {
    bytes[i / 2] = parseInt(hex.substr(i, 2), 16);
  }
  return bytes;
}

// Keccak-256 implementation (legacy Keccak, as used by Go's sha3 package)
// Extracted from js-sha3 v0.9.3 (MIT License) by Chen, Yi-Cyuan
// https://github.com/emn178/js-sha3
function keccak256(message) {
  const KECCAK_PADDING = [1, 256, 65536, 16777216];
  const SHIFT = [0, 8, 16, 24];
  const RC = [1, 0, 32898, 0, 32906, 2147483648, 2147516416, 2147483648, 32907, 0, 2147483649,
    0, 2147516545, 2147483648, 32777, 2147483648, 138, 0, 136, 0, 2147516425, 0,
    2147483658, 0, 2147516555, 0, 139, 2147483648, 32905, 2147483648, 32771,
    2147483648, 32770, 2147483648, 128, 2147483648, 32778, 0, 2147483658, 2147483648,
    2147516545, 2147483648, 32896, 2147483648, 2147483649, 0, 2147516424, 2147483648];

  const bits = 256;
  const padding = KECCAK_PADDING;
  const outputBits = 256;

  const blocks = [];
  const s = [];
  for (let idx = 0; idx < 50; ++idx) {
    s[idx] = 0;
  }

  let block = 0;
  let start = 0;
  let reset = true;
  const blockCount = (1600 - (bits << 1)) >> 5;
  const byteCount = blockCount << 2;
  const outputBlocks = outputBits >> 5;
  const extraBytes = (outputBits & 31) >> 3;

  // Update phase
  const length = message.length;
  let index = 0;
  let lastByteIndex = 0;

  while (index < length) {
    if (reset) {
      reset = false;
      blocks[0] = block;
      for (let idx = 1; idx < blockCount + 1; ++idx) {
        blocks[idx] = 0;
      }
    }

    let i;
    for (i = start; index < length && i < byteCount; ++index) {
      blocks[i >> 2] |= message[index] << SHIFT[i++ & 3];
    }

    lastByteIndex = i;
    if (i >= byteCount) {
      start = i - byteCount;
      block = blocks[blockCount];
      for (let idx = 0; idx < blockCount; ++idx) {
        s[idx] ^= blocks[idx];
      }
      f(s);
      reset = true;
    } else {
      start = i;
    }
  }

  // Finalize phase
  let i = lastByteIndex;
  blocks[i >> 2] |= padding[i & 3];
  if (lastByteIndex === byteCount) {
    blocks[0] = blocks[blockCount];
    for (i = 1; i < blockCount + 1; ++i) {
      blocks[i] = 0;
    }
  }
  blocks[blockCount - 1] |= 0x80000000;
  for (i = 0; i < blockCount; ++i) {
    s[i] ^= blocks[i];
  }
  f(s);

  // Digest phase
  const array = [];
  let offset, block2;
  i = 0;
  let j = 0;
  while (j < outputBlocks) {
    for (i = 0; i < blockCount && j < outputBlocks; ++i, ++j) {
      offset = j << 2;
      block2 = s[i];
      array[offset] = block2 & 0xFF;
      array[offset + 1] = (block2 >> 8) & 0xFF;
      array[offset + 2] = (block2 >> 16) & 0xFF;
      array[offset + 3] = (block2 >> 24) & 0xFF;
    }
    if (j % blockCount === 0) {
      f(s);
    }
  }
  if (extraBytes) {
    offset = j << 2;
    block2 = s[i];
    array[offset] = block2 & 0xFF;
    if (extraBytes > 1) {
      array[offset + 1] = (block2 >> 8) & 0xFF;
    }
    if (extraBytes > 2) {
      array[offset + 2] = (block2 >> 16) & 0xFF;
    }
  }

  function f(s) {
    let h, l, n, c0, c1, c2, c3, c4, c5, c6, c7, c8, c9;
    let b0, b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11, b12, b13, b14, b15, b16, b17;
    let b18, b19, b20, b21, b22, b23, b24, b25, b26, b27, b28, b29, b30, b31, b32, b33;
    let b34, b35, b36, b37, b38, b39, b40, b41, b42, b43, b44, b45, b46, b47, b48, b49;

    for (n = 0; n < 48; n += 2) {
      c0 = s[0] ^ s[10] ^ s[20] ^ s[30] ^ s[40];
      c1 = s[1] ^ s[11] ^ s[21] ^ s[31] ^ s[41];
      c2 = s[2] ^ s[12] ^ s[22] ^ s[32] ^ s[42];
      c3 = s[3] ^ s[13] ^ s[23] ^ s[33] ^ s[43];
      c4 = s[4] ^ s[14] ^ s[24] ^ s[34] ^ s[44];
      c5 = s[5] ^ s[15] ^ s[25] ^ s[35] ^ s[45];
      c6 = s[6] ^ s[16] ^ s[26] ^ s[36] ^ s[46];
      c7 = s[7] ^ s[17] ^ s[27] ^ s[37] ^ s[47];
      c8 = s[8] ^ s[18] ^ s[28] ^ s[38] ^ s[48];
      c9 = s[9] ^ s[19] ^ s[29] ^ s[39] ^ s[49];

      h = c8 ^ ((c2 << 1) | (c3 >>> 31));
      l = c9 ^ ((c3 << 1) | (c2 >>> 31));
      s[0] ^= h;
      s[1] ^= l;
      s[10] ^= h;
      s[11] ^= l;
      s[20] ^= h;
      s[21] ^= l;
      s[30] ^= h;
      s[31] ^= l;
      s[40] ^= h;
      s[41] ^= l;
      h = c0 ^ ((c4 << 1) | (c5 >>> 31));
      l = c1 ^ ((c5 << 1) | (c4 >>> 31));
      s[2] ^= h;
      s[3] ^= l;
      s[12] ^= h;
      s[13] ^= l;
      s[22] ^= h;
      s[23] ^= l;
      s[32] ^= h;
      s[33] ^= l;
      s[42] ^= h;
      s[43] ^= l;
      h = c2 ^ ((c6 << 1) | (c7 >>> 31));
      l = c3 ^ ((c7 << 1) | (c6 >>> 31));
      s[4] ^= h;
      s[5] ^= l;
      s[14] ^= h;
      s[15] ^= l;
      s[24] ^= h;
      s[25] ^= l;
      s[34] ^= h;
      s[35] ^= l;
      s[44] ^= h;
      s[45] ^= l;
      h = c4 ^ ((c8 << 1) | (c9 >>> 31));
      l = c5 ^ ((c9 << 1) | (c8 >>> 31));
      s[6] ^= h;
      s[7] ^= l;
      s[16] ^= h;
      s[17] ^= l;
      s[26] ^= h;
      s[27] ^= l;
      s[36] ^= h;
      s[37] ^= l;
      s[46] ^= h;
      s[47] ^= l;
      h = c6 ^ ((c0 << 1) | (c1 >>> 31));
      l = c7 ^ ((c1 << 1) | (c0 >>> 31));
      s[8] ^= h;
      s[9] ^= l;
      s[18] ^= h;
      s[19] ^= l;
      s[28] ^= h;
      s[29] ^= l;
      s[38] ^= h;
      s[39] ^= l;
      s[48] ^= h;
      s[49] ^= l;

      b0 = s[0];
      b1 = s[1];
      b32 = (s[11] << 4) | (s[10] >>> 28);
      b33 = (s[10] << 4) | (s[11] >>> 28);
      b14 = (s[20] << 3) | (s[21] >>> 29);
      b15 = (s[21] << 3) | (s[20] >>> 29);
      b46 = (s[31] << 9) | (s[30] >>> 23);
      b47 = (s[30] << 9) | (s[31] >>> 23);
      b28 = (s[40] << 18) | (s[41] >>> 14);
      b29 = (s[41] << 18) | (s[40] >>> 14);
      b20 = (s[2] << 1) | (s[3] >>> 31);
      b21 = (s[3] << 1) | (s[2] >>> 31);
      b2 = (s[13] << 12) | (s[12] >>> 20);
      b3 = (s[12] << 12) | (s[13] >>> 20);
      b34 = (s[22] << 10) | (s[23] >>> 22);
      b35 = (s[23] << 10) | (s[22] >>> 22);
      b16 = (s[33] << 13) | (s[32] >>> 19);
      b17 = (s[32] << 13) | (s[33] >>> 19);
      b48 = (s[42] << 2) | (s[43] >>> 30);
      b49 = (s[43] << 2) | (s[42] >>> 30);
      b40 = (s[5] << 30) | (s[4] >>> 2);
      b41 = (s[4] << 30) | (s[5] >>> 2);
      b22 = (s[14] << 6) | (s[15] >>> 26);
      b23 = (s[15] << 6) | (s[14] >>> 26);
      b4 = (s[25] << 11) | (s[24] >>> 21);
      b5 = (s[24] << 11) | (s[25] >>> 21);
      b36 = (s[34] << 15) | (s[35] >>> 17);
      b37 = (s[35] << 15) | (s[34] >>> 17);
      b18 = (s[45] << 29) | (s[44] >>> 3);
      b19 = (s[44] << 29) | (s[45] >>> 3);
      b10 = (s[6] << 28) | (s[7] >>> 4);
      b11 = (s[7] << 28) | (s[6] >>> 4);
      b42 = (s[17] << 23) | (s[16] >>> 9);
      b43 = (s[16] << 23) | (s[17] >>> 9);
      b24 = (s[26] << 25) | (s[27] >>> 7);
      b25 = (s[27] << 25) | (s[26] >>> 7);
      b6 = (s[36] << 21) | (s[37] >>> 11);
      b7 = (s[37] << 21) | (s[36] >>> 11);
      b38 = (s[47] << 24) | (s[46] >>> 8);
      b39 = (s[46] << 24) | (s[47] >>> 8);
      b30 = (s[8] << 27) | (s[9] >>> 5);
      b31 = (s[9] << 27) | (s[8] >>> 5);
      b12 = (s[18] << 20) | (s[19] >>> 12);
      b13 = (s[19] << 20) | (s[18] >>> 12);
      b44 = (s[29] << 7) | (s[28] >>> 25);
      b45 = (s[28] << 7) | (s[29] >>> 25);
      b26 = (s[38] << 8) | (s[39] >>> 24);
      b27 = (s[39] << 8) | (s[38] >>> 24);
      b8 = (s[48] << 14) | (s[49] >>> 18);
      b9 = (s[49] << 14) | (s[48] >>> 18);

      s[0] = b0 ^ (~b2 & b4);
      s[1] = b1 ^ (~b3 & b5);
      s[10] = b10 ^ (~b12 & b14);
      s[11] = b11 ^ (~b13 & b15);
      s[20] = b20 ^ (~b22 & b24);
      s[21] = b21 ^ (~b23 & b25);
      s[30] = b30 ^ (~b32 & b34);
      s[31] = b31 ^ (~b33 & b35);
      s[40] = b40 ^ (~b42 & b44);
      s[41] = b41 ^ (~b43 & b45);
      s[2] = b2 ^ (~b4 & b6);
      s[3] = b3 ^ (~b5 & b7);
      s[12] = b12 ^ (~b14 & b16);
      s[13] = b13 ^ (~b15 & b17);
      s[22] = b22 ^ (~b24 & b26);
      s[23] = b23 ^ (~b25 & b27);
      s[32] = b32 ^ (~b34 & b36);
      s[33] = b33 ^ (~b35 & b37);
      s[42] = b42 ^ (~b44 & b46);
      s[43] = b43 ^ (~b45 & b47);
      s[4] = b4 ^ (~b6 & b8);
      s[5] = b5 ^ (~b7 & b9);
      s[14] = b14 ^ (~b16 & b18);
      s[15] = b15 ^ (~b17 & b19);
      s[24] = b24 ^ (~b26 & b28);
      s[25] = b25 ^ (~b27 & b29);
      s[34] = b34 ^ (~b36 & b38);
      s[35] = b35 ^ (~b37 & b39);
      s[44] = b44 ^ (~b46 & b48);
      s[45] = b45 ^ (~b47 & b49);
      s[6] = b6 ^ (~b8 & b0);
      s[7] = b7 ^ (~b9 & b1);
      s[16] = b16 ^ (~b18 & b10);
      s[17] = b17 ^ (~b19 & b11);
      s[26] = b26 ^ (~b28 & b20);
      s[27] = b27 ^ (~b29 & b21);
      s[36] = b36 ^ (~b38 & b30);
      s[37] = b37 ^ (~b39 & b31);
      s[46] = b46 ^ (~b48 & b40);
      s[47] = b47 ^ (~b49 & b41);
      s[8] = b8 ^ (~b0 & b2);
      s[9] = b9 ^ (~b1 & b3);
      s[18] = b18 ^ (~b10 & b12);
      s[19] = b19 ^ (~b11 & b13);
      s[28] = b28 ^ (~b20 & b22);
      s[29] = b29 ^ (~b21 & b23);
      s[38] = b38 ^ (~b30 & b32);
      s[39] = b39 ^ (~b31 & b33);
      s[48] = b48 ^ (~b40 & b42);
      s[49] = b49 ^ (~b41 & b43);

      s[0] ^= RC[n];
      s[1] ^= RC[n + 1];
    }
  }

  return new Uint8Array(array);
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

// Generate secp256k1 key pair
async function generateKeyPair() {
  // Generate a random 32-byte private key
  const privateKeyBytes = new Uint8Array(32);
  crypto.getRandomValues(privateKeyBytes);

  // Convert to BigInt to check validity
  const privateKeyD = BigInt('0x' + Array.from(privateKeyBytes).map(b => b.toString(16).padStart(2, '0')).join(''));

  // Ensure it's in valid range [1, n-1]
  if (privateKeyD === 0n || privateKeyD >= SECP256K1_N) {
    // Rare case, try again
    return generateKeyPair();
  }

  // Convert to hex string
  const privateKeyHex = Array.from(privateKeyBytes).map(b => b.toString(16).padStart(2, '0')).join('');

  // Derive public key
  const publicKeyHex = derivePublicKey(privateKeyHex);

  return {
    privateKey: privateKeyHex,
    publicKey: publicKeyHex
  };
}

// secp256k1 curve parameters (Bitcoin/Ethereum curve)
// We use secp256k1 for all cryptographic operations (signing, encryption, key generation)
const SECP256K1_P = BigInt('0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F');
const SECP256K1_A = BigInt(0); // a = 0 for secp256k1
const SECP256K1_N = BigInt('0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141'); // curve order
const SECP256K1_GX = BigInt('0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798'); // generator point X
const SECP256K1_GY = BigInt('0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8'); // generator point Y

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

// EC point addition on secp256k1
function pointAddSecp256k1(p1, p2) {
  if (p1 === null) return p2;
  if (p2 === null) return p1;

  const [x1, y1] = p1;
  const [x2, y2] = p2;

  if (x1 === x2 && y1 === y2) {
    // Point doubling
    // For secp256k1, a = 0, so the formula simplifies: s = 3*x1^2 / (2*y1)
    const s = (3n * x1 * x1 + SECP256K1_A) * modInverse(2n * y1, SECP256K1_P) % SECP256K1_P;
    const x3 = (s * s - 2n * x1) % SECP256K1_P;
    const y3 = (s * (x1 - x3) - y1) % SECP256K1_P;
    return [(x3 + SECP256K1_P) % SECP256K1_P, (y3 + SECP256K1_P) % SECP256K1_P];
  }

  const s = (y2 - y1) * modInverse((x2 - x1 + SECP256K1_P) % SECP256K1_P, SECP256K1_P) % SECP256K1_P;
  const x3 = (s * s - x1 - x2) % SECP256K1_P;
  const y3 = (s * (x1 - x3) - y1) % SECP256K1_P;
  return [(x3 + SECP256K1_P) % SECP256K1_P, (y3 + SECP256K1_P) % SECP256K1_P];
}

// Scalar multiplication on secp256k1
function scalarMultSecp256k1(k, point) {
  let result = null;
  let addend = point;

  while (k > 0n) {
    if (k & 1n) {
      result = pointAddSecp256k1(result, addend);
    }
    addend = pointAddSecp256k1(addend, addend);
    k >>= 1n;
  }

  return result;
}

// HMAC-SHA256 implementation
async function hmacSHA256(key, message) {
  const cryptoKey = await crypto.subtle.importKey(
    'raw',
    key,
    { name: 'HMAC', hash: 'SHA-256' },
    false,
    ['sign']
  );
  const signature = await crypto.subtle.sign('HMAC', cryptoKey, message);
  return new Uint8Array(signature);
}

// RFC 6979 deterministic k-value generation for ECDSA
async function generateDeterministicK(privateKeyBytes, messageHash) {
  // RFC 6979 implementation
  const hlen = 32; // SHA-256 hash length
  const qlen = 32; // secp256k1 order length in bytes

  // Step a: h1 = H(m) (already done, messageHash is the hash)
  const h1 = new Uint8Array(messageHash);

  // Step b: V = 0x01 0x01 0x01 ... 0x01 (hlen bytes)
  let V = new Uint8Array(hlen);
  V.fill(0x01);

  // Step c: K = 0x00 0x00 0x00 ... 0x00 (hlen bytes)
  let K = new Uint8Array(hlen);
  K.fill(0x00);

  // Step d: K = HMAC_K(V || 0x00 || int2octets(x) || bits2octets(h1))
  let data = new Uint8Array(V.length + 1 + privateKeyBytes.length + h1.length);
  data.set(V, 0);
  data[V.length] = 0x00;
  data.set(privateKeyBytes, V.length + 1);
  data.set(h1, V.length + 1 + privateKeyBytes.length);
  K = await hmacSHA256(K, data);

  // Step e: V = HMAC_K(V)
  V = await hmacSHA256(K, V);

  // Step f: K = HMAC_K(V || 0x01 || int2octets(x) || bits2octets(h1))
  data = new Uint8Array(V.length + 1 + privateKeyBytes.length + h1.length);
  data.set(V, 0);
  data[V.length] = 0x01;
  data.set(privateKeyBytes, V.length + 1);
  data.set(h1, V.length + 1 + privateKeyBytes.length);
  K = await hmacSHA256(K, data);

  // Step g: V = HMAC_K(V)
  V = await hmacSHA256(K, V);

  // Step h: Generate k
  while (true) {
    // 1. Set T to empty
    let T = new Uint8Array(0);

    // 2. While tlen < qlen, do: V = HMAC_K(V), T = T || V
    while (T.length < qlen) {
      V = await hmacSHA256(K, V);
      const newT = new Uint8Array(T.length + V.length);
      newT.set(T);
      newT.set(V, T.length);
      T = newT;
    }

    // 3. k = bits2int(T)
    const kBytes = T.slice(0, qlen);
    const k = BigInt('0x' + Array.from(kBytes).map(b => b.toString(16).padStart(2, '0')).join(''));

    // 4. If k is in [1, q-1], return it
    if (k >= 1n && k < SECP256K1_N) {
      return k;
    }

    // Otherwise, update K and V and try again
    data = new Uint8Array(V.length + 1);
    data.set(V, 0);
    data[V.length] = 0x00;
    K = await hmacSHA256(K, data);
    V = await hmacSHA256(K, V);
  }
}

// ECDSA signing with secp256k1
async function signSecp256k1(privateKeyHex, messageHash) {
  const privateKeyD = BigInt('0x' + privateKeyHex);

  // Ensure private key is in valid range
  if (privateKeyD <= 0n || privateKeyD >= SECP256K1_N) {
    throw new Error('Invalid private key');
  }

  // Convert private key and message hash to bytes
  const privateKeyBytes = hexToBytes(privateKeyHex);
  const messageHashBytes = new Uint8Array(messageHash);

  // Generate deterministic k using RFC 6979
  const k = await generateDeterministicK(privateKeyBytes, messageHashBytes);

  // Calculate r = (k * G).x mod n
  const [Rx, Ry] = scalarMultSecp256k1(k, [SECP256K1_GX, SECP256K1_GY]);
  const r = Rx % SECP256K1_N;

  if (r === 0n) {
    throw new Error('Invalid r value (should never happen with RFC 6979)');
  }

  // Calculate e = H(m) as BigInt
  const e = BigInt('0x' + Array.from(messageHashBytes).map(b => b.toString(16).padStart(2, '0')).join(''));

  // Calculate s = k^-1 * (e + r * privateKey) mod n
  const kInv = modInverse(k, SECP256K1_N);
  let s = (kInv * (e + r * privateKeyD)) % SECP256K1_N;

  // Ensure s is in lower half (BIP 62: use low-s values)
  if (s > SECP256K1_N / 2n) {
    s = SECP256K1_N - s;
  }

  if (s === 0n) {
    throw new Error('Invalid s value (should never happen with RFC 6979)');
  }

  return { r, s };
}

// Encode ECDSA signature to ASN.1 DER format (to match Go backend)
function encodeSignatureDER(r, s) {
  // Convert r and s to byte arrays (big-endian, minimal encoding)
  function bigIntToBytes(value) {
    let hex = value.toString(16);
    // Ensure even length
    if (hex.length % 2 !== 0) {
      hex = '0' + hex;
    }
    // Add leading 0x00 if high bit is set (to indicate positive number)
    if (parseInt(hex.substring(0, 2), 16) >= 0x80) {
      hex = '00' + hex;
    }
    return hexToBytes(hex);
  }

  const rBytes = bigIntToBytes(r);
  const sBytes = bigIntToBytes(s);

  // Build ASN.1 DER structure
  // SEQUENCE { INTEGER r, INTEGER s }
  const rEncoded = new Uint8Array(2 + rBytes.length);
  rEncoded[0] = 0x02; // INTEGER tag
  rEncoded[1] = rBytes.length;
  rEncoded.set(rBytes, 2);

  const sEncoded = new Uint8Array(2 + sBytes.length);
  sEncoded[0] = 0x02; // INTEGER tag
  sEncoded[1] = sBytes.length;
  sEncoded.set(sBytes, 2);

  const sequence = new Uint8Array(2 + rEncoded.length + sEncoded.length);
  sequence[0] = 0x30; // SEQUENCE tag
  sequence[1] = rEncoded.length + sEncoded.length;
  sequence.set(rEncoded, 2);
  sequence.set(sEncoded, 2 + rEncoded.length);

  return sequence;
}

// Compute public key from private key on secp256k1
function derivePublicKey(privateKeyHex) {
  const d = BigInt('0x' + privateKeyHex);

  // Ensure private key is in valid range
  if (d <= 0n || d >= SECP256K1_N) {
    throw new Error('Invalid private key');
  }

  const [x, y] = scalarMultSecp256k1(d, [SECP256K1_GX, SECP256K1_GY]);

  // Convert to uncompressed format (0x04 || X || Y) - 65 bytes total
  const xHex = x.toString(16).padStart(64, '0');
  const yHex = y.toString(16).padStart(64, '0');

  // Return as hex string in uncompressed format
  return '0x04' + xHex + yHex;
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

// Marshal EC point to bytes (uncompressed format)
// Format: 0x04 | X (32 bytes) | Y (32 bytes)
function marshalECPoint(x, y) {
  const xHex = x.toString(16).padStart(64, '0');
  const yHex = y.toString(16).padStart(64, '0');
  const pointHex = '04' + xHex + yHex;
  return hexToBytes(pointHex);
}

// Generate a cryptographically secure random BigInt in range [1, max-1]
function generateRandomBigInt(max) {
  const bytes = new Uint8Array(32);
  crypto.getRandomValues(bytes);
  let value = 0n;
  for (let i = 0; i < bytes.length; i++) {
    value = (value << 8n) | BigInt(bytes[i]);
  }
  // Ensure value is in range [1, max-1]
  value = (value % (max - 1n)) + 1n;
  return value;
}

// Encrypt ElGamal - mirrors Go's EncryptElGamal function
// This encrypts a message for a recipient using their public key
function encryptElGamal(message, recipientPubKeyHex, r) {
  // Parse recipient public key (format: 0x04 or 04 || X || Y)
  let pubKeyHex = recipientPubKeyHex;
  if (pubKeyHex.startsWith('0x')) {
    pubKeyHex = pubKeyHex.slice(2);
  }
  if (!pubKeyHex.startsWith('04')) {
    throw new Error('Invalid public key format - expected uncompressed format');
  }

  const pubKeyBytes = hexToBytes(pubKeyHex);
  const [recipientPubX, recipientPubY] = parseECPoint(pubKeyBytes);

  // Ensure r is valid
  if (r <= 0n || r >= SECP256K1_N) {
    throw new Error('Invalid r value');
  }

  // C1 = g^r (ephemeral public key)
  const [c1x, c1y] = scalarMultSecp256k1(r, [SECP256K1_GX, SECP256K1_GY]);
  const c1Bytes = marshalECPoint(c1x, c1y);

  // Shared secret: S = recipientPub^r
  const [sx, sy] = scalarMultSecp256k1(r, [recipientPubX, recipientPubY]);
  const sharedSecretBytes = marshalECPoint(sx, sy);

  // Derive symmetric key from r AND the shared secret
  // key = Keccak256(r || sx)
  const rBytes = bigIntToBytes(r);
  const sxBytes = bigIntToBytes(sx);
  const keyInput = new Uint8Array(rBytes.length + sxBytes.length);
  keyInput.set(rBytes, 0);
  keyInput.set(sxBytes, rBytes.length);
  const key = keccak256(keyInput);

  // XOR message with key to get C2
  const messageBytes = typeof message === 'string'
    ? new TextEncoder().encode(message)
    : new Uint8Array(message);
  const c2 = new Uint8Array(messageBytes.length);
  for (let i = 0; i < messageBytes.length; i++) {
    c2[i] = messageBytes[i] ^ key[i % key.length];
  }

  return {
    c1: c1Bytes,
    c2: c2,
    sharedSecret: sharedSecretBytes
  };
}

// Marshal ElGamalCiphertext to bytes
// Format: len(C1) | C1 | len(C2) | C2 | len(SharedSecret) | SharedSecret
function marshalElGamalCiphertext(ciphertext) {
  const result = [];

  // C1 length (4 bytes, big endian)
  result.push((ciphertext.c1.length >> 24) & 0xff);
  result.push((ciphertext.c1.length >> 16) & 0xff);
  result.push((ciphertext.c1.length >> 8) & 0xff);
  result.push(ciphertext.c1.length & 0xff);
  // C1 data
  for (let i = 0; i < ciphertext.c1.length; i++) {
    result.push(ciphertext.c1[i]);
  }

  // C2 length (4 bytes, big endian)
  result.push((ciphertext.c2.length >> 24) & 0xff);
  result.push((ciphertext.c2.length >> 16) & 0xff);
  result.push((ciphertext.c2.length >> 8) & 0xff);
  result.push(ciphertext.c2.length & 0xff);
  // C2 data
  for (let i = 0; i < ciphertext.c2.length; i++) {
    result.push(ciphertext.c2[i]);
  }

  // SharedSecret length (4 bytes, big endian)
  result.push((ciphertext.sharedSecret.length >> 24) & 0xff);
  result.push((ciphertext.sharedSecret.length >> 16) & 0xff);
  result.push((ciphertext.sharedSecret.length >> 8) & 0xff);
  result.push(ciphertext.sharedSecret.length & 0xff);
  // SharedSecret data
  for (let i = 0; i < ciphertext.sharedSecret.length; i++) {
    result.push(ciphertext.sharedSecret[i]);
  }

  return new Uint8Array(result);
}

// Generate DLEQ proof
// Proves that log_g(C1) == log_sellerPub(S_seller) == log_buyerPub(S_buyer)
// i.e., the same r was used for all operations
function generateDLEQProof(r, sellerPubKeyHex, buyerPubKeyHex, sellerCipher, buyerCipher) {
  // Parse seller public key
  let sellerPubHex = sellerPubKeyHex;
  if (sellerPubHex.startsWith('0x')) {
    sellerPubHex = sellerPubHex.slice(2);
  }
  const sellerPubBytes = hexToBytes(sellerPubHex);
  const [sellerPubX, sellerPubY] = parseECPoint(sellerPubBytes);

  // Parse buyer public key
  let buyerPubHex = buyerPubKeyHex;
  if (buyerPubHex.startsWith('0x')) {
    buyerPubHex = buyerPubHex.slice(2);
  }
  const buyerPubBytes = hexToBytes(buyerPubHex);
  const [buyerPubX, buyerPubY] = parseECPoint(buyerPubBytes);

  // Generate random w
  const w = generateRandomBigInt(SECP256K1_N);

  // A1 = g^w
  const [a1x, a1y] = scalarMultSecp256k1(w, [SECP256K1_GX, SECP256K1_GY]);
  const a1Bytes = marshalECPoint(a1x, a1y);

  // A2 = sellerPub^w
  const [a2x, a2y] = scalarMultSecp256k1(w, [sellerPubX, sellerPubY]);
  const a2Bytes = marshalECPoint(a2x, a2y);

  // A3 = buyerPub^w
  const [a3x, a3y] = scalarMultSecp256k1(w, [buyerPubX, buyerPubY]);
  const a3Bytes = marshalECPoint(a3x, a3y);

  // Compute r hash commitment: rHash = Keccak256(r)
  const rBytes = bigIntToBytes(r);
  const rHash = keccak256(rBytes);

  // Generate challenge: c = Keccak256(sellerPub || buyerPub || A1 || A2 || A3 || rHash)
  const challengeInput = new Uint8Array(
    sellerPubBytes.length + buyerPubBytes.length +
    a1Bytes.length + a2Bytes.length + a3Bytes.length + 32
  );
  let offset = 0;
  challengeInput.set(sellerPubBytes, offset); offset += sellerPubBytes.length;
  challengeInput.set(buyerPubBytes, offset); offset += buyerPubBytes.length;
  challengeInput.set(a1Bytes, offset); offset += a1Bytes.length;
  challengeInput.set(a2Bytes, offset); offset += a2Bytes.length;
  challengeInput.set(a3Bytes, offset); offset += a3Bytes.length;
  challengeInput.set(rHash, offset);

  const challengeBytes = keccak256(challengeInput);
  let c = 0n;
  for (let i = 0; i < challengeBytes.length; i++) {
    c = (c << 8n) | BigInt(challengeBytes[i]);
  }
  c = c % SECP256K1_N;

  // Response: z = w + c*r mod n
  let z = (w + c * r) % SECP256K1_N;

  return {
    a1: a1Bytes,
    a2: a2Bytes,
    a3: a3Bytes,
    z: z,
    c: c,
    rHash: rHash
  };
}

// Marshal DLEQ proof to bytes
// Format: len(A1) | A1 | len(A2) | A2 | len(A3) | A3 | len(Z) | Z | len(C) | C | RHash
function marshalDLEQProof(proof) {
  const result = [];

  // Helper to add length-prefixed data
  const addLengthPrefixed = (data) => {
    result.push((data.length >> 24) & 0xff);
    result.push((data.length >> 16) & 0xff);
    result.push((data.length >> 8) & 0xff);
    result.push(data.length & 0xff);
    for (let i = 0; i < data.length; i++) {
      result.push(data[i]);
    }
  };

  // A1
  addLengthPrefixed(proof.a1);

  // A2
  addLengthPrefixed(proof.a2);

  // A3
  addLengthPrefixed(proof.a3);

  // Z (BigInt to bytes)
  let zHex = proof.z.toString(16);
  if (zHex.length % 2 !== 0) zHex = '0' + zHex;
  const zBytes = hexToBytes(zHex);
  addLengthPrefixed(zBytes);

  // C (BigInt to bytes)
  let cHex = proof.c.toString(16);
  if (cHex.length % 2 !== 0) cHex = '0' + cHex;
  const cBytes = hexToBytes(cHex);
  addLengthPrefixed(cBytes);

  // RHash (fixed 32 bytes)
  for (let i = 0; i < proof.rHash.length; i++) {
    result.push(proof.rHash[i]);
  }

  return new Uint8Array(result);
}

// Generate plaintext equality proof
// This proves that the buyer cipher encrypts the same plaintext as the original cipher
// by committing to KeyBuyer = Hash(newR || S_buyer.X)
function generatePlaintextEqualityProof(newR, buyerCipher) {
  // Extract S_buyer.X from buyer cipher's shared secret
  // The shared secret is in the format: 0x04 | X (32 bytes) | Y (32 bytes)
  const [sBuyerX, ] = parseECPoint(buyerCipher.sharedSecret);

  // Compute KeyBuyer = Keccak256(newR || S_buyer.X)
  // This is the actual key used to encrypt buyerCipher.c2
  const newRBytes = bigIntToBytes(newR);
  const sBuyerXBytes = bigIntToBytes(sBuyerX);
  const keyInput = new Uint8Array(newRBytes.length + sBuyerXBytes.length);
  keyInput.set(newRBytes, 0);
  keyInput.set(sBuyerXBytes, newRBytes.length);
  const keyBuyer = keccak256(keyInput);

  return {
    keyBuyer: keyBuyer  // 32 bytes
  };
}

// Marshal plaintext equality proof to bytes
// Format: KeyBuyer (fixed 32 bytes)
function marshalPlaintextProof(proof) {
  return new Uint8Array(proof.keyBuyer);
}

// Marshal TransferProof to bytes (DLEQ + Plaintext)
// Format: len(DLEQ) | DLEQ | Plaintext (32 bytes)
function marshalTransferProof(proof) {
  const result = [];

  // Marshal DLEQ proof
  const dleqBytes = marshalDLEQProof(proof.dleq);

  // DLEQ length (4 bytes, big endian)
  result.push((dleqBytes.length >> 24) & 0xff);
  result.push((dleqBytes.length >> 16) & 0xff);
  result.push((dleqBytes.length >> 8) & 0xff);
  result.push(dleqBytes.length & 0xff);

  // DLEQ data
  for (let i = 0; i < dleqBytes.length; i++) {
    result.push(dleqBytes[i]);
  }

  // Plaintext proof (fixed 32 bytes)
  const plaintextBytes = marshalPlaintextProof(proof.plaintext);
  for (let i = 0; i < plaintextBytes.length; i++) {
    result.push(plaintextBytes[i]);
  }

  return new Uint8Array(result);
}

// Generate verifiable ElGamal transfer
// Creates two ciphertexts (one for seller, one for buyer) and a unified TransferProof
// containing both DLEQ proof (same r for both ciphers) and plaintext equality proof
async function generateVerifiableElGamalTransfer(plaintext, sellerPrivKeyHex, buyerPubKeyHex) {
  // Derive seller's public key from private key
  const sellerPubKeyHex = derivePublicKey(sellerPrivKeyHex);

  // Generate random salt for commitment
  const salt = new Uint8Array(32);
  crypto.getRandomValues(salt);

  // Create commitment = Keccak256(plaintext || salt)
  const plaintextBytes = typeof plaintext === 'string'
    ? new TextEncoder().encode(plaintext)
    : new Uint8Array(plaintext);
  const commitmentInput = new Uint8Array(plaintextBytes.length + salt.length);
  commitmentInput.set(plaintextBytes, 0);
  commitmentInput.set(salt, plaintextBytes.length);
  const commitment = keccak256(commitmentInput);

  // Generate random r (used for BOTH encryptions!)
  const r = generateRandomBigInt(SECP256K1_N);

  // Encrypt for seller
  const sellerCipher = encryptElGamal(plaintextBytes, sellerPubKeyHex, r);

  // Encrypt for buyer (SAME r value!)
  const buyerCipher = encryptElGamal(plaintextBytes, buyerPubKeyHex, r);

  // Generate DLEQ proof that both ciphers use same r
  const dleqProof = generateDLEQProof(r, sellerPubKeyHex, buyerPubKeyHex, sellerCipher, buyerCipher);

  // Generate plaintext equality proof
  // This proves buyerCipher encrypts the same content as the original on-chain cipher
  const plaintextProof = generatePlaintextEqualityProof(r, buyerCipher);

  // Create consolidated proof object matching Go structure
  const proof = {
    dleq: dleqProof,
    plaintext: plaintextProof
  };

  // Parse public keys for return
  let sellerPubHex = sellerPubKeyHex;
  if (sellerPubHex.startsWith('0x')) {
    sellerPubHex = sellerPubHex.slice(2);
  }
  const sellerPubBytes = hexToBytes(sellerPubHex);

  let buyerPubHex = buyerPubKeyHex;
  if (buyerPubHex.startsWith('0x')) {
    buyerPubHex = buyerPubHex.slice(2);
  }
  const buyerPubBytes = hexToBytes(buyerPubHex);

  return {
    sellerCipher: sellerCipher,
    buyerCipher: buyerCipher,
    proof: proof,  // Consolidated proof with DLEQ and Plaintext
    commitment: commitment,
    salt: salt,
    sellerPubKey: sellerPubBytes,
    buyerPubKey: buyerPubBytes,
    sharedR: r  // Seller keeps this secret until after payment!
  };
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
    // Uses secp256k1 curve to match Go backend encryption
    const [sx, sy] = scalarMultSecp256k1(privateKeyD, [c1x, c1y]);

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
        // keyPair now contains { privateKey: hex string, publicKey: hex string (uncompressed secp256k1) }

        await storeKeys({
          privateKey: keyPair.privateKey,
          publicKey: keyPair.publicKey
        }, password);

        // Always create session so webapp can access the public key
        await createSession(password);

        return { success: true, publicKey: keyPair.publicKey };
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
          // Derive secp256k1 public key from private key
          const publicKeyHex = derivePublicKey(rawPrivateKeyHex);

          // Store both keys
          await storeKeys({
            privateKey: rawPrivateKeyHex,
            publicKey: publicKeyHex
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

      case 'signMessage': {
        if (!password) {
          return { success: false, error: 'Password required or session expired' };
        }

        if (!request.message) {
          return { success: false, error: 'Missing required parameter: message' };
        }

        const keys = await retrieveKeys(password);
        if (!keys) {
          return { success: false, error: 'Invalid password or no keys found' };
        }

        try {
          // Hash the message with SHA-256
          const encoder = new TextEncoder();
          const messageBytes = encoder.encode(request.message);
          const hashBuffer = await crypto.subtle.digest('SHA-256', messageBytes);

          // Sign using secp256k1 ECDSA
          const { r, s } = await signSecp256k1(keys.privateKey, hashBuffer);

          // Encode signature in ASN.1 DER format (to match Go backend)
          const signatureDER = encodeSignatureDER(r, s);

          // Convert to hex
          const signatureHex = '0x' + bufferToHex(signatureDER);

          return { success: true, signature: signatureHex };
        } catch (error) {
          return { success: false, error: 'Failed to sign message: ' + error.message };
        }
      }

      case 'generateTransferProof': {
        // Generates a verifiable ElGamal transfer proof for selling a clue
        // Requires: password, plaintext (the clue content), buyerPublicKey
        if (!password) {
          return { success: false, error: 'Password required or session expired' };
        }

        if (!request.plaintext) {
          return { success: false, error: 'Missing required parameter: plaintext' };
        }

        if (!request.buyerPublicKey) {
          return { success: false, error: 'Missing required parameter: buyerPublicKey' };
        }

        const keys = await retrieveKeys(password);
        if (!keys) {
          return { success: false, error: 'Invalid password or no keys found' };
        }

        try {
          console.log('=== GENERATE TRANSFER PROOF ===');
          console.log('Buyer public key:', request.buyerPublicKey);

          // Generate the verifiable transfer
          const transfer = await generateVerifiableElGamalTransfer(
            request.plaintext,
            keys.privateKey,
            request.buyerPublicKey
          );

          // Marshal the ciphertexts and proof
          const sellerCiphertextBytes = marshalElGamalCiphertext(transfer.sellerCipher);
          const buyerCiphertextBytes = marshalElGamalCiphertext(transfer.buyerCipher);

          // Marshal the unified TransferProof (DLEQ + Plaintext)
          const proofBytes = marshalTransferProof(transfer.proof);

          // Convert sharedR to hex string for storage (will be revealed later)
          let sharedRHex = transfer.sharedR.toString(16);
          if (sharedRHex.length % 2 !== 0) sharedRHex = '0' + sharedRHex;

          console.log('Seller ciphertext length:', sellerCiphertextBytes.length);
          console.log('Buyer ciphertext length:', buyerCiphertextBytes.length);
          console.log('Proof length:', proofBytes.length);
          console.log('=== END GENERATE TRANSFER PROOF ===');

          return {
            success: true,
            sellerCiphertext: '0x' + bufferToHex(sellerCiphertextBytes),
            buyerCiphertext: '0x' + bufferToHex(buyerCiphertextBytes),
            proof: '0x' + bufferToHex(proofBytes),
            sellerPublicKey: '0x' + bufferToHex(transfer.sellerPubKey),
            buyerPublicKey: '0x' + bufferToHex(transfer.buyerPubKey),
            sharedR: '0x' + sharedRHex,
            commitment: '0x' + bufferToHex(transfer.commitment)
          };
        } catch (error) {
          console.error('Failed to generate transfer proof:', error);
          return { success: false, error: 'Failed to generate transfer proof: ' + error.message };
        }
      }

      default:
        return { success: false, error: 'Unknown action' };
    }
  } catch (error) {
    return { success: false, error: error.message };
  }
}
