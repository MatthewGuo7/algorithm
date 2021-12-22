package binary_tree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flattenHelper(root)
}

func flattenHelper(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftLast := flattenHelper(root.Left)
	rightLast := flattenHelper(root.Right)

	if leftLast != nil {
		leftLast.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	if rightLast != nil {
		return rightLast
	}

	if leftLast != nil {
		return leftLast
	}

	return root
}
