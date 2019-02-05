package skiena

import "testing"

func TestLinkedList(t *testing.T) {
	l := &List{}

	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)

	res := l.Search(1)

	if res.Item == nil {
		t.Errorf("Expected 1, got nil")
	}

	res = l.Search(2)

	if res.Item == nil {
		t.Errorf("Expected 2, got nil")
	}

	res = l.Search(3)

	if res.Item == nil {
		t.Errorf("Expected 3, got nil")
	}

	res = l.Search(4)

	if res.Item == nil {
		t.Errorf("Expected 4, got nil")
	}

	l.Delete(4)
	l.Delete(3)
	l.Delete(2)
	l.Delete(1)

	if l.Item != nil {
		t.Errorf("Expected no more values")
	}
}
