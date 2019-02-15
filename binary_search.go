package skiena

// import "fmt"

func BinarySearch(s []int, key int) int {
	return binarySearchHelper(s, key, 0, len(s)-1)
}

func binarySearchHelper(s []int, key, low, high int) int {
	if low > high {
		return -1
	}
	middle := (low + high) / 2

	if s[middle] == key {
		return middle
	}
	if s[middle] < key {
		return binarySearchHelper(s, key, middle+1, high)
	} else {
		return binarySearchHelper(s, key, low, middle-1)
	}
}
