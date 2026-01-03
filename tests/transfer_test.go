// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/deelawn/skavenge/eth/bindings"
	"github.com/deelawn/skavenge/tests/util"
	"github.com/deelawn/skavenge/zkproof"
)

var (
	deployer = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	minter   = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	buyer    = "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
	other    = "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"
)

// TestAlteredTransfer tests an attack where the seller provides an altered content to the buyer.
// The attack should be detected by the plaintext equality proof.
func TestAlteredTransfer(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, address, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Create event listener
	listener, err := util.NewEventListener(client, contract, address)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAddr := minterAuth.From

	// Update authorized minter to minter address.
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Sample data for the clue
	clueContent := []byte("Find the hidden treasure in the old oak tree")
	solution := "Oak tree"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the clue content using ElGamal for the minter
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()

	// Mint a new clue with ElGamal ciphertext and r value
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Verify the clue is minted successfully
	owner, err := contract.OwnerOf(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, crypto.PubkeyToAddress(minterPrivKey.PublicKey), owner, "Minter should be the owner")

	// Set sale price
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH
	timeout := big.NewInt(180)                   // 3 minutes
	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice, timeout)
	require.NoError(t, err)
	salePriceReceipt, err := util.WaitForTransaction(client, salePriceTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), salePriceReceipt.Status, "Set sale price transaction failed")

	// Initiate purchase from buyer account
	buyerAuth.Value = big.NewInt(1000000000000000000) // 1 ETH
	buyTx, err := contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	buyReceipt, err := util.WaitForTransaction(client, buyTx)
	require.NoError(t, err)

	// Verify the TransferInitiated event is emitted
	transferInitiatedFound, err := listener.CheckEvent(buyReceipt, "TransferInitiated")
	require.NoError(t, err)
	require.True(t, transferInitiatedFound, "TransferInitiated event not found")

	// Generate transfer ID
	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// ATTACK SCENARIO: Seller attempts to provide altered content
	// After our ZK proof implementation, this attack should be DETECTED!
	newClueContent := []byte("Content altered!")

	// Generate ElGamal verifiable transfer with the ALTERED content
	// The proof system now requires the original cipher and mintR to verify against
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		newClueContent,
		encryptedCipher, // Original on-chain cipher
		mintR,           // Public r value from mint transaction
		minterPrivKey,
		&buyerPrivKey.PublicKey,
	)
	require.NoError(t, err, "Failed to generate verifiable transfer")

	// Marshal the buyer ciphertext for on-chain storage
	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)

	// Provide proof to the contract (r value hash is extracted from proof)
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Marshal entire transfer proof (includes both DLEQ and Plaintext proofs)
	proofBytes := transfer.Proof.Marshal()

	proofTx, err := contract.ProvideProof(minterAuth, transferId, proofBytes, buyerCiphertextHash)
	require.NoError(t, err)
	proofReceipt, err := util.WaitForTransaction(client, proofTx)
	require.NoError(t, err)

	// Verify the ProofProvided event is emitted
	proofProvidedFound, err := listener.CheckEvent(proofReceipt, "ProofProvided")
	require.NoError(t, err)
	require.True(t, proofProvidedFound, "ProofProvided event not found")

	// Buyer verifies the proof off-chain
	transferData, err := contract.Transfers(&bind.CallOpts{}, transferId)
	require.NoError(t, err)

	// Unmarshal the entire transfer proof (DLEQ + Plaintext)
	proof := &zkproof.TransferProof{}
	err = proof.Unmarshal(transferData.Proof)
	require.NoError(t, err)

	// For the buyer to verify, they need:
	// 1. Original on-chain cipher (encryptedCipher)
	// 2. New seller cipher (for DLEQ)
	// 3. New buyer cipher
	// 4. Complete proof (DLEQ + Plaintext proofs)
	// 5. mintR (public from mint transaction)
	sellerCipher := transfer.SellerCipher
	buyerCipher := transfer.BuyerCipher

	// ATTACK DETECTION: Since content was altered, verification should FAIL!
	// The plaintext equality proof will detect that buyerCipher encrypts
	// different content than the original on-chain cipher.
	valid := ps.VerifyElGamalTransfer(
		encryptedCipher, // Original on-chain cipher
		sellerCipher,
		buyerCipher,
		proof,
		mintR, // Public r from mint tx
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)

	// THIS IS THE KEY ASSERTION: Verification should FAIL because content was altered!
	require.False(t, valid, "Proof verification should FAIL when content is altered - attack detected!")

	t.Log("✅ ATTACK PREVENTED: Altered content was detected by plaintext equality proof!")
	t.Log("   Buyer's off-chain verification correctly rejected the fraudulent transfer")

	// Since the attack was detected, buyer would cancel the transfer
	// Cancel the transfer
	buyerAuth, err = util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	cancelTx, err := contract.CancelTransfer(buyerAuth, transferId)
	require.NoError(t, err)
	cancelReceipt, err := util.WaitForTransaction(client, cancelTx)
	require.NoError(t, err)

	// Verify the TransferCancelled event is emitted
	transferCancelledFound, err := listener.CheckEvent(cancelReceipt, "TransferCancelled")
	require.NoError(t, err)
	require.True(t, transferCancelledFound, "TransferCancelled event not found")

	// Verify the TransferCancelled event parameters
	cancelEvents, err := listener.GetEventsByName(cancelReceipt, "TransferCancelled")
	require.NoError(t, err)
	require.Len(t, cancelEvents, 1, "Expected exactly one TransferCancelled event")

	cancelEventData := cancelEvents[0].(map[string]interface{})
	emittedTransferId := cancelEventData["transferId"].([32]byte)
	emittedCancelledBy := cancelEventData["cancelledBy"].(common.Address)

	require.Equal(t, transferId, emittedTransferId, "TransferCancelled event should contain correct transferId")
	require.Equal(t, buyerAddr, emittedCancelledBy, "TransferCancelled event should contain buyer address as cancelledBy")

	t.Log("✅ Buyer successfully cancelled the fraudulent transfer and received refund")
}

