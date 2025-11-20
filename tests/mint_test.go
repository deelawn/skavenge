// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/deelawn/skavenge/tests/util"
	"github.com/deelawn/skavenge/zkproof"
)

// TestSuccessfulMint tests the successful minting of a new clue.
func TestSuccessfulMint(t *testing.T) {
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
	minterAddr := crypto.PubkeyToAddress(minterPrivKey.PublicKey)

	// Update authorized minter to minter address.
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Sample data for the clue
	clueContent := []byte("Find the hidden treasure in the old oak tree")
	solution := "Oak tree"
	solutionHash := sha256.Sum256([]byte(solution))

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the clue content using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()
	require.NotEmpty(t, encryptedClueContent, "Encrypted clue content should not be empty")

	// Since deployer is the authorized minter in our test setup, we need to use deployer account to mint
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Mint a new clue with the encrypted content, but to the minter address
	tx, err = contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Verify the appropriate events are emitted
	transferFound, err := listener.CheckEvent(receipt, "Transfer")
	require.NoError(t, err)
	require.True(t, transferFound, "Transfer event not found")

	// Check that the token with ID 1 has been minted to the minter
	owner, err := contract.OwnerOf(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, minterAddr, owner, "Minter is not the owner of the token")

	// Check the clue content is correctly stored (encrypted)
	clueContents, err := contract.GetClueContents(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, encryptedClueContent, clueContents, "Stored encrypted content does not match")

	// Verify the clue struct in the mapping
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, encryptedClueContent, clueData.EncryptedContents, "Encrypted contents do not match")
	require.Equal(t, solutionHash, clueData.SolutionHash, "Solution hash does not match")
	require.False(t, clueData.IsSolved, "Clue should not be marked as solved")
	require.Equal(t, uint64(0), clueData.SolveAttempts.Uint64(), "Solve attempts should be 0")

	// Decrypt the clue to verify it matches the original content
	decryptedClueBytes, err := ps.DecryptMessage(encryptedClueContent, minterPrivKey)
	require.NoError(t, err, "Failed to decrypt clue content")
	require.Equal(t, clueContent, string(decryptedClueBytes), "Decrypted content does not match original")
}

// TestMintWithEmptySolutionHash tests minting a clue with empty solution hash.
func TestMintWithEmptySolutionHash(t *testing.T) {
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

	// Update authorized minter to minter address.
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err := util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Sample data for the clue
	clueContent := []byte("Find the hidden treasure in the old oak tree")
	var emptySolutionHash [32]byte

	// Generate random r value for ElGamal encryption
	mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the clue content using ElGamal
	encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Marshal to bytes for on-chain storage
	encryptedClueContent := encryptedCipher.Marshal()

	// Get deployer auth for minting
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Try to mint with empty solution hash
	tx, err = contract.MintClue(minterAuth, encryptedClueContent, emptySolutionHash, mintR)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	receipt, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Verify that the transaction succeeded
	require.Equal(t, uint64(1), receipt.Status, "Transaction should succeed")

	tokenId, err := getLastMintedTokenID(contract)
	require.NoError(t, err)

	// Check the solution hash
	clueData, err := contract.Clues(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, emptySolutionHash, clueData.SolutionHash, "Solution hash should be empty")
}

