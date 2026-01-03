// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/deelawn/skavenge/tests/util"
	"github.com/deelawn/skavenge/zkproof"
)

// TestTokenByIndex tests the TokenByIndex method from ERC721Enumerable.
func TestTokenByIndex(t *testing.T) {
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

	// Update authorized minter
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Verify initial total supply is 0
	totalSupply, err := contract.TotalSupply(nil)
	require.NoError(t, err)
	require.Equal(t, uint64(0), totalSupply.Uint64(), "Initial total supply should be 0")

	// Mint three clues
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	tokenIds := make([]*big.Int, 3)
	for i := 0; i < 3; i++ {
		clueContent := []byte("Test clue content")
		solution := "Solution"
		solutionHash := sha256.Sum256([]byte(solution))

		mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
		require.NoError(t, err)

		encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
		require.NoError(t, err)
		encryptedClueContent := encryptedCipher.Marshal()

		tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR, uint8(1))
		require.NoError(t, err)
		_, err = util.WaitForTransaction(client, tx)
		require.NoError(t, err)

		tokenId, err := getLastMintedTokenID(contract)
		require.NoError(t, err)
		tokenIds[i] = tokenId

		// Update auth for next mint
		minterAuth, err = util.NewTransactOpts(client, minter)
		require.NoError(t, err)
	}

	// Verify total supply is 3
	totalSupply, err = contract.TotalSupply(nil)
	require.NoError(t, err)
	require.Equal(t, uint64(3), totalSupply.Uint64(), "Total supply should be 3 after minting 3 tokens")

	// Test TokenByIndex for each minted token
	for i := 0; i < 3; i++ {
		tokenId, err := contract.TokenByIndex(nil, big.NewInt(int64(i)))
		require.NoError(t, err)
		require.Equal(t, tokenIds[i].Uint64(), tokenId.Uint64(), "TokenByIndex should return the correct token ID")
	}

	// Test TokenByIndex with out-of-bounds index should fail
	_, err = contract.TokenByIndex(nil, big.NewInt(3))
	require.Error(t, err, "TokenByIndex should fail for out-of-bounds index")

	_, err = contract.TokenByIndex(nil, big.NewInt(100))
	require.Error(t, err, "TokenByIndex should fail for large out-of-bounds index")
}

