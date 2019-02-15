package problems

import "testing"

func TestC4p2a(t *testing.T) {
	n := []int{2, 1, 3, 5}

	if x, y := C4p2a(n); x != 1 || y != 5 {
		t.Errorf("Expected (1, 5), got (%d, %d)", x, y)
	}
}

func TestC4p2b(t *testing.T) {
	n := []int{1, 2, 3, 5}

	if x, y := C4p2b(n); x != 1 || y != 5 {
		t.Errorf("Expected (1, 5), got (%d, %d)", x, y)
	}
}

func TestC4p2c(t *testing.T) {
	n := []int{2, 1, 3, 1, 5}

	if x, y := C4p2c(n); x != 1 || y != 2 {
		t.Errorf("Expected (1, 5), got (%d, %d)", x, y)
	}
}

func TestC4p2d(t *testing.T) {
	n := []int{1, 1, 2, 3, 5}

	if x, y := C4p2d(n); x != 1 || y != 2 {
		t.Errorf("Expected (1, 5), got (%d, %d)", x, y)
	}
}
