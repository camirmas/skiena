package skiena

import (
	"testing"
)

func TestFindSubstring(t *testing.T) {
	res := FindSubstring("Son Goku", "Goku")

	if res != 4 {
		t.Errorf("Expected %d, got %d", 4, res)
	}

	res = FindSubstring("Son Goku", "Gohan")

	if res != -1 {
		t.Errorf("Expected %d, got %d", -1, res)
	}
}
