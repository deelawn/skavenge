package main

import (
	"errors"
	"sync"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

// Storage defines the interface for storing linked account information
type Storage interface {
	// Set stores a key-value pair, overwriting any existing value.
	Set(key, value string)
	// Get retrieves the value associated with a key. Returns ErrKeyNotFound if the key doesn't exist.
	Get(key string) (string, error)
}

// TransferCiphertextStorage defines the interface for storing transfer ciphertext data
type TransferCiphertextStorage interface {
	// SetTransferCiphertext stores buyer and seller ciphertext for a transfer ID
	SetTransferCiphertext(transferID, buyerCiphertext, sellerCiphertext string)
	// GetTransferCiphertext retrieves ciphertext for a transfer ID
	GetTransferCiphertext(transferID string) (buyerCiphertext string, sellerCiphertext string, err error)
}

// InMemoryStorage implements Storage using an in-memory map
type InMemoryStorage struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewInMemoryStorage creates a new InMemoryStorage instance
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		data: make(map[string]string),
	}
}

// Set stores a key-value pair, overwriting any existing value.
func (s *InMemoryStorage) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

// Get retrieves the value associated with a key. Returns ErrKeyNotFound if the key doesn't exist.
func (s *InMemoryStorage) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, exists := s.data[key]
	if !exists {
		return "", ErrKeyNotFound
	}

	return value, nil
}

// TransferCiphertextData holds the ciphertext for a transfer
type TransferCiphertextData struct {
	BuyerCiphertext  string
	SellerCiphertext string
}

// InMemoryTransferStorage implements TransferCiphertextStorage using an in-memory map
type InMemoryTransferStorage struct {
	mu   sync.RWMutex
	data map[string]TransferCiphertextData
}

// NewInMemoryTransferStorage creates a new InMemoryTransferStorage instance
func NewInMemoryTransferStorage() *InMemoryTransferStorage {
	return &InMemoryTransferStorage{
		data: make(map[string]TransferCiphertextData),
	}
}

// SetTransferCiphertext stores buyer and seller ciphertext for a transfer ID
func (s *InMemoryTransferStorage) SetTransferCiphertext(transferID, buyerCiphertext, sellerCiphertext string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[transferID] = TransferCiphertextData{
		BuyerCiphertext:  buyerCiphertext,
		SellerCiphertext: sellerCiphertext,
	}
}

// GetTransferCiphertext retrieves ciphertext for a transfer ID
func (s *InMemoryTransferStorage) GetTransferCiphertext(transferID string) (string, string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, exists := s.data[transferID]
	if !exists {
		return "", "", ErrKeyNotFound
	}

	return data.BuyerCiphertext, data.SellerCiphertext, nil
}