// TestSuccessfulTransferWithCorrectContent tests a legitimate transfer with correct content
func TestSuccessfulTransferWithCorrectContent(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, address, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Create event listener
	listener, err := util.NewEventListener(client, contract, address)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAddr := minterAuth.From

	// Update authorized minter to minter address.
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Sample data for the clue
	clueContent := []byte("Find the hidden treasure in the old oak tree")
	solution := "Oak tree"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the clue content using ElGamal for the minter
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()

	// Mint a new clue with ElGamal ciphertext and r value
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Verify the clue is minted successfully
	owner, err := contract.OwnerOf(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, crypto.PubkeyToAddress(minterPrivKey.PublicKey), owner, "Minter should be the owner")

	// Set sale price
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH
	timeout := big.NewInt(180)                   // 3 minutes
	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice, timeout)
	require.NoError(t, err)
	salePriceReceipt, err := util.WaitForTransaction(client, salePriceTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), salePriceReceipt.Status, "Set sale price transaction failed")

	// Initiate purchase from buyer account
	buyerAuth.Value = big.NewInt(1000000000000000000) // 1 ETH
	buyTx, err := contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	buyReceipt, err := util.WaitForTransaction(client, buyTx)
	require.NoError(t, err)

	// Verify the TransferInitiated event is emitted
	transferInitiatedFound, err := listener.CheckEvent(buyReceipt, "TransferInitiated")
	require.NoError(t, err)
	require.True(t, transferInitiatedFound, "TransferInitiated event not found")

	// Generate transfer ID
	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// LEGITIMATE TRANSFER: Use the CORRECT original content
	// This should pass all verifications!
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		clueContent,     // CORRECT content (same as original)
		encryptedCipher, // Original on-chain cipher
		mintR,           // Public r value from mint transaction
		minterPrivKey,
		&buyerPrivKey.PublicKey,
	)
	require.NoError(t, err, "Failed to generate verifiable transfer")

	// Marshal the buyer ciphertext for on-chain storage
	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)

	// Provide proof to the contract (r value hash is extracted from proof)
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Marshal entire transfer proof (includes both DLEQ and Plaintext proofs)
	proofBytes := transfer.Proof.Marshal()

	proofTx, err := contract.ProvideProof(minterAuth, transferId, proofBytes, buyerCiphertextHash)
	require.NoError(t, err)
	proofReceipt, err := util.WaitForTransaction(client, proofTx)
	require.NoError(t, err)

	// Verify the ProofProvided event is emitted
	proofProvidedFound, err := listener.CheckEvent(proofReceipt, "ProofProvided")
	require.NoError(t, err)
	require.True(t, proofProvidedFound, "ProofProvided event not found")

	// Buyer verifies the proof off-chain
	transferData, err := contract.Transfers(&bind.CallOpts{}, transferId)
	require.NoError(t, err)

	// Unmarshal the entire transfer proof (DLEQ + Plaintext)
	proof := &zkproof.TransferProof{}
	err = proof.Unmarshal(transferData.Proof)
	require.NoError(t, err)

	// Buyer verifies both DLEQ and plaintext equality
	sellerCipher := transfer.SellerCipher
	buyerCipher := transfer.BuyerCipher

	valid := ps.VerifyElGamalTransfer(
		encryptedCipher, // Original on-chain cipher
		sellerCipher,
		buyerCipher,
		proof,
		mintR, // Public r from mint tx
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)

	// With correct content, verification should PASS!
	require.True(t, valid, "Proof verification should PASS when content is correct")

	t.Log("✅ Legitimate transfer: Plaintext equality proof verified successfully!")

	// Buyer verifies proof by calling smart contract
	buyerAuth, err = util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	verifyTx, err := contract.VerifyProof(buyerAuth, transferId)
	require.NoError(t, err)
	verifyReceipt, err := util.WaitForTransaction(client, verifyTx)
	require.NoError(t, err)

	// Verify the ProofVerified event is emitted
	proofVerifiedFound, err := listener.CheckEvent(verifyReceipt, "ProofVerified")
	require.NoError(t, err)
	require.True(t, proofVerifiedFound, "ProofVerified event not found")

	// Complete transfer with new encrypted clue and r value
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	completeTx, err := contract.CompleteTransfer(minterAuth, transferId, buyerCiphertextBytes, transfer.SharedR)
	require.NoError(t, err)
	completeReceipt, err := util.WaitForTransaction(client, completeTx)
	require.NoError(t, err)

	// Verify the TransferCompleted event is emitted
	transferCompletedFound, err := listener.CheckEvent(completeReceipt, "TransferCompleted")
	require.NoError(t, err)
	require.True(t, transferCompletedFound, "TransferCompleted event not found")

	// Verify ownership has changed
	newOwner, err := contract.OwnerOf(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, buyerAddr, newOwner, "Buyer should be the new owner")

	// Verify the clue content is updated
	newClueContents, err := contract.GetClueContents(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, buyerCiphertextBytes, newClueContents, "Clue content should be updated for buyer")

	// Verify the r value is updated on-chain
	buyerAuth, err = util.NewTransactOpts(client, buyer)
	require.NoError(t, err)
	storedRValue, err := contract.GetRValue(&bind.CallOpts{From: buyerAuth.From}, tokenId)
	require.NoError(t, err, "Buyer should be able to retrieve r value as new owner")
	require.Equal(t, transfer.SharedR, storedRValue, "Stored r value should match the revealed r")

	// Buyer can now decrypt the clue using the r value from the contract
	// Unmarshal the ciphertext from contract storage
	retrievedCipher := &zkproof.ElGamalCiphertext{}
	err = retrievedCipher.Unmarshal(newClueContents)
	require.NoError(t, err, "Should be able to unmarshal stored ciphertext")

	// Decrypt using stored r value and buyer's private key
	decryptedClueBytes, err := ps.DecryptElGamal(retrievedCipher, storedRValue, buyerPrivKey)
	require.NoError(t, err, "Buyer should be able to decrypt using on-chain r value")
	require.Equal(t, string(clueContent), string(decryptedClueBytes), "Decrypted content should match original")

	t.Log("✅ Transfer completed successfully! Buyer received correct content.")
}

