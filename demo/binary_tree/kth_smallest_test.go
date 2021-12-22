package binary_tree

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0)
	count := 0

	for len(stack) > 0 || root != nil {
		for root.Left != nil {
			stack = append(stack, root.Left)
			root = root.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		count++

		if count == k {
			return top.Val
		}

		root = root.Right
	}

	return 0
}
