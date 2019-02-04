package skiena

import "testing"

func TestFastExponentiate(t *testing.T) {
	res := FastExponentiate(2, 3)

	if res != 8 {
		t.Errorf("Expected %d, got %d", 8, res)
	}

	res = FastExponentiate(2, 4)

	if res != 16 {
		t.Errorf("Expected %d, got %d", 16, res)
	}
}
