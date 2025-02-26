// Package tests contains integration tests for the Skavenge contract.
package tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	proto "github.com/deelawn/skavenge-zk-proof/api/zkproof"
	"github.com/deelawn/skavenge/tests/util"
)

var (
	other = "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"
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

	// Create event listener
	listener, err := util.NewEventListener(client, contract, address)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	// Create API client for ZK proof operations
	apiClient, err := util.NewGRPCClient()
	require.NoError(t, err)

	// Sample data for the clue
	clueContent := "Find the hidden treasure in the old oak tree"
	solution := "Oak tree"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Encrypt the clue content for the minter
	encryptedClueContent, err := apiClient.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Mint a new clue
	tx, err := contract.MintClue(minterAuth, encryptedClueContent, solutionHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	// Verify the clue is minted successfully
	tokenId := big.NewInt(1)
	owner, err := contract.OwnerOf(nil, tokenId)
	require.NoError(t, err)
	require.Equal(t, crypto.PubkeyToAddress(minterPrivKey.PublicKey), owner, "Minter should be the owner")

	// Initiate purchase from buyer account
	// Set value to 1 ETH
	buyerAuth.Value = big.NewInt(1000000000000000000)
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

	// Generate proof and encrypted clue for buyer
	partialProof, err := apiClient.GeneratePartialProof(clueContent, encryptedClueContent, &buyerPrivKey.PublicKey, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to generate proof")

	// Calculate S value for the proof
	c, ok := new(big.Int).SetString(partialProof.C, 10)
	require.True(t, ok, "Failed to parse C value")
	k, ok := new(big.Int).SetString(partialProof.K, 10)
	require.True(t, ok, "Failed to parse K value")
	curveN, ok := new(big.Int).SetString(partialProof.CurveN, 10)
	require.True(t, ok, "Failed to parse CurveN value")

	s := computeS(c, k, minterPrivKey.D, curveN)
	proof := &proto.Proof{
		C:               c.String(),
		S:               s.String(),
		R1:              partialProof.R1,
		R2:              partialProof.R2,
		BuyerPubKey:     partialProof.BuyerPubKey,
		SellerPubKey:    partialProof.SellerPubKey,
		BuyerCipherHash: partialProof.BuyerCipherHash,
	}

	marshalRes, err := apiClient.MarshalProof(proof)
	require.NoError(t, err)

	// Provide proof to the contract
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	proofTx, err := contract.ProvideProof(minterAuth, transferId, marshalRes, [32]byte(partialProof.BuyerCipherHash))
	require.NoError(t, err)
	proofReceipt, err := util.WaitForTransaction(client, proofTx)
	require.NoError(t, err)

	// Verify the ProofProvided event is emitted
	proofProvidedFound, err := listener.CheckEvent(proofReceipt, "ProofProvided")
	require.NoError(t, err)
	require.True(t, proofProvidedFound, "ProofProvided event not found")

	transfer, err := contract.Transfers(&bind.CallOpts{}, transferId)
	require.NoError(t, err)

	clue, err := contract.Clues(&bind.CallOpts{}, transfer.TokenId)
	require.NoError(t, err)

	// Verify the proof using the API.
	ok, err = apiClient.VerifyProof(transfer.Proof, clue.EncryptedContents)
	require.NoError(t, err)
	require.True(t, ok, "Proof verification failed")

	// Verify proof by buyer
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

	// Complete transfer with new encrypted clue
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	completeTx, err := contract.CompleteTransfer(minterAuth, transferId, partialProof.BuyerCipherText)
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
	require.Equal(t, partialProof.BuyerCipherText, newClueContents, "Clue content should be updated for buyer")

	// Verify the buyer can decrypt the clue
	decryptedClue, err := apiClient.DecryptMessage(newClueContents, buyerPrivKey)
	require.NoError(t, err, "Buyer should be able to decrypt the clue")
	require.Equal(t, clueContent, decryptedClue, "Decrypted content should match original")
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

	// Create event listener
	listener, err := util.NewEventListener(client, contract, address)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	// Create API client for ZK proof operations
	apiClient, err := util.NewGRPCClient()
	require.NoError(t, err)

	// Mint a clue
	clueContent := "Find the hidden treasure in the forest"
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Encrypt the clue content
	encryptedClueContent, err := apiClient.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Mint the clue
	tx, err := contract.MintClue(minterAuth, []byte(encryptedClueContent), solutionHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId := big.NewInt(1)

	// Initiate purchase from buyer
	buyerAuth.Value = big.NewInt(1000000000000000000) // 1 ETH
	buyTx, err := contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, buyTx)
	require.NoError(t, err)

	// Generate transfer ID
	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// Generate proof and encrypted clue for buyer
	partialProof, err := apiClient.GeneratePartialProof(clueContent, encryptedClueContent, &buyerPrivKey.PublicKey, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to generate proof")

	// Calculate S value for the proof
	c, ok := new(big.Int).SetString(partialProof.C, 10)
	require.True(t, ok, "Failed to parse C value")
	k, ok := new(big.Int).SetString(partialProof.K, 10)
	require.True(t, ok, "Failed to parse K value")
	curveN, ok := new(big.Int).SetString(partialProof.CurveN, 10)
	require.True(t, ok, "Failed to parse CurveN value")
	s := computeS(c, k, minterPrivKey.D, curveN)

	proof := &proto.Proof{
		C:               c.String(),
		S:               s.String(),
		R1:              partialProof.R1,
		R2:              partialProof.R2,
		BuyerPubKey:     partialProof.BuyerPubKey,
		SellerPubKey:    partialProof.SellerPubKey,
		BuyerCipherHash: partialProof.BuyerCipherHash,
	}

	marshalRes, err := apiClient.MarshalProof(proof)
	require.NoError(t, err)

	// Modify proof to make it invalid (just corrupt the first byte)
	invalidProof := append([]byte{'X'}, marshalRes[1:]...)

	// Provide the invalid proof to the contract
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	proofTx, err := contract.ProvideProof(minterAuth, transferId, []byte(invalidProof), [32]byte(partialProof.BuyerCipherHash))
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, proofTx)
	require.NoError(t, err)

	transfer, err := contract.Transfers(&bind.CallOpts{}, transferId)
	require.NoError(t, err)

	clue, err := contract.Clues(&bind.CallOpts{}, transfer.TokenId)
	require.NoError(t, err)

	// Verify the proof using the API.
	ok, err = apiClient.VerifyProof(transfer.Proof, clue.EncryptedContents)
	require.NoError(t, err)
	require.False(t, ok, "Proof verification didn't fail")

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

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	// Create API client for ZK proof operations
	apiClient, err := util.NewGRPCClient()
	require.NoError(t, err)

	// Mint a clue
	clueContent := "Find the hidden treasure in the forest"
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Encrypt the clue content
	encryptedClueContent, err := apiClient.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Mint the clue
	tx, err := contract.MintClue(minterAuth, []byte(encryptedClueContent), solutionHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId := big.NewInt(1)

	// Initiate purchase from buyer
	buyerAuth.Value = big.NewInt(1000000000000000000) // 1 ETH
	buyTx, err := contract.InitiatePurchase(buyerAuth, tokenId)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, buyTx)
	require.NoError(t, err)

	// Generate transfer ID
	transferId, err := contract.GenerateTransferId(nil, buyerAddr, tokenId)
	require.NoError(t, err)

	// Generate proof and encrypted clue for buyer
	partialProof, err := apiClient.GeneratePartialProof(clueContent, encryptedClueContent, &buyerPrivKey.PublicKey, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to generate proof")

	// Calculate S value for the proof
	c, ok := new(big.Int).SetString(partialProof.C, 10)
	require.True(t, ok, "Failed to parse C value")
	k, ok := new(big.Int).SetString(partialProof.K, 10)
	require.True(t, ok, "Failed to parse K value")
	curveN, ok := new(big.Int).SetString(partialProof.CurveN, 10)
	require.True(t, ok, "Failed to parse CurveN value")
	s := computeS(c, k, minterPrivKey.D, curveN)

	proof := &proto.Proof{
		C:               c.String(),
		S:               s.String(),
		R1:              partialProof.R1,
		R2:              partialProof.R2,
		BuyerPubKey:     partialProof.BuyerPubKey,
		SellerPubKey:    partialProof.SellerPubKey,
		BuyerCipherHash: partialProof.BuyerCipherHash,
	}

	marshalRes, err := apiClient.MarshalProof(proof)
	require.NoError(t, err)

	// Provide proof to the contract
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	proofTx, err := contract.ProvideProof(minterAuth, transferId, marshalRes, [32]byte(partialProof.BuyerCipherHash))
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, proofTx)
	require.NoError(t, err)

	// Skip the verification step
	// Attempt to complete transfer without verification
	minterAuth, err = util.NewTransactOpts(client, minter)
	require.NoError(t, err)
	minterAuth.GasLimit = 300000 // Higher gas limit for failing transaction

	_, err = contract.CompleteTransfer(minterAuth, transferId, partialProof.BuyerCipherText)
	require.Error(t, err, "Transaction should fail")
	// require.Contains(t, err.Error(), "execution reverted", "Transaction should revert")
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

	// Create event listener
	listener, err := util.NewEventListener(client, contract, address)
	require.NoError(t, err)

	// Setup minter account and keys
	minterPrivKey, err := crypto.HexToECDSA(minter)
	require.NoError(t, err)
	minterAuth, err := util.NewTransactOpts(client, minter)
	require.NoError(t, err)

	// Setup buyer account and keys
	buyerPrivKey, err := crypto.HexToECDSA(buyer)
	require.NoError(t, err)
	buyerAddr := crypto.PubkeyToAddress(buyerPrivKey.PublicKey)
	buyerAuth, err := util.NewTransactOpts(client, buyer)
	require.NoError(t, err)

	// Create API client for ZK proof operations
	apiClient, err := util.NewGRPCClient()
	require.NoError(t, err)

	// Get initial buyer balance
	initialBuyerBalance, err := client.BalanceAt(context.Background(), buyerAddr, nil)
	require.NoError(t, err)

	// Mint a clue
	clueContent := "Find the hidden treasure in the forest"
	solution := "Behind the waterfall"
	solutionHash := crypto.Keccak256Hash([]byte(solution))

	// Encrypt the clue content
	encryptedClueContent, err := apiClient.EncryptMessage(clueContent, &minterPrivKey.PublicKey)
	require.NoError(t, err, "Failed to encrypt clue content")

	// Mint the clue
	tx, err := contract.MintClue(minterAuth, []byte(encryptedClueContent), solutionHash)
	require.NoError(t, err)
	_, err = util.WaitForTransaction(client, tx)
	require.NoError(t, err)

	tokenId := big.NewInt(1)

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

	// Verify buyer received refund (checking balance increase is challenging due to gas costs,
	// so we'll just check the transfer object is deleted)
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

func computeS(c, k, sellerKeyD, curveN *big.Int) *big.Int {
	s := new(big.Int).Mul(c, sellerKeyD)
	s.Sub(k, s)
	s.Mod(s, curveN)
	return s
}
