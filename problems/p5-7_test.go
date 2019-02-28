package problems

import "testing"

func TestC5p7(t *testing.T) {
	preorder := []int{1, 2, 4, 5, 3}
	inorder := []int{4, 2, 5, 1, 3}

	expected := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}

	tree := &TreeNode{}
	Construct(tree, preorder, inorder)
	if !equal(tree, expected) {
		t.Errorf("Failed to construct proper tree")
	}
}

func equal(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1.Val != root2.Val {
		return false
	}
	return equal(root1.Left, root2.Left) && equal(root1.Right, root2.Right)
}
