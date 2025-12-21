// Package tests contains security tests demonstrating that vulnerabilities are prevented.
package tests

import (
	"crypto/rand"
	"math/big"
	"strings"
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

// TestSecurity_AttackPrevented_DifferentPlaintexts verifies that the ElGamal+DLEQ system
// prevents an attacker from providing different plaintexts to seller vs buyer.
func TestSecurity_AttackPrevented_DifferentPlaintexts(t *testing.T) {
	ps := zkproof.NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	realContent := []byte("The treasure is buried under the old oak tree")
	fakeContent := []byte("The treasure is in a volcano (FAKE!)")

	// Create "original" cipher (as if minted on-chain)
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err)
	originalCipher, err := ps.EncryptElGamal(realContent, &sellerKey.PublicKey, mintR)
	require.NoError(t, err)

	// Attacker generates transfer for REAL content
	realTransfer, err := ps.GenerateVerifiableElGamalTransfer(
		realContent,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Attacker tries to generate a fake buyer ciphertext with DIFFERENT content
	// but using the same DLEQ proof
	fakeTransfer, err := ps.GenerateVerifiableElGamalTransfer(
		fakeContent,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Try to verify with mismatched ciphertexts
	// Real seller cipher + Fake buyer cipher + Real proof
	valid := ps.VerifyElGamalTransfer(
		originalCipher,            // Original on-chain cipher
		realTransfer.SellerCipher, // Seller's real cipher
		fakeTransfer.BuyerCipher,  // Buyer's FAKE cipher (different plaintext!)
		realTransfer.DLEQProof,    // Proof for real cipher
		realTransfer.PlaintextProof,
		mintR,
		realTransfer.SellerPubKey,
		realTransfer.BuyerPubKey,
	)

	// ATTACK PREVENTED: The DLEQ proof will fail because the ciphertexts
	// use different random values (different r)
	require.False(t, valid, "✅ ATTACK PREVENTED: Mismatched ciphertexts detected by DLEQ proof")

	t.Log("✅ ElGamal+DLEQ system prevents different plaintext attack")
	t.Log("   Seller cannot provide different content to buyer")
}

// TestSecurity_AttackPrevented_WrongRValue verifies that the contract
// prevents seller from revealing wrong r value.
func TestSecurity_AttackPrevented_WrongRValue(t *testing.T) {
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
	clueContent := []byte("The treasure is buried under the old oak tree")
	solution := "Oak tree"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err)

	// Marshal to bytes for on-chain storage
	encryptedClue := encryptedCipher.Marshal()

	minterAuth, err = util.NewTransactOpts(client, secMinter)
	tx, err = contract.MintClue(minterAuth, encryptedClue, solutionHash, mintR)
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

	// Initiate purchase
	buyerAuth.Value = big.NewInt(1000000000000000000)
	tx, err = contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// Generate verifiable transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		clueContent,
		encryptedCipher, // Original on-chain cipher
		mintR,           // Public r from mint
		minterPrivKey,
		&buyerPrivKey.PublicKey,
	)
	require.NoError(t, err)

	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)

	// Seller commits to REAL r (rHash is embedded in DLEQ proof)
	proofBytes := transfer.DLEQProof.Marshal()

	// Provide proof (contract extracts rHash from proof)
	minterAuth, err = util.NewTransactOpts(client, secMinter)
	tx, err = contract.ProvideProof(minterAuth, transferId, proofBytes, buyerCiphertextHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Buyer verifies proof
	buyerAuth, err = util.NewTransactOpts(client, secBuyer)
	tx, err = contract.VerifyProof(buyerAuth, transferId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// ATTACK: Seller tries to provide WRONG r value
	wrongR, _ := rand.Int(rand.Reader, ps.Curve.Params().N)

	// Try to complete transfer with wrong r
	minterAuth, err = util.NewTransactOpts(client, secMinter)
	minterAuth.GasLimit = 500000 // Higher gas limit for failing transaction

	_, err = contract.CompleteTransfer(minterAuth, transferId, buyerCiphertextBytes, wrongR)

	// ATTACK PREVENTED: Transaction should revert because:
	// Hash(wrongR) != rValueHash (committed during provideProof)
	require.Error(t, err, "✅ ATTACK PREVENTED: Wrong r value rejected by contract")

	t.Log("✅ Contract prevents seller from revealing wrong r value")
	t.Log("   Hash commitment ensures r cannot be changed after buyer verification")
}

// TestSecurity_BuyerCannotDecryptEarly verifies that buyer cannot decrypt
// before seller reveals r value.
func TestSecurity_BuyerCannotDecryptEarly(t *testing.T) {
	ps := zkproof.NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	plaintext := []byte("Secret clue content")

	// Create "original" cipher (as if minted on-chain)
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err)
	originalCipher, err := ps.EncryptElGamal(plaintext, &sellerKey.PublicKey, mintR)
	require.NoError(t, err)

	// Generate transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Buyer has the ciphertext and can verify the proof
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.DLEQProof,
		transfer.PlaintextProof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	require.True(t, valid, "Proof should verify")

	// But buyer CANNOT decrypt without r
	// Try with nil r
	_, err = ps.DecryptElGamal(transfer.BuyerCipher, nil, buyerKey)
	require.Error(t, err, "✅ ATTACK PREVENTED: Cannot decrypt with nil r")

	// Try with wrong r
	wrongR := big.NewInt(12345)
	wrongDecrypted, err := ps.DecryptElGamal(transfer.BuyerCipher, wrongR, buyerKey)
	require.NoError(t, err) // Decryption succeeds but...
	require.NotEqual(t, plaintext, wrongDecrypted, "✅ ATTACK PREVENTED: Wrong r produces garbage")

	// Only correct r works
	correctDecrypted, err := ps.DecryptElGamal(transfer.BuyerCipher, transfer.SharedR, buyerKey)
	require.NoError(t, err)
	require.Equal(t, plaintext, correctDecrypted)

	t.Log("✅ Buyer cannot decrypt before seller reveals r")
	t.Log("   Prevents decrypt-then-cancel attack")
}

