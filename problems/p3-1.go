package skiena

func CheckParens(parens string) int {
	stack := []int{}

	if parens != "" && string(parens[0]) == ")" {
		return 0
	}

	for i, char := range parens {
		if string(char) == "(" {
			stack = append(stack, i)
		}
		if string(char) == ")" {
			if len(stack) == 0 {
				return i
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return -1
	}
	return stack[0]
}
