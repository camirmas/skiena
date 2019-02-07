package problems

type C3p11a [][]int

func NewC3p11a(list []int) C3p11a {
	d := make(C3p11a, len(list))
	for i, _ := range d {
		d[i] = make([]int, len(list))
	}
	for i, _ := range list {
		for j, _ := range list {
			if i < j {
				d[i][j] = min(list[i:j])
			}
		}
	}

	return d
}

func (d C3p11a) Min(i, j int) int {
	return d[i][j]
}

func min(list []int) int {
	min := list[0]

	for _, v := range list {
		if v < min {
			min = v
		}
	}

	return min
}
