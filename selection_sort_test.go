package skiena

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	res := SelectionSort([]int{2, 1, 3})
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("Expected %v, got %v", expected, res)
	}

	res = SelectionSort([]int{1, 2, 3})
	expected = []int{1, 2, 3}

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}
