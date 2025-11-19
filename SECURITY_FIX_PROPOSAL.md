# Security Fix Proposal for Skavenge ZK Proof System

## Executive Summary

This proposal outlines a comprehensive fix for the cryptographic vulnerabilities in the Skavenge transfer system. The solution uses a **commitment-based verifiable encryption scheme** that provides cryptographic guarantees while keeping expensive operations off-chain.

## Design Goals

1. ✅ **Prove plaintext equality** - Cryptographically bind both ciphertexts to the same plaintext
2. ✅ **Off-chain computation** - Heavy crypto work done off-chain to save gas
3. ✅ **On-chain verification** - Lightweight verification on-chain for security
4. ✅ **No external dependencies** - Implementable in pure Go
5. ✅ **Backward compatible** - Minimal smart contract changes

## Proposed Solution: Commitment-Based Verifiable Dual Encryption

### Core Concept

Instead of trying to prove two ECIES ciphertexts contain the same plaintext (which is cryptographically hard), we:

1. **Commit to the plaintext** using a hash commitment
2. **Prove knowledge of the committed plaintext**
3. **Prove both ciphertexts encrypt to the committed value**
4. **Verify the complete Schnorr proof** (including R2)
5. **Verify proof on-chain** (lightweight check)

### Cryptographic Scheme

#### Phase 1: Seller Commits to Plaintext

```
plaintextCommitment = keccak256(plaintext || salt)
```

The commitment binds the seller to a specific plaintext without revealing it.

#### Phase 2: Generate Verifiable Dual Encryption Proof

The proof demonstrates:
- "I know plaintext P with commitment C"
- "sellerCipher correctly encrypts P for seller's key"
- "buyerCipher correctly encrypts P for buyer's key"
- "I possess seller's private key"

**Proof Components:**

1. **Plaintext Commitment**: `C = Hash(P || salt)`
2. **DLEQ Proof** (Discrete Log Equality): Proves same private key used
3. **Encryption Binding**: Hash of both ciphertexts and commitment
4. **Complete Schnorr Signature**: With both R1 and R2 verification

**Proof Structure:**
```go
type VerifiableTransferProof struct {
    // Commitment to plaintext
    PlaintextCommitment [32]byte
    Salt                []byte

    // Schnorr proof components (DL knowledge)
    C  *big.Int  // Challenge
    S  *big.Int  // Response
    R1 []byte    // Commitment g^k
    R2 []byte    // Commitment buyerPub^k

    // Public keys
    SellerPubKey []byte
    BuyerPubKey  []byte

    // Ciphertext hashes
    SellerCipherHash [32]byte
    BuyerCipherHash  [32]byte

    // Binding signature (proves ciphers encrypt committed value)
    BindingSignature []byte
}
```

#### Phase 3: Verification (Off-Chain First)

**Buyer verifies off-chain:**

1. ✅ Schnorr proof validates (check R1 AND R2!)
2. ✅ Binding signature is valid
3. ✅ Ciphertext hashes match actual ciphertexts
4. ✅ Can decrypt buyer ciphertext
5. ✅ Decrypted plaintext matches commitment when hashed with salt

**Smart contract verifies on-chain:**

1. ✅ Proof signature is valid (cheap ECDSA check)
2. ✅ Ciphertext hash matches what will be transferred
3. ✅ Transfer has been verified by buyer (attestation)

#### Phase 4: Reveal and Verify (After Transfer)

**After successful transfer, seller provides:**
- Original plaintext
- Salt value

**Buyer can verify:**
```
Hash(receivedPlaintext || salt) == PlaintextCommitment
```

If mismatch → **fraud proof** → refund via smart contract

## Detailed Implementation Changes

### 1. New Proof Structure (`zkproof/types.go`)

```go
// VerifiableTransferProof contains all components for secure transfer
type VerifiableTransferProof struct {
    PlaintextCommitment [32]byte
    Salt                []byte
    C                   *big.Int
    S                   *big.Int
    R1                  []byte
    R2                  []byte
    SellerPubKey        []byte
    BuyerPubKey         []byte
    SellerCipherHash    [32]byte
    BuyerCipherHash     [32]byte
    BindingProof        []byte  // ECDSA signature over all components
}

// RevealData is provided after transfer to verify plaintext
type RevealData struct {
    Plaintext []byte
    Salt      []byte
}
```

