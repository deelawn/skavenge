package indexer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// EventParser handles parsing of contract events
type EventParser struct {
	contractABI abi.ABI
	eventSigs   map[common.Hash]string
}

// NewEventParser creates a new event parser
func NewEventParser() (*EventParser, error) {
	// Define the ABI for the Skavenge contract events
	// This is a simplified ABI containing only the events we care about
	abiJSON := `[
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "tokenId", "type": "uint256"},
				{"indexed": false, "name": "minter", "type": "address"}
			],
			"name": "ClueMinted",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "tokenId", "type": "uint256"},
				{"indexed": false, "name": "remainingAttempts", "type": "uint256"}
			],
			"name": "ClueAttempted",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "tokenId", "type": "uint256"},
				{"indexed": false, "name": "solution", "type": "string"}
			],
			"name": "ClueSolved",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "tokenId", "type": "uint256"},
				{"indexed": false, "name": "price", "type": "uint256"}
			],
			"name": "SalePriceSet",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "tokenId", "type": "uint256"}
			],
			"name": "SalePriceRemoved",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "transferId", "type": "bytes32"},
				{"indexed": true, "name": "buyer", "type": "address"},
				{"indexed": true, "name": "tokenId", "type": "uint256"}
			],
			"name": "TransferInitiated",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "transferId", "type": "bytes32"},
				{"indexed": false, "name": "proof", "type": "bytes"},
				{"indexed": false, "name": "newClueHash", "type": "bytes32"},
				{"indexed": false, "name": "rValueHash", "type": "bytes32"}
			],
			"name": "ProofProvided",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "transferId", "type": "bytes32"}
			],
			"name": "ProofVerified",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "transferId", "type": "bytes32"},
				{"indexed": false, "name": "rValue", "type": "uint256"}
			],
			"name": "TransferCompleted",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "transferId", "type": "bytes32"}
			],
			"name": "TransferCancelled",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "oldMinter", "type": "address"},
				{"indexed": true, "name": "newMinter", "type": "address"}
			],
			"name": "AuthorizedMinterUpdated",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "from", "type": "address"},
				{"indexed": true, "name": "to", "type": "address"},
				{"indexed": true, "name": "tokenId", "type": "uint256"}
			],
			"name": "Transfer",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "owner", "type": "address"},
				{"indexed": true, "name": "approved", "type": "address"},
				{"indexed": true, "name": "tokenId", "type": "uint256"}
			],
			"name": "Approval",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "name": "owner", "type": "address"},
				{"indexed": true, "name": "operator", "type": "address"},
				{"indexed": false, "name": "approved", "type": "bool"}
			],
			"name": "ApprovalForAll",
			"type": "event"
		}
	]`

	contractABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	// Build event signature map
	eventSigs := make(map[common.Hash]string)
	for _, event := range contractABI.Events {
		sig := crypto.Keccak256Hash([]byte(event.Sig))
		eventSigs[sig] = event.Name
	}

	return &EventParser{
		contractABI: contractABI,
		eventSigs:   eventSigs,
	}, nil
}

