package skiena

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	s := []int{2, 1, 3}
	expected := []int{1, 2, 3}

	if res := Quicksort(s); !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v, got %v", expected, s)
	}
}
