package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/deelawn/skavenge/eth/bindings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ContractClientInterface defines the interface for interacting with the Skavenge contract
type ContractClientInterface interface {
	GetTransferInfo(ctx context.Context, transferID [32]byte) (*TransferInfo, error)
	GetTokenOwner(ctx context.Context, tokenID *big.Int) (common.Address, error)
	GetAddress() common.Address
	Close()
}

// ContractClient provides methods to interact with the Skavenge contract
type ContractClient struct {
	client   *ethclient.Client
	contract *bindings.Skavenge
	address  common.Address
}

// NewContractClient creates a new contract client
func NewContractClient(rpcURL, contractAddress string) (*ContractClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to blockchain: %w", err)
	}

	address := common.HexToAddress(contractAddress)
	contract, err := bindings.NewSkavenge(address, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate contract: %w", err)
	}

	return &ContractClient{
		client:   client,
		contract: contract,
		address:  address,
	}, nil
}

// Close closes the client connection
func (c *ContractClient) Close() {
	if c.client != nil {
		c.client.Close()
	}
}

// GetAddress returns the contract address
func (c *ContractClient) GetAddress() common.Address {
	return c.address
}

// GetTokenOwner retrieves the owner of a token
func (c *ContractClient) GetTokenOwner(ctx context.Context, tokenID *big.Int) (common.Address, error) {
	owner, err := c.contract.OwnerOf(nil, tokenID)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get token owner: %w", err)
	}
	return owner, nil
}

// TransferInfo holds the parsed transfer information
type TransferInfo struct {
	Buyer           common.Address
	TokenID         *big.Int
	Value           *big.Int
	InitiatedAt     *big.Int
	Proof           []byte
	NewClueHash     [32]byte
	RValueHash      [32]byte
	ProofVerified   bool
	ProofProvidedAt *big.Int
	VerifiedAt      *big.Int
}

// GetTransferInfo retrieves and parses transfer information
func (c *ContractClient) GetTransferInfo(ctx context.Context, transferID [32]byte) (*TransferInfo, error) {
	transfer, err := c.contract.Transfers(nil, transferID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transfer: %w", err)
	}

	// Check if transfer exists
	if transfer.Buyer == (common.Address{}) {
		return nil, fmt.Errorf("transfer does not exist")
	}

	return &TransferInfo{
		Buyer:           transfer.Buyer,
		TokenID:         transfer.TokenId,
		Value:           transfer.Value,
		InitiatedAt:     transfer.InitiatedAt,
		Proof:           transfer.Proof,
		NewClueHash:     transfer.NewClueHash,
		RValueHash:      transfer.RValueHash,
		ProofVerified:   transfer.ProofVerified,
		ProofProvidedAt: transfer.ProofProvidedAt,
		VerifiedAt:      transfer.VerifiedAt,
	}, nil
}

