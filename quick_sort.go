package skiena

// import "fmt"

func Quicksort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	p := partition(s)
	Quicksort(s[:p])
	Quicksort(s[p+1:])

	return s
}

func partition(s []int) int {
	p := len(s) - 1

	firstHigh := 0
	for i, v := range s {
		if v < s[p] {
			s[i], s[firstHigh] = s[firstHigh], s[i]
			firstHigh++
		}
	}
	s[p], s[firstHigh] = s[firstHigh], s[p]

	return firstHigh
}
