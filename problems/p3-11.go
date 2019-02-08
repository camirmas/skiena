package problems

import "errors"

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

type TreeC3p11b struct {
	root *Node
}

type Node struct {
	Min        int
	LowerBound int
	UpperBound int
	Left       *Node
	Right      *Node
}

func NewTreeC3p11b(list []int) TreeC3p11b {
	root := &Node{min(list), 0, len(list) - 1, nil, nil}
	t := TreeC3p11b{root}

	t.root.build(list)

	return t
}

func (node *Node) build(list []int) {
	if node == nil {
		return
	}
	if len(list) == 0 {
		return
	}
	if len(list) == 1 {
		i := list[0]
		if i > node.Min {
			node.Right = &Node{i, node.LowerBound, node.LowerBound + 1, nil, nil}
			return
		} else {
			node.Left = &Node{i, node.LowerBound - 1, node.LowerBound, nil, nil}
			return
		}
	}
	midpoint := len(list) / 2
	left, right := list[:midpoint], list[midpoint:]

	node.Left = &Node{min(left), 0, midpoint, nil, nil}
	node.Right = &Node{min(right), midpoint, len(list) - 1, nil, nil}

	node.Left.build(left)
	node.Right.build(right)
}

func (t *TreeC3p11b) Min(i, j int) (int, error) {
	return t.root.min(i, j)
}

func (node *Node) min(i, j int) (int, error) {
	if i >= j {
		return 0, errors.New("i must be < j")
	}
	if node == nil {
		return 0, errors.New("node is empty")
	}
	if i < node.LowerBound || j > node.UpperBound {
		return 0, errors.New("Out of bounds")
	} else if i > node.LowerBound {
		return node.Right.min(i, j)
	} else if j < node.UpperBound {
		return node.Left.min(i, j)
	} else {
		return node.Min, nil
	}
}
