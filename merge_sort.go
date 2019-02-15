package skiena

// import "fmt"

func Mergesort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	middle := len(s) / 2
	return merge(Mergesort(s[:middle]), Mergesort(s[middle:]))
}

func merge(left, right []int) []int {
	i := 0
	s := make([]int, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			var x int
			x, left = left[0], left[1:]
			s[i] = x
			i++
		} else {
			var x int
			x, right = right[0], right[1:]
			s[i] = x
			i++
		}
	}

	for len(left) > 0 {
		var x int
		x, left = left[0], left[1:]
		s[i] = x
		i++
	}
	for len(right) > 0 {
		var x int
		x, right = right[0], right[1:]
		s[i] = x
		i++
	}
	return s
}
