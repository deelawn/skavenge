// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/deelawn/skavenge/tests/util"
)

var (
	deployer = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	minter   = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
)

// TestSuccessfulMint tests the successful minting of a new clue.
func TestSuccessfulMint(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, address, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// TODO: Mint a new clue using the minter account
	// TODO: Verify the clue exists and the minter is the owner
	// TODO: Verify the appropriate events are emitted
}

// TestMintWithEmptyContent tests minting a clue with empty content.
func TestMintWithEmptyContent(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, address, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// TODO: Attempt to mint a clue with empty content
	// TODO: Verify the appropriate error is returned
}

// TestMintWithEmptySolutionHash tests minting a clue with empty solution hash.
func TestMintWithEmptySolutionHash(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, address, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// TODO: Attempt to mint a clue with empty solution hash
	// TODO: Verify the appropriate error is returned
}
