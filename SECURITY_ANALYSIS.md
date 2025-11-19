# Security Analysis of Skavenge ZK Proof System

## Executive Summary

This analysis examines the zero-knowledge proof system used in the Skavenge NFT scavenger hunt, focusing on the transfer mechanism that is supposed to securely transfer encrypted clue content from seller to buyer.

**Critical Finding**: The current ZK proof implementation **does not actually prove** that the buyer will receive the same plaintext that the seller possesses. Multiple fundamental cryptographic flaws allow a malicious seller to provide invalid or different encrypted content to the buyer while still passing all verification checks.

## Transfer Flow Validation

### Expected Flow (from README and Smart Contract)

1. Seller sets a sale price for a clue NFT
2. Buyer initiates purchase by sending ETH (`InitiatePurchase`)
3. Seller generates and provides a ZK proof that they know the decrypted clue content (`ProvideProof`)
4. Buyer verifies the proof (`VerifyProof`)
5. Seller completes the transfer with the re-encrypted clue for the buyer (`CompleteTransfer`)
6. NFT ownership transfers to buyer, seller gets paid

### Actual Test Implementation

The test in `tests/transfer_test.go` correctly follows the smart contract flow:

✅ Lines 93-101: Sets sale price
✅ Lines 104-114: Buyer initiates purchase
✅ Lines 116-142: Seller provides ZK proof
✅ Lines 158-169: Buyer verifies proof (off-chain and on-chain)
✅ Lines 172-183: Seller completes transfer
✅ Lines 186-198: Verifies ownership transfer and buyer can decrypt

The smart contract calls match the expected flow in `eth/skavenge.sol`.

## Critical ZK Proof Vulnerabilities

### Issue 1: Incomplete Proof Verification - Missing R2 Check ⚠️ CRITICAL

**Location**: `zkproof/proof.go:179-226` (`VerifyProof` function)

**Problem**: The proof verification only checks one of two required equations:
- ✅ Verifies: `g^s * Y^c = R1` (line 220-225)
- ❌ **NEVER verifies R2!**

