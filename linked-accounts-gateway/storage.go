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
