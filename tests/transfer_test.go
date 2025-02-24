// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/deelawn/skavenge/tests/util"
)

var (
	buyer = "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
	other = "0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"
)

// TestSuccessfulTransfer tests the successful transfer of a clue.
func TestSuccessfulTransfer(t *testing.T) {
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
	// TODO: Initiate purchase from buyer account
	// TODO: Generate transfer proof
	// TODO: Encrypt clue for buyer
	// TODO: Provide proof to the contract
	// TODO: Verify proof by buyer
	// TODO: Complete transfer with new encrypted clue
	// TODO: Verify ownership has changed
	// TODO: Verify all expected events are emitted
}

// TestTransferSolvedClue tests attempting to transfer a solved clue.
func TestTransferSolvedClue(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, address, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// TODO: Mint and solve a clue
	// TODO: Attempt to initiate purchase of the solved clue
	// TODO: Verify the transaction fails with SolvedClueTransferNotAllowed error
}

// TestInvalidProofVerification tests verification of an invalid proof.
func TestInvalidProofVerification(t *testing.T) {
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
	// TODO: Initiate purchase
	// TODO: Generate proof
	// TODO: Modify proof to make it invalid
	// TODO: Attempt to verify the invalid proof
	// TODO: Verify the verification fails
	// TODO: Cancel the transfer
	// TODO: Verify the TransferCancelled event is emitted
}

// TestCompletingTransferWithoutVerification tests completing a transfer without verification.
func TestCompletingTransferWithoutVerification(t *testing.T) {
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
	// TODO: Initiate purchase
	// TODO: Generate proof and provide it
	// TODO: Attempt to complete transfer without verification
	// TODO: Verify the appropriate error is returned
}

// TestCancelTransfer tests cancelling a transfer.
func TestCancelTransfer(t *testing.T) {
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
	// TODO: Initiate purchase
	// TODO: Wait for the timeout period
	// TODO: Cancel the transfer
	// TODO: Verify the TransferCancelled event is emitted
	// TODO: Verify the buyer is refunded
}
