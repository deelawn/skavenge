# Skavenge Security Review

**Review Date:** 2025-11-20
**Reviewers:** Security Audit Team
**Scope:** Smart Contract (`eth/skavenge.sol`) and ZK Proof System (`zkproof/*.go`)

## Executive Summary

This security review examines the Skavenge NFT scavenger hunt system, which uses ElGamal encryption and DLEQ (Discrete Log Equality) proofs to enable secure, verifiable transfers of encrypted clues. The system demonstrates strong cryptographic foundations with proper implementation of zero-knowledge proofs. However, we identified **1 CRITICAL vulnerability** and several **MEDIUM/LOW severity issues** that should be addressed.

### Overall Security Assessment
- **Cryptographic Design:** ✅ Strong
- **ZK Proof Implementation:** ✅ Correct
- **Smart Contract Logic:** ⚠️ Good with caveats
- **Transfer Process Security:** ❌ Critical issue identified

---

## Architecture Overview

The system implements a three-phase transfer protocol:

1. **Proof Phase:** Seller generates DLEQ proof demonstrating that seller and buyer ciphertexts encrypt the same plaintext (without revealing the plaintext or decryption key)
2. **Verification Phase:** Buyer verifies the proof off-chain, then commits on-chain by calling `verifyProof()`
3. **Completion Phase:** Seller reveals the decryption key `r` by calling `completeTransfer()`, which atomically transfers ownership and payment

### Encryption Scheme

The system uses a modified ElGamal encryption where the decryption key is derived as:
```
key = Keccak256(r || sharedSecret)
```
where:
- `r` is the random ephemeral value used in ElGamal
- `sharedSecret = buyerPrivKey * C1 = buyerPrivKey * g^r`

This design ensures the buyer needs BOTH:
- Their private key (to compute `sharedSecret`)
- The revealed `r` value (from seller)

---

## High Severity Issues

### 🟠 HIGH: No On-Chain Verification of C1 = g^r

**Severity:** HIGH
**Component:** Smart Contract - Transfer Completion
**Location:** `eth/skavenge.sol:445-449`

#### Description

The smart contract does NOT verify on-chain that the provided `r` value correctly generates the stored C1 value (i.e., `g^r == C1`). The code explicitly documents this design decision (lines 445-449):

```solidity
// Note: We don't verify g^r == C1 on-chain because:
// 1. The ecmul precompile (0x07) only supports alt_bn128, not secp256k1
// 2. The hash commitment already prevents r value tampering
// 3. The buyer verified the DLEQ proof off-chain before calling verifyProof()
// 4. Implementing secp256k1 multiplication in Solidity is prohibitively expensive
```

#### Security Implications

This creates a trust assumption: **The buyer MUST verify the DLEQ proof correctly off-chain before calling `verifyProof()`**. If a buyer:
- Skips verification
- Uses buggy verification code
- Is tricked into accepting an invalid proof

They could call `verifyProof()` for an invalid proof, and the seller could provide a garbage `r` value that:
- Passes the hash check (because the hash was embedded in the invalid proof)
- Cannot decrypt to the correct plaintext

#### Mitigation

The current design is acceptable given the constraints, BUT requires:

1. **Client-Side Safeguards:** Ensure all buyer clients ALWAYS verify proofs correctly
2. **UI Warnings:** Display clear warnings that verification is REQUIRED
3. **Reference Implementation:** Provide audited, tested verification code
4. **Test Coverage:** Comprehensive tests demonstrating that skipping verification leads to loss

#### Recommended Enhancement

Add on-chain verification using `ecrecover` and precompile tricks, or accept higher gas costs. Example approaches:
- Use `alt_bn128` curve instead of secp256k1 (breaking change)
- Implement Solidity-based secp256k1 verification (expensive but possible)
- Use optimistic verification with challenge periods

---

## Medium Severity Issues

### 🟡 MEDIUM: Unbounded Array Iteration in _removeFromForSaleList

**Severity:** MEDIUM
**Component:** Smart Contract - Sale Management
**Location:** `eth/skavenge.sol:544-563`

#### Description

The `_removeFromForSaleList()` function iterates through the entire `_cluesForSaleList` array to find and remove an item (lines 552-562):

```solidity
for (uint256 i = 0; i < _cluesForSaleList.length; i++) {
    if (_cluesForSaleList[i] == tokenId) {
        _cluesForSaleList[i] = _cluesForSaleList[_cluesForSaleList.length - 1];
        _cluesForSaleList.pop();
        break;
    }
}
```