// TestInvalidProofVerification tests verification of an invalid proof.
func TestInvalidProofVerification(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, address, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Create event listener
	listener, err := util.NewEventListener(client, contract, address)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAddr := minterAuth.From

	// Update authorized minter to minter address.
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Mint a clue
	clueContent := []byte("Find the hidden treasure in the forest")
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the clue content using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()

	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Set sale price
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH
	timeout := big.NewInt(180)                   // 3 minutes
	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice, timeout)
	require.NoError(t, err)
	salePriceReceipt, err := util.WaitForTransaction(client, salePriceTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), salePriceReceipt.Status, "Set sale price transaction failed")

	// Initiate purchase from buyer
	buyerAuth.Value = big.NewInt(1000000000000000000) // 1 ETH
	buyTx, err := contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, buyTx)
	require.NoError(t, err)

	// Generate transfer ID
	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// Generate valid transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		clueContent,
		encryptedCipher, // Original on-chain cipher
		mintR,           // Public r from mint
		minterPrivKey,
		&buyerPrivKey.PublicKey,
	)
	require.NoError(t, err, "Failed to generate verifiable transfer")

	// Marshal the buyer ciphertext
	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)

	// Marshal entire transfer proof and CORRUPT it
	validProof := transfer.Proof.Marshal()
	invalidProof := append([]byte{0xFF}, validProof[1:]...) // Corrupt first byte (length field)

	// Try to provide the invalid proof to the contract
	// The transaction should revert because the contract validates the proof structure
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	_, err = contract.ProvideProof(minterAuth, transferId, invalidProof, buyerCiphertextHash)
	require.Error(t, err, "Contract should reject invalid proof structure")
	require.Contains(t, err.Error(), "Invalid proof structure", "Should fail with proof structure error")

	t.Log("✅ Contract correctly rejected proof with corrupted structure")

	// The transfer should still be in ProofPending state since the invalid proof was rejected
	// We can verify the buyer can cancel since no valid proof was provided

	// Cancel the transfer
	buyerAuth, err = util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	cancelTx, err := contract.CancelTransfer(buyerAuth, transferId)
	require.NoError(t, err)
	cancelReceipt, err := util.WaitForTransaction(client, cancelTx)
	require.NoError(t, err)

	// Verify the TransferCancelled event is emitted
	transferCancelledFound, err := listener.CheckEvent(cancelReceipt, "TransferCancelled")
	require.NoError(t, err)
	require.True(t, transferCancelledFound, "TransferCancelled event not found")

	// Verify the TransferCancelled event parameters
	cancelEvents, err := listener.GetEventsByName(cancelReceipt, "TransferCancelled")
	require.NoError(t, err)
	require.Len(t, cancelEvents, 1, "Expected exactly one TransferCancelled event")

	cancelEventData := cancelEvents[0].(map[string]interface{})
	emittedTransferId := cancelEventData["transferId"].([32]byte)
	emittedCancelledBy := cancelEventData["cancelledBy"].(common.Address)

	require.Equal(t, transferId, emittedTransferId, "TransferCancelled event should contain correct transferId")
	require.Equal(t, buyerAddr, emittedCancelledBy, "TransferCancelled event should contain buyer address as cancelledBy")
}

