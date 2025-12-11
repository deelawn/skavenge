package main

import (
	"errors"
	"sync"
)

var (
	ErrKeyExists   = errors.New("key already exists")
	ErrKeyNotFound = errors.New("key not found")
)

// Storage defines the interface for storing linked account information
type Storage interface {
	// Set stores a key-value pair. Returns ErrKeyExists if the key already exists.
	Set(key, value string) error
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

// Set stores a key-value pair. Returns ErrKeyExists if the key already exists.
func (s *InMemoryStorage) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if key already exists
	if _, exists := s.data[key]; exists {
		return ErrKeyExists
	}

	s.data[key] = value
	return nil
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
