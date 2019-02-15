package problems

import "github.com/camirmas/skiena"

// 4-6
// O(nlogn)
func C4p6(s1, s2 []int, sum int) bool {
	sorted := skiena.Mergesort(s1)
	unsorted := s2

	for _, v := range unsorted {
		toSearch := sum - v

		if i := skiena.BinarySearch(sorted, toSearch); i != -1 {
			return true
		}
	}
	return false
}

// 4-6 (other)
// O(n)
func C4p6v2(s1, s2 []int, sum int) bool {
	seen := make(map[int]bool, 0)

	for i, _ := range s1 {
		toSearch := sum - s1[i]
		if _, ok := seen[toSearch]; ok {
			return true
		} else {
			seen[s1[i]] = true
		}

		toSearch = sum - s2[i]
		if _, ok := seen[toSearch]; ok {
			return true
		} else {
			seen[s2[i]] = true
		}
	}
	return false
}