// TestMintMultipleClues tests minting multiple clues.
func TestMintMultipleClues(t *testing.T) {
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

	// Update authorized minter to minter address.
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	// Wait for the transaction to be mined
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Update authorized minter to minter address.
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	startTokenId, err := contract.GetCurrentTokenId(nil)
	require.NoError(t, err)

	expTokenId := startTokenId.Add(startTokenId, big.NewInt(2))

	// Mint first clue
	firstClueContent := []byte("First clue content")
	firstSolution := "First solution"
	firstSolutionHash := sha256.Sum256([]byte(firstSolution))

	// Generate random r value for first ElGamal encryption
	firstMintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the first clue content using ElGamal
	firstEncryptedCipher, err := ps.EncryptElGamal(firstClueContent, &minterPrivKey.PublicKey, firstMintR)
	require.NoError(t, err, "Failed to encrypt first clue content")

	// Marshal to bytes for on-chain storage
	firstEncryptedClueContent := firstEncryptedCipher.Marshal()

	tx1, err := contract.MintClue(minterAuth, firstEncryptedClueContent, firstSolutionHash, firstMintR)
	require.NoError(t, err)
	receipt1, err := util.WaitForTransaction(client, tx1)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt1.Status, "First transaction should succeed")

	// Mint second clue
	secondClueContent := []byte("Second clue content")
	secondSolution := "Second solution"
	secondSolutionHash := sha256.Sum256([]byte(secondSolution))

	// Generate random r value for second ElGamal encryption
	secondMintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
	require.NoError(t, err, "Failed to generate r value")

	// Encrypt the second clue content using ElGamal
	secondEncryptedCipher, err := ps.EncryptElGamal(secondClueContent, &minterPrivKey.PublicKey, secondMintR)
	require.NoError(t, err, "Failed to encrypt second clue content")

	// Marshal to bytes for on-chain storage
	secondEncryptedClueContent := secondEncryptedCipher.Marshal()

	// Need to use deployer for second mint as well
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	tx2, err := contract.MintClue(minterAuth, secondEncryptedClueContent, secondSolutionHash, secondMintR)
	require.NoError(t, err)
	receipt2, err := util.WaitForTransaction(client, tx2)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt2.Status, "Second transaction should succeed")

	// Check that getCurrentTokenId returns 2
	currentTokenId, err := contract.GetCurrentTokenId(nil)
	require.NoError(t, err)
	require.Equal(t, expTokenId.Uint64(), currentTokenId.Uint64(), "Current token ID should be 2 more than start")

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
	decryptedFirstClueBytes, err := ps.DecryptMessage(firstEncryptedClueContent, minterPrivKey)
	require.NoError(t, err, "Failed to decrypt first clue")
	require.Equal(t, firstClueContent, string(decryptedFirstClueBytes), "Decrypted first clue doesn't match original")

	decryptedSecondClueBytes, err := ps.DecryptMessage(secondEncryptedClueContent, minterPrivKey)
	require.NoError(t, err, "Failed to decrypt second clue")
	require.Equal(t, secondClueContent, string(decryptedSecondClueBytes), "Decrypted second clue doesn't match original")
}

// TestUpdateAuthorizedMinter tests the authorized minter update functionality
func TestUpdateAuthorizedMinter(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account (initial authorized minter)
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, address, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Create event listener
	listener, err := util.NewEventListener(client, contract, address)
	require.NoError(t, err)

	// Verify initial authorized minter is the deployer
	initialMinter, err := contract.AuthorizedMinter(nil)
	require.NoError(t, err)
	require.Equal(t, deployerAuth.From, initialMinter, "Initial minter should be the deployer")

	// Setup minter (who will become the new authorized minter)
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAddr := crypto.PubkeyToAddress(minterPrivKey.PublicKey)

	// Update the authorized minter to the minter account
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	updateTx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)

	updateReceipt, err := util.WaitForTransaction(client, updateTx)
	require.NoError(t, err)

	// Verify the AuthorizedMinterUpdated event was emitted
	updateEventFound, err := listener.CheckEvent(updateReceipt, "AuthorizedMinterUpdated")
	require.NoError(t, err)
	require.True(t, updateEventFound, "AuthorizedMinterUpdated event not found")

	// Verify the authorized minter has been updated
	newMinter, err := contract.AuthorizedMinter(nil)
	require.NoError(t, err)
	require.Equal(t, minterAddr, newMinter, "Authorized minter should be updated to the minter address")

	// Try to update the authorized minter using the old minter (deployer) - should fail
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)
	deployerAuth.GasLimit = 300000 // Higher gas limit for failing transaction

	_, err = contract.UpdateAuthorizedMinter(deployerAuth, deployerAuth.From)
	require.Error(t, err, "Only current authorized minter should be able to update")

	// Try to mint a clue with the deployer - should fail as it's no longer authorized
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)
	deployerAuth.GasLimit = 300000 // Higher gas limit for failing transaction

	_, err = contract.MintClue(deployerAuth, []byte{1, 2, 3}, [32]byte{}, big.NewInt(1))
	require.Error(t, err, "Non-authorized account should not be able to mint")

	// Mint a clue with the new authorized minter
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	mintTx, err := contract.MintClue(minterAuth, []byte{1, 2, 3}, [32]byte{}, big.NewInt(1))
	require.NoError(t, err)

	_, err = util.WaitForTransaction(client, mintTx)
	require.NoError(t, err, "New authorized minter should be able to mint clues")
}
