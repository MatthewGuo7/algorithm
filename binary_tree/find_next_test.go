package binary_tree

func FindNext(node *TreeNode, root *TreeNode) *TreeNode  {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		return MostLeftNode(root)
	}

	curNode := root
	for curNode != nil {
		if curNode.Val > node.Val {
			curNode = curNode.Left
		} else {
			curNode = curNode.Right
		}
	}

	return curNode
}

func MostLeftNode(root *TreeNode) *TreeNode  {
	if root == nil {
		return nil
	}

	for root.Left != nil {
		root = root.Left
	}

	return root
}