func TestInvalidProofContentAltered(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, address, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Create event listener
	listener, err := util.NewEventListener(client, contract, address)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAddr := minterAuth.From

	// Update authorized minter to minter address.
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Mint a clue
	clueContent := []byte("Find the hidden treasure in the forest")
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the clue content using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()

	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Set sale price
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH
	timeout := big.NewInt(180)                   // 3 minutes
	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice, timeout)
	require.NoError(t, err)
	salePriceReceipt, err := util.WaitForTransaction(client, salePriceTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), salePriceReceipt.Status, "Set sale price transaction failed")

	// Initiate purchase from buyer
	buyerAuth.Value = big.NewInt(1000000000000000000) // 1 ETH
	buyTx, err := contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, buyTx)
	require.NoError(t, err)

	// Generate transfer ID
	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// ATTACK SCENARIO: Seller attempts to provide altered content
	alteredContent := []byte("Content altered!")

	// Generate transfer with ALTERED content
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		alteredContent,  // ATTACK: different content
		encryptedCipher, // Original on-chain cipher
		mintR,           // Public r from mint
		minterPrivKey,
		&buyerPrivKey.PublicKey,
	)
	require.NoError(t, err, "Failed to generate verifiable transfer")

	// Marshal the buyer ciphertext
	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)

	// Marshal entire transfer proof
	validProof := transfer.Proof.Marshal()

	// Provide the proof to the contract
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	proofTx, err := contract.ProvideProof(minterAuth, transferId, validProof, buyerCiphertextHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, proofTx)
	require.NoError(t, err)

	// Buyer verifies the proof off-chain
	transferData, err := contract.Transfers(&bind.CallOpts{}, transferId)
	require.NoError(t, err)

	// Unmarshal the entire transfer proof
	proof := &zkproof.TransferProof{}
	err = proof.Unmarshal(transferData.Proof)
	require.NoError(t, err)

	// ATTACK DETECTION: Verification should FAIL because content was altered!
	valid := ps.VerifyElGamalTransfer(
		encryptedCipher, // Original on-chain cipher
		transfer.SellerCipher,
		transfer.BuyerCipher,
		proof,
		mintR, // Public r from mint
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	require.False(t, valid, "Altered content should be detected by plaintext equality proof!")

	t.Log("✅ ATTACK PREVENTED: Altered content was detected!")

	// Cancel the transfer
	buyerAuth, err = util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	cancelTx, err := contract.CancelTransfer(buyerAuth, transferId)
	require.NoError(t, err)
	cancelReceipt, err := util.WaitForTransaction(client, cancelTx)
	require.NoError(t, err)

	// Verify the TransferCancelled event is emitted
	transferCancelledFound, err := listener.CheckEvent(cancelReceipt, "TransferCancelled")
	require.NoError(t, err)
	require.True(t, transferCancelledFound, "TransferCancelled event not found")

	// Verify the TransferCancelled event parameters
	cancelEvents, err := listener.GetEventsByName(cancelReceipt, "TransferCancelled")
	require.NoError(t, err)
	require.Len(t, cancelEvents, 1, "Expected exactly one TransferCancelled event")

	cancelEventData := cancelEvents[0].(map[string]interface{})
	emittedTransferId := cancelEventData["transferId"].([32]byte)
	emittedCancelledBy := cancelEventData["cancelledBy"].(common.Address)

	require.Equal(t, transferId, emittedTransferId, "TransferCancelled event should contain correct transferId")
	require.Equal(t, buyerAddr, emittedCancelledBy, "TransferCancelled event should contain buyer address as cancelledBy")
}

