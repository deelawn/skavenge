package minting

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/deelawn/skavenge/eth/bindings"
	"github.com/deelawn/skavenge/zkproof"
)

// Minter handles clue minting operations
type Minter struct {
	config           *Config
	client           *ethclient.Client
	contract         *bindings.Skavenge
	minterPrivateKey *ecdsa.PrivateKey
	skavengePrivKey  *ecdsa.PrivateKey
	proofSystem      *zkproof.ProofSystem
	gatewayClient    *GatewayClient
}

// NewMinter creates a new Minter instance
func NewMinter(config *Config) (*Minter, error) {
	// Connect to blockchain
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to blockchain: %w", err)
	}

	// Load contract
	contractAddress := common.HexToAddress(config.ContractAddress)
	contract, err := bindings.NewSkavenge(contractAddress, client)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to load contract: %w", err)
	}

	// Load minter private key
	minterPrivKey, err := crypto.HexToECDSA(strings.TrimPrefix(config.MinterPrivateKey, "0x"))
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to load minter private key: %w", err)
	}

	// Load Skavenge private key for encryption (optional - only needed when minting to self)
	var skavengePrivKey *ecdsa.PrivateKey
	if config.SkavengePrivateKey != "" {
		skavengePrivKey, err = crypto.HexToECDSA(strings.TrimPrefix(config.SkavengePrivateKey, "0x"))
		if err != nil {
			client.Close()
			return nil, fmt.Errorf("failed to load skavenge private key: %w", err)
		}
	}

	// Create proof system
	ps := zkproof.NewProofSystem()

	// Create gateway client
	gatewayClient := NewGatewayClient(config.GatewayURL)

	return &Minter{
		config:           config,
		client:           client,
		contract:         contract,
		minterPrivateKey: minterPrivKey,
		skavengePrivKey:  skavengePrivKey,
		proofSystem:      ps,
		gatewayClient:    gatewayClient,
	}, nil
}

// Close closes the blockchain connection
func (m *Minter) Close() {
	if m.client != nil {
		m.client.Close()
	}
}

// MintClue mints a new clue with the given data and options
func (m *Minter) MintClue(ctx context.Context, clueData *ClueData, options *MintOptions) (*MintResult, error) {
	result := &MintResult{
		ClueContent: clueData.Content,
		Solution:    clueData.Solution,
		PointValue:  clueData.PointValue,
	}

	// Determine the recipient (self or another address)
	var recipientPubKey *ecdsa.PublicKey
	var encryptionR *big.Int

	if options != nil && options.RecipientAddress != "" {
		// Minting to another address - retrieve their public key
		recipientInfo, err := m.getRecipientInfo(options.RecipientAddress)
		if err != nil {
			result.Error = fmt.Errorf("failed to get recipient info: %w", err)
			return result, result.Error
		}
		recipientPubKey = recipientInfo.SkavengePublicKey

		// Generate new r value for the recipient
		encryptionR, err = rand.Int(rand.Reader, m.proofSystem.Curve.Params().N)
		if err != nil {
			result.Error = fmt.Errorf("failed to generate r value: %w", err)
			return result, result.Error
		}
	} else {
		// Minting to self - use our own public key
		if m.skavengePrivKey == nil {
			result.Error = fmt.Errorf("skavenge private key is required when minting to self")
			return result, result.Error
		}
		recipientPubKey = &m.skavengePrivKey.PublicKey

		// Generate r value
		encryptionR, err = rand.Int(rand.Reader, m.proofSystem.Curve.Params().N)
		if err != nil {
			result.Error = fmt.Errorf("failed to generate r value: %w", err)
			return result, result.Error
		}
	}

	// Encrypt the clue content
	encryptedCipher, err := m.proofSystem.EncryptElGamal([]byte(clueData.Content), recipientPubKey, encryptionR)
	if err != nil {
		result.Error = fmt.Errorf("failed to encrypt clue content: %w", err)
		return result, result.Error
	}

	// Marshal encrypted content
	encryptedClueContent := encryptedCipher.Marshal()

	// Hash the solution
	solutionHash := crypto.Keccak256Hash([]byte(clueData.Solution))

	// Prepare transaction options
	auth, err := m.createTransactOpts(ctx)
	if err != nil {
		result.Error = fmt.Errorf("failed to create transaction options: %w", err)
		return result, result.Error
	}

	// Set solve reward if provided
	if clueData.SolveReward != nil && clueData.SolveReward.Cmp(big.NewInt(0)) > 0 {
		auth.Value = clueData.SolveReward
	}

	// Mint the clue
	tx, err := m.contract.MintClue(auth, encryptedClueContent, solutionHash, encryptionR, clueData.PointValue)
	if err != nil {
		result.Error = fmt.Errorf("failed to mint clue: %w", err)
		return result, result.Error
	}

	result.TxHash = tx.Hash().Hex()

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(ctx, m.client, tx)
	if err != nil {
		result.Error = fmt.Errorf("failed to wait for transaction: %w", err)
		return result, result.Error
	}

	if receipt.Status == 0 {
		result.Error = fmt.Errorf("transaction failed")
		return result, result.Error
	}

	// Get the minted token ID
	tokenID, err := m.contract.GetCurrentTokenId(&bind.CallOpts{Context: ctx})
	if err != nil {
		result.Error = fmt.Errorf("failed to get token ID: %w", err)
		return result, result.Error
	}

	// CurrentTokenId is the next token ID to be minted, so subtract 1
	result.TokenID = new(big.Int).Sub(tokenID, big.NewInt(1))

	// If sale price is set, list the clue for sale
	if options != nil && options.SalePrice != nil && options.SalePrice.Cmp(big.NewInt(0)) > 0 {
		if err := m.setSalePrice(ctx, result.TokenID, options.SalePrice, options.Timeout); err != nil {
			result.Error = fmt.Errorf("minted successfully but failed to list for sale: %w", err)
			return result, result.Error
		}
	}

	return result, nil
}

