/**
 * ElGamal cryptography utilities for Skavenge webapp
 * Implements verification of ElGamal transfer proofs (DLEQ proofs)
 */

/* global BigInt */

// secp256k1 curve parameters
const SECP256K1_P = BigInt('0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F');
const SECP256K1_A = BigInt(0);
const SECP256K1_N = BigInt('0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141');
const SECP256K1_GX = BigInt('0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798');
const SECP256K1_GY = BigInt('0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8');

/**
 * Convert hex string to Uint8Array
 */
function hexToBytes(hex) {
  // Remove 0x prefix if present
  const cleanHex = hex.startsWith('0x') ? hex.slice(2) : hex;
  const bytes = new Uint8Array(cleanHex.length / 2);
  for (let i = 0; i < cleanHex.length; i += 2) {
    bytes[i / 2] = parseInt(cleanHex.substr(i, 2), 16);
  }
  return bytes;
}

/**
 * Keccak-256 implementation (legacy Keccak, as used by Go's sha3 package)
 * Extracted from js-sha3 v0.9.3 (MIT License) by Chen, Yi-Cyuan
 */
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

/**
 * Modular inverse using Extended Euclidean Algorithm
 */
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

/**
 * EC point addition on secp256k1
 */
function pointAddSecp256k1(p1, p2) {
  if (p1 === null) return p2;
  if (p2 === null) return p1;

  const [x1, y1] = p1;
  const [x2, y2] = p2;

  if (x1 === x2 && y1 === y2) {
    // Point doubling
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

/**
 * Scalar multiplication on secp256k1
 */
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

/**
 * Parse elliptic curve point from bytes (uncompressed format)
 * Format: 0x04 | X (32 bytes) | Y (32 bytes)
 */
function parseECPoint(bytes) {
  if (bytes.length !== 65 || bytes[0] !== 0x04) {
    throw new Error('Invalid EC point format');
  }

  const xBytes = bytes.slice(1, 33);
  const yBytes = bytes.slice(33, 65);

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

/**
 * Unmarshal ElGamal ciphertext from bytes
 * Format: len(C1) | C1 | len(C2) | C2 | len(SharedSecret) | SharedSecret
 */
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

  const c1 = bytes.slice(offset, offset + c1Len);
  offset += c1Len;

  if (offset + 4 > bytes.length) {
    throw new Error('Data too short for C2 length');
  }

  // Read C2 length
  const c2Len = (bytes[offset] << 24) | (bytes[offset + 1] << 16) | (bytes[offset + 2] << 8) | bytes[offset + 3];
  offset += 4;

  if (offset + c2Len > bytes.length) {
    throw new Error('Invalid C2 length');
  }

  const c2 = bytes.slice(offset, offset + c2Len);
  offset += c2Len;

  if (offset + 4 > bytes.length) {
    throw new Error('Data too short for SharedSecret length');
  }

  // Read SharedSecret length
  const ssLen = (bytes[offset] << 24) | (bytes[offset + 1] << 16) | (bytes[offset + 2] << 8) | bytes[offset + 3];
  offset += 4;

  if (offset + ssLen > bytes.length) {
    throw new Error('Invalid SharedSecret length');
  }

  const sharedSecret = bytes.slice(offset, offset + ssLen);

  return { c1, c2, sharedSecret };
}

/**
 * Unmarshal DLEQ proof from bytes
 * Format: len(A1) | A1 | len(A2) | A2 | len(A3) | A3 | len(Z) | Z | len(C) | C | RHash
 */
function unmarshalDLEQProof(data) {
  const bytes = new Uint8Array(data);
  let offset = 0;

  if (bytes.length < 20) {
    throw new Error('Proof data too short');
  }

  // Helper to read length-prefixed data
  const readLengthPrefixed = () => {
    if (offset + 4 > bytes.length) {
      throw new Error('Data too short for length prefix');
    }
    const len = (bytes[offset] << 24) | (bytes[offset + 1] << 16) | (bytes[offset + 2] << 8) | bytes[offset + 3];
    offset += 4;
    if (offset + len > bytes.length) {
      throw new Error('Invalid length');
    }
    const data = bytes.slice(offset, offset + len);
    offset += len;
    return data;
  };

  // Read A1
  const a1 = readLengthPrefixed();

  // Read A2
  const a2 = readLengthPrefixed();

  // Read A3
  const a3 = readLengthPrefixed();

  // Read Z
  const zBytes = readLengthPrefixed();
  let z = 0n;
  for (let i = 0; i < zBytes.length; i++) {
    z = (z << 8n) | BigInt(zBytes[i]);
  }

  // Read C
  const cBytes = readLengthPrefixed();
  let c = 0n;
  for (let i = 0; i < cBytes.length; i++) {
    c = (c << 8n) | BigInt(cBytes[i]);
  }

  // Read RHash (fixed 32 bytes)
  if (offset + 32 > bytes.length) {
    throw new Error('Data too short for RHash');
  }
  const rHash = bytes.slice(offset, offset + 32);

  return { a1, a2, a3, z, c, rHash };
}

/**
 * Verify ElGamal Transfer
 * This is the JavaScript equivalent of the Go function VerifyElGamalTransfer
 *
 * @param {string} sellerCipherHex - Seller ciphertext (hex string with 0x prefix)
 * @param {string} buyerCipherHex - Buyer ciphertext (hex string with 0x prefix)
 * @param {string} proofHex - DLEQ proof (hex string with 0x prefix)
 * @param {string} sellerPubKeyHex - Seller's public key (hex string)
 * @param {string} buyerPubKeyHex - Buyer's public key (hex string)
 * @returns {boolean} - True if verification passes, false otherwise
 */
export function verifyElGamalTransfer(
  sellerCipherHex,
  buyerCipherHex,
  proofHex,
  sellerPubKeyHex,
  buyerPubKeyHex
) {
  try {
    console.log('=== VERIFY ELGAMAL TRANSFER ===');
    console.log('Seller cipher hex:', sellerCipherHex);
    console.log('Buyer cipher hex:', buyerCipherHex);
    console.log('Proof hex:', proofHex);
    console.log('Seller pub key:', sellerPubKeyHex);
    console.log('Buyer pub key:', buyerPubKeyHex);

    // Parse ciphertexts
    const sellerCipher = unmarshalElGamalCiphertext(hexToBytes(sellerCipherHex));
    const buyerCipher = unmarshalElGamalCiphertext(hexToBytes(buyerCipherHex));

    // Parse proof
    const proof = unmarshalDLEQProof(hexToBytes(proofHex));

    // Parse public keys
    const sellerPubKeyBytes = hexToBytes(sellerPubKeyHex);
    const buyerPubKeyBytes = hexToBytes(buyerPubKeyHex);

    const [sellerX, sellerY] = parseECPoint(sellerPubKeyBytes);
    const [buyerX, buyerY] = parseECPoint(buyerPubKeyBytes);

    console.log('Parsed seller pub key:', sellerX.toString(16), sellerY.toString(16));
    console.log('Parsed buyer pub key:', buyerX.toString(16), buyerY.toString(16));

    // Parse proof commitments
    const [a1x, a1y] = parseECPoint(proof.a1);
    const [a2x, a2y] = parseECPoint(proof.a2);
    const [a3x, a3y] = parseECPoint(proof.a3);

    // Parse C1 values from ciphertexts (should be identical since same r)
    const [c1SellerX, c1SellerY] = parseECPoint(sellerCipher.c1);
    const [c1BuyerX, c1BuyerY] = parseECPoint(buyerCipher.c1);

    console.log('Seller C1:', c1SellerX.toString(16), c1SellerY.toString(16));
    console.log('Buyer C1:', c1BuyerX.toString(16), c1BuyerY.toString(16));

    // First, verify that C1 is the same for both (same r was used)
    if (c1SellerX !== c1BuyerX || c1SellerY !== c1BuyerY) {
      console.error('Different r values used!');
      return false;
    }

    // Parse shared secrets
    const [sSellerX, sSellerY] = parseECPoint(sellerCipher.sharedSecret);
    const [sBuyerX, sBuyerY] = parseECPoint(buyerCipher.sharedSecret);

    // Verify challenge (must include rHash to match proof generation)
    const challengeInput = new Uint8Array(
      sellerPubKeyBytes.length + buyerPubKeyBytes.length +
      proof.a1.length + proof.a2.length + proof.a3.length + 32
    );
    let offset = 0;
    challengeInput.set(sellerPubKeyBytes, offset); offset += sellerPubKeyBytes.length;
    challengeInput.set(buyerPubKeyBytes, offset); offset += buyerPubKeyBytes.length;
    challengeInput.set(proof.a1, offset); offset += proof.a1.length;
    challengeInput.set(proof.a2, offset); offset += proof.a2.length;
    challengeInput.set(proof.a3, offset); offset += proof.a3.length;
    challengeInput.set(proof.rHash, offset);

    const challengeBytes = keccak256(challengeInput);
    let cCheck = 0n;
    for (let i = 0; i < challengeBytes.length; i++) {
      cCheck = (cCheck << 8n) | BigInt(challengeBytes[i]);
    }
    cCheck = cCheck % SECP256K1_N;

    console.log('Challenge check:', cCheck.toString(16));
    console.log('Proof challenge:', proof.c.toString(16));

    if (cCheck !== proof.c) {
      console.error('Challenge mismatch - proof is invalid');
      return false;
    }

    // Verify equation 1: g^z = A1 * C1^c
    const [gzX, gzY] = scalarMultSecp256k1(proof.z, [SECP256K1_GX, SECP256K1_GY]);
    const [c1cX, c1cY] = scalarMultSecp256k1(proof.c, [c1SellerX, c1SellerY]);
    const [right1X, right1Y] = pointAddSecp256k1([a1x, a1y], [c1cX, c1cY]);

    if (gzX !== right1X || gzY !== right1Y) {
      console.error('Equation 1 failed');
      return false;
    }

    // Verify equation 2: sellerPub^z = A2 * S_seller^c
    const [sellerPubZX, sellerPubZY] = scalarMultSecp256k1(proof.z, [sellerX, sellerY]);
    const [sSellerCX, sSellerCY] = scalarMultSecp256k1(proof.c, [sSellerX, sSellerY]);
    const [right2X, right2Y] = pointAddSecp256k1([a2x, a2y], [sSellerCX, sSellerCY]);

    if (sellerPubZX !== right2X || sellerPubZY !== right2Y) {
      console.error('Equation 2 failed');
      return false;
    }

    // Verify equation 3: buyerPub^z = A3 * S_buyer^c
    const [buyerPubZX, buyerPubZY] = scalarMultSecp256k1(proof.z, [buyerX, buyerY]);
    const [sBuyerCX, sBuyerCY] = scalarMultSecp256k1(proof.c, [sBuyerX, sBuyerY]);
    const [right3X, right3Y] = pointAddSecp256k1([a3x, a3y], [sBuyerCX, sBuyerCY]);

    if (buyerPubZX !== right3X || buyerPubZY !== right3Y) {
      console.error('Equation 3 failed');
      return false;
    }

    console.log('All three equations verified!');
    console.log('=== END VERIFY ELGAMAL TRANSFER ===');

    // All three equations verified!
    return true;
  } catch (error) {
    console.error('Verification error:', error);
    return false;
  }
}
