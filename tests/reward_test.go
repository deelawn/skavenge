// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"context"
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

// TestMintClueWithReward tests minting a clue with an ETH reward.
func TestMintClueWithReward(t *testing.T) {
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
	minterAddr := crypto.PubkeyToAddress(minterPrivKey.PublicKey)

	// Update authorized minter to minter address
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
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

	// Encrypt the clue content using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()
	require.NotEmpty(t, encryptedClueContent, "Encrypted clue content should not be empty")

	// Set reward value (0.5 ETH)
	rewardValue := big.NewInt(500000000000000000)

	// Mint a new clue with the encrypted content and reward
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAuth.Value = rewardValue

	pointValue := uint8(3)
	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR, pointValue)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Check that the solve reward is correctly stored
	storedReward, err := contract.GetSolveReward(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, rewardValue.String(), storedReward.String(), "Solve reward should match the sent value")

	// Verify the clue struct in the mapping
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, rewardValue.String(), clueData.SolveReward.String(), "Solve reward in clue data should match")
	require.False(t, clueData.IsSolved, "Clue should not be marked as solved")
}

// TestMintClueWithoutReward tests minting a clue without an ETH reward (zero value).
func TestMintClueWithoutReward(t *testing.T) {
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
	minterAddr := crypto.PubkeyToAddress(minterPrivKey.PublicKey)

	// Update authorized minter to minter address
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Sample data for the clue
	clueContent := []byte("Find the hidden treasure")
	solution := "Treasure"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the clue content using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()

	// Mint a new clue without a reward (no Value set)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	pointValue := uint8(2)
	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR, pointValue)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Check that the solve reward is zero
	storedReward, err := contract.GetSolveReward(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(0).String(), storedReward.String(), "Solve reward should be zero")

	// Verify the clue struct in the mapping
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(0).String(), clueData.SolveReward.String(), "Solve reward should be zero")
}

// TestSolveClueWithReward tests solving a clue with an ETH reward and verifying the reward is dispersed.
func TestSolveClueWithReward(t *testing.T) {
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
	minterAddr := crypto.PubkeyToAddress(minterPrivKey.PublicKey)

	// Update authorized minter to minter address
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
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

	// Encrypt the clue content using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()

	// Set reward value (1 ETH)
	rewardValue := big.NewInt(1000000000000000000)

	// Mint a new clue with the encrypted content and reward
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAuth.Value = rewardValue

	pointValue := uint8(4)
	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR, pointValue)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Verify the reward is stored
	storedReward, err := contract.GetSolveReward(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, rewardValue.String(), storedReward.String(), "Solve reward should match the sent value")

	// Get the minter's balance before solving
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	balanceBefore, err := client.BalanceAt(ctx, minterAddr, nil)
	require.NoError(t, err)

	// Solve the clue
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	solveTx, err := contract.AttemptSolution(minterAuth, tokenId, solution)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	solveReceipt, err := util.WaitForTransaction(client, solveTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), solveReceipt.Status, "Solution attempt should succeed")

	// Calculate gas cost
	gasCost := new(big.Int).Mul(big.NewInt(int64(solveReceipt.GasUsed)), solveReceipt.EffectiveGasPrice)

	// Get the minter's balance after solving
	balanceAfter, err := client.BalanceAt(ctx, minterAddr, nil)
	require.NoError(t, err)

	// Calculate expected balance: balanceBefore + rewardValue - gasCost
	expectedBalance := new(big.Int).Sub(balanceBefore, gasCost)
	expectedBalance.Add(expectedBalance, rewardValue)

	require.Equal(t, expectedBalance.String(), balanceAfter.String(), "Solver should receive the reward minus gas costs")

	// Verify the clue is marked as solved
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.True(t, clueData.IsSolved, "Clue should be marked as solved")

	// Verify the reward has been cleared
	storedRewardAfter, err := contract.GetSolveReward(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(0).String(), storedRewardAfter.String(), "Solve reward should be cleared after solving")
}

// TestSolveClueWithoutReward tests solving a clue without a reward.
func TestSolveClueWithoutReward(t *testing.T) {
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
	minterAddr := crypto.PubkeyToAddress(minterPrivKey.PublicKey)

	// Update authorized minter to minter address
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Sample data for the clue
	clueContent := []byte("Find the hidden treasure")
	solution := "Treasure"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the clue content using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()

	// Mint a new clue without a reward
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	pointValue := uint8(1)
	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR, pointValue)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Verify the reward is zero
	storedReward, err := contract.GetSolveReward(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(0).String(), storedReward.String(), "Solve reward should be zero")

	// Get the minter's balance before solving
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	balanceBefore, err := client.BalanceAt(ctx, minterAddr, nil)
	require.NoError(t, err)

	// Solve the clue
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	solveTx, err := contract.AttemptSolution(minterAuth, tokenId, solution)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	solveReceipt, err := util.WaitForTransaction(client, solveTx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), solveReceipt.Status, "Solution attempt should succeed")

	// Calculate gas cost
	gasCost := new(big.Int).Mul(big.NewInt(int64(solveReceipt.GasUsed)), solveReceipt.EffectiveGasPrice)

	// Get the minter's balance after solving
	balanceAfter, err := client.BalanceAt(ctx, minterAddr, nil)
	require.NoError(t, err)

	// Calculate expected balance: balanceBefore - gasCost (no reward)
	expectedBalance := new(big.Int).Sub(balanceBefore, gasCost)

	require.Equal(t, expectedBalance.String(), balanceAfter.String(), "Solver balance should only decrease by gas costs")

	// Verify the clue is marked as solved
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.True(t, clueData.IsSolved, "Clue should be marked as solved")
}

