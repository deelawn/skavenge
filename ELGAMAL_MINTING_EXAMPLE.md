# ElGamal Minting Update Guide

## Contract Changes Summary

The smart contract has been updated to store the ElGamal `r` value on-chain for each clue. This allows owners to always decrypt their clues using just their private key and the on-chain r value.

### Contract Updates:
1. **Clue struct** now includes `uint256 rValue`
2. **mintClue()** signature changed from:
   ```solidity
   function mintClue(bytes calldata encryptedContents, bytes32 solutionHash)
   ```
   to:
   ```solidity
   function mintClue(bytes calldata encryptedContents, bytes32 solutionHash, uint256 rValue)
   ```

3. **New function** `getRValue(uint256 tokenId)` allows owners to retrieve their r value
4. **completeTransfer()** now updates the stored r value when transferring

## Test Updates Needed

After regenerating the Go bindings with `make compile`, update your tests as follows:

### Old Minting Code (ECIES):
```go
// OLD - Don't use this anymore
encryptedClueContent, err := ps.EncryptMessage([]byte(clueContent), &minterPrivKey.PublicKey)
require.NoError(t, err)

tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash)
require.NoError(t, err)
```

### New Minting Code (ElGamal):
```go
// Generate random r value for ElGamal encryption
r, err := rand.Int(rand.Reader, ps.Curve.Params().N)
require.NoError(t, err)

// Encrypt using ElGamal
encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, r)
require.NoError(t, err)

// Marshal to bytes for on-chain storage
encryptedClueContent := encryptedCipher.Marshal()

// Mint with ElGamal ciphertext and r value
tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, r)
require.NoError(t, err)
_, err = util.WaitForTransaction(client, tx)
require.NoError(t, err)
```

### Decrypting Minted Clues:
```go
// Owner can retrieve their clue like this:
tokenId := big.NewInt(1)

// Get the encrypted contents
encryptedBytes, err := contract.GetClueContents(nil, tokenId)
require.NoError(t, err)

// Get the r value (only owner can call this)
rValue, err := contract.GetRValue(&bind.CallOpts{From: ownerAddr}, tokenId)
require.NoError(t, err)

// Unmarshal the ciphertext
cipher := &zkproof.ElGamalCiphertext{}
err = cipher.Unmarshal(encryptedBytes)
require.NoError(t, err)

// Decrypt using r value and private key
plaintext, err := ps.DecryptElGamal(cipher, rValue, ownerPrivKey)
require.NoError(t, err)

fmt.Printf("Decrypted clue: %s\n", string(plaintext))
```

## Transfer Flow Changes

The transfer flow is now simpler because the seller's ciphertext comes from the transfer generation:

```go
// Generate verifiable transfer (creates NEW seller and buyer ciphertexts with same plaintext)
transfer, err := ps.GenerateVerifiableElGamalTransfer(
    clueContent,
    minterPrivKey,
    &buyerPrivKey.PublicKey,
)
require.NoError(t, err)

// The seller's ciphertext for DLEQ verification comes from the transfer
sellerCipher := transfer.SellerCipher  // Use this for verification
buyerCipher := transfer.BuyerCipher

// Buyer verifies DLEQ proof
valid := ps.VerifyElGamalTransfer(
    sellerCipher,
    buyerCipher,
    dleqProof,
    transfer.SellerPubKey,
    transfer.BuyerPubKey,
)
require.True(t, valid, "DLEQ proof verification failed")

// Complete transfer (updates both ciphertext AND r value on-chain)
completeTx, err := contract.CompleteTransfer(
    minterAuth,
    transferId,
    buyerCiphertextBytes,
    transfer.SharedR,  // New r value
)
```

## Security Properties

With r stored on-chain, decryption still requires BOTH:
1. **r value** (public, stored on-chain)
2. **Private key** (secret, owned by token holder)

The encryption key derivation is:
```
sharedSecret = C1^privKey  // Requires private key
key = Hash(r || sharedSecret)  // Requires both r and private key
```

So even though r is public:
- ✅ Owner can decrypt (has r from chain + their privKey)
- ❌ Non-owner cannot decrypt (has r but not the privKey)
- ✅ No risk of losing r value (always on-chain)
- ✅ Transfer updates both ciphertext and r atomically

## Files to Update

After regenerating bindings:
1. `tests/transfer_test.go` - All 4 test functions
2. `tests/mint_test.go` - Update MintClue calls
3. `tests/solve_test.go` - Update MintClue calls
4. `tests/security_test.go` - Update test setup

## Regenerating Bindings

Run this command to regenerate the Go bindings:
```bash
make compile
```

Or manually:
```bash
# Compile contract
solcjs --abi --bin --base-path eth/node_modules --include-path eth/node_modules --optimize -o eth/build eth/skavenge.sol

# Rename outputs
cd eth/build
cp _home_user_skavenge_eth_skavenge_sol_Skavenge.abi Skavenge.abi
cp _home_user_skavenge_eth_skavenge_sol_Skavenge.bin Skavenge.bin

# Generate Go bindings
abigen --bin Skavenge.bin --abi Skavenge.abi --pkg bindings --type Skavenge --out ../bindings/bindings.go
```