// TestCompletingTransferWithoutVerification tests completing a transfer without verification.
func TestCompletingTransferWithoutVerification(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAddr := minterAuth.From

	// Update authorized minter to minter address.
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Mint a clue
	clueContent := []byte("Find the hidden treasure in the forest")
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the clue content using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()

	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Set sale price
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH
	timeout := big.NewInt(180)                   // 3 minutes
	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice, timeout)
	require.NoError(t, err)
	salePriceReceipt, err := util.WaitForTransaction(client, salePriceTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), salePriceReceipt.Status, "Set sale price transaction failed")

	// Initiate purchase from buyer
	buyerAuth.Value = big.NewInt(1000000000000000000) // 1 ETH
	buyTx, err := contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, buyTx)
	require.NoError(t, err)

	// Generate transfer ID
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
	require.NoError(t, err, "Failed to generate verifiable transfer")

	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)
	proofBytes := transfer.Proof.Marshal()

	// Provide proof to the contract
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	proofTx, err := contract.ProvideProof(minterAuth, transferId, proofBytes, buyerCiphertextHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, proofTx)
	require.NoError(t, err)

	// Skip the verification step
	// Attempt to complete transfer without verification
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAuth.GasLimit = 300000 // Higher gas limit for failing transaction

	_, err = contract.CompleteTransfer(minterAuth, transferId, buyerCiphertextBytes, transfer.SharedR)
	require.Error(t, err, "Transaction should fail")
}

