package skiena

func shitExponentiate(base, power int) int {
	if power == 0 {
		return 1
	}
	return base * shitExponentiate(base, power-1)
}

func FastExponentiate(base, power int) int {
	if power == 0 {
		return 1
	}
	x := FastExponentiate(base, power/2)
	if power%2 == 0 {
		return x * x
	} else {
		return base * x * x
	}
}
