package structures

import (
	"testing"
	// "fmt"
)

func TestHashTable(t *testing.T) {
	table := make(HashTable, 2^32)

	table.Insert("hello", 1234)

	if res := table.Get("hello"); res == nil {
		t.Errorf("Expected 1234, got nil")
	}

	if res := table.Get("hello123"); res != nil {
		t.Errorf("Expected nil, got %v", res)
	}

	table.Delete("hello")

	if res := table.Get("hello"); res != nil {
		t.Errorf("Expected nil, got %v", res)
	}
}
