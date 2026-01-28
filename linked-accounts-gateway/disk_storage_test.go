package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

// TestDiskStorage_Interface ensures DiskStorage implements Storage interface
func TestDiskStorage_Interface(t *testing.T) {
	var _ Storage = (*DiskStorage)(nil)
}

// TestDiskTransferStorage_Interface ensures DiskTransferStorage implements TransferCiphertextStorage interface
func TestDiskTransferStorage_Interface(t *testing.T) {
	var _ TransferCiphertextStorage = (*DiskTransferStorage)(nil)
}

// TestDiskStorage_SetAndGet tests basic Set and Get operations
func TestDiskStorage_SetAndGet(t *testing.T) {
	tmpDir := t.TempDir()

	store, err := NewDiskStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskStorage() failed: %v", err)
	}

	key := "0x742d35cc6634c0532925a3b844bc9e7595f0beb"
	value := "skv1abc123"

	store.Set(key, value)

	got, err := store.Get(key)
	if err != nil {
		t.Errorf("Get() after Set() failed: %v", err)
	}
	if got != value {
		t.Errorf("Get() = %v, want %v", got, value)
	}
}

// TestDiskStorage_Persistence tests that data persists across storage instances
func TestDiskStorage_Persistence(t *testing.T) {
	tmpDir := t.TempDir()

	// Create first storage instance and set some values
	store1, err := NewDiskStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskStorage() failed: %v", err)
	}

	testData := map[string]string{
		"0x1234567890abcdef1234567890abcdef12345678": "skv1key1",
		"0xabcdef1234567890abcdef1234567890abcdef12": "skv1key2",
		"0x9876543210fedcba9876543210fedcba98765432": "skv1key3",
	}

	for k, v := range testData {
		store1.Set(k, v)
	}

	// Create a new storage instance pointing to the same directory
	store2, err := NewDiskStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskStorage() second instance failed: %v", err)
	}

	// Verify all values were loaded from disk
	for k, expectedValue := range testData {
		got, err := store2.Get(k)
		if err != nil {
			t.Errorf("Get(%s) after reload failed: %v", k, err)
		}
		if got != expectedValue {
			t.Errorf("Get(%s) = %v, want %v", k, got, expectedValue)
		}
	}
}

// TestDiskStorage_FileFormat tests that the file is written in valid JSON format
func TestDiskStorage_FileFormat(t *testing.T) {
	tmpDir := t.TempDir()

	store, err := NewDiskStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskStorage() failed: %v", err)
	}

	key := "0x742d35cc6634c0532925a3b844bc9e7595f0beb"
	value := "skv1abc123"
	store.Set(key, value)

	// Read the file directly and verify it's valid JSON
	filePath := filepath.Join(tmpDir, linkedAccountsFile)
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read storage file: %v", err)
	}

	var kvPairs map[string]string
	if err := json.Unmarshal(data, &kvPairs); err != nil {
		t.Errorf("File content is not valid JSON: %v", err)
	}

	if kvPairs[key] != value {
		t.Errorf("File content: got %v, want %v", kvPairs[key], value)
	}
}

// TestDiskStorage_NonExistentKey tests Get with a non-existent key
func TestDiskStorage_NonExistentKey(t *testing.T) {
	tmpDir := t.TempDir()

	store, err := NewDiskStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskStorage() failed: %v", err)
	}

	_, err = store.Get("nonexistent")
	if !errors.Is(err, ErrKeyNotFound) {
		t.Errorf("Get() error = %v, wantError %v", err, ErrKeyNotFound)
	}
}

// TestDiskStorage_Overwrite tests that Set overwrites existing values
func TestDiskStorage_Overwrite(t *testing.T) {
	tmpDir := t.TempDir()

	store, err := NewDiskStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskStorage() failed: %v", err)
	}

	key := "0x742d35cc6634c0532925a3b844bc9e7595f0beb"
	value1 := "skv1abc123"
	value2 := "skv1xyz789"

	store.Set(key, value1)
	store.Set(key, value2)

	got, err := store.Get(key)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}
	if got != value2 {
		t.Errorf("Get() = %v, want %v (value should be overwritten)", got, value2)
	}

	// Verify persistence of overwritten value
	store2, err := NewDiskStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskStorage() second instance failed: %v", err)
	}

	got, err = store2.Get(key)
	if err != nil {
		t.Errorf("Get() after reload failed: %v", err)
	}
	if got != value2 {
		t.Errorf("Get() after reload = %v, want %v", got, value2)
	}
}

// TestDiskStorage_Concurrency tests thread-safety of DiskStorage
func TestDiskStorage_Concurrency(t *testing.T) {
	tmpDir := t.TempDir()

	store, err := NewDiskStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskStorage() failed: %v", err)
	}

	numGoroutines := 50
	var wg sync.WaitGroup

	// Concurrently set different keys
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			key := string(rune('a' + index))
			value := string(rune('A' + index))
			store.Set(key, value)
		}(i)
	}
	wg.Wait()

	// Verify all values were stored correctly
	for i := 0; i < numGoroutines; i++ {
		key := string(rune('a' + i))
		expectedValue := string(rune('A' + i))
		got, err := store.Get(key)
		if err != nil {
			t.Errorf("Get() failed for key %s: %v", key, err)
		}
		if got != expectedValue {
			t.Errorf("Get(%s) = %v, want %v", key, got, expectedValue)
		}
	}
}