#### Impact

If the array grows large (e.g., 1000+ items), this operation becomes expensive. An attacker could:
1. Mint many tokens
2. List them all for sale
3. Remove them from sale individually
4. Cause high gas costs for legitimate users

#### Recommended Fix

Use a mapping-based index to avoid iteration:

```solidity
mapping(uint256 => uint256) private _forSaleIndex; // tokenId => index in array

function _addToForSaleList(uint256 tokenId) private {
    _cluesForSaleList.push(tokenId);
    _forSaleIndex[tokenId] = _cluesForSaleList.length - 1;
}

function _removeFromForSaleList(uint256 tokenId) private {
    if (!cluesForSale[tokenId]) return;

    uint256 index = _forSaleIndex[tokenId];
    uint256 lastIndex = _cluesForSaleList.length - 1;

    if (index != lastIndex) {
        uint256 lastTokenId = _cluesForSaleList[lastIndex];
        _cluesForSaleList[index] = lastTokenId;
        _forSaleIndex[lastTokenId] = index;
    }

    _cluesForSaleList.pop();
    delete _forSaleIndex[tokenId];
    cluesForSale[tokenId] = false;
}
```

---

### 🟡 MEDIUM: Insufficient Randomness Protection

**Severity:** MEDIUM
**Component:** ZK Proof System
**Location:** `zkproof/elgamal.go:163`, `zkproof/proof.go:39`

#### Description

The system uses `crypto/rand.Int()` for generating random values `r`, `w`, and `k`. While `crypto/rand` is cryptographically secure, there's no explicit check that the generated values are non-zero and within the valid range.

#### Current Code

```go
r, err := rand.Int(rand.Reader, curveN)
if err != nil {
    return nil, fmt.Errorf("failed to generate r: %v", err)
}
// No check that r != 0
```

#### Recommended Fix

```go
r, err := rand.Int(rand.Reader, new(big.Int).Sub(curveN, big.NewInt(1)))
if err != nil {
    return nil, fmt.Errorf("failed to generate r: %v", err)
}
r.Add(r, big.NewInt(1)) // Ensure r is in [1, curveN-1]

// Additional sanity check
if r.Sign() <= 0 || r.Cmp(curveN) >= 0 {
    return nil, fmt.Errorf("invalid r value generated")
}
```

---

## Low Severity Issues

### 🟢 LOW: Timestamp Manipulation

**Severity:** LOW
**Location:** `eth/skavenge.sol` - Various timeout checks

#### Description

The contract uses `block.timestamp` for timeout enforcement. Miners can manipulate timestamps by approximately ±15 seconds.

#### Impact

With a 180-second timeout, ±15 seconds manipulation (~8% variance) is acceptable but should be documented.

#### Recommendation

Document the acceptable timestamp variance in comments.

---

### 🟢 LOW: No Event for r Value Reveal

**Severity:** LOW
**Location:** `eth/skavenge.sol:478`

#### Description

The `TransferCompleted` event includes the r value (line 478), but there's no dedicated event for r value reveal. This makes it harder to track when decryption becomes possible.

#### Recommendation

```solidity
event DecryptionKeyRevealed(uint256 indexed tokenId, uint256 rValue, address indexed buyer);

// In completeTransfer, before transferring ownership:
emit DecryptionKeyRevealed(transfer.tokenId, rValue, transfer.buyer);
```

---

### 🟢 LOW: Missing Input Validation

**Severity:** LOW
**Location:** Multiple functions

#### Description

Several functions lack input validation:

1. `mintClue()` - No check for empty `encryptedContents`
2. `setSalePrice()` - No maximum price limit
3. `provideProof()` - No minimum proof length check (only `>= 32` on line 379)

#### Recommendations

```solidity
function mintClue(...) external returns (uint256) {
    require(encryptedContents.length > 0, "Empty encrypted contents");
    require(encryptedContents.length <= MAX_CIPHERTEXT_SIZE, "Ciphertext too large");
    // ...
}
```

---

## ZK Proof Implementation Analysis

### ✅ DLEQ Proof Correctness

The DLEQ (Discrete Log Equality) proof implementation is **mathematically correct**. It properly implements the Chaum-Pedersen protocol:

