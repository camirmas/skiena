package structures

import (
	// "fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := PriorityQueue{}

	i1 := &PQItem{1, "first"}
	i2 := &PQItem{2, "second"}
	i3 := &PQItem{3, "third"}
	i4 := &PQItem{4, "fourth"}
	pq.Push(i4)
	pq.Push(i2)
	pq.Push(i1)
	pq.Push(i3)

	if l := pq.Len(); l != 4 {
		t.Errorf("expected 4, got %d", l)
	}

	if i := pq.Pop(); i.(*PQItem).priority != 1 {
		t.Errorf("error: %v", i)
	}

	if i := pq.Pop(); i.(*PQItem).priority != 2 {
		t.Errorf("error: %v", i)
	}

	if i := pq.Pop(); i.(*PQItem).priority != 3 {
		t.Errorf("error: %v", i)
	}

	if i := pq.Pop(); i.(*PQItem).priority != 4 {
		t.Errorf("error: %v", i)
	}

	if i := pq.Pop(); i != nil {
		t.Errorf("Expected nil, got %v", i)
	}

	if l := pq.Len(); l != 0 {
		t.Errorf("expected 0, got %d", l)
	}
}