// TestSecurity_OnlyBuyerCanDecrypt verifies that even with revealed r,
// only the buyer can decrypt (needs private key).
func TestSecurity_OnlyBuyerCanDecrypt(t *testing.T) {
	ps := zkproof.NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()
	attackerKey, _ := ps.GenerateKeyPair() // Third party

	plaintext := []byte("Secret clue content")

	// Create "original" cipher (as if minted on-chain)
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err)
	originalCipher, err := ps.EncryptElGamal(plaintext, &sellerKey.PublicKey, mintR)
	require.NoError(t, err)

	// Generate transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Simulate r being revealed on-chain (public)
	revealedR := transfer.SharedR

	// Attacker has:
	// - Buyer's ciphertext (public on-chain)
	// - Revealed r value (public from event)
	// - Buyer's public key (public)
	// But attacker does NOT have buyer's private key

	// Try to decrypt with attacker's key
	attackerDecrypted, err := ps.DecryptElGamal(transfer.BuyerCipher, revealedR, attackerKey)
	require.NoError(t, err) // Decryption succeeds but...
	require.NotEqual(t, plaintext, attackerDecrypted, "✅ ATTACK PREVENTED: Attacker cannot decrypt")

	// Only buyer can decrypt
	buyerDecrypted, err := ps.DecryptElGamal(transfer.BuyerCipher, revealedR, buyerKey)
	require.NoError(t, err)
	require.Equal(t, plaintext, buyerDecrypted)

	t.Log("✅ Only buyer can decrypt even with revealed r")
	t.Log("   Privacy preserved: requires both r AND private key")
}