// TestTokenOfOwnerByIndex tests the TokenOfOwnerByIndex method from ERC721Enumerable.
func TestTokenOfOwnerByIndex(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Setup accounts
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAddr := crypto.PubkeyToAddress(minterPrivKey.PublicKey)

	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)

	// Update authorized minter
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Verify initial balance for both accounts is 0
	minterBalance, err := contract.BalanceOf(nil, minterAddr)
	require.NoError(t, err)
	require.Equal(t, uint64(0), minterBalance.Uint64(), "Minter balance should start at 0")

	buyerBalance, err := contract.BalanceOf(nil, buyerAddr)
	require.NoError(t, err)
	require.Equal(t, uint64(0), buyerBalance.Uint64(), "Buyer balance should start at 0")

	// Mint 3 tokens to minter and 2 to buyer (by transferring after mint)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	minterTokens := make([]*big.Int, 3)
	buyerTokens := make([]*big.Int, 2)

	// Mint first 3 tokens to minter
	for i := 0; i < 3; i++ {
		clueContent := []byte("Minter's clue")
		solution := "Solution"
		solutionHash := sha256.Sum256([]byte(solution))

		mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
		require.NoError(t, err)

		encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
		require.NoError(t, err)
		encryptedClueContent := encryptedCipher.Marshal()

		tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR, uint8(1))
		require.NoError(t, err)
		_, err = util.WaitForTransaction(client, tx)
		require.NoError(t, err)

		tokenId, err := getLastMintedTokenID(contract)
		require.NoError(t, err)
		minterTokens[i] = tokenId

		minterAuth, err = util.NewTransactOpts(client, minter)
		require.NoError(t, err)
	}

	// Mint 2 more tokens and transfer them to buyer
	for i := 0; i < 2; i++ {
		clueContent := []byte("Buyer's clue")
		solution := "Solution"
		solutionHash := sha256.Sum256([]byte(solution))

		mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
		require.NoError(t, err)

		encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
		require.NoError(t, err)
		encryptedClueContent := encryptedCipher.Marshal()

		tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR, uint8(1))
		require.NoError(t, err)
		_, err = util.WaitForTransaction(client, tx)
		require.NoError(t, err)

		tokenId, err := getLastMintedTokenID(contract)
		require.NoError(t, err)
		buyerTokens[i] = tokenId

		// Transfer to buyer
		minterAuth, err = util.NewTransactOpts(client, minter)
		require.NoError(t, err)
		tx, err = contract.TransferFrom(minterAuth, minterAddr, buyerAddr, tokenId)
		require.NoError(t, err)
		_, err = util.WaitForTransaction(client, tx)
		require.NoError(t, err)

		minterAuth, err = util.NewTransactOpts(client, minter)
		require.NoError(t, err)
	}

	// Verify balances
	minterBalance, err = contract.BalanceOf(nil, minterAddr)
	require.NoError(t, err)
	require.Equal(t, uint64(3), minterBalance.Uint64(), "Minter should own 3 tokens")

	buyerBalance, err = contract.BalanceOf(nil, buyerAddr)
	require.NoError(t, err)
	require.Equal(t, uint64(2), buyerBalance.Uint64(), "Buyer should own 2 tokens")

	// Test TokenOfOwnerByIndex for minter
	for i := 0; i < 3; i++ {
		tokenId, err := contract.TokenOfOwnerByIndex(nil, minterAddr, big.NewInt(int64(i)))
		require.NoError(t, err)
		require.Equal(t, minterTokens[i].Uint64(), tokenId.Uint64(),
			"TokenOfOwnerByIndex should return correct token ID for minter at index %d", i)
	}

	// Test TokenOfOwnerByIndex for buyer
	for i := 0; i < 2; i++ {
		tokenId, err := contract.TokenOfOwnerByIndex(nil, buyerAddr, big.NewInt(int64(i)))
		require.NoError(t, err)
		require.Equal(t, buyerTokens[i].Uint64(), tokenId.Uint64(),
			"TokenOfOwnerByIndex should return correct token ID for buyer at index %d", i)
	}

	// Test TokenOfOwnerByIndex with out-of-bounds index should fail
	_, err = contract.TokenOfOwnerByIndex(nil, minterAddr, big.NewInt(3))
	require.Error(t, err, "TokenOfOwnerByIndex should fail for out-of-bounds index")

	_, err = contract.TokenOfOwnerByIndex(nil, buyerAddr, big.NewInt(2))
	require.Error(t, err, "TokenOfOwnerByIndex should fail for out-of-bounds index")

	// Test TokenOfOwnerByIndex for address with no tokens
	_, err = contract.TokenOfOwnerByIndex(nil, common.HexToAddress("0x1234567890123456789012345678901234567890"), big.NewInt(0))
	require.Error(t, err, "TokenOfOwnerByIndex should fail for address with no tokens")
}

