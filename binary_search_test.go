package skiena

import "testing"

func TestBinarySearch(t *testing.T) {
	s := []int{1, 2, 3}

	if i := BinarySearch(s, 1); i != 0 {
		t.Errorf("Expected 0, got %d", i)
	}
	if i := BinarySearch(s, 2); i != 1 {
		t.Errorf("Expected 1, got %d", i)
	}
	if i := BinarySearch(s, 3); i != 2 {
		t.Errorf("Expected 2, got %d", i)
	}
	if i := BinarySearch(s, 4); i != -1 {
		t.Errorf("Expected -1, got %d", i)
	}
}
