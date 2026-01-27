package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

const (
	linkedAccountsFile      = "linked_accounts.json"
	transferCiphertextsFile = "transfer_ciphertexts.json"
)

// DiskStorage implements Storage by wrapping InMemoryStorage and persisting to disk
type DiskStorage struct {
	*InMemoryStorage
	filePath string
	mu       sync.Mutex // mutex for file operations
}

// NewDiskStorage creates a new DiskStorage instance that persists data to the specified directory.
// It loads any existing data from disk on startup.
func NewDiskStorage(storageDir string) (*DiskStorage, error) {
	// Ensure storage directory exists
	if err := os.MkdirAll(storageDir, 0755); err != nil {
		return nil, err
	}

	ds := &DiskStorage{
		InMemoryStorage: NewInMemoryStorage(),
		filePath:        filepath.Join(storageDir, linkedAccountsFile),
	}

	// Load existing data from disk
	if err := ds.loadFromDisk(); err != nil {
		return nil, err
	}

	return ds, nil
}

// loadFromDisk loads existing data from the JSON file into memory
func (ds *DiskStorage) loadFromDisk() error {
	data, err := os.ReadFile(ds.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist yet, that's okay
			return nil
		}
		return err
	}

	var kvPairs map[string]string
	if err := json.Unmarshal(data, &kvPairs); err != nil {
		return err
	}

	// Load into memory storage
	ds.InMemoryStorage.mu.Lock()
	defer ds.InMemoryStorage.mu.Unlock()
	for k, v := range kvPairs {
		ds.InMemoryStorage.data[k] = v
	}

	return nil
}

// saveToDisk writes all data from memory to the JSON file
func (ds *DiskStorage) saveToDisk() error {
	ds.InMemoryStorage.mu.RLock()
	dataCopy := make(map[string]string, len(ds.InMemoryStorage.data))
	for k, v := range ds.InMemoryStorage.data {
		dataCopy[k] = v
	}
	ds.InMemoryStorage.mu.RUnlock()

	jsonData, err := json.MarshalIndent(dataCopy, "", "  ")
	if err != nil {
		return err
	}

	// Write atomically using a temporary file
	tmpFile := ds.filePath + ".tmp"
	if err := os.WriteFile(tmpFile, jsonData, 0644); err != nil {
		return err
	}

	return os.Rename(tmpFile, ds.filePath)
}

// Set stores a key-value pair in memory and persists to disk
func (ds *DiskStorage) Set(key, value string) {
	// Store in memory first
	ds.InMemoryStorage.Set(key, value)

	// Persist to disk
	ds.mu.Lock()
	defer ds.mu.Unlock()
	if err := ds.saveToDisk(); err != nil {
		// Log the error but don't fail the operation
		// The data is still in memory
		println("Warning: failed to persist to disk:", err.Error())
	}
}

// DiskTransferStorage implements TransferCiphertextStorage by wrapping InMemoryTransferStorage and persisting to disk
type DiskTransferStorage struct {
	*InMemoryTransferStorage
	filePath string
	mu       sync.Mutex // mutex for file operations
}

// NewDiskTransferStorage creates a new DiskTransferStorage instance that persists data to the specified directory.
// It loads any existing data from disk on startup.
func NewDiskTransferStorage(storageDir string) (*DiskTransferStorage, error) {
	// Ensure storage directory exists
	if err := os.MkdirAll(storageDir, 0755); err != nil {
		return nil, err
	}

	dts := &DiskTransferStorage{
		InMemoryTransferStorage: NewInMemoryTransferStorage(),
		filePath:                filepath.Join(storageDir, transferCiphertextsFile),
	}

	// Load existing data from disk
	if err := dts.loadFromDisk(); err != nil {
		return nil, err
	}

	return dts, nil
}

// loadFromDisk loads existing data from the JSON file into memory
func (dts *DiskTransferStorage) loadFromDisk() error {
	data, err := os.ReadFile(dts.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist yet, that's okay
			return nil
		}
		return err
	}

	var transfers map[string]TransferCiphertextData
	if err := json.Unmarshal(data, &transfers); err != nil {
		return err
	}

	// Load into memory storage
	dts.InMemoryTransferStorage.mu.Lock()
	defer dts.InMemoryTransferStorage.mu.Unlock()
	for k, v := range transfers {
		dts.InMemoryTransferStorage.data[k] = v
	}

	return nil
}

// saveToDisk writes all data from memory to the JSON file
func (dts *DiskTransferStorage) saveToDisk() error {
	dts.InMemoryTransferStorage.mu.RLock()
	dataCopy := make(map[string]TransferCiphertextData, len(dts.InMemoryTransferStorage.data))
	for k, v := range dts.InMemoryTransferStorage.data {
		dataCopy[k] = v
	}
	dts.InMemoryTransferStorage.mu.RUnlock()

	jsonData, err := json.MarshalIndent(dataCopy, "", "  ")
	if err != nil {
		return err
	}

	// Write atomically using a temporary file
	tmpFile := dts.filePath + ".tmp"
	if err := os.WriteFile(tmpFile, jsonData, 0644); err != nil {
		return err
	}

	return os.Rename(tmpFile, dts.filePath)
}

// SetTransferCiphertext stores buyer and seller ciphertext for a transfer ID in memory and persists to disk
func (dts *DiskTransferStorage) SetTransferCiphertext(transferID, buyerCiphertext, sellerCiphertext string) {
	// Store in memory first
	dts.InMemoryTransferStorage.SetTransferCiphertext(transferID, buyerCiphertext, sellerCiphertext)

	// Persist to disk
	dts.mu.Lock()
	defer dts.mu.Unlock()
	if err := dts.saveToDisk(); err != nil {
		// Log the error but don't fail the operation
		// The data is still in memory
		println("Warning: failed to persist to disk:", err.Error())
	}
}
