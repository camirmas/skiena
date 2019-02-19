package problems

func C4p20(s []int) []int {
	p := len(s) - 1

	firstHigh := 0
	for i, v := range s {
		if v < 0 {
			s[i], s[firstHigh] = s[firstHigh], s[i]
			firstHigh++
		}
	}
	s[p], s[firstHigh] = s[firstHigh], s[p]

	return s
}
