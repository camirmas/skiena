package structures

// import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	Item   int
	Left   *Node
	Right  *Node
	Parent *Node
}

func NewTree() Tree {
	return Tree{}
}

func (t *Tree) Insert(item int) {
	if t == nil {
		return
	}

	if t.root == nil {
		newTree := &Node{item, nil, nil, nil}
		t.root = newTree
	} else {
		t.root.insert(item)
	}
}

func (t *Tree) Delete(item int) {
	if t == nil || t.root == nil {
		return
	}
	t.root.delete(item)
}

func (node *Node) delete(item int) {
	if node == nil {
		return
	}
	if item > node.Item {
		node.Right.delete(item)
	} else if item < node.Item {
		node.Left.delete(item)
	} else {
		if node.Left == nil && node.Right == nil {
			if node.Parent.Left == node {
				node.Parent.Left = nil
			} else {
				node.Parent.Right = nil
			}
		} else if node.Left == nil && node.Right != nil {
			node.replace(node.Right)
		} else if node.Left != nil && node.Right == nil {
			node.replace(node.Left)
		} else {
			min := node.Right.Min()
			min.Left = node.Left
			min.Left.Parent = min
			node.replace(min)
		}
	}
}

func (node *Node) replace(replace *Node) {
	if node.Parent.Left == node {
		node.Parent.Left = replace
	} else {
		node.Parent.Right = replace
	}
	replace.Parent = node.Parent
}

func (node *Node) insert(item int) {
	if item > node.Item {
		if node.Right == nil {
			newNode := &Node{item, nil, nil, node}
			node.Right = newNode
		} else {
			node.Right.insert(item)
		}
	} else {
		if node.Left == nil {
			newNode := &Node{item, nil, nil, node}
			node.Left = newNode
		} else {
			node.Left.insert(item)
		}
	}
}

func (t *Tree) Search(item int) *Node {
	if t.root == nil {
		return nil
	}

	return t.root.search(item)
}

func (node *Node) search(item int) *Node {
	if node == nil {
		return nil
	}
	if node.Item == item {
		return node
	}

	if item < node.Item {
		return node.Left.search(item)
	} else {
		return node.Right.search(item)
	}
}

func (t *Tree) Min() *Node {
	if t.root == nil {
		return nil
	}

	return t.root.Min()
}

func (node *Node) Min() *Node {
	if node == nil {
		return nil
	}
	min := node

	for min.Left != nil {
		if min.Left.Item < min.Item {
			min = min.Left
		}
	}

	return min
}

func (t *Tree) Max() *Node {
	if t.root == nil {
		return nil
	}

	max := t.root

	for max.Right != nil {
		if max.Right.Item > max.Item {
			max = max.Right
		}
	}

	return max
}

func (t *Tree) Traverse(handler func(int)) {
	if t == nil {
		return
	}

	t.root.traverse(handler)
}

func (node *Node) traverse(handler func(int)) {
	if node != nil {
		node.Left.traverse(handler)
		handler(node.Item)
		node.Right.traverse(handler)
	}
}

func (t *Tree) GetValues() []int {
	items := &[]int{}

	f := func(items *[]int) func(int) {
		return func(i int) {
			*items = append(*items, i)
		}
	}

	t.Traverse(f(items))

	return *items
}