**What Should Be Verified**:
In a proper Schnorr-like proof for Diffie-Hellman knowledge, you need to verify both:
1. `g^s * seller_pub^c = R1` (proves knowledge of seller's private key)
2. `buyer_pub^s * (some_point)^c = R2` (proves the same secret was used with buyer's key)

**Impact**: Without the R2 verification, the proof is mathematically incomplete and doesn't bind the buyer's ciphertext to the seller's knowledge.

### Issue 2: R2 Doesn't Bind to Buyer Ciphertext ⚠️ CRITICAL

**Location**: `zkproof/proof.go:45-57` and `zkproof/message.go:15-65`

**Problem**: The proof's R2 commitment is computed as:
```go
// In GenerateProof (proof.go:53)
r2x, r2y := ps.Curve.ScalarMult(buyerPubKey.X, buyerPubKey.Y, k.Bytes())
```

But the buyer's ciphertext is created via ECIES with a **completely independent** random ephemeral key:
```go
// In EncryptMessage (message.go:18)
ephemeral, err := ecdsa.GenerateKey(ps.Curve, rand.Reader)
```

**This means**:
- The proof's random value `k` has NO cryptographic relationship to the buyer's ciphertext
- The seller can generate a valid proof for one message
- But encrypt a **completely different message** for the buyer
- The proof still verifies because they are cryptographically independent!

**Example Attack**:
```go
// Seller generates proof claiming to send "Hidden treasure location"
proof := GenerateProof([]byte("Hidden treasure location"), sellerKey, buyerKey, sellerCipher)

// But actually encrypts "Fake location" for buyer
buyerCipher := EncryptMessage([]byte("Fake location"), buyerPubKey)

// Complete transfer with the fake ciphertext
CompleteTransfer(transferId, buyerCipher) // ✅ SUCCEEDS!
```

### Issue 3: No Proof of Plaintext Equality ⚠️ CRITICAL

**Location**: `zkproof/proof.go` (entire proof system)

**Problem**: The proof system is a Schnorr-like proof of knowledge of the seller's private key. It does NOT prove:
1. That the seller can decrypt `sellerCipherText`
2. That `buyerCipherText` contains the same plaintext as `sellerCipherText`
3. That the seller even knows what plaintext is in either ciphertext

**What the proof actually proves**: "I know the discrete logarithm of the seller's public key" (i.e., the seller's private key)

**What it SHOULD prove**: "The ciphertext encrypted for the buyer contains the same plaintext that I can decrypt from the seller's ciphertext"

### Issue 4: Non-Deterministic Encryption Breaks Binding

**Location**: `zkproof/message.go:15-65`

**Problem**: The ECIES encryption scheme uses:
- Random ephemeral keys (line 18)
- Random nonces (lines 47-50)

This means:
- Encrypting the same message twice produces completely different ciphertexts
- You cannot compare ciphertexts to verify they contain the same plaintext
- The proof cannot bind to the actual ciphertext bits

**Consequence**: Even with a perfect ZK proof, you cannot prove that two ECIES ciphertexts contain the same message without revealing the message or using specialized cryptographic schemes (e.g., proxy re-encryption, homomorphic encryption, or ZK-SNARKs).

### Issue 5: Smart Contract Doesn't Verify Proof On-Chain ⚠️ CRITICAL

**Location**: `eth/skavenge.sol:346-386`

**Problem**:
- `ProvideProof` (line 346) just **stores** the proof bytes, doesn't verify them
- `VerifyProof` (line 372) just **sets a boolean flag** when the buyer calls it
- No cryptographic verification happens on-chain

**Code Analysis**:
```solidity
// skavenge.sol:361-363
transfer.proof = proof;
transfer.newClueHash = newClueHash;
transfer.proofProvidedAt = block.timestamp;
// No verification! Just storage!

// skavenge.sol:383-384
transfer.proofVerified = true;
transfer.verifiedAt = block.timestamp;
// Just sets a flag! No crypto verification!
```

**Impact**: A malicious seller can:
1. Provide a completely invalid proof
2. Buyer doesn't verify off-chain (or is colluding)
3. Buyer calls `VerifyProof` anyway (just sets the flag to true)
4. Transfer completes with garbage data
5. Buyer receives worthless encrypted junk

### Issue 6: Hash Check Provides No Security

**Location**: `eth/skavenge.sol:409-413`

**Problem**: The contract verifies:
```solidity
require(
    keccak256(newEncryptedContents) == transfer.newClueHash,
    "Content hash mismatch"
);
```

This only proves that the seller is providing the same ciphertext they committed to in `ProvideProof`. It does NOT prove:
- The ciphertext is valid
- The ciphertext contains the correct plaintext
- The buyer can decrypt it
- It matches the seller's ciphertext content

### Issue 7: Off-Chain Verification is Optional and Not Enforced

**Location**: `tests/transfer_test.go:150-155`

**Problem**: The test does verify the proof off-chain:
```go
valid := ps.VerifyProof(verifyProof, clue.EncryptedContents)
require.True(t, valid, "Proof verification failed")
```

But this is in the TEST, not enforced by the protocol. A real buyer could:
- Skip the off-chain verification
- Call the on-chain `VerifyProof` anyway
- Complete the transfer with an invalid proof

## Fundamental Cryptographic Design Flaws

The core issue is that this system attempts to prove plaintext equality across two ciphertexts without using proper cryptographic primitives designed for this purpose.

### What's Needed for Secure Plaintext Equality Proofs:

1. **zk-SNARKs/STARKs**: Prove "I know plaintext P such that Decrypt(sellerCipher, sellerKey) = P AND Encrypt(P, buyerKey) = buyerCipher"

2. **Proxy Re-Encryption**: Use schemes like BBS98 or AFGH where a seller can transform their ciphertext into one encrypted for the buyer without decrypting

3. **Commitment-Based Schemes**: Commit to the plaintext, prove both ciphertexts decrypt to the committed value

4. **Homomorphic Properties**: Use encryption schemes where you can prove properties about encrypted data without decryption

### Current Approach is Fundamentally Broken:

The current Schnorr-like proof proves knowledge of a discrete logarithm (private key), but has no way to bind that to the plaintext content of independently-encrypted ECIES ciphertexts.

## Test Files Demonstrating Vulnerabilities

See `tests/security_test.go` for executable proof-of-concept exploits demonstrating each vulnerability.

## Recommendations

### Immediate Actions (Proof of Concept):

1. **Do not deploy this system to production** - The ZK proof provides no security
2. **Document clearly** that this is a proof-of-concept with known vulnerabilities
3. **Add warnings** to the README about the security limitations

### For Production:

1. **Use zk-SNARKs** - Implement a proper zero-knowledge proof using libraries like:
   - gnark (Go-based zk-SNARK library)
   - circom + snarkjs
   - Implement a circuit proving: Decrypt(sellerCipher, sellerKey) = Encrypt^(-1)(buyerCipher, buyerKey)

2. **Use Proxy Re-Encryption** - Replace the custom scheme with a proven proxy re-encryption library

3. **On-Chain Verification** - Move proof verification to the smart contract (though this is expensive for complex proofs)

4. **Use Trusted Oracle** - Have a trusted third party verify the proof off-chain and attest on-chain

5. **Simplified Model** - Consider if the complexity is necessary - perhaps reveal the clue after purchase with a deposit/penalty system for fraud

## Conclusion

The current implementation provides **NO cryptographic security** for the transfer mechanism. A malicious seller can:
- Send different encrypted content than claimed
- Send garbage data
- Complete transfers with invalid proofs

The buyer has no cryptographic guarantee that they will receive the correct plaintext, making the entire ZK proof system ineffective for its stated purpose.
