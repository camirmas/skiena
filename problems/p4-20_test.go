package problems

import (
	"reflect"
	"testing"
)

func TestC4p20(t *testing.T) {
	s := []int{-1, 0, 1, -2, -3, 2}
	expected := []int{-1, -2, -3, 2, 1, 0}

	if re := C4p20(s); !reflect.DeepEqual(re, expected) {
		t.Errorf("Expected %v, got %v", expected, re)
	}
}