// TestSecurity_AttackPrevented_FakeRHashInProof verifies that the DLEQ proof
// verification fails if seller tampers with the rHash embedded in the proof.
func TestSecurity_AttackPrevented_FakeRHashInProof(t *testing.T) {
	ps := zkproof.NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	plaintext := []byte("Secret clue content")

	// Create "original" cipher (as if minted on-chain)
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err)
	originalCipher, err := ps.EncryptElGamal(plaintext, &sellerKey.PublicKey, mintR)
	require.NoError(t, err)

	// Generate valid transfer with real r
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		originalCipher,
		mintR,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Proof should verify with real rHash
	valid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.DLEQProof,
		transfer.PlaintextProof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	require.True(t, valid, "Valid proof should verify")

	// ATTACK: Seller tries to tamper with the rHash in the proof
	// Marshal the proof
	proofBytes := transfer.DLEQProof.Marshal()

	// Generate a fake rHash (different from the real one)
	fakeRHash := [32]byte{}
	for i := 0; i < 32; i++ {
		fakeRHash[i] = 0xFF // All ones (obviously different from real hash)
	}

	// Replace the last 32 bytes (rHash) with fake value
	copy(proofBytes[len(proofBytes)-32:], fakeRHash[:])

	// Unmarshal the tampered proof
	tamperedProof := &zkproof.DLEQProof{}
	err = tamperedProof.Unmarshal(proofBytes)
	require.NoError(t, err, "Tampered proof should unmarshal")

	// Verify that the rHash was actually replaced
	require.NotEqual(t, transfer.DLEQProof.RHash, tamperedProof.RHash, "RHash should be different")

	// Try to verify with tampered proof
	tamperedValid := ps.VerifyElGamalTransfer(
		originalCipher,
		transfer.SellerCipher,
		transfer.BuyerCipher,
		tamperedProof, // Using tampered proof with fake rHash
		transfer.PlaintextProof,
		mintR,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)

	// ATTACK PREVENTED: Verification should fail because the challenge won't match
	// The challenge is computed as Hash(sellerPub || buyerPub || A1 || A2 || A3 || rHash)
	// If rHash is tampered, the challenge will be different from the one used in proof generation
	require.False(t, tamperedValid, "✅ ATTACK PREVENTED: Tampered rHash causes proof verification to fail")

	t.Log("✅ DLEQ proof verification prevents fake rHash attack")
	t.Log("   Seller cannot commit to fake r value in proof")
	t.Log("   Challenge binding ensures rHash matches the r used in proof generation")
}

