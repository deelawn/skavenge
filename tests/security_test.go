// Package tests contains security tests demonstrating that vulnerabilities are prevented.
package tests

import (
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

// TestSecurity_AttackPrevented_DifferentPlaintexts verifies that the ElGamal+DLEQ system
// prevents an attacker from providing different plaintexts to seller vs buyer.
func TestSecurity_AttackPrevented_DifferentPlaintexts(t *testing.T) {
	ps := zkproof.NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	realContent := []byte("The treasure is buried under the old oak tree")
	fakeContent := []byte("The treasure is in a volcano (FAKE!)")

	// Attacker generates transfer for REAL content
	realTransfer, err := ps.GenerateVerifiableElGamalTransfer(
		realContent,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Attacker tries to generate a fake buyer ciphertext with DIFFERENT content
	// but using the same DLEQ proof
	fakeTransfer, err := ps.GenerateVerifiableElGamalTransfer(
		fakeContent,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Try to verify with mismatched ciphertexts
	// Real seller cipher + Fake buyer cipher + Real proof
	valid := ps.VerifyElGamalTransfer(
		realTransfer.SellerCipher,   // Seller's real cipher
		fakeTransfer.BuyerCipher,    // Buyer's FAKE cipher (different plaintext!)
		realTransfer.DLEQProof,      // Proof for real cipher
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

	encryptedClue, err := ps.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
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
		minterPrivKey,
		&buyerPrivKey.PublicKey,
	)
	require.NoError(t, err)

	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)

	// Seller commits to REAL r
	rValueHash := crypto.Keccak256Hash(transfer.SharedR.Bytes())
	proofBytes := transfer.DLEQProof.Marshal()

	// Provide proof with correct r hash
	minterAuth, err = util.NewTransactOpts(client, secMinter)
	tx, err = contract.ProvideProof(minterAuth, transferId, proofBytes, buyerCiphertextHash, rValueHash)
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
	// 1. Hash(wrongR) != rValueHash, OR
	// 2. g^wrongR != C1
	require.Error(t, err, "✅ ATTACK PREVENTED: Wrong r value rejected by contract")

	t.Log("✅ Contract prevents seller from revealing wrong r value")
	t.Log("   Two-layer verification: hash commitment + cryptographic proof")
}

// TestSecurity_BuyerCannotDecryptEarly verifies that buyer cannot decrypt
// before seller reveals r value.
func TestSecurity_BuyerCannotDecryptEarly(t *testing.T) {
	ps := zkproof.NewProofSystem()

	sellerKey, _ := ps.GenerateKeyPair()
	buyerKey, _ := ps.GenerateKeyPair()

	plaintext := []byte("Secret clue content")

	// Generate transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
		sellerKey,
		&buyerKey.PublicKey,
	)
	require.NoError(t, err)

	// Buyer has the ciphertext and can verify the proof
	valid := ps.VerifyElGamalTransfer(
		transfer.SellerCipher,
		transfer.BuyerCipher,
		transfer.DLEQProof,
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

	// Generate transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		plaintext,
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