// ParseLog parses a contract log into an Event
func (p *EventParser) ParseLog(log types.Log, txSender common.Address, txIndex uint, timestamp time.Time) (*Event, error) {
	if len(log.Topics) == 0 {
		return nil, fmt.Errorf("log has no topics")
	}

	eventName, exists := p.eventSigs[log.Topics[0]]
	if !exists {
		return nil, fmt.Errorf("unknown event signature: %s", log.Topics[0].Hex())
	}

	event := &Event{
		BlockNumber:      log.BlockNumber,
		TransactionIndex: txIndex,
		EventIndex:       log.Index,
		TransactionHash:  log.TxHash.Hex(),
		InitiatedBy:      strings.ToLower(txSender.Hex()),
		Timestamp:        timestamp.Unix(),
		BlockHash:        log.BlockHash,
		Removed:          log.Removed,
	}

	// Parse event-specific data
	var eventData EventData
	var clueID uint64
	var err error

	switch eventName {
	case "ClueMinted":
		data, parseErr := p.parseClueMinted(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = data.TokenID.Uint64()
		event.EventType = string(EventTypeClueMinted)

	case "ClueAttempted":
		data, parseErr := p.parseClueAttempted(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = data.TokenID.Uint64()
		event.EventType = string(EventTypeClueAttempted)

	case "ClueSolved":
		data, parseErr := p.parseClueSolved(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = data.TokenID.Uint64()
		event.EventType = string(EventTypeClueSolved)

	case "SalePriceSet":
		data, parseErr := p.parseSalePriceSet(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = data.TokenID.Uint64()
		event.EventType = string(EventTypeSalePriceSet)

	case "SalePriceRemoved":
		data, parseErr := p.parseSalePriceRemoved(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = data.TokenID.Uint64()
		event.EventType = string(EventTypeSalePriceRemoved)

	case "TransferInitiated":
		data, parseErr := p.parseTransferInitiated(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = data.TokenID.Uint64()
		event.EventType = string(EventTypeTransferInitiated)

	case "ProofProvided":
		data, parseErr := p.parseProofProvided(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = 0 // No token ID for this event
		event.EventType = string(EventTypeProofProvided)

	case "ProofVerified":
		data, parseErr := p.parseProofVerified(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = 0 // No token ID for this event
		event.EventType = string(EventTypeProofVerified)

	case "TransferCompleted":
		data, parseErr := p.parseTransferCompleted(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = 0 // No token ID for this event
		event.EventType = string(EventTypeTransferCompleted)

	case "TransferCancelled":
		data, parseErr := p.parseTransferCancelled(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = 0 // No token ID for this event
		event.EventType = string(EventTypeTransferCancelled)

	case "AuthorizedMinterUpdated":
		data, parseErr := p.parseAuthorizedMinterUpdated(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = 0 // No token ID for this event
		event.EventType = string(EventTypeAuthorizedMinterUpdated)

	case "Transfer":
		data, parseErr := p.parseTransfer(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = data.TokenID.Uint64()
		event.EventType = string(EventTypeTransfer)

	case "Approval":
		data, parseErr := p.parseApproval(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = data.TokenID.Uint64()
		event.EventType = string(EventTypeApproval)

	case "ApprovalForAll":
		data, parseErr := p.parseApprovalForAll(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = 0 // No token ID for this event
		event.EventType = string(EventTypeApprovalForAll)

	default:
		return nil, fmt.Errorf("unhandled event: %s", eventName)
	}

	// Set ClueID
	event.ClueID = clueID

	// Serialize event data to JSON for metadata
	event.Metadata, err = json.Marshal(eventData)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize event data: %w", err)
	}

	return event, nil
}

func (p *EventParser) parseClueMinted(log types.Log) (ClueMintedData, error) {
	var data ClueMintedData
	data.TokenID = new(big.Int).SetBytes(log.Topics[1].Bytes())

	var decoded struct {
		Minter common.Address
	}
	err := p.contractABI.UnpackIntoInterface(&decoded, "ClueMinted", log.Data)
	if err != nil {
		return data, err
	}
	data.Minter = decoded.Minter
	return data, nil
}

func (p *EventParser) parseClueAttempted(log types.Log) (ClueAttemptedData, error) {
	var data ClueAttemptedData
	data.TokenID = new(big.Int).SetBytes(log.Topics[1].Bytes())

	var decoded struct {
		RemainingAttempts *big.Int
	}
	err := p.contractABI.UnpackIntoInterface(&decoded, "ClueAttempted", log.Data)
	if err != nil {
		return data, err
	}
	data.RemainingAttempts = decoded.RemainingAttempts
	return data, nil
}

func (p *EventParser) parseClueSolved(log types.Log) (ClueSolvedData, error) {
	var data ClueSolvedData
	data.TokenID = new(big.Int).SetBytes(log.Topics[1].Bytes())

	var decoded struct {
		Solution string
	}
	err := p.contractABI.UnpackIntoInterface(&decoded, "ClueSolved", log.Data)
	if err != nil {
		return data, err
	}
	data.Solution = decoded.Solution
	return data, nil
}

func (p *EventParser) parseSalePriceSet(log types.Log) (SalePriceSetData, error) {
	var data SalePriceSetData
	data.TokenID = new(big.Int).SetBytes(log.Topics[1].Bytes())

	var decoded struct {
		Price *big.Int
	}
	err := p.contractABI.UnpackIntoInterface(&decoded, "SalePriceSet", log.Data)
	if err != nil {
		return data, err
	}
	data.Price = decoded.Price
	return data, nil
}

func (p *EventParser) parseSalePriceRemoved(log types.Log) (SalePriceRemovedData, error) {
	var data SalePriceRemovedData
	data.TokenID = new(big.Int).SetBytes(log.Topics[1].Bytes())
	return data, nil
}

func (p *EventParser) parseTransferInitiated(log types.Log) (TransferInitiatedData, error) {
	var data TransferInitiatedData
	copy(data.TransferID[:], log.Topics[1].Bytes())
	data.Buyer = common.BytesToAddress(log.Topics[2].Bytes())
	data.TokenID = new(big.Int).SetBytes(log.Topics[3].Bytes())
	return data, nil
}

func (p *EventParser) parseProofProvided(log types.Log) (ProofProvidedData, error) {
	var data ProofProvidedData
	copy(data.TransferID[:], log.Topics[1].Bytes())

	var decoded struct {
		Proof       []byte
		NewClueHash [32]byte
		RValueHash  [32]byte
	}
	err := p.contractABI.UnpackIntoInterface(&decoded, "ProofProvided", log.Data)
	if err != nil {
		return data, err
	}
	data.Proof = decoded.Proof
	data.NewClueHash = decoded.NewClueHash
	data.RValueHash = decoded.RValueHash
	return data, nil
}

func (p *EventParser) parseProofVerified(log types.Log) (ProofVerifiedData, error) {
	var data ProofVerifiedData
	copy(data.TransferID[:], log.Topics[1].Bytes())
	return data, nil
}

func (p *EventParser) parseTransferCompleted(log types.Log) (TransferCompletedData, error) {
	var data TransferCompletedData
	copy(data.TransferID[:], log.Topics[1].Bytes())

	var decoded struct {
		RValue *big.Int
	}
	err := p.contractABI.UnpackIntoInterface(&decoded, "TransferCompleted", log.Data)
	if err != nil {
		return data, err
	}
	data.RValue = decoded.RValue
	return data, nil
}

func (p *EventParser) parseTransferCancelled(log types.Log) (TransferCancelledData, error) {
	var data TransferCancelledData
	copy(data.TransferID[:], log.Topics[1].Bytes())
	return data, nil
}

func (p *EventParser) parseAuthorizedMinterUpdated(log types.Log) (AuthorizedMinterUpdatedData, error) {
	var data AuthorizedMinterUpdatedData
	data.OldMinter = common.BytesToAddress(log.Topics[1].Bytes())
	data.NewMinter = common.BytesToAddress(log.Topics[2].Bytes())
	return data, nil
}

func (p *EventParser) parseTransfer(log types.Log) (TransferData, error) {
	var data TransferData
	data.From = common.BytesToAddress(log.Topics[1].Bytes())
	data.To = common.BytesToAddress(log.Topics[2].Bytes())
	data.TokenID = new(big.Int).SetBytes(log.Topics[3].Bytes())
	return data, nil
}

func (p *EventParser) parseApproval(log types.Log) (ApprovalData, error) {
	var data ApprovalData
	data.Owner = common.BytesToAddress(log.Topics[1].Bytes())
	data.Approved = common.BytesToAddress(log.Topics[2].Bytes())
	data.TokenID = new(big.Int).SetBytes(log.Topics[3].Bytes())
	return data, nil
}

func (p *EventParser) parseApprovalForAll(log types.Log) (ApprovalForAllData, error) {
	var data ApprovalForAllData
	data.Owner = common.BytesToAddress(log.Topics[1].Bytes())
	data.Operator = common.BytesToAddress(log.Topics[2].Bytes())

	var decoded struct {
		Approved bool
	}
	err := p.contractABI.UnpackIntoInterface(&decoded, "ApprovalForAll", log.Data)
	if err != nil {
		return data, err
	}
	data.Approved = decoded.Approved
	return data, nil
}
