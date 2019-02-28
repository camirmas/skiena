package problems

// import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Construct(root *TreeNode, preorder, inorder []int) {
	var v int
	v, preorder = preorder[0], preorder[1:]

	root.Val = v

	idx := getIdx(inorder, v)
	left, right := inorder[:idx], inorder[idx+1:]

	if len(left) == 1 {
		root.Left = &TreeNode{left[0], nil, nil}
	} else if len(left) > 0 {
		root.Left = &TreeNode{}
		Construct(root.Left, preorder, left)
	}

	if len(right) == 1 {
		root.Right = &TreeNode{right[0], nil, nil}
	} else if len(right) > 0 {
		root.Right = &TreeNode{}
		Construct(root.Right, preorder, right)
	}
}

func getIdx(items []int, item int) int {
	for i, v := range items {
		if v == item {
			return i
		}
	}

	return -1
}
