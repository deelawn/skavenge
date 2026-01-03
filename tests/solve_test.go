// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"crypto/rand"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/deelawn/skavenge/tests/util"
	"github.com/deelawn/skavenge/zkproof"
)

// TestSuccessfulSolve tests the successful solving of a clue.
func TestSuccessfulSolve(t *testing.T) {
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

	// Setup minter account
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

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Mint a clue with a known solution
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

	// Mint the clue using deployer as authorized minter
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "Mint transaction failed")

	// Verify the clue is not solved yet
	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.False(t, clueData.IsSolved, "Clue should not be solved yet")

	// Set sale price for the clue
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH

	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice)
	require.NoError(t, err)
	salePriceReceipt, err := util.WaitForTransaction(client, salePriceTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), salePriceReceipt.Status, "Set sale price transaction failed")

	// Verify sale price was set
	clueData, err = contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, salePrice.String(), clueData.SalePrice.String(), "Sale price should be set to 1 ETH")

	// Verify the clue is marked for sale
	isForSale, err := contract.CluesForSale(nil, tokenId)
	require.NoError(t, err)
	require.True(t, isForSale, "Clue should be marked for sale")

	// Solve the clue with the correct solution
	// Update auth to ensure correct nonce
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	solveTx, err := contract.AttemptSolution(minterAuth, tokenId, solution)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	solveReceipt, err := util.WaitForTransaction(client, solveTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), solveReceipt.Status, "Solution attempt transaction failed")

	time.Sleep(time.Second * 2)

	// Verify the ClueSolved event was emitted
	clueSolvedFound, err := listener.CheckEvent(solveReceipt, "ClueSolved")
	require.NoError(t, err)
	require.True(t, clueSolvedFound, "ClueSolved event not found")

	// Verify the SalePriceRemoved event was emitted
	salePriceRemovedFound, err := listener.CheckEvent(solveReceipt, "SalePriceRemoved")
	require.NoError(t, err)
	require.True(t, salePriceRemovedFound, "SalePriceRemoved event not found")

	// Verify the clue is now marked as solved
	clueData, err = contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.True(t, clueData.IsSolved, "Clue should be marked as solved")

	// Verify the sale price was reset to 0
	require.Equal(t, big.NewInt(0).String(), clueData.SalePrice.String(), "Sale price should be reset to 0")

	// Verify the clue is no longer for sale
	isForSale, err = contract.CluesForSale(nil, tokenId)
	require.NoError(t, err)
	require.False(t, isForSale, "Clue should no longer be marked for sale")
}

// TestFailedSolveAttempt tests a failed attempt to solve a clue.
func TestFailedSolveAttempt(t *testing.T) {
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

	// Setup minter account
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

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Mint a clue with a known solution
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

	// Mint the clue using deployer as authorized minter
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "Mint transaction failed")

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Set sale price for the clue
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH

	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice)
	require.NoError(t, err)
	salePriceReceipt, err := util.WaitForTransaction(client, salePriceTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), salePriceReceipt.Status, "Set sale price transaction failed")

	// Attempt to solve with incorrect solution
	incorrectSolution := "In the cave"

	// Update auth to ensure correct nonce
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	solveTx, err := contract.AttemptSolution(minterAuth, tokenId, incorrectSolution)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	solveReceipt, err := util.WaitForTransaction(client, solveTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), solveReceipt.Status, "Solution attempt transaction failed")

	// Verify the clue is still not solved
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.False(t, clueData.IsSolved, "Clue should still not be solved")

	// Verify the sale price is still set
	require.Equal(t, salePrice.String(), clueData.SalePrice.String(), "Sale price should still be set")

	// Verify the clue is still for sale
	isForSale, err := contract.CluesForSale(nil, tokenId)
	require.NoError(t, err)
	require.True(t, isForSale, "Clue should still be marked for sale")
}

// TestSetSalePriceOnSolvedClue tests attempting to set a sale price on a solved clue.
func TestSetSalePriceOnSolvedClue(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Create event listener
	// listener, err := util.NewEventListener(client, contract, address)
	// require.NoError(t, err)

	// Setup minter account
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

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Mint a clue with a known solution
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

	// Mint the clue using deployer as authorized minter
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "Mint transaction failed")

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Solve the clue first
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	solveTx, err := contract.AttemptSolution(minterAuth, tokenId, solution)
	require.NoError(t, err)
	solveReceipt, err := util.WaitForTransaction(client, solveTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), solveReceipt.Status, "Solution attempt should succeed")

	// Verify the clue is solved
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.True(t, clueData.IsSolved, "Clue should be marked as solved")

	// Now try to set a sale price on the solved clue
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAuth.GasLimit = 300000 // Set higher gas limit for failing transaction

	salePrice := big.NewInt(1000000000000000000) // 1 ETH

	// This should fail because the clue is solved
	_, err = contract.SetSalePrice(minterAuth, tokenId, salePrice)
	require.Error(t, err, "Setting sale price on solved clue should fail")
}

// TestRemoveSalePrice tests the function to remove a sale price.
func TestRemoveSalePrice(t *testing.T) {
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

	// Setup minter account
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

	// Mint the clue
	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "Mint transaction failed")

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Set sale price for the clue
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	salePrice := big.NewInt(1000000000000000000) // 1 ETH

	salePriceTx, err := contract.SetSalePrice(minterAuth, tokenId, salePrice)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, salePriceTx)
	require.NoError(t, err)

	// Verify the sale price is set
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, salePrice.String(), clueData.SalePrice.String(), "Sale price should be set to 1 ETH")

	// Verify the clue is marked for sale
	isForSale, err := contract.CluesForSale(nil, tokenId)
	require.NoError(t, err)
	require.True(t, isForSale, "Clue should be marked for sale")

	// Now remove the sale price
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	removeTx, err := contract.RemoveSalePrice(minterAuth, tokenId)
	require.NoError(t, err)
	removeReceipt, err := util.WaitForTransaction(client, removeTx)
	require.NoError(t, err)

	// Verify the SalePriceRemoved event was emitted
	salePriceRemovedFound, err := listener.CheckEvent(removeReceipt, "SalePriceRemoved")
	require.NoError(t, err)
	require.True(t, salePriceRemovedFound, "SalePriceRemoved event not found")

	// Verify the sale price was reset to 0
	clueData, err = contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(0).String(), clueData.SalePrice.String(), "Sale price should be reset to 0")

	// Verify the clue is no longer for sale
	isForSale, err = contract.CluesForSale(nil, tokenId)
	require.NoError(t, err)
	require.False(t, isForSale, "Clue should no longer be marked for sale")
}
