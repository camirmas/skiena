package problems

// 4-16
func C4p16(s []int) int {
	return medianHelper(s, len(s)/2)
}

func medianHelper(s []int, k int) int {
	if len(s) < 2 {
		return s[0]
	}
	p := partition(s)
	larger := s[p+1:]

	if len(larger) == k-1 {
		return s[p]
	} else if len(larger) >= k {
		return medianHelper(larger, k)
	} else {
		smaller := s[:p]
		nk := len(smaller) / 2
		return medianHelper(smaller, nk)
	}
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