**Proof Generation (zkproof/elgamal.go:201-262):**
```
1. Generate random w
2. Compute commitments: A1 = g^w, A2 = sellerPub^w, A3 = buyerPub^w
3. Compute challenge: c = Hash(sellerPub || buyerPub || A1 || A2 || A3 || rHash)
4. Compute response: z = w + c*r mod n
```

**Proof Verification (zkproof/elgamal.go:264-372):**
```
1. Verify C1_seller == C1_buyer (same r used)
2. Verify g^z == A1 * C1^c (proves C1 = g^r)
3. Verify sellerPub^z == A2 * S_seller^c (proves S_seller = sellerPub^r)
4. Verify buyerPub^z == A3 * S_buyer^c (proves S_buyer = buyerPub^r)
```

This correctly proves that both ciphertexts use the same `r` value, which means they encrypt the same plaintext (given the same plaintext was used to derive the encryption keys).

### ✅ rHash Binding

The inclusion of `rHash` in the challenge computation (line 244) is **critical for security**:

```go
h.Write(proof.RHash[:]) // Include r hash in challenge
```

This binds the proof to a specific `r` value, preventing the seller from:
- Generating a proof with r₁
- Providing r₂ ≠ r₁ during completion

The contract verifies `Hash(provided_r) == stored_rHash`, ensuring the seller cannot change r after proof generation.

### ✅ Encryption Key Derivation

The key derivation in `EncryptElGamal()` (lines 83-86) is **well-designed**:

```go
keyHash := sha3.NewLegacyKeccak256()
keyHash.Write(r.Bytes())      // Prevents early decryption (needs r)
keyHash.Write(sharedSecret)   // Prevents public decryption (needs private key)
key := keyHash.Sum(nil)
```

This ensures the buyer needs BOTH:
- `r` (revealed by seller after payment)
- Their private key (to compute `sharedSecret`)

This prevents:
- **Early decryption:** Buyer cannot decrypt before r is revealed
- **Public decryption:** Third parties cannot decrypt even after r is revealed

---

## Positive Security Features

### ✅ Reentrancy Protection

The contract properly uses OpenZeppelin's `ReentrancyGuard` on all state-changing external functions with token/value transfers.

### ✅ Checks-Effects-Interactions Pattern

Payment transfers follow the correct pattern:
```solidity
// 1. Checks
require(transfer.verifiedAt > 0, "Proof not verified");

// 2. Effects
clues[transfer.tokenId].encryptedContents = newEncryptedContents;
_safeTransfer(msg.sender, transfer.buyer, transfer.tokenId, "");

// 3. Interactions
(bool sent, ) = payable(msg.sender).call{value: transfer.value}("");
require(sent, "Failed to send Ether");
```

### ✅ Integer Overflow Protection

Uses Solidity 0.8.19 with built-in overflow/underflow protection.

### ✅ Access Control

Proper ownership checks throughout:
- `mintClue()` - Only authorized minter
- `getRValue()` - Only token owner
- `completeTransfer()` - Only token owner

### ✅ Solved Clue Protection

Prevents transfer of solved clues (lines 316-318, 452-454), maintaining game integrity.

---

## Smart Contract Best Practices

### ✅ Well Implemented

- Uses OpenZeppelin battle-tested contracts (ERC721, ReentrancyGuard)
- Follows Solidity style guide
- Clear error messages with custom errors
- Comprehensive event emission
- Gas-efficient storage patterns

### ⚠️ Could Be Improved

- Missing NatSpec documentation on some functions
- No circuit breaker / pause mechanism
- No upgradability (acceptable for immutability, but document)

---

## Test Coverage Analysis

### Existing Tests

The test suite demonstrates good coverage:

**Transfer Tests (`tests/transfer_test.go`):**
- ✅ Successful transfer flow
- ✅ Invalid proof rejection
- ✅ Completing without verification (correctly fails)
- ✅ Cancel transfer with refund
- ✅ Corrupted r value rejection

**Security Tests (`tests/security_test.go`):**
- ✅ Different plaintext attack prevention
- ✅ Wrong r value attack prevention
- ✅ Early decryption prevention
- ✅ Public key requirement
- ✅ Fake rHash attack prevention

### Missing Tests

**Critical:**
- ❌ Mempool frontrunning attack simulation
- ❌ Buyer cancellation after verification
- ❌ Race condition with multiple buyers

