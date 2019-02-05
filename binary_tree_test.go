package skiena

import (
	"reflect"
	"testing"
	// "fmt"
)

func TestBinaryTree(t *testing.T) {
	tree := &Tree{}

	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)

	res := tree.Search(2)

	if res == nil || res.Item != 2 {
		t.Fatalf("Expected 2, got %v", res)
	}

	res = tree.Search(20)

	if res != nil {
		t.Fatalf("Expected nil, got %v", res)
	}

	min := tree.Min()

	if min == nil || min.Item != 1 {
		t.Fatalf("Expected 1, got %v", min)
	}

	max := tree.Max()

	if max == nil || max.Item != 3 {
		t.Fatalf("Expected 3, got %v", max)
	}

	expected := []int{1, 2, 3}

	if res := tree.GetValues(); !reflect.DeepEqual(res, expected) {
		t.Fatalf("Expected %v, got %v", expected, res)
	}

	// no children
	tree.Delete(4)

	res = tree.Search(4)

	if res != nil {
		t.Fatalf("Expected nil, got %v", res)
	}

	// left child
	tree.Insert(5)
	tree.Insert(4)

	tree.Delete(5)

	res = tree.Search(5)

	if res != nil {
		t.Fatalf("Expected nil, got %v", res)
	}

	res = tree.Search(4)

	if res == nil || res.Item != 4 {
		t.Fatalf("Expected 4, got %v", res)
	}

	if res.Parent == nil || res.Parent.Item != 3 {
		t.Fatalf("Expected 3, got %v", res.Parent)
	}

	tree.Delete(4)

	// right child
	tree.Insert(5)
	tree.Insert(6)

	tree.Delete(5)

	res = tree.Search(5)

	if res != nil {
		t.Fatalf("Expected nil, got %v", res)
	}

	res = tree.Search(6)

	if res == nil || res.Item != 6 {
		t.Fatalf("Expected 4, got %v", res)
	}

	if res.Parent == nil || res.Parent.Item != 3 {
		t.Fatalf("Expected 3, got %v", res.Parent)
	}

	tree.Delete(6)

	// two children
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)

	tree.Delete(5)

	res = tree.Search(5)

	if res != nil {
		t.Fatalf("Expected nil, got %v", res)
	}

	res = tree.Search(4)

	if res == nil || res.Item != 4 {
		t.Fatalf("Expected 4, got %v", res)
	}

	if res.Parent == nil || res.Parent.Item != 6 {
		t.Fatalf("Expected 6, got nil")
	}
}
