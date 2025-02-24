// Package util provides utility functions for the Skavenge contract tests.
package util

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/deelawn/skavenge/eth/bindings"
)

// EventListener is a utility for listening to contract events.
type EventListener struct {
	client      *ethclient.Client
	contract    *bindings.Skavenge
	contractABI abi.ABI
	address     common.Address
}

// NewEventListener creates a new event listener for the given contract.
func NewEventListener(client *ethclient.Client, contract *bindings.Skavenge, address common.Address) (*EventListener, error) {
	// TODO: Initialize the event listener with the contract details
	return &EventListener{}, nil
}

// WatchEvent watches for a specific event from the contract.
func (e *EventListener) WatchEvent(ctx context.Context, eventName string) (chan types.Log, ethereum.Subscription, error) {
	// TODO: Implement event watching logic
	return nil, nil, nil
}

// CheckEvent checks if a specific event was emitted in a transaction receipt.
func (e *EventListener) CheckEvent(receipt *types.Receipt, eventName string) (bool, error) {
	// TODO: Implement event checking logic
	return false, nil
}

// ParseEvent parses an event from a log.
func (e *EventListener) ParseEvent(log types.Log, eventName string) (interface{}, error) {
	// TODO: Implement event parsing logic
	return nil, nil
}
