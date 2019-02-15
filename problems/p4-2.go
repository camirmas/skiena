package problems

import "github.com/camirmas/skiena"

// 4-2 (a)
func C4p2a(n []int) (x, y int) {
	max := n[0]
	min := n[0]

	for _, v := range n {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return min, max
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 4-2 (b)
func C4p2b(sorted []int) (x, y int) {
	return sorted[0], sorted[len(sorted)-1]
}

// 4-2 (c)
func C4p2c(s []int) (x, y int) {
	sorted := skiena.Mergesort(s)

	min1 := sorted[0]

	for _, v := range sorted {
		if v != min1 {
			return min1, v
		}
	}

	return min1, sorted[1]
}

// 4-2 (d)
func C4p2d(sorted []int) (x, y int) {
	min1 := sorted[0]

	for _, v := range sorted {
		if v != min1 {
			return min1, v
		}
	}

	return min1, sorted[1]
}