// TestDiskStorage_EmptyDirectory tests that storage works with a fresh directory
func TestDiskStorage_EmptyDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	store, err := NewDiskStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskStorage() failed: %v", err)
	}

	// Should return error for non-existent key
	_, err = store.Get("anykey")
	if !errors.Is(err, ErrKeyNotFound) {
		t.Errorf("Get() on empty storage error = %v, wantError %v", err, ErrKeyNotFound)
	}
}

// TestDiskStorage_CreateDirectory tests that storage creates directory if it doesn't exist
func TestDiskStorage_CreateDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	nestedDir := filepath.Join(tmpDir, "nested", "storage", "dir")

	store, err := NewDiskStorage(nestedDir)
	if err != nil {
		t.Fatalf("NewDiskStorage() with nested dir failed: %v", err)
	}

	key := "testkey"
	value := "testvalue"
	store.Set(key, value)

	got, err := store.Get(key)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}
	if got != value {
		t.Errorf("Get() = %v, want %v", got, value)
	}
}

// TestDiskTransferStorage_SetAndGet tests basic operations
func TestDiskTransferStorage_SetAndGet(t *testing.T) {
	tmpDir := t.TempDir()

	store, err := NewDiskTransferStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskTransferStorage() failed: %v", err)
	}

	transferID := "0x1234567890abcdef"
	buyerCiphertext := "buyer_encrypted_data"
	sellerCiphertext := "seller_encrypted_data"

	store.SetTransferCiphertext(transferID, buyerCiphertext, sellerCiphertext)

	gotBuyer, gotSeller, err := store.GetTransferCiphertext(transferID)
	if err != nil {
		t.Errorf("GetTransferCiphertext() failed: %v", err)
	}
	if gotBuyer != buyerCiphertext {
		t.Errorf("buyer ciphertext = %v, want %v", gotBuyer, buyerCiphertext)
	}
	if gotSeller != sellerCiphertext {
		t.Errorf("seller ciphertext = %v, want %v", gotSeller, sellerCiphertext)
	}
}

// TestDiskTransferStorage_Persistence tests that transfer data persists across instances
func TestDiskTransferStorage_Persistence(t *testing.T) {
	tmpDir := t.TempDir()

	// Create first storage instance
	store1, err := NewDiskTransferStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskTransferStorage() failed: %v", err)
	}

	transferID := "0xabcdef1234567890"
	buyerCiphertext := "buyer_data"
	sellerCiphertext := "seller_data"

	store1.SetTransferCiphertext(transferID, buyerCiphertext, sellerCiphertext)

	// Create new storage instance
	store2, err := NewDiskTransferStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskTransferStorage() second instance failed: %v", err)
	}

	gotBuyer, gotSeller, err := store2.GetTransferCiphertext(transferID)
	if err != nil {
		t.Errorf("GetTransferCiphertext() after reload failed: %v", err)
	}
	if gotBuyer != buyerCiphertext {
		t.Errorf("buyer ciphertext after reload = %v, want %v", gotBuyer, buyerCiphertext)
	}
	if gotSeller != sellerCiphertext {
		t.Errorf("seller ciphertext after reload = %v, want %v", gotSeller, sellerCiphertext)
	}
}

// TestDiskTransferStorage_NonExistentTransfer tests Get with a non-existent transfer
func TestDiskTransferStorage_NonExistentTransfer(t *testing.T) {
	tmpDir := t.TempDir()

	store, err := NewDiskTransferStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskTransferStorage() failed: %v", err)
	}

	_, _, err = store.GetTransferCiphertext("nonexistent")
	if !errors.Is(err, ErrKeyNotFound) {
		t.Errorf("GetTransferCiphertext() error = %v, wantError %v", err, ErrKeyNotFound)
	}
}

// TestDiskTransferStorage_Concurrency tests thread-safety
func TestDiskTransferStorage_Concurrency(t *testing.T) {
	tmpDir := t.TempDir()

	store, err := NewDiskTransferStorage(tmpDir)
	if err != nil {
		t.Fatalf("NewDiskTransferStorage() failed: %v", err)
	}

	numGoroutines := 50
	var wg sync.WaitGroup

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			transferID := string(rune('a' + index))
			buyer := "buyer_" + transferID
			seller := "seller_" + transferID
			store.SetTransferCiphertext(transferID, buyer, seller)
		}(i)
	}
	wg.Wait()

	// Verify all values were stored correctly
	for i := 0; i < numGoroutines; i++ {
		transferID := string(rune('a' + i))
		expectedBuyer := "buyer_" + transferID
		expectedSeller := "seller_" + transferID

		gotBuyer, gotSeller, err := store.GetTransferCiphertext(transferID)
		if err != nil {
			t.Errorf("GetTransferCiphertext() failed for %s: %v", transferID, err)
		}
		if gotBuyer != expectedBuyer {
			t.Errorf("buyer for %s = %v, want %v", transferID, gotBuyer, expectedBuyer)
		}
		if gotSeller != expectedSeller {
			t.Errorf("seller for %s = %v, want %v", transferID, gotSeller, expectedSeller)
		}
	}
}
