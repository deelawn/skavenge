// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/deelawn/skavenge/tests/util"
)

var buyer string = "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"

// TestSuccessfulSolve tests the successful solving of a clue.
func TestSuccessfulSolve(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial("http://localhost:8545")
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

	// Create API client for encryption
	apiClient, err := util.NewGRPCClient()
	require.NoError(t, err)

	// Mint a clue with a known solution
	clueContent := "Find the hidden treasure in the forest"
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Encrypt the clue content
	encryptedClueContent, err := apiClient.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Mint the clue
	tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "Mint transaction failed")

	// Verify the clue is not solved yet
	tokenId := big.NewInt(1)
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.False(t, clueData.IsSolved, "Clue should not be solved yet")

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

	// Verify the clue is now marked as solved
	clueData, err = contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.True(t, clueData.IsSolved, "Clue should be marked as solved")
}

// TestFailedSolveAttempt tests a failed attempt to solve a clue.
func TestFailedSolveAttempt(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial("http://localhost:8545")
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

	// Create API client for encryption
	apiClient, err := util.NewGRPCClient()
	require.NoError(t, err)

	// Mint a clue with a known solution
	clueContent := "Find the hidden treasure in the forest"
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Encrypt the clue content
	encryptedClueContent, err := apiClient.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Mint the clue
	tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "Mint transaction failed")

	// Verify initial solve attempts
	tokenId := big.NewInt(1)
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, uint64(0), clueData.SolveAttempts.Uint64(), "Initial solve attempts should be 0")

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

	// Verify the ClueAttempted event was emitted
	clueAttemptedFound, err := listener.CheckEvent(solveReceipt, "ClueAttempted")
	require.NoError(t, err)
	require.True(t, clueAttemptedFound, "ClueAttempted event not found")

	// Verify the clue is still not solved
	clueData, err = contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.False(t, clueData.IsSolved, "Clue should still not be solved")

	// Verify solve attempts was incremented
	require.Equal(t, uint64(1), clueData.SolveAttempts.Uint64(), "Solve attempts should be incremented to 1")
}

// TestMaximumSolveAttempts tests reaching the maximum number of solve attempts.
func TestMaximumSolveAttempts(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial("http://localhost:8545")
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

	// Create API client for encryption
	apiClient, err := util.NewGRPCClient()
	require.NoError(t, err)

	// Mint a clue with a known solution
	clueContent := "Find the hidden treasure in the forest"
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Encrypt the clue content
	encryptedClueContent, err := apiClient.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Mint the clue
	tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "Mint transaction failed")

	tokenId := big.NewInt(1)

	// Make three incorrect solution attempts
	incorrectSolutions := []string{
		"In the cave",
		"Under the bridge",
		"Behind the tree",
	}

	for i, incorrectSolution := range incorrectSolutions {
		// Update auth to ensure correct nonce
		minterAuth, err = util.NewTransactOpts(client, minter)
		require.NoError(t, err)

		solveTx, err := contract.AttemptSolution(minterAuth, tokenId, incorrectSolution)
		require.NoError(t, err)

		// Wait for the transaction to be mined
		solveReceipt, err := util.WaitForTransaction(client, solveTx)
		require.NoError(t, err)
		require.Equal(t, uint64(1), solveReceipt.Status, "Solution attempt transaction failed")

		// Verify the ClueAttempted event was emitted
		clueAttemptedFound, err := listener.CheckEvent(solveReceipt, "ClueAttempted")
		require.NoError(t, err)
		require.True(t, clueAttemptedFound, "ClueAttempted event not found")

		// Verify solve attempts was incremented
		clueData, err := contract.Clues(nil, tokenId)
		require.NoError(t, err)
		require.Equal(t, uint64(i+1), clueData.SolveAttempts.Uint64(),
			"Solve attempts should be incremented to %d", i+1)
	}

	// Attempt a 4th solution (should fail)
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// We need to set the gas limit higher as the transaction will revert
	minterAuth.GasLimit = 300000

	// Try one more attempt which should fail with "No attempts remaining"
	_, err = contract.AttemptSolution(minterAuth, tokenId, "In the hollow tree")
	require.Error(t, err, "Transaction should be sent but will revert")

	// TODO: take another look at why this doesn't work.

	// // Wait for the transaction to be mined
	// fourthSolveReceipt, err := bind.WaitMined(context.Background(), client, fourthSolveTx)
	// require.NoError(t, err)

	// // Transaction should fail (status 0)
	// require.Equal(t, uint64(0), fourthSolveReceipt.Status, "Fourth solution attempt should fail")

	// Verify solve attempts is still 3
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, uint64(3), clueData.SolveAttempts.Uint64(), "Solve attempts should still be 3")
}

// TestNonOwnerSolveAttempt tests a non-owner attempting to solve a clue.
func TestNonOwnerSolveAttempt(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Setup minter account
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Setup non-owner account
	nonOwnerAuth, err := util.NewTransactOpts(client, buyer) // Using buyer as non-owner
	require.NoError(t, err)

	// Create API client for encryption
	apiClient, err := util.NewGRPCClient()
	require.NoError(t, err)

	// Mint a clue with a known solution
	clueContent := "Find the hidden treasure in the forest"
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Encrypt the clue content
	encryptedClueContent, err := apiClient.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Mint the clue
	tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "Mint transaction failed")

	tokenId := big.NewInt(1)

	// Set higher gas limit for the failing transaction
	nonOwnerAuth.GasLimit = 300000

	// Non-owner attempts to solve (should fail)
	_, err = contract.AttemptSolution(nonOwnerAuth, tokenId, solution)
	require.Error(t, err, "Transaction should be sent but will revert")

	// // Wait for the transaction to be mined
	// solveReceipt, err := bind.WaitMined(context.Background(), client, solveTx)
	// require.NoError(t, err)

	// // Transaction should fail (status 0)
	// require.Equal(t, uint64(0), solveReceipt.Status, "Non-owner solution attempt should fail")

	// Verify clue is still not solved
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.False(t, clueData.IsSolved, "Clue should not be solved")

	// Verify solve attempts is still 0
	require.Equal(t, uint64(0), clueData.SolveAttempts.Uint64(), "Solve attempts should still be 0")
}

// TestSolveAlreadySolvedClue tests attempting to solve an already solved clue.
func TestSolveAlreadySolvedClue(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Setup minter account
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Create API client for encryption
	apiClient, err := util.NewGRPCClient()
	require.NoError(t, err)

	// Mint a clue with a known solution
	clueContent := "Find the hidden treasure in the forest"
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Encrypt the clue content
	encryptedClueContent, err := apiClient.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Mint the clue
	tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "Mint transaction failed")

	tokenId := big.NewInt(1)

	// Solve the clue first time
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	solveTx, err := contract.AttemptSolution(minterAuth, tokenId, solution)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	solveReceipt, err := util.WaitForTransaction(client, solveTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), solveReceipt.Status, "First solution attempt should succeed")

	// Verify the clue is now solved
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.True(t, clueData.IsSolved, "Clue should be marked as solved")

	// Try to solve again with the same solution
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAuth.GasLimit = 300000 // Set higher gas limit for failing transaction

	_, err = contract.AttemptSolution(minterAuth, tokenId, solution)
	require.Error(t, err, "Transaction should be sent but will revert")

	// // Wait for the transaction to be mined
	// secondSolveReceipt, err := bind.WaitMined(context.Background(), client, secondSolveTx)
	// require.NoError(t, err)

	// // Transaction should fail (status 0)
	// require.Equal(t, uint64(0), secondSolveReceipt.Status, "Second solution attempt should fail")
}
