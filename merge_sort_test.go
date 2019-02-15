package skiena

import (
	"reflect"
	"testing"
)

func TestMergesort(t *testing.T) {
	s := []int{2, 1, 3}
	expected := []int{1, 2, 3}

	if res := Mergesort(s); !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v, got %v", expected, s)
	}
}
