package skiena

func InsertionSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		j := i

		for j > 0 && s[j] < s[j-1] {
			s[j], s[j-1] = s[j-1], s[j]
			j--
		}
	}

	return s
}
