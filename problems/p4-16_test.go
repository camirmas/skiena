package problems

import "testing"

func TestC4p16(t *testing.T) {
	s := []int{2, 3, 1, 4, 5}

	if med := C4p16(s); med != 3 {
		t.Errorf("Expected median 3, got %d", med)
	}
}