// TestCancelTransfer tests cancelling a transfer.
func TestCancelTransfer(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, address, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Create event listener
	listener, err := util.NewEventListener(client, contract, address)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAddr := minterAuth.From

	// Update authorized minter to minter address.
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Get initial buyer balance
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	initialBuyerBalance, err := client.BalanceAt(ctx, buyerAddr, nil)
	require.NoError(t, err)

	// Mint a clue
	clueContent := []byte("Find the hidden treasure in the forest")
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the clue content using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()

	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Set sale price
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH
	timeout := big.NewInt(180)                   // 3 minutes
	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice, timeout)
	require.NoError(t, err)
	salePriceReceipt, err := util.WaitForTransaction(client, salePriceTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), salePriceReceipt.Status, "Set sale price transaction failed")

	// Initiate purchase from buyer with 1 ETH
	paymentAmount := big.NewInt(1000000000000000000) // 1 ETH
	buyerAuth.Value = paymentAmount
	buyTx, err := contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, buyTx)
	require.NoError(t, err)

	// Generate transfer ID
	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// Cancel the transfer
	buyerAuth, err = util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	cancelTx, err := contract.CancelTransfer(buyerAuth, transferId)
	require.NoError(t, err)
	cancelReceipt, err := util.WaitForTransaction(client, cancelTx)
	require.NoError(t, err)

	// Verify the TransferCancelled event is emitted
	transferCancelledFound, err := listener.CheckEvent(cancelReceipt, "TransferCancelled")
	require.NoError(t, err)
	require.True(t, transferCancelledFound, "TransferCancelled event not found")

	// Verify the TransferCancelled event parameters
	cancelEvents, err := listener.GetEventsByName(cancelReceipt, "TransferCancelled")
	require.NoError(t, err)
	require.Len(t, cancelEvents, 1, "Expected exactly one TransferCancelled event")

	cancelEventData := cancelEvents[0].(map[string]interface{})
	emittedTransferId := cancelEventData["transferId"].([32]byte)
	emittedCancelledBy := cancelEventData["cancelledBy"].(common.Address)

	require.Equal(t, transferId, emittedTransferId, "TransferCancelled event should contain correct transferId")
	require.Equal(t, buyerAddr, emittedCancelledBy, "TransferCancelled event should contain buyer address as cancelledBy")

	// Verify buyer received refund
	transferData, err := contract.Transfers(nil, transferId)
	require.NoError(t, err)
	require.Equal(t, common.Address{}, transferData.Buyer, "Transfer should be deleted after cancellation")
	require.Equal(t, big.NewInt(0).String(), transferData.Value.String(), "Transfer value should be 0 after cancellation")

	// Final buyer balance should be close to initial balance minus gas costs
	finalBuyerBalance, err := client.BalanceAt(ctx, buyerAddr, nil)
	require.NoError(t, err)

	// Calculate difference (should be just gas costs)
	balanceDiff := new(big.Int).Sub(initialBuyerBalance, finalBuyerBalance)

	// Check that the difference is less than 0.1 ETH (meaning the 1 ETH was refunded, with just gas costs deducted)
	maxGasCost := big.NewInt(100000000000000000) // 0.1 ETH
	require.True(t, balanceDiff.Cmp(maxGasCost) < 0, "Balance difference should be small (just gas costs)")
}

