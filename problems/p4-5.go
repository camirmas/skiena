package problems

import (
	"github.com/camirmas/skiena"
	// "fmt"
)

// 4-5
func C4p5(s []int) int {
	sorted := skiena.Mergesort(s)

	current := sorted[0]
	count := 0
	mode := [2]int{sorted[0], count}

	for _, v := range sorted {
		if v == current {
			count++
		} else {
			if count > mode[1] {
				mode = [2]int{current, count}
			}
			count = 1
			current = v
		}
	}

	return mode[0]
}
