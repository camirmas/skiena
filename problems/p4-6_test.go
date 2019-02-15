package problems

import "testing"

func TestC4p6(t *testing.T) {
	s1 := []int{2, 1, 3}
	s2 := []int{8, 0, 0}

	if sumExists := C4p6(s1, s2, 10); !sumExists {
		t.Error("Expected true, got false")
	}

	if sumExists := C4p6(s1, s2, 22); sumExists {
		t.Error("Expected false, got true")
	}
}

func TestC4p6v2(t *testing.T) {
	s1 := []int{2, 1, 3}
	s2 := []int{8, 0, 0}

	if sumExists := C4p6v2(s1, s2, 10); !sumExists {
		t.Error("Expected true, got false")
	}

	if sumExists := C4p6v2(s1, s2, 22); sumExists {
		t.Error("Expected false, got true")
	}
}