**Medium:**
- ❌ Large _cluesForSaleList DoS scenario
- ❌ Timeout edge cases (expired proofs)
- ❌ Gas limit attacks

**Recommended Additional Tests:**

```go
// Test buyer cannot cancel after verification
func TestCannotCancelAfterVerification(t *testing.T) {
    // Setup: initiate, provide proof, buyer verifies
    // Attempt: buyer tries to cancel
    // Assert: transaction reverts
}

// Test frontrunning attack is prevented
func TestFrontrunningPrevention(t *testing.T) {
    // Setup: seller submits completeTransfer
    // Simulate: buyer extracts r from pending tx
    // Attempt: buyer tries to cancel with higher gas
    // Assert: either cancel fails or complete succeeds first
}
```

---

## Recommendations Summary

### Immediate Actions (Critical)

1. **FIX CRITICAL:** Prevent buyer cancellation after `verifyProof()` to eliminate mempool frontrunning attack
2. **Add tests** for the frontrunning scenario
3. **Deploy fix** before mainnet launch

### High Priority (Pre-Mainnet)

4. **Document trust assumptions** clearly: Buyers MUST verify proofs correctly
5. **Provide audited client libraries** with proper verification
6. **Add UI safeguards** to prevent skipping verification

### Medium Priority (Performance/UX)

7. **Optimize** `_removeFromForSaleList` with index mapping
8. **Consider** first-come-first-served for purchase initiation
9. **Improve** randomness validation in zkproof

### Low Priority (Nice-to-Have)

10. **Add** dedicated DecryptionKeyRevealed event
11. **Document** timestamp manipulation tolerance
12. **Enhance** input validation

---

## Cryptographic Security Assessment

### Overall: STRONG ✅

The cryptographic primitives and protocols are well-chosen and correctly implemented:

- **Curve:** secp256k1 (Bitcoin/Ethereum standard, well-studied)
- **Hash:** Keccak256 (SHA3, Ethereum standard)
- **Proof System:** Chaum-Pedersen DLEQ with Fiat-Shamir (standard, secure)
- **Encryption:** Modified ElGamal with proper key derivation
- **Commitment:** Hash-based commitment scheme (binding and hiding)

No cryptographic vulnerabilities identified in the zero-knowledge proof or encryption implementations.

---

## Conclusion

The Skavenge system demonstrates **strong cryptographic design** with a well-implemented DLEQ proof system that correctly ensures both seller and buyer ciphertexts contain the same plaintext. The smart contract follows security best practices with proper reentrancy protection and access control.

However, the **critical mempool frontrunning vulnerability** must be addressed before mainnet deployment. The fix is straightforward: prevent buyer cancellation after verification commitment. With this fix and the recommended enhancements, the system will provide robust security guarantees for verifiable encrypted clue transfers.

### Risk Rating

**Current State:** HIGH RISK (due to critical frontrunning issue)
**After Fixes:** LOW RISK (with documented trust assumptions)

---

## Appendix: Transfer Protocol Flow

```
┌─────────┐                                   ┌─────────┐
│  Seller │                                   │  Buyer  │
└────┬────┘                                   └────┬────┘
     │                                             │
     │ 1. Generate DLEQ proof                      │
     │    Encrypt for both parties                 │
     │    Commit to r via Hash(r)                  │
     │                                             │
     │ 2. Call initiatePurchase()                  │
     │    <─────────────────────────────────────── │
     │                                             │
     │ 3. Call provideProof()                      │
     │    ──────────────────────────────────────> │
     │                                             │
     │                                4. Verify    │
     │                                   DLEQ      │
     │                                   proof     │
     │                                   off-chain │
     │                                             │
     │ 5. Call verifyProof() [COMMITMENT POINT]    │
     │    <─────────────────────────────────────── │
     │                                             │
     │ 6. Call completeTransfer(r)                 │
     │    - Verify Hash(r) == stored hash          │
     │    - Transfer ownership                     │
     │    - Send payment                           │
     │    ──────────────────────────────────────> │
     │                                             │
     │                                   7. Buyer  │
     │                                      can    │
     │                                      decrypt│
     │                                      with r │
     ▼                                             ▼
```

**Vulnerability:** Between steps 6 and 7, buyer can extract r from mempool and cancel before transaction mines (if cancellation after step 5 is not prevented).

---

**Report Version:** 1.0
**Classification:** CONFIDENTIAL
**Distribution:** Internal Security Review