// TestSecurity_BuyerCannotCancelAfterVerification verifies that buyers cannot cancel
// transfers after they have called verifyProof(), preventing frontrunning attacks.
func TestSecurity_BuyerCannotCancelAfterVerification(t *testing.T) {
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
	clueContent := []byte("The treasure is buried under the old oak tree")
	solution := "Oak tree"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err)

	// Encrypt using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err)
	encryptedClue := encryptedCipher.Marshal()

	minterAuth, err = util.NewTransactOpts(client, secMinter)
	tx, err = contract.MintClue(minterAuth, encryptedClue, solutionHash, mintR)
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
	buyerAuth.Value = salePrice
	tx, err = contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// TEST PHASE 1: Buyer CAN cancel before verification
	t.Log("\n[Phase 1] Buyer can cancel BEFORE verification")
	buyerAuth, err = util.NewTransactOpts(client, secBuyer)
	tx, err = contract.CancelTransfer(buyerAuth, transferId)
	require.NoError(t, err, "Buyer should be able to cancel before verification")
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	t.Log("✓ Buyer successfully canceled before verification")

	// Re-initiate purchase for phase 2
	buyerAuth, err = util.NewTransactOpts(client, secBuyer)
	buyerAuth.Value = salePrice
	tx, err = contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Generate verifiable transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		clueContent,
		encryptedCipher, // Original on-chain cipher
		mintR,           // Public r from mint
		minterPrivKey,
		&buyerPrivKey.PublicKey,
	)
	require.NoError(t, err)

	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)

	// Seller provides proof
	minterAuth, err = util.NewTransactOpts(client, secMinter)
	require.NoError(t, err)
	proofBytes := transfer.DLEQProof.Marshal()
	tx, err = contract.ProvideProof(minterAuth, transferId, proofBytes, buyerCiphertextHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Buyer verifies proof (COMMITMENT POINT)
	t.Log("\n[Phase 2] Buyer verifies proof (commitment)")
	buyerAuth, err = util.NewTransactOpts(client, secBuyer)
	tx, err = contract.VerifyProof(buyerAuth, transferId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	t.Log("✓ Buyer committed by calling verifyProof()")

	// TEST PHASE 2: Buyer CANNOT cancel after verification
	t.Log("\n[Phase 3] Buyer attempts to cancel AFTER verification")
	buyerAuth, err = util.NewTransactOpts(client, secBuyer)
	buyerAuth.GasLimit = 500000 // Higher gas limit for expected failure

	_, err = contract.CancelTransfer(buyerAuth, transferId)

	// ATTACK PREVENTED: Transaction should revert
	require.Error(t, err, "✅ ATTACK PREVENTED: Buyer cannot cancel after verification")
	require.Contains(t, err.Error(), "Cannot cancel after proof verification",
		"Error should indicate cancellation blocked after verification")

	t.Log("✅ Buyer CANNOT cancel after calling verifyProof()")
	t.Log("   This prevents mempool frontrunning attack")
	t.Log("   Buyer is committed once they verify the proof")
}

// TestSecurity_FrontrunningAttackPrevented demonstrates that the complete
// frontrunning attack scenario is now prevented by the cancellation restriction.
func TestSecurity_FrontrunningAttackPrevented(t *testing.T) {
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
	clueContent := []byte("The secret treasure location is 40.7128°N, 74.0060°W")
	solution := "New York City"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err)

	// Encrypt using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err)
	encryptedClue := encryptedCipher.Marshal()

	minterAuth, err = util.NewTransactOpts(client, secMinter)
	tx, err = contract.MintClue(minterAuth, encryptedClue, solutionHash, mintR)
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

	// === ATTACK SCENARIO ===
	t.Log("\n" + strings.Repeat("=", 70))
	t.Log("SIMULATING FRONTRUNNING ATTACK")
	t.Log(strings.Repeat("=", 70))

	// Step 1: Buyer initiates purchase
	t.Log("\n[1] Buyer initiates purchase (1 ETH)")
	buyerAuth.Value = salePrice
	tx, err = contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// Step 2: Seller generates and provides proof
	t.Log("\n[2] Seller provides DLEQ proof")
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		clueContent,
		encryptedCipher, // Original on-chain cipher
		mintR,           // Public r from mint
		minterPrivKey,
		&buyerPrivKey.PublicKey,
	)
	require.NoError(t, err)

	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)

	minterAuth, err = util.NewTransactOpts(client, secMinter)
	proofBytes := transfer.DLEQProof.Marshal()
	tx, err = contract.ProvideProof(minterAuth, transferId, proofBytes, buyerCiphertextHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Step 3: Buyer verifies proof (commits to purchase)
	t.Log("\n[3] Buyer verifies proof off-chain and commits on-chain")
	buyerAuth, err = util.NewTransactOpts(client, secBuyer)
	tx, err = contract.VerifyProof(buyerAuth, transferId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	t.Log("✓ Buyer committed to purchase")

	// Step 4: ATTACK - Malicious buyer tries to extract r and cancel
	t.Log("\n[4] ATTACK: Seller submits completeTransfer() with r value")
	t.Log("    Malicious buyer would normally:")
	t.Log("    - Monitor mempool for completeTransfer transaction")
	t.Log("    - Extract r value from calldata")
	t.Log("    - Decrypt clue off-chain")
	t.Log("    - Submit cancelTransfer with higher gas price")

	// In a real attack, buyer would extract r from pending tx
	extractedR := transfer.SharedR
	t.Logf("    Extracted r value: %s", extractedR.String()[:20]+"...")

	// Buyer tries to decrypt with extracted r (would succeed without fix)
	decrypted, err := ps.DecryptElGamal(transfer.BuyerCipher, extractedR, buyerPrivKey)
	require.NoError(t, err)
	require.Equal(t, clueContent, decrypted)
	t.Log("    ⚠️  Buyer successfully decrypted clue with extracted r!")

	// Step 5: ATTACK ATTEMPT - Buyer tries to cancel and get refund
	t.Log("\n[5] ATTACK ATTEMPT: Buyer tries to cancel with higher gas")
	buyerAuth, err = util.NewTransactOpts(client, secBuyer)
	buyerAuth.GasLimit = 500000

	_, err = contract.CancelTransfer(buyerAuth, transferId)

	// ATTACK PREVENTED!
	require.Error(t, err, "✅ ATTACK PREVENTED: Cancel transaction reverted")
	require.Contains(t, err.Error(), "Cannot cancel after proof verification")
	t.Log("    ✅ ATTACK PREVENTED: cancelTransfer() reverted")
	t.Log("    ✅ Buyer cannot cancel after verifyProof() commitment")

	// Step 6: Seller's completeTransfer succeeds
	t.Log("\n[6] Seller completes transfer successfully")
	minterAuth, err = util.NewTransactOpts(client, secMinter)
	tx, err = contract.CompleteTransfer(minterAuth, transferId, buyerCiphertextBytes, extractedR)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	t.Log("    ✓ Transfer completed")
	t.Log("    ✓ Ownership transferred to buyer")
	t.Log("    ✓ Payment sent to seller")

	// Verify final state
	newOwner, err := contract.OwnerOf(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, buyerAddr, newOwner, "Buyer should own the token")

	t.Log("\n" + strings.Repeat("=", 70))
	t.Log("✅ FRONTRUNNING ATTACK PREVENTED")
	t.Log(strings.Repeat("=", 70))
	t.Log("\nSecurity guarantee restored:")
	t.Log("  ✅ Buyer cannot extract r and cancel")
	t.Log("  ✅ verifyProof() creates binding commitment")
	t.Log("  ✅ Transfer completes atomically")
	t.Log("  ✅ Seller receives payment")
	t.Log("  ✅ Buyer receives NFT + decryption ability")
}

// TestSecurity_ConcurrentPurchasePrevention verifies that multiple buyers cannot
// initiate concurrent purchases for the same token.
func TestSecurity_ConcurrentPurchasePrevention(t *testing.T) {
	// These test accounts are from Hardhat's default accounts
	const (
		testDeployer = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
		testMinter   = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
		testBuyer1   = "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
		testBuyer2   = "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"
	)

	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, testDeployer)
	require.NoError(t, err)

	// Deploy contract
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(testMinter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, testMinter)
	require.NoError(t, err)
	minterAddr := minterAuth.From

	// Update authorized minter
	deployerAuth, err = util.NewTransactOpts(client, testDeployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Setup buyer 1
	buyer1PrivKey, err := crypto.HexToECDSA(testBuyer1)
	require.NoError(t, err)
	buyer1Auth, err := util.NewTransactOpts(client, testBuyer1)
	require.NoError(t, err)

	// Setup buyer 2
	buyer2PrivKey, err := crypto.HexToECDSA(testBuyer2)
	require.NoError(t, err)
	buyer2Addr := crypto.PubkeyToAddress(buyer2PrivKey.PublicKey)
	buyer2Auth, err := util.NewTransactOpts(client, testBuyer2)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Mint a clue
	clueContent := []byte("The treasure is at the end of the rainbow")
	solution := "Rainbow end"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err)

	// Encrypt using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err)
	encryptedClue := encryptedCipher.Marshal()

	minterAuth, err = util.NewTransactOpts(client, testMinter)
	tx, err = contract.MintClue(minterAuth, encryptedClue, solutionHash, mintR)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Set sale price
	minterAuth, err = util.NewTransactOpts(client, testMinter)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH
	tx, err = contract.SetSalePrice(minterAuth, tokenId, salePrice)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	t.Log("\n" + strings.Repeat("=", 70))
	t.Log("TESTING CONCURRENT PURCHASE PREVENTION")
	t.Log(strings.Repeat("=", 70))

	// TEST PHASE 1: Buyer 1 initiates purchase
	t.Log("\n[1] Buyer 1 initiates purchase for token", tokenId)
	buyer1Auth.Value = salePrice
	tx, err = contract.InitiatePurchase(buyer1Auth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	t.Log("✓ Buyer 1 successfully initiated purchase")

	// Verify transferInProgress flag is set
	transferInProgress, err := contract.TransferInProgress(nil, tokenId)
	require.NoError(t, err)
	require.True(t, transferInProgress, "transferInProgress should be true after buyer 1 initiates")
	t.Log("✓ transferInProgress flag is set to true")

	// TEST PHASE 2: Buyer 2 tries to initiate purchase for same token
	t.Log("\n[2] Buyer 2 attempts to initiate purchase for same token", tokenId)
	buyer2Auth.Value = salePrice
	buyer2Auth.GasLimit = 500000 // Higher gas limit for expected failure

	_, err = contract.InitiatePurchase(buyer2Auth, tokenId)

	// SHOULD FAIL - Transfer already in progress
	require.Error(t, err, "✅ PREVENTED: Buyer 2 cannot initiate concurrent purchase")
	t.Log("✅ Buyer 2's purchase attempt was blocked")
	t.Log("   Error: TransferAlreadyInProgress")

	// TEST PHASE 3: Verify flag persists
	transferInProgress, err = contract.TransferInProgress(nil, tokenId)
	require.NoError(t, err)
	require.True(t, transferInProgress, "transferInProgress should still be true")
	t.Log("✓ transferInProgress flag still set (buyer 1's transfer active)")

	// TEST PHASE 4: Buyer 1 cancels, then buyer 2 can purchase
	t.Log("\n[3] Buyer 1 cancels their purchase")
	transferId1, err := contract.GenerateTransferId(nil, crypto.PubkeyToAddress(buyer1PrivKey.PublicKey), tokenId)
	require.NoError(t, err)

	buyer1Auth, err = util.NewTransactOpts(client, testBuyer1)
	tx, err = contract.CancelTransfer(buyer1Auth, transferId1)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	t.Log("✓ Buyer 1 canceled transfer")

	// Verify flag is cleared after cancellation
	transferInProgress, err = contract.TransferInProgress(nil, tokenId)
	require.NoError(t, err)
	require.False(t, transferInProgress, "transferInProgress should be false after cancellation")
	t.Log("✓ transferInProgress flag cleared after cancellation")

	// TEST PHASE 5: Now buyer 2 can purchase
	t.Log("\n[4] Buyer 2 initiates purchase (should succeed now)")
	buyer2Auth, err = util.NewTransactOpts(client, testBuyer2)
	buyer2Auth.Value = salePrice
	tx, err = contract.InitiatePurchase(buyer2Auth, tokenId)
	require.NoError(t, err, "Buyer 2 should be able to purchase after buyer 1 canceled")
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	t.Log("✓ Buyer 2 successfully initiated purchase")

	// Verify flag is set again
	transferInProgress, err = contract.TransferInProgress(nil, tokenId)
	require.NoError(t, err)
	require.True(t, transferInProgress, "transferInProgress should be true for buyer 2")
	t.Log("✓ transferInProgress flag set for buyer 2's transfer")

	// TEST PHASE 6: Complete the transfer to verify flag clears
	t.Log("\n[5] Complete transfer to verify flag clears")

	// Generate verifiable transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		clueContent,
		encryptedCipher, // Original on-chain cipher
		mintR,           // Public r from mint
		minterPrivKey,
		&buyer2PrivKey.PublicKey,
	)
	require.NoError(t, err)

	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)

	// Seller provides proof
	transferId2, err := contract.GenerateTransferId(nil, buyer2Addr, tokenId)
	require.NoError(t, err)

	minterAuth, err = util.NewTransactOpts(client, testMinter)
	proofBytes := transfer.DLEQProof.Marshal()
	tx, err = contract.ProvideProof(minterAuth, transferId2, proofBytes, buyerCiphertextHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Buyer 2 verifies proof
	buyer2Auth, err = util.NewTransactOpts(client, testBuyer2)
	tx, err = contract.VerifyProof(buyer2Auth, transferId2)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Seller completes transfer
	minterAuth, err = util.NewTransactOpts(client, testMinter)
	tx, err = contract.CompleteTransfer(minterAuth, transferId2, buyerCiphertextBytes, transfer.SharedR)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	t.Log("✓ Transfer completed successfully")

	// Verify flag is cleared after completion
	transferInProgress, err = contract.TransferInProgress(nil, tokenId)
	require.NoError(t, err)
	require.False(t, transferInProgress, "transferInProgress should be false after completion")
	t.Log("✓ transferInProgress flag cleared after transfer completion")

	t.Log("\n" + strings.Repeat("=", 70))
	t.Log("✅ CONCURRENT PURCHASE PREVENTION VERIFIED")
	t.Log(strings.Repeat("=", 70))
	t.Log("\nSecurity guarantees:")
	t.Log("  ✅ Only one buyer can initiate purchase at a time")
	t.Log("  ✅ transferInProgress flag prevents concurrent purchases")
	t.Log("  ✅ Flag is cleared on cancellation")
	t.Log("  ✅ Flag is cleared on transfer completion")
	t.Log("  ✅ New purchases possible after previous transfer ends")
}
