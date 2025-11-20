// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"

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

// TestSuccessfulTransfer tests the successful transfer of a clue using ElGamal encryption.
func TestSuccessfulTransfer(t *testing.T) {
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
	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice)
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

	// Generate ElGamal verifiable transfer
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		clueContent,
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

	// Marshal DLEQ proof (includes rHash in last 32 bytes)
	proofBytes := transfer.DLEQProof.Marshal()

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

	// Unmarshal and verify DLEQ proof
	dleqProof := &zkproof.DLEQProof{}
	err = dleqProof.Unmarshal(transferData.Proof)
	require.NoError(t, err)

	// For the buyer to verify, they need both ciphertexts from the transfer
	// Both come from GenerateVerifiableElGamalTransfer with same plaintext
	sellerCipher := transfer.SellerCipher
	buyerCipher := transfer.BuyerCipher

	valid := ps.VerifyElGamalTransfer(
		sellerCipher,
		buyerCipher,
		dleqProof,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	require.True(t, valid, "DLEQ proof verification failed")

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
	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice)
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
		minterPrivKey,
		&buyerPrivKey.PublicKey,
	)
	require.NoError(t, err, "Failed to generate verifiable transfer")

	// Marshal the buyer ciphertext
	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)

	// Marshal proof and CORRUPT it
	validProof := transfer.DLEQProof.Marshal()
	invalidProof := append([]byte{0xFF}, validProof[1:]...) // Corrupt first byte

	// Provide the invalid proof to the contract
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	proofTx, err := contract.ProvideProof(minterAuth, transferId, invalidProof, buyerCiphertextHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, proofTx)
	require.NoError(t, err)

	// Buyer verifies the proof off-chain
	transferData, err := contract.Transfers(&bind.CallOpts{}, transferId)
	require.NoError(t, err)

	// Try to unmarshal and verify - should fail
	dleqProof := &zkproof.DLEQProof{}
	err = dleqProof.Unmarshal(transferData.Proof)
	// Unmarshaling might fail or succeed depending on corruption
	// If it succeeds, verification should fail
	if err == nil {
		sellerCipher := &zkproof.ElGamalCiphertext{}
		err = sellerCipher.Unmarshal(encryptedClueContent)
		require.NoError(t, err)

		valid := ps.VerifyElGamalTransfer(
			sellerCipher,
			transfer.BuyerCipher,
			dleqProof,
			transfer.SellerPubKey,
			transfer.BuyerPubKey,
		)
		require.False(t, valid, "Invalid proof should not verify")
	}

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
	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice)
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
		minterPrivKey,
		&buyerPrivKey.PublicKey,
	)
	require.NoError(t, err, "Failed to generate verifiable transfer")

	buyerCiphertextBytes := transfer.BuyerCipher.Marshal()
	buyerCiphertextHash := crypto.Keccak256Hash(buyerCiphertextBytes)
	proofBytes := transfer.DLEQProof.Marshal()

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
	initialBuyerBalance, err := client.BalanceAt(context.Background(), buyerAddr, nil)
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
	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice)
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

	// Verify buyer received refund
	transferData, err := contract.Transfers(nil, transferId)
	require.NoError(t, err)
	require.Equal(t, common.Address{}, transferData.Buyer, "Transfer should be deleted after cancellation")
	require.Equal(t, big.NewInt(0).String(), transferData.Value.String(), "Transfer value should be 0 after cancellation")

	// Final buyer balance should be close to initial balance minus gas costs
	finalBuyerBalance, err := client.BalanceAt(context.Background(), buyerAddr, nil)
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

	// Generate verifiable transfer with REAL r
	transfer, err := ps.GenerateVerifiableElGamalTransfer(
		clueContent,
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

	proofBytes := transfer.DLEQProof.Marshal()
	proofTx, err := contract.ProvideProof(minterAuth, transferId, proofBytes, buyerCiphertextHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, proofTx)
	require.NoError(t, err)

	// Buyer verifies the proof off-chain
	transferData, err := contract.Transfers(&bind.CallOpts{}, transferId)
	require.NoError(t, err)

	// Unmarshal and verify DLEQ proof
	dleqProof := &zkproof.DLEQProof{}
	err = dleqProof.Unmarshal(transferData.Proof)
	require.NoError(t, err)

	// For the buyer to verify, they need both ciphertexts from the transfer
	sellerCipher := transfer.SellerCipher
	buyerCipher := transfer.BuyerCipher

	// Verify the DLEQ proof
	valid := ps.VerifyElGamalTransfer(
		sellerCipher,
		buyerCipher,
		dleqProof,
		transfer.SellerPubKey,
		transfer.BuyerPubKey,
	)
	require.True(t, valid, "DLEQ proof verification should succeed")

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
