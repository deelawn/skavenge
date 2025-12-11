package main

import (
	"errors"
	"sync"
	"testing"
)

// TestInMemoryStorage_Set tests the Set method of InMemoryStorage
func TestInMemoryStorage_Set(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		wantError error
	}{
		{
			name:      "set new key-value pair",
			key:       "0x742d35cc6634c0532925a3b844bc9e7595f0beb",
			value:     "skv1abc123",
			wantError: nil,
		},
		{
			name:      "set another key-value pair",
			key:       "0x1234567890abcdef",
			value:     "skv1xyz789",
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewInMemoryStorage()
			err := store.Set(tt.key, tt.value)

			if !errors.Is(err, tt.wantError) {
				t.Errorf("Set() error = %v, wantError %v", err, tt.wantError)
			}

			// Verify the value was actually stored
			if err == nil {
				got, err := store.Get(tt.key)
				if err != nil {
					t.Errorf("Get() after Set() failed: %v", err)
				}
				if got != tt.value {
					t.Errorf("Get() = %v, want %v", got, tt.value)
				}
			}
		})
	}
}

// TestInMemoryStorage_Set_DuplicateKey tests that Set returns error for duplicate keys
func TestInMemoryStorage_Set_DuplicateKey(t *testing.T) {
	store := NewInMemoryStorage()
	key := "0x742d35cc6634c0532925a3b844bc9e7595f0beb"
	value1 := "skv1abc123"
	value2 := "skv1xyz789"

	// First set should succeed
	err := store.Set(key, value1)
	if err != nil {
		t.Fatalf("First Set() failed: %v", err)
	}

	// Second set with same key should fail
	err = store.Set(key, value2)
	if !errors.Is(err, ErrKeyExists) {
		t.Errorf("Set() with duplicate key error = %v, want ErrKeyExists", err)
	}

	// Verify the original value is still stored
	got, err := store.Get(key)
	if err != nil {
		t.Fatalf("Get() failed: %v", err)
	}
	if got != value1 {
		t.Errorf("Get() = %v, want %v (original value should be preserved)", got, value1)
	}
}

// TestInMemoryStorage_Get tests the Get method of InMemoryStorage
func TestInMemoryStorage_Get(t *testing.T) {
	store := NewInMemoryStorage()
	key := "0x742d35cc6634c0532925a3b844bc9e7595f0beb"
	value := "skv1abc123"

	// Set a value first
	err := store.Set(key, value)
	if err != nil {
		t.Fatalf("Set() failed: %v", err)
	}

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
			err := store.Set(key, value)
			if err != nil {
				t.Errorf("Concurrent Set() failed for key %s: %v", key, err)
			}
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
	err := store.Set(key, value)
	if err != nil {
		t.Fatalf("Set() failed: %v", err)
	}

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

// TestInMemoryStorage_EmptyKey tests behavior with empty key
func TestInMemoryStorage_EmptyKey(t *testing.T) {
	store := NewInMemoryStorage()

	// Empty key should be allowed (storage doesn't validate keys)
	err := store.Set("", "value")
	if err != nil {
		t.Errorf("Set() with empty key failed: %v", err)
	}

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
	err := store.Set(key, "")
	if err != nil {
		t.Errorf("Set() with empty value failed: %v", err)
	}

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
