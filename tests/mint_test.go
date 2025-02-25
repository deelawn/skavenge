// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"crypto/sha256"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/deelawn/skavenge/tests/util"
)

var (
	deployer = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	minter   = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
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

	// Create event listener
	listener, err := util.NewEventListener(client, contract, address)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAddr := crypto.PubkeyToAddress(minterPrivKey.PublicKey)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Create API client for ZK proof operations
	apiClient := util.NewAPIClient()

	// Sample data for the clue
	clueContent := "Find the hidden treasure in the old oak tree"
	solution := "Oak tree"
	solutionHash := sha256.Sum256([]byte(solution))

	// Encrypt the clue content using the ZK proof API
	// Pass the ECDSA public key directly to the API client
	encryptedClueContent, err := apiClient.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt clue content")
	require.NotEmpty(t, encryptedClueContent, "Encrypted clue content should not be empty")

	// Mint a new clue with the encrypted content
	tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Verify the appropriate events are emitted
	transferFound, err := listener.CheckEvent(receipt, "Transfer")
	require.NoError(t, err)
	require.True(t, transferFound, "Transfer event not found")

	// Check that the token with ID 1 has been minted to the minter
	tokenId := uint64(1) // First token is ID 1
	owner, err := contract.OwnerOf(nil, big.NewInt(int64(tokenId)))
	require.NoError(t, err)
	require.Equal(t, minterAddr, owner, "Minter is not the owner of the token")

	// Check the clue content is correctly stored (encrypted)
	clueContents, err := contract.GetClueContents(nil, big.NewInt(int64(tokenId)))
	require.NoError(t, err)
	require.Equal(t, encryptedClueContent, clueContents, "Stored encrypted content does not match")

	// Verify the clue struct in the mapping
	clueData, err := contract.Clues(nil, big.NewInt(int64(tokenId)))
	require.NoError(t, err)
	require.Equal(t, encryptedClueContent, clueData.EncryptedContents, "Encrypted contents do not match")
	require.Equal(t, solutionHash, clueData.SolutionHash, "Solution hash does not match")
	require.False(t, clueData.IsSolved, "Clue should not be marked as solved")
	require.Equal(t, uint64(0), clueData.SolveAttempts.Uint64(), "Solve attempts should be 0")

	// Decrypt the clue to verify it matches the original content
	// Pass the ECDSA private key directly to the API client
	decryptedClue, err := apiClient.DecryptMessage(encryptedClueContent, minterPrivKey)
	require.NoError(t, err, "Failed to decrypt clue content")
	require.Equal(t, clueContent, decryptedClue, "Decrypted content does not match original")
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
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Create API client for ZK proof operations
	apiClient := util.NewAPIClient()

	// Sample data for the clue
	clueContent := "Find the hidden treasure in the old oak tree"
	var emptySolutionHash [32]byte

	// Encrypt the clue content using the ZK proof API
	// Pass the ECDSA public key directly to the API client
	encryptedClueContent, err := apiClient.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Try to mint with empty solution hash
	tx, err := contract.MintClue(minterAuth, encryptedClueContent, emptySolutionHash)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Verify that the transaction succeeded
	require.Equal(t, uint64(1), receipt.Status, "Transaction should succeed")

	// Get the token ID
	tokenId := uint64(1)

	// Check the solution hash
	clueData, err := contract.Clues(nil, big.NewInt(int64(tokenId)))
	require.NoError(t, err)
	require.Equal(t, emptySolutionHash, clueData.SolutionHash, "Solution hash should be empty")
}

// TestMintMultipleClues tests minting multiple clues.
func TestMintMultipleClues(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial("http://localhost:8545")
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

	// Create API client for ZK proof operations
	apiClient := util.NewAPIClient()

	// Mint first clue
	firstClueContent := "First clue content"
	firstSolution := "First solution"
	firstSolutionHash := sha256.Sum256([]byte(firstSolution))

	// Encrypt the first clue content
	// Pass the ECDSA public key directly to the API client
	firstEncryptedClueContent, err := apiClient.EncryptMessage(firstClueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt first clue content")

	tx1, err := contract.MintClue(minterAuth, firstEncryptedClueContent, firstSolutionHash)
	require.NoError(t, err)
	receipt1, err := util.WaitForTransaction(client, tx1)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt1.Status, "First transaction should succeed")

	// Mint second clue
	secondClueContent := "Second clue content"
	secondSolution := "Second solution"
	secondSolutionHash := sha256.Sum256([]byte(secondSolution))

	// Encrypt the second clue content
	// Pass the ECDSA public key directly to the API client
	secondEncryptedClueContent, err := apiClient.EncryptMessage(secondClueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt second clue content")

	// Need to update nonce for second transaction
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	tx2, err := contract.MintClue(minterAuth, secondEncryptedClueContent, secondSolutionHash)
	require.NoError(t, err)
	receipt2, err := util.WaitForTransaction(client, tx2)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt2.Status, "Second transaction should succeed")

	// Check that getCurrentTokenId returns 2
	currentTokenId, err := contract.GetCurrentTokenId(nil)
	require.NoError(t, err)
	require.Equal(t, uint64(2), currentTokenId.Uint64(), "Current token ID should be 2")

	// Verify contents of first clue (encrypted)
	firstClueData, err := contract.Clues(nil, big.NewInt(1))
	require.NoError(t, err)
	require.Equal(t, firstEncryptedClueContent, firstClueData.EncryptedContents, "First clue encrypted content does not match")
	require.Equal(t, firstSolutionHash, firstClueData.SolutionHash, "First solution hash does not match")

	// Verify contents of second clue (encrypted)
	secondClueData, err := contract.Clues(nil, big.NewInt(2))
	require.NoError(t, err)
	require.Equal(t, secondEncryptedClueContent, secondClueData.EncryptedContents, "Second clue encrypted content does not match")
	require.Equal(t, secondSolutionHash, secondClueData.SolutionHash, "Second solution hash does not match")

	// Verify we can decrypt both clues
	// Pass the ECDSA private key directly to the API client
	decryptedFirstClue, err := apiClient.DecryptMessage(firstEncryptedClueContent, minterPrivKey)
	require.NoError(t, err, "Failed to decrypt first clue")
	require.Equal(t, firstClueContent, decryptedFirstClue, "Decrypted first clue doesn't match original")

	decryptedSecondClue, err := apiClient.DecryptMessage(secondEncryptedClueContent, minterPrivKey)
	require.NoError(t, err, "Failed to decrypt second clue")
	require.Equal(t, secondClueContent, decryptedSecondClue, "Decrypted second clue doesn't match original")
}