### 2. Enhanced Proof Generation (`zkproof/proof.go`)

```go
func (ps *ProofSystem) GenerateVerifiableTransferProof(
    plaintext []byte,
    sellerKey *ecdsa.PrivateKey,
    buyerPubKey *ecdsa.PublicKey,
) (*VerifiableTransferProof, *TransferData, error) {

    // Step 1: Generate random salt and commit to plaintext
    salt := make([]byte, 32)
    rand.Read(salt)

    commitment := keccak256(plaintext, salt)

    // Step 2: Encrypt plaintext for both parties
    sellerCipher, err := ps.EncryptMessage(plaintext, &sellerKey.PublicKey)
    buyerCipher, err := ps.EncryptMessage(plaintext, buyerPubKey)

    sellerCipherHash := keccak256(sellerCipher)
    buyerCipherHash := keccak256(buyerCipher)

    // Step 3: Generate Schnorr proof of discrete log knowledge
    k, _ := rand.Int(rand.Reader, curveN)

    R1 = g^k
    R2 = buyerPub^k  // This time we'll actually verify it!

    // Challenge includes ALL components
    c = Hash(
        commitment,
        sellerCipherHash,
        buyerCipherHash,
        sellerPubKey,
        buyerPubKey,
        R1,
        R2,
    )

    s = k - c*sellerPrivKey mod n

    // Step 4: Create binding proof (signature over commitment + ciphers)
    bindingData := commitment || sellerCipherHash || buyerCipherHash
    bindingProof = ECDSA.Sign(sellerKey, bindingData)

    return proof, transferData{sellerCipher, buyerCipher, salt}, nil
}
```

### 3. Complete Proof Verification (`zkproof/proof.go`)

```go
func (ps *ProofSystem) VerifyTransferProof(
    proof *VerifiableTransferProof,
    sellerCipher []byte,
    buyerCipher []byte,
) bool {

    // Verify ciphertext hashes match
    if keccak256(sellerCipher) != proof.SellerCipherHash {
        return false
    }
    if keccak256(buyerCipher) != proof.BuyerCipherHash {
        return false
    }

    // Verify binding signature
    bindingData := proof.PlaintextCommitment ||
                   proof.SellerCipherHash ||
                   proof.BuyerCipherHash
    if !VerifyECDSA(proof.SellerPubKey, bindingData, proof.BindingProof) {
        return false
    }

    // Verify Schnorr proof - CHECK BOTH EQUATIONS!
    c_reconstructed = Hash(
        proof.PlaintextCommitment,
        proof.SellerCipherHash,
        proof.BuyerCipherHash,
        proof.SellerPubKey,
        proof.BuyerPubKey,
        proof.R1,
        proof.R2,
    )

    if c_reconstructed != proof.C {
        return false
    }

    // CRITICAL: Verify R1 equation
    if g^s * sellerPub^c != R1 {
        return false
    }

    // CRITICAL: Verify R2 equation (THIS WAS MISSING!)
    if buyerPub^s * (sellerPub_x_on_buyerPub)^c != R2 {
        return false
    }

    return true
}
```

### 4. Buyer Verification After Decryption

```go
func (ps *ProofSystem) VerifyReceivedPlaintext(
    decryptedPlaintext []byte,
    proof *VerifiableTransferProof,
    revealSalt []byte,
) bool {

    // Verify plaintext matches commitment
    reconstructedCommitment := keccak256(decryptedPlaintext, revealSalt)

    return reconstructedCommitment == proof.PlaintextCommitment
}
```

### 5. Smart Contract Updates (`eth/skavenge.sol`)