// TestCorruptedRValueRejected tests that seller cannot complete transfer with wrong r value.
// This demonstrates that the rHash commitment prevents seller from providing fake r.
func TestCorruptedRValueRejected(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAddr := minterAuth.From

	// Update authorized minter
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, buyer)
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
	encryptedClueContent := encryptedCipher.Marshal()

	// Mint clue
	minterAuth, err = util.NewTransactOpts(client, minter)
	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Set sale price
	minterAuth, err = util.NewTransactOpts(client, minter)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH
	timeout := big.NewInt(180)                   // 3 minutes
	tx, err = contract.SetSalePrice(minterAuth, tokenId, salePrice, timeout)
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

	// Generate verifiable transfer with REAL r
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		clueContent,
		encryptedCipher, // Original on-chain cipher
		mintR,           // Public r from mint
		minterPrivKey,
		&buyerPrivKey.PublicKey,
	)
	require.NoError(t, err, "Failed to generate verifiable transfer")

	// Marshal the buyer ciphertext
	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)

	// Provide proof to the contract (includes Hash(r_real) in DLEQ proof)
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	proofBytes := transfer.Proof.Marshal()
	proofTx, err := contract.ProvideProof(minterAuth, transferId, proofBytes, buyerCiphertextHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, proofTx)
	require.NoError(t, err)

	// Buyer verifies the proof off-chain
	transferData, err := contract.Transfers(&bind.CallOpts{}, transferId)
	require.NoError(t, err)

	// Unmarshal the entire transfer proof
	proof := &zkproof.TransferProof{}
	err = proof.Unmarshal(transferData.Proof)
	require.NoError(t, err)

	// For the buyer to verify, they need:
	// - Original on-chain cipher
	// - New seller/buyer ciphers
	// - Complete proof (DLEQ + Plaintext)
	// - mintR
	sellerCipher := transfer.SellerCipher
	buyerCipher := transfer.BuyerCipher

	// Verify both DLEQ and plaintext equality proofs
	valid := ps.VerifyElGamalTransfer(
		encryptedCipher, // Original on-chain cipher
		sellerCipher,
		buyerCipher,
		proof,
		mintR, // Public r from mint
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	require.True(t, valid, "Proof verification should succeed with correct content")

	// Buyer calls verifyProof on contract
	buyerAuth, err = util.NewTransactOpts(client, buyer)
	verifyTx, err := contract.VerifyProof(buyerAuth, transferId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, verifyTx)
	require.NoError(t, err)

	// ATTACK: Seller tries to complete transfer with CORRUPTED r value
	corruptedR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err)
	require.NotEqual(t, transfer.SharedR, corruptedR, "Corrupted r should be different from real r")

	// Attempt to complete transfer with corrupted r
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAuth.GasLimit = 500000 // Higher gas limit for failing transaction

	_, err = contract.CompleteTransfer(minterAuth, transferId, buyerCiphertextBytes, corruptedR)

	// ATTACK PREVENTED: Transaction should fail because Hash(corruptedR) != rValueHash
	// The contract stored rValueHash = Hash(r_real) extracted from the DLEQ proof
	require.Error(t, err, "CompleteTransfer should fail with corrupted r value")
	require.Contains(t, err.Error(), "R value hash mismatch", "Error should indicate r value hash mismatch")

	t.Log("✅ Contract prevents seller from completing transfer with corrupted r value")
	t.Log("   The rHash commitment (embedded in DLEQ proof) ensures seller cannot change r")
}

// Helper function to get the last minted token ID
func getLastMintedTokenID(contract *bindings.Skavenge) (*big.Int, error) {
	tokenId, err := contract.GetCurrentTokenId(nil)
	if err != nil {
		return nil, err
	}
	// Subtract 1 because the counter has been incremented after minting
	return new(big.Int).Sub(tokenId, big.NewInt(1)), nil
}
