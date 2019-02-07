package skiena

type C3p2 []int

func NewC3p2(size int) C3p2 {
	return make(C3p2, size)
}

func (d C3p2) Search(i int) int {
	return d[i]
}

func (d C3p2) Insert(i int) {
	d[i] = 1
}

func (d C3p2) Delete(i int) {
	d[i] = 0
}
