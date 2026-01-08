package indexer

import (
	"bytes"
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
	contractABI, err := abi.JSON(bytes.NewReader(abiJSON))
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

	case "ClueAttemptFailed":
		data, parseErr := p.parseClueAttemptFailed(log)
		if parseErr != nil {
			return nil, parseErr
		}
		eventData = data
		clueID = data.TokenID.Uint64()
		event.EventType = string(EventTypeClueAttemptFailed)

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
	// All fields are indexed, so they're in the topics array
	data.TokenID = new(big.Int).SetBytes(log.Topics[1].Bytes())
	data.Minter = common.BytesToAddress(log.Topics[2].Bytes())
	data.Recipient = common.BytesToAddress(log.Topics[3].Bytes())
	return data, nil
}

func (p *EventParser) parseClueAttemptFailed(log types.Log) (ClueAttemptFailedData, error) {
	var data ClueAttemptFailedData
	data.TokenID = new(big.Int).SetBytes(log.Topics[1].Bytes())

	var decoded struct {
		AttemptedSolution string
	}
	err := p.contractABI.UnpackIntoInterface(&decoded, "ClueAttemptFailed", log.Data)
	if err != nil {
		return data, err
	}
	data.AttemptedSolution = decoded.AttemptedSolution
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
	data.CancelledBy = common.BytesToAddress(log.Topics[2].Bytes())
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