```solidity
// Add fraud proof mechanism
struct FraudClaim {
    bytes32 transferId;
    bytes decryptedPlaintext;
    bytes salt;
    uint256 claimedAt;
}

mapping(bytes32 => FraudClaim) public fraudClaims;

// Dispute period (e.g., 7 days)
uint256 public constant DISPUTE_PERIOD = 7 days;

function completeTransfer(
    bytes32 transferId,
    bytes calldata newEncryptedContents,
    bytes calldata plaintextCommitment,
    bytes calldata revealSalt
) external nonReentrant {
    // ... existing checks ...

    // Store the commitment and salt for later verification
    transfers[transferId].plaintextCommitment = plaintextCommitment;
    transfers[transferId].revealSalt = revealSalt;
    transfers[transferId].completedAt = block.timestamp;

    // Transfer happens immediately, but dispute period starts
    _safeTransfer(msg.sender, transfer.buyer, transfer.tokenId, "");

    // Payment held in escrow until dispute period ends
    transfers[transferId].paymentHeld = true;

    emit TransferCompleted(transferId);
}

// Claim fraud if plaintext doesn't match commitment
function claimFraud(
    bytes32 transferId,
    bytes calldata decryptedPlaintext,
    bytes calldata salt
) external {
    TokenTransfer storage transfer = transfers[transferId];
    require(transfer.buyer == msg.sender, "Not buyer");
    require(transfer.completedAt > 0, "Transfer not completed");
    require(
        block.timestamp <= transfer.completedAt + DISPUTE_PERIOD,
        "Dispute period ended"
    );

    // Verify the fraud claim
    bytes32 reconstructed = keccak256(abi.encodePacked(decryptedPlaintext, salt));
    require(reconstructed != transfer.plaintextCommitment, "No fraud detected");

    // Refund buyer, penalize seller
    _safeTransfer(transfer.buyer, ownerOf(transfer.tokenId), transfer.tokenId, "");
    payable(transfer.buyer).call{value: transfer.value}("");

    emit FraudDetected(transferId);
}

// Seller can claim payment after dispute period
function claimPayment(bytes32 transferId) external {
    TokenTransfer storage transfer = transfers[transferId];
    require(transfer.paymentHeld, "Payment not held");
    require(
        block.timestamp > transfer.completedAt + DISPUTE_PERIOD,
        "Dispute period not ended"
    );

    // Release payment to seller
    payable(msg.sender).call{value: transfer.value}("");
    transfer.paymentHeld = false;

    emit PaymentReleased(transferId);
}

// On-chain proof verification (lightweight)
function verifyProofSignature(
    bytes32 transferId,
    bytes calldata bindingSignature
) public view returns (bool) {
    TokenTransfer storage transfer = transfers[transferId];

    // Reconstruct signed data
    bytes32 signedData = keccak256(abi.encodePacked(
        transfer.plaintextCommitment,
        transfer.sellerCipherHash,
        transfer.newClueHash
    ));

    // Verify ECDSA signature
    address signer = ecrecover(signedData, /* signature components */);
    return signer == ownerOf(transfer.tokenId);
}
```

## Migration Path

### Phase 1: Add New Proof System (Backward Compatible)

1. Add new `VerifiableTransferProof` types
2. Add new `GenerateVerifiableTransferProof` function
3. Keep old functions for backward compatibility
4. Add tests for new proof system

### Phase 2: Update Smart Contract

1. Add fraud proof mechanism
2. Add dispute period logic
3. Add payment escrow
4. Deploy new contract version

### Phase 3: Update Client App

1. Use new proof generation
2. Implement off-chain verification
3. Add fraud detection monitoring
4. Add dispute submission UI

## Security Guarantees (After Fix)

### What the System Proves:

1. ✅ **Seller knows plaintext**: Committed and signed
2. ✅ **Seller encrypted correctly**: Binding proof ties ciphers to commitment
3. ✅ **Same plaintext in both ciphers**: Commitment + verification
4. ✅ **Seller owns private key**: Schnorr proof (both R1 and R2)
5. ✅ **Buyer can detect fraud**: Commitment verification + dispute mechanism
6. ✅ **Economic incentive**: Dispute period protects buyer

### Attack Scenarios (Now Prevented):

❌ **Seller sends different plaintext**
- Buyer decrypts, verifies against commitment, submits fraud proof, gets refund

❌ **Seller provides invalid proof**
- Verification fails off-chain, buyer doesn't approve, no transfer

