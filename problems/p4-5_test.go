package problems

import "testing"

func TestC4p5(t *testing.T) {
	s := []int{4, 6, 2, 4, 3, 1}

	if mode := C4p5(s); mode != 4 {
		t.Errorf("Expected 4, got %d", mode)
	}
}
