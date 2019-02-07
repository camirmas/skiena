package skiena

import "testing"

func TestC3p1(t *testing.T) {
	if res := CheckParens(")()("); res != 0 {
		t.Errorf("Expected error at position 0")
	}
	if res := CheckParens("(()"); res != 0 {
		t.Errorf("Expected error at position 0, got %d", res)
	}
	if res := CheckParens("())"); res != 2 {
		t.Errorf("Expected error at position 3, got %d", res)
	}

	if res := CheckParens("((())())()"); res != -1 {
		t.Errorf("Expected no errors, got %d", res)
	}
}
