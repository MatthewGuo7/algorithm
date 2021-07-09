package binary_tree

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	
	if root == nil {
		return false
	}
	
	if root.Val == subRoot.Val {
		return isMatchedTree(root, subRoot)
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isMatchedTree(root1 *TreeNode, root2 *TreeNode) bool  {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return isMatchedTree(root1.Left, root2.Left) && isMatchedTree(root1.Right, root2.Right)
}
