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

type TreeC3p11b struct {
	root *Node
	list []int
}

type Node struct {
	Key    int
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func NewTreeC3p11b(list []int) TreeC3p11b {
	t := TreeC3p11b{nil, list}

	for idx, key := range list {
		t.Insert(key, idx)
	}

	return t
}

func (t *TreeC3p11b) Insert(key, val int) {
	if t == nil {
		return
	}
	if t.root == nil {
		t.root = &Node{key, val, nil, nil, nil}
	} else {
		t.root.insert(key, val)
	}
}

func (node *Node) insert(key, val int) {
	if node == nil {
		return
	}
	if key > node.Key {
		if node.Right == nil {
			node.Right = &Node{key, val, nil, nil, node}
		} else {
			node.Right.insert(key, val)
		}
	} else {
		if node.Left == nil {
			node.Left = &Node{key, val, nil, nil, node}
		} else {
			node.Left.insert(key, val)
		}
	}
}

func (t *TreeC3p11b) Search(key int) *Node {
	if t == nil || t.root == nil {
		return nil
	}
	return t.root.search(key)
}

func (node *Node) search(key int) *Node {
	if node == nil {
		return nil
	}

	if node.Key == key {
		return node
	}
	if key < node.Key {
		return node.Left.search(key)
	} else {
		return node.Right.search(key)
	}
}

func (t *TreeC3p11b) Min(i, j int) int {
	iNode := t.Search(t.list[i])

	min := iNode
	current := iNode

	for current != nil {
		if current.Left != nil && current.Val <= j {
			min = current.Left
		}
		current = current.Left
	}

	return min.Key
}
