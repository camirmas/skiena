package problems

import "testing"

func TestC3p2(t *testing.T) {
	d := NewC3p2(10)

	d.Insert(1)
	d.Insert(2)
	d.Insert(3)

	if res := d.Search(1); res == 0 {
		t.Errorf("Expected key %d to exist", 1)
	}
	if res := d.Search(2); res == 0 {
		t.Errorf("Expected key %d to exist", 2)
	}
	if res := d.Search(3); res == 0 {
		t.Errorf("Expected key %d to exist", 3)
	}

	d.Delete(1)
	d.Delete(2)
	d.Delete(3)

	if res := d.Search(1); res == 1 {
		t.Errorf("Expected key %d to not exist", 1)
	}
	if res := d.Search(2); res == 1 {
		t.Errorf("Expected key %d to not exist", 2)
	}
	if res := d.Search(3); res == 1 {
		t.Errorf("Expected key %d to not exist", 3)
	}

}
