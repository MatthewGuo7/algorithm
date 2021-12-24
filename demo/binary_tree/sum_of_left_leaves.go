package binary_tree

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	answer := 0

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if front != nil && front.Left == nil && front.Right == nil {
			answer += front.Val
		}

		if front != nil && front.Left != nil {
			queue = append(queue, front.Left)
		}

		if front != nil && front.Right != nil {
			queue = append(queue, front.Right)
		}
	}

	return answer
}

func sumOfLeftLeaves1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	sumOfLeftLeavesHelper(root, &sum)

	return sum
}

func sumOfLeftLeavesHelper(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	sumOfLeftLeavesHelper(root.Left, sum)
	sumOfLeftLeavesHelper(root.Right, sum)

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*sum += root.Val
	}
}