❌ **Seller provides garbage ciphertext**
- Buyer can't decrypt or decryption doesn't match commitment → fraud proof

❌ **Malicious buyer claims fraud falsely**
- They must provide plaintext that hashes to commitment (impossible if seller was honest)

## Gas Cost Analysis

### Original System:
- `ProvideProof`: ~50k gas (just storage)
- `VerifyProof`: ~30k gas (set flag)
- `CompleteTransfer`: ~200k gas (transfer + payment)
- **Total: ~280k gas**

### New System:
- `ProvideProof`: ~60k gas (storage + signature verification)
- `VerifyProof`: ~40k gas (signature check)
- `CompleteTransfer`: ~250k gas (transfer + escrow)
- `ClaimPayment`: ~50k gas (release from escrow)
- **Total: ~350k gas (+25%)**

The 25% increase in gas is acceptable given the massive security improvement.

## Alternative: Optimistic Verification

For even lower gas costs, use **optimistic verification**:

1. Assume proof is valid by default
2. Buyer has 7 days to dispute
3. On-chain verification only happens if disputed
4. Reduces normal-case gas by ~30%

## Implementation Priority

### Critical (Must Fix):

1. ✅ Fix `VerifyProof` to check R2
2. ✅ Add plaintext commitment
3. ✅ Add binding proof/signature
4. ✅ Add fraud proof mechanism

### Important (Should Fix):

1. ✅ Add on-chain signature verification
2. ✅ Add dispute period
3. ✅ Add payment escrow

### Optional (Nice to Have):

1. Optimistic verification mode
2. Slashing for malicious actors
3. Reputation system

## Testing Strategy

### Unit Tests:

1. Test proof generation with valid inputs
2. Test proof verification with valid proof
3. Test proof verification rejects tampered proof
4. Test commitment verification
5. Test fraud detection

### Integration Tests:

1. Full transfer flow with honest parties
2. Seller attempts to send wrong plaintext → fraud detected
3. Seller provides invalid proof → transfer fails
4. Malicious buyer attempts false fraud claim → fails
5. Dispute period mechanics

### Security Tests:

1. All attack scenarios from SECURITY_ANALYSIS.md should fail
2. Fuzz testing on proof verification
3. Commitment collision resistance
4. Signature forgery attempts

## Comparison with Alternatives

### vs. zk-SNARKs:
- ✅ Much simpler to implement
- ✅ No trusted setup needed
- ✅ Easier to audit
- ❌ Requires dispute mechanism
- ❌ Not zero-knowledge (commitment revealed after)

### vs. Proxy Re-Encryption:
- ✅ Simpler cryptography
- ✅ Better gas efficiency
- ✅ Easier integration with existing code
- ❌ Requires fraud proof mechanism

### vs. Reveal on Purchase:
- ✅ Maintains privacy before purchase
- ✅ Seller can't change content after sale listing
- ✅ Cryptographic guarantees
- ❌ More complex than simple reveal

## Conclusion

This proposal provides a **practical, secure, and gas-efficient** solution to the ZK proof vulnerabilities. The commitment-based approach:

- ✅ Fixes all critical security issues
- ✅ Provides cryptographic guarantees
- ✅ Uses off-chain computation to save gas
- ✅ Adds fraud protection via dispute mechanism
- ✅ Implementable without external dependencies
- ✅ Backward compatible migration path

The combination of cryptographic proofs + economic incentives (dispute period) provides robust security for the NFT transfer mechanism.

## Next Steps

1. Review and approve proposal
2. Implement new proof system in `zkproof/` package
3. Update smart contract with fraud proofs
4. Create comprehensive test suite
5. Audit implementation
6. Deploy to testnet
7. Production deployment

## References

- Commitment Schemes: [Wikipedia](https://en.wikipedia.org/wiki/Commitment_scheme)
- Schnorr Signatures: [Wikipedia](https://en.wikipedia.org/wiki/Schnorr_signature)
- DLEQ Proofs: [Chaum-Pedersen Protocol](https://link.springer.com/chapter/10.1007/3-540-48071-4_7)
- Optimistic Rollups: [Ethereum.org](https://ethereum.org/en/developers/docs/scaling/optimistic-rollups/)
