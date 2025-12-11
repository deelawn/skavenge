package main

import (
	"errors"
	"sync"
	"testing"
)

// TestInMemoryStorage_Set tests the Set method of InMemoryStorage
func TestInMemoryStorage_Set(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value string
	}{
		{
			name:  "set new key-value pair",
			key:   "0x742d35cc6634c0532925a3b844bc9e7595f0beb",
			value: "skv1abc123",
		},
		{
			name:  "set another key-value pair",
			key:   "0x1234567890abcdef",
			value: "skv1xyz789",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewInMemoryStorage()
			store.Set(tt.key, tt.value)

			// Verify the value was actually stored
			got, err := store.Get(tt.key)
			if err != nil {
				t.Errorf("Get() after Set() failed: %v", err)
			}
			if got != tt.value {
				t.Errorf("Get() = %v, want %v", got, tt.value)
			}
		})
	}
}

// TestInMemoryStorage_Set_Overwrite tests that Set overwrites existing values
func TestInMemoryStorage_Set_Overwrite(t *testing.T) {
	store := NewInMemoryStorage()
	key := "0x742d35cc6634c0532925a3b844bc9e7595f0beb"
	value1 := "skv1abc123"
	value2 := "skv1xyz789"

	// First set
	store.Set(key, value1)

	// Verify first value is stored
	got, err := store.Get(key)
	if err != nil {
		t.Fatalf("Get() after first Set() failed: %v", err)
	}
	if got != value1 {
		t.Errorf("Get() = %v, want %v", got, value1)
	}

	// Second set with same key should overwrite
	store.Set(key, value2)

	// Verify the new value overwrote the old one
	got, err = store.Get(key)
	if err != nil {
		t.Fatalf("Get() after second Set() failed: %v", err)
	}
	if got != value2 {
		t.Errorf("Get() = %v, want %v (value should be overwritten)", got, value2)
	}
}

// TestInMemoryStorage_Get tests the Get method of InMemoryStorage
func TestInMemoryStorage_Get(t *testing.T) {
	store := NewInMemoryStorage()
	key := "0x742d35cc6634c0532925a3b844bc9e7595f0beb"
	value := "skv1abc123"

	// Set a value first
	store.Set(key, value)

	tests := []struct {
		name      string
		key       string
		wantValue string
		wantError error
	}{
		{
			name:      "get existing key",
			key:       key,
			wantValue: value,
			wantError: nil,
		},
		{
			name:      "get non-existent key",
			key:       "0xnonexistent",
			wantValue: "",
			wantError: ErrKeyNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.Get(tt.key)

			if !errors.Is(err, tt.wantError) {
				t.Errorf("Get() error = %v, wantError %v", err, tt.wantError)
			}

			if got != tt.wantValue {
				t.Errorf("Get() = %v, want %v", got, tt.wantValue)
			}
		})
	}
}

// TestInMemoryStorage_Concurrency tests thread-safety of InMemoryStorage
func TestInMemoryStorage_Concurrency(t *testing.T) {
	store := NewInMemoryStorage()
	numGoroutines := 100
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

// TestInMemoryStorage_ConcurrentReads tests concurrent read operations
func TestInMemoryStorage_ConcurrentReads(t *testing.T) {
	store := NewInMemoryStorage()
	key := "0x742d35cc6634c0532925a3b844bc9e7595f0beb"
	value := "skv1abc123"

	// Set a value
	store.Set(key, value)

	// Concurrently read the same key
	numGoroutines := 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			got, err := store.Get(key)
			if err != nil {
				t.Errorf("Concurrent Get() failed: %v", err)
			}
			if got != value {
				t.Errorf("Concurrent Get() = %v, want %v", got, value)
			}
		}()
	}
	wg.Wait()
}

// TestInMemoryStorage_ConcurrentWrites tests concurrent write operations to the same key
func TestInMemoryStorage_ConcurrentWrites(t *testing.T) {
	store := NewInMemoryStorage()
	key := "test-key"
	numGoroutines := 100
	var wg sync.WaitGroup

	// Concurrently write to the same key
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			value := string(rune('A' + (index % 26)))
			store.Set(key, value)
		}(i)
	}
	wg.Wait()

	// After all writes complete, there should be a valid value
	// (we don't care which one, just that it doesn't crash or corrupt data)
	got, err := store.Get(key)
	if err != nil {
		t.Errorf("Get() after concurrent writes failed: %v", err)
	}
	// Value should be one of the written values (A-Z)
	if len(got) != 1 || got[0] < 'A' || got[0] > 'Z' {
		t.Errorf("Get() = %v, want a letter A-Z", got)
	}
}

// TestInMemoryStorage_EmptyKey tests behavior with empty key
func TestInMemoryStorage_EmptyKey(t *testing.T) {
	store := NewInMemoryStorage()

	// Empty key should be allowed (storage doesn't validate keys)
	store.Set("", "value")

	got, err := store.Get("")
	if err != nil {
		t.Errorf("Get() with empty key failed: %v", err)
	}
	if got != "value" {
		t.Errorf("Get() = %v, want 'value'", got)
	}
}

// TestInMemoryStorage_EmptyValue tests behavior with empty value
func TestInMemoryStorage_EmptyValue(t *testing.T) {
	store := NewInMemoryStorage()
	key := "0x742d35cc6634c0532925a3b844bc9e7595f0beb"

	// Empty value should be allowed
	store.Set(key, "")

	got, err := store.Get(key)
	if err != nil {
		t.Errorf("Get() with empty value failed: %v", err)
	}
	if got != "" {
		t.Errorf("Get() = %v, want empty string", got)
	}
}

// TestStorage_Interface ensures InMemoryStorage implements Storage interface
func TestStorage_Interface(t *testing.T) {
	var _ Storage = (*InMemoryStorage)(nil)
}
