package skiena

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	res := InsertionSort([]int{2, 1, 3})
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("Expected %v, got %v", expected, res)
	}

	res = InsertionSort([]int{1, 2, 3})
	expected = []int{1, 2, 3}

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}