// MintClues mints multiple clues in batch
func (m *Minter) MintClues(ctx context.Context, clues []ClueData, options []MintOptions) ([]*MintResult, error) {
	if len(options) > 0 && len(clues) != len(options) {
		return nil, fmt.Errorf("number of options must match number of clues")
	}

	results := make([]*MintResult, len(clues))

	for i, clueData := range clues {
		var opts *MintOptions
		if len(options) > i {
			opts = &options[i]
		}

		result, err := m.MintClue(ctx, &clueData, opts)
		if err != nil {
			result.Error = err
		}
		results[i] = result
	}

	return results, nil
}

// getRecipientInfo retrieves the recipient's public key from the gateway
func (m *Minter) getRecipientInfo(ethereumAddress string) (*RecipientInfo, error) {
	pubKey, err := m.gatewayClient.GetPublicKey(ethereumAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve public key from gateway: %w", err)
	}

	return &RecipientInfo{
		EthereumAddress:   ethereumAddress,
		SkavengePublicKey: pubKey,
	}, nil
}

// setSalePrice lists a clue for sale
func (m *Minter) setSalePrice(ctx context.Context, tokenID *big.Int, price *big.Int, timeout uint64) error {
	auth, err := m.createTransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed to create transaction options: %w", err)
	}

	// Convert timeout to big.Int
	timeoutBig := new(big.Int).SetUint64(timeout)

	tx, err := m.contract.SetSalePrice(auth, tokenID, price, timeoutBig)
	if err != nil {
		return fmt.Errorf("failed to set sale price: %w", err)
	}

	// Wait for transaction
	receipt, err := bind.WaitMined(ctx, m.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for sale price transaction: %w", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("set sale price transaction failed")
	}

	return nil
}

// createTransactOpts creates transaction options for the minter
func (m *Minter) createTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	chainID, err := m.client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(m.minterPrivateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	// Get nonce
	nonce, err := m.client.PendingNonceAt(ctx, auth.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// Set gas limit
	auth.GasLimit = uint64(300000)

	// Get gas price
	gasPrice, err := m.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}
	auth.GasPrice = gasPrice

	auth.Context = ctx

	return auth, nil
}

// GetMinterAddress returns the minter's Ethereum address
func (m *Minter) GetMinterAddress() common.Address {
	return crypto.PubkeyToAddress(m.minterPrivateKey.PublicKey)
}
