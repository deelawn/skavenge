# Smart Contract Update Proposal: ElGamal Transfer with R Reveal

## Problem Statement

The current smart contract transfer flow doesn't include a mechanism for the seller to reveal the `r` value, which is now **required** for the buyer to decrypt the ciphertext.

With the updated ElGamal implementation:
- Encryption key = `Hash(r || buyerPubKey)`
- Buyer CANNOT decrypt with just their private key
- Buyer MUST have the `r` value revealed by the seller after payment

## Current Flow (Insufficient)

```
1. initiatePurchase()  - Buyer commits to purchase
2. provideProof()      - Seller provides DLEQ proof + buyer ciphertext hash
3. verifyProof()       - Buyer verifies proof off-chain, calls this to confirm
4. completeTransfer()  - Seller provides buyer ciphertext
```

**Problem**: No mechanism to reveal `r`, so buyer can't decrypt!

## Updated Flow (With R Reveal)

```
1. initiatePurchase()  - Buyer commits to purchase
2. provideProof()      - Seller provides DLEQ proof + buyer ciphertext hash
3. verifyProof()       - Buyer verifies proof off-chain, calls this to confirm
4. completeTransfer()  - Seller provides buyer ciphertext + r value
5. Buyer decrypts      - Off-chain using revealed r
```

## Smart Contract Changes Required

### Update TokenTransfer struct

```solidity
struct TokenTransfer {
    address buyer;
    uint256 tokenId;
    uint256 value;
    uint256 initiatedAt;
    bytes proof;                    // DLEQ proof
    bytes32 newClueHash;            // Hash of buyer ciphertext
    bytes32 rValueHash;             // NEW: Hash commitment to r value
    bool proofVerified;
    uint256 proofProvidedAt;
    uint256 verifiedAt;
    bool rRevealed;                 // NEW: Whether r has been revealed
}
```

### Update provideProof function

```solidity
function provideProof(
    bytes32 transferId,
    bytes calldata proof,
    bytes32 newClueHash,
    bytes32 rValueHash          // NEW: Commitment to r
) external nonReentrant {
    TokenTransfer storage transfer = transfers[transferId];
    require(transfer.buyer != address(0), "Transfer does not exist");
    require(ownerOf(transfer.tokenId) == msg.sender, "Not token owner");

    require(
        block.timestamp - transfer.initiatedAt <= TRANSFER_TIMEOUT,
        "Transfer expired"
    );

    transfer.proof = proof;
    transfer.newClueHash = newClueHash;
    transfer.rValueHash = rValueHash;    // Store commitment
    transfer.proofProvidedAt = block.timestamp;

    emit ProofProvided(transferId, proof, newClueHash, rValueHash);
}
```

### Update completeTransfer function

```solidity
function completeTransfer(
    bytes32 transferId,
    bytes calldata newEncryptedContents,
    uint256 rValue              // NEW: The actual r value
) external nonReentrant {
    TokenTransfer storage transfer = transfers[transferId];
    require(transfer.buyer != address(0), "Transfer does not exist");
    require(ownerOf(transfer.tokenId) == msg.sender, "Not token owner");
    require(transfer.verifiedAt > 0, "Proof not verified");
    require(
        block.timestamp - transfer.verifiedAt <= TRANSFER_TIMEOUT,
        "Transfer completion expired"
    );

    // Verify the hash of the new encrypted contents
    require(
        keccak256(newEncryptedContents) == transfer.newClueHash,
        "Content hash mismatch"
    );

    // NEW: Verify r value matches commitment
    require(
        keccak256(abi.encodePacked(rValue)) == transfer.rValueHash,
        "R value mismatch"
    );

    require(!clues[transfer.tokenId].isSolved, "Solved clue cannot be transferred");

    // Update the clue contents
    clues[transfer.tokenId].encryptedContents = newEncryptedContents;
    clues[transfer.tokenId].solveAttempts = 0;

    // Mark r as revealed
    transfer.rRevealed = true;

    // Transfer ownership
    _safeTransfer(msg.sender, transfer.buyer, transfer.tokenId, "");

    // Send payment to seller
    (bool sent, ) = payable(msg.sender).call{value: transfer.value}("");
    require(sent, "Failed to send Ether");

    if (cluesForSale[transfer.tokenId]) {
        _removeFromForSaleList(transfer.tokenId);
        clues[transfer.tokenId].salePrice = 0;
        emit SalePriceRemoved(transfer.tokenId);
    }

    delete transfers[transferId];

    emit TransferCompleted(transferId, rValue);  // Include r in event
}
```

### Update Events

```solidity
event ProofProvided(
    bytes32 indexed transferId,
    bytes proof,
    bytes32 newClueHash,
    bytes32 rValueHash           // NEW
);

event TransferCompleted(
    bytes32 indexed transferId,
    uint256 rValue               // NEW: Buyer can get r from event
);
```

## Security Properties

1. **Seller commits to r** - Through `rValueHash` in `provideProof()`
2. **Buyer verifies DLEQ proof** - Before calling `verifyProof()` (locks payment)
3. **Buyer receives payment** - Only after revealing correct `r` and ciphertext
4. **R is revealed on-chain** - Buyer can decrypt using revealed `r`
5. **Buyer can verify** - Check that revealed `r` matches commitment

## Attack Prevention

### Attack: Seller provides wrong r
- **Prevention**: `keccak256(abi.encodePacked(rValue)) == transfer.rValueHash`
- **Result**: Transaction reverts, buyer keeps payment

### Attack: Buyer decrypts before paying
- **Prevention**: Encryption key = `Hash(r || buyerPubKey)`, buyer doesn't have `r` yet
- **Result**: Buyer cannot decrypt until seller reveals `r` in `completeTransfer()`

### Attack: Seller doesn't reveal r after payment
- **Prevention**: Payment released AFTER valid `r` is revealed
- **Result**: If seller doesn't call `completeTransfer()` with valid `r`, they don't get paid

## Gas Optimization

- **r value**: Store as `uint256` (32 bytes) - cheaper than `bytes`
- **Commitment**: Use `keccak256(abi.encodePacked(rValue))` for efficient hashing
- **Event emission**: Include `r` in event so buyer doesn't need separate call

## Migration from Current Contract

### Option 1: Deploy new contract
- Deploy updated contract with new transfer flow
- Migrate existing NFTs if needed

### Option 2: Upgrade existing contract (if using proxy pattern)
- Add storage slots for new fields
- Update functions with new logic
- Ensure backward compatibility

## Next Steps

1. Review and approve this proposal
2. Update `skavenge.sol` with the changes
3. Update tests in `transfer_test.go` to match new flow
4. Test complete end-to-end flow
5. Deploy to testnet for integration testing
