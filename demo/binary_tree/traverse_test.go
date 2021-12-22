package binary_tree

func Inorder(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		for root.Left != nil {
			stack = append(stack, root.Left)
			root = root.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ret = append(ret, top.Val)

		root = root.Right
	}

	return ret
}