// TestMultipleCluesWithDifferentRewards tests minting and solving multiple clues with different reward values.
func TestMultipleCluesWithDifferentRewards(t *testing.T) {
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
	minterAddr := crypto.PubkeyToAddress(minterPrivKey.PublicKey)

	// Update authorized minter to minter address
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Define test cases with different reward values
	testCases := []struct {
		name        string
		clueContent []byte
		solution    string
		rewardValue *big.Int
		pointValue  uint8
	}{
		{
			name:        "No reward",
			clueContent: []byte("First clue"),
			solution:    "First solution",
			rewardValue: big.NewInt(0),
			pointValue:  1,
		},
		{
			name:        "Small reward",
			clueContent: []byte("Second clue"),
			solution:    "Second solution",
			rewardValue: big.NewInt(100000000000000000), // 0.1 ETH
			pointValue:  2,
		},
		{
			name:        "Medium reward",
			clueContent: []byte("Third clue"),
			solution:    "Third solution",
			rewardValue: big.NewInt(500000000000000000), // 0.5 ETH
			pointValue:  3,
		},
		{
			name:        "Large reward",
			clueContent: []byte("Fourth clue"),
			solution:    "Fourth solution",
			rewardValue: big.NewInt(2000000000000000000), // 2 ETH
			pointValue:  5,
		},
	}

	tokenIds := make([]*big.Int, len(testCases))

	// Mint all clues
	for i, tc := range testCases {
		t.Run("Mint_"+tc.name, func(t *testing.T) {
			// Generate solution hash
			solutionHash := crypto.Keccak256Hash([]byte(tc.solution))

			// Generate random r value for ElGamal encryption
			mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
			require.NoError(t, err, "Failed to generate r value")

			// Encrypt the clue content using ElGamal
			encryptedCipher, err := ps.EncryptElGamal(tc.clueContent, &minterPrivKey.PublicKey, mintR)
			require.NoError(t, err, "Failed to encrypt clue content")

			// Marshal to bytes for on-chain storage
			encryptedClueContent := encryptedCipher.Marshal()

			// Mint the clue
			minterAuth, err := util.NewTransactOpts(client, minter)
			require.NoError(t, err)
			minterAuth.Value = tc.rewardValue

			tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR, tc.pointValue)
			require.NoError(t, err)

			// Wait for the transaction to be mined
			_, err = util.WaitForTransaction(client, tx)
			require.NoError(t, err)

			tokenId, err := getLastMintedTokenID(contract)
			require.NoError(t, err)
			tokenIds[i] = tokenId

			// Verify the reward is stored correctly
			storedReward, err := contract.GetSolveReward(nil, tokenId)
			require.NoError(t, err)
			require.Equal(t, tc.rewardValue.String(), storedReward.String(), "Solve reward should match")
		})
	}

	// Solve all clues and verify rewards
	for i, tc := range testCases {
		t.Run("Solve_"+tc.name, func(t *testing.T) {
			tokenId := tokenIds[i]

			// Get the minter's balance before solving
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			balanceBefore, err := client.BalanceAt(ctx, minterAddr, nil)
			require.NoError(t, err)

			// Solve the clue
			minterAuth, err := util.NewTransactOpts(client, minter)
			require.NoError(t, err)

			solveTx, err := contract.AttemptSolution(minterAuth, tokenId, tc.solution)
			require.NoError(t, err)

			// Wait for the transaction to be mined
			solveReceipt, err := util.WaitForTransaction(client, solveTx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), solveReceipt.Status, "Solution attempt should succeed")

			// Calculate gas cost
			gasCost := new(big.Int).Mul(big.NewInt(int64(solveReceipt.GasUsed)), solveReceipt.EffectiveGasPrice)

			// Get the minter's balance after solving
			balanceAfter, err := client.BalanceAt(ctx, minterAddr, nil)
			require.NoError(t, err)

			// Calculate expected balance: balanceBefore + rewardValue - gasCost
			expectedBalance := new(big.Int).Sub(balanceBefore, gasCost)
			expectedBalance.Add(expectedBalance, tc.rewardValue)

			require.Equal(t, expectedBalance.String(), balanceAfter.String(), "Solver should receive the correct reward")

			// Verify the clue is marked as solved
			clueData, err := contract.Clues(nil, tokenId)
			require.NoError(t, err)
			require.True(t, clueData.IsSolved, "Clue should be marked as solved")

			// Verify the reward has been cleared
			storedRewardAfter, err := contract.GetSolveReward(nil, tokenId)
			require.NoError(t, err)
			require.Equal(t, big.NewInt(0).String(), storedRewardAfter.String(), "Solve reward should be cleared")
		})
	}
}
