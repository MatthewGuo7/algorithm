package binary2

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}

	flattenHelper(root)
}

func flattenHelper(root *TreeNode) *TreeNode  {
	if root == nil {
		return nil
	}

	leftLast := flattenHelper(root.Left)
	rightLast := flattenHelper(root.Right)

	if root.Left != nil {
		leftLast.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	if rightLast != nil {
		return rightLast
	} else if leftLast != nil {
		return leftLast
	}

	return root
}
