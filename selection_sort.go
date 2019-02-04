package skiena

func SelectionSort(s []int) []int {
	for i, _ := range s {
		var min = i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}

	return s
}
