package problems

import "testing"

func TestC3p11a(t *testing.T) {
	d := NewC3p11a([]int{1, 2, 3})

	if min := d.Min(0, 2); min != 1 {
		t.Errorf("Expected 1, got %d", min)
	}

	if min := d.Min(1, 2); min != 2 {
		t.Errorf("Expected 2, got %d", min)
	}
}

func TestC3p11b(t *testing.T) {
	d := NewTreeC3p11b([]int{1, 2, 3})

	if min := d.Min(0, 2); min != 1 {
		t.Errorf("Expected 1, got %d", min)
	}

	if min := d.Min(1, 2); min != 2 {
		t.Errorf("Expected 2, got %d", min)
	}

	d = NewTreeC3p11b([]int{1, 3, 2})

	if min := d.Min(0, 2); min != 1 {
		t.Errorf("Expected 1, got %d", min)
	}

	if min := d.Min(1, 2); min != 2 {
		t.Errorf("Expected 2, got %d", min)
	}
}
