package skiena

func FindSubstring(text, match string) int {
	textLen := len(text)
	matchLen := len(match)

	for i := 0; i <= textLen-matchLen; i++ {
		var j int
		for j < matchLen {
			if text[i+j] != match[j] {
				break
			}
			j++
		}
		if j == matchLen {
			return i
		}
	}

	return -1
}