// TestEnumerableAfterTransfer tests that enumeration works correctly after transfers.
func TestEnumerableAfterTransfer(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Setup accounts
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAddr := crypto.PubkeyToAddress(minterPrivKey.PublicKey)

	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)

	// Update authorized minter
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Mint 2 tokens
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	tokenIds := make([]*big.Int, 2)
	for i := 0; i < 2; i++ {
		clueContent := []byte("Test clue")
		solution := "Solution"
		solutionHash := sha256.Sum256([]byte(solution))

		mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
		require.NoError(t, err)

		encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
		require.NoError(t, err)
		encryptedClueContent := encryptedCipher.Marshal()

		tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR, uint8(1))
		require.NoError(t, err)
		_, err = util.WaitForTransaction(client, tx)
		require.NoError(t, err)

		tokenId, err := getLastMintedTokenID(contract)
		require.NoError(t, err)
		tokenIds[i] = tokenId

		minterAuth, err = util.NewTransactOpts(client, minter)
		require.NoError(t, err)
	}

	// Verify minter owns both tokens
	minterBalance, err := contract.BalanceOf(nil, minterAddr)
	require.NoError(t, err)
	require.Equal(t, uint64(2), minterBalance.Uint64(), "Minter should own 2 tokens initially")

	// Verify buyer owns no tokens
	buyerBalance, err := contract.BalanceOf(nil, buyerAddr)
	require.NoError(t, err)
	require.Equal(t, uint64(0), buyerBalance.Uint64(), "Buyer should own 0 tokens initially")

	// Check TokenOfOwnerByIndex for minter before transfer
	token0, err := contract.TokenOfOwnerByIndex(nil, minterAddr, big.NewInt(0))
	require.NoError(t, err)
	require.Equal(t, tokenIds[0].Uint64(), token0.Uint64())

	token1, err := contract.TokenOfOwnerByIndex(nil, minterAddr, big.NewInt(1))
	require.NoError(t, err)
	require.Equal(t, tokenIds[1].Uint64(), token1.Uint64())

	// Transfer first token to buyer
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	tx, err = contract.TransferFrom(minterAuth, minterAddr, buyerAddr, tokenIds[0])
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Verify balances after transfer
	minterBalance, err = contract.BalanceOf(nil, minterAddr)
	require.NoError(t, err)
	require.Equal(t, uint64(1), minterBalance.Uint64(), "Minter should own 1 token after transfer")

	buyerBalance, err = contract.BalanceOf(nil, buyerAddr)
	require.NoError(t, err)
	require.Equal(t, uint64(1), buyerBalance.Uint64(), "Buyer should own 1 token after transfer")

	// Check TokenOfOwnerByIndex for minter after transfer (should only have one token)
	remainingToken, err := contract.TokenOfOwnerByIndex(nil, minterAddr, big.NewInt(0))
	require.NoError(t, err)
	require.Equal(t, tokenIds[1].Uint64(), remainingToken.Uint64(), "Minter should still own the second token")

	// Check that index 1 is now out of bounds for minter
	_, err = contract.TokenOfOwnerByIndex(nil, minterAddr, big.NewInt(1))
	require.Error(t, err, "Minter should only have 1 token, index 1 should be out of bounds")

	// Check TokenOfOwnerByIndex for buyer after transfer
	buyerToken, err := contract.TokenOfOwnerByIndex(nil, buyerAddr, big.NewInt(0))
	require.NoError(t, err)
	require.Equal(t, tokenIds[0].Uint64(), buyerToken.Uint64(), "Buyer should own the transferred token")

	// Verify total supply remains the same
	totalSupply, err := contract.TotalSupply(nil)
	require.NoError(t, err)
	require.Equal(t, uint64(2), totalSupply.Uint64(), "Total supply should remain 2 after transfer")

	// Verify TokenByIndex still works for both tokens
	globalToken0, err := contract.TokenByIndex(nil, big.NewInt(0))
	require.NoError(t, err)
	require.Equal(t, tokenIds[0].Uint64(), globalToken0.Uint64())

	globalToken1, err := contract.TokenByIndex(nil, big.NewInt(1))
	require.NoError(t, err)
	require.Equal(t, tokenIds[1].Uint64(), globalToken1.Uint64())
}

// TestTotalSupply tests the TotalSupply method from ERC721Enumerable.
func TestTotalSupply(t *testing.T) {
	// Connect to Hardhat network
	client, err := ethclient.Dial(util.GetHardhatURL())
	require.NoError(t, err)

	// Setup deployer account
	deployerAuth, err := util.NewTransactOpts(client, deployer)
	require.NoError(t, err)

	// Deploy contract
	contract, _, err := util.DeployContract(client, deployerAuth)
	require.NoError(t, err)

	// Setup minter
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAddr := crypto.PubkeyToAddress(minterPrivKey.PublicKey)

	// Update authorized minter
	deployerAuth, err = util.NewTransactOpts(client, deployer)
	require.NoError(t, err)
	tx, err := contract.UpdateAuthorizedMinter(deployerAuth, minterAddr)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Create ZK proof system
	ps := zkproof.NewProofSystem()

	// Check initial supply
	totalSupply, err := contract.TotalSupply(nil)
	require.NoError(t, err)
	require.Equal(t, uint64(0), totalSupply.Uint64(), "Initial supply should be 0")

	// Mint tokens and check supply increases
	expectedSupply := 0
	for i := 1; i <= 5; i++ {
		minterAuth, err := util.NewTransactOpts(client, minter)
		require.NoError(t, err)

		clueContent := []byte("Test clue")
		solution := "Solution"
		solutionHash := sha256.Sum256([]byte(solution))

		mintR, err := rand.Int(rand.Reader, ps.Curve.Params().N)
		require.NoError(t, err)

		encryptedCipher, err := ps.EncryptElGamal(clueContent, &minterPrivKey.PublicKey, mintR)
		require.NoError(t, err)
		encryptedClueContent := encryptedCipher.Marshal()

		tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash, mintR, uint8(1))
		require.NoError(t, err)
		_, err = util.WaitForTransaction(client, tx)
		require.NoError(t, err)

		expectedSupply++
		totalSupply, err = contract.TotalSupply(nil)
		require.NoError(t, err)
		require.Equal(t, uint64(expectedSupply), totalSupply.Uint64(),
			"Total supply should be %d after minting %d tokens", expectedSupply, expectedSupply)
	}
}
