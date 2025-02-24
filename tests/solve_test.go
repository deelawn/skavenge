// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/deelawn/skavenge/tests/util"
)

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

	// TODO: Mint a new clue with a known solution
	// TODO: Successfully solve the clue
	// TODO: Verify the clue is marked as solved
	// TODO: Verify the ClueSolved event is emitted
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

	// TODO: Mint a new clue
	// TODO: Attempt to solve with incorrect solution
	// TODO: Verify the clue is not marked as solved
	// TODO: Verify the ClueAttempted event is emitted with the correct remaining attempts
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

	// TODO: Mint a new clue
	// TODO: Make 3 failed solve attempts
	// TODO: Verify the 4th attempt fails with the "No attempts remaining" error
	// TODO: Verify the appropriate events are emitted
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
	contract, address, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// TODO: Mint a new clue
	// TODO: Attempt to solve from a non-owner account
	// TODO: Verify the appropriate error is returned
}
