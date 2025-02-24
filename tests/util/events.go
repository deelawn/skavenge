// Package util provides utility functions for the Skavenge contract tests.
package util

import (
	"context"
	"fmt"
	"strings"

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
	parsedABI, err := abi.JSON(strings.NewReader(bindings.SkavengeMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	return &EventListener{
		client:      client,
		contract:    contract,
		contractABI: parsedABI,
		address:     address,
	}, nil
}

// WatchEvent watches for a specific event from the contract.
func (e *EventListener) WatchEvent(ctx context.Context, eventName string) (chan types.Log, ethereum.Subscription, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{e.address},
		Topics:    [][]common.Hash{{e.contractABI.Events[eventName].ID}},
	}

	logs := make(chan types.Log)
	sub, err := e.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to subscribe to %s events: %w", eventName, err)
	}

	return logs, sub, nil
}

// CheckEvent checks if a specific event was emitted in a transaction receipt.
func (e *EventListener) CheckEvent(receipt *types.Receipt, eventName string) (bool, error) {
	eventID := e.contractABI.Events[eventName].ID

	for _, log := range receipt.Logs {
		if log.Address == e.address && len(log.Topics) > 0 && log.Topics[0] == eventID {
			return true, nil
		}
	}

	return false, nil
}

// ParseEvent parses an event from a log.
func (e *EventListener) ParseEvent(log types.Log, eventName string) (interface{}, error) {
	event, ok := e.contractABI.Events[eventName]
	if !ok {
		return nil, fmt.Errorf("event %s not found in ABI", eventName)
	}

	// Create a map to hold the event data
	data := make(map[string]interface{})

	// Unpack the event data
	if err := e.contractABI.UnpackIntoMap(data, eventName, log.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack event data: %w", err)
	}

	// Add indexed parameters
	for i := 0; i < len(event.Inputs) && i < len(log.Topics)-1; i++ {
		if event.Inputs[i].Indexed {
			// +1 because first topic is event signature
			if err := abi.ParseTopicsIntoMap(data, event.Inputs, log.Topics[1:]); err != nil {
				return nil, fmt.Errorf("failed to parse topic: %w", err)
			}
		}
	}

	return data, nil
}

// GetEventsByName gets all events of a specific type from a transaction receipt.
func (e *EventListener) GetEventsByName(receipt *types.Receipt, eventName string) ([]interface{}, error) {
	eventID := e.contractABI.Events[eventName].ID
	var events []interface{}

	for _, log := range receipt.Logs {
		if log.Address == e.address && len(log.Topics) > 0 && log.Topics[0] == eventID {
			event, err := e.ParseEvent(*log, eventName)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}

	return events, nil
}
