package binary_tree

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	rootValue := postorder[len(postorder)-1]
	root := &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   rootValue,
	}

	for pos, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(inorder[:pos], postorder[:pos])
			root.Right = buildTree(inorder[pos+1:], postorder[pos:len(postorder)])
		}
	}

	return root
}
