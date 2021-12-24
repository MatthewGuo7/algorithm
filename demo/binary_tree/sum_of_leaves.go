package binary_tree

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	answer := 0
	sumOfNumbersDFS(root, 0, &answer)

	return answer
}

func sumOfNumbersDFS(root *TreeNode, sum int, answer *int) {
	if root == nil {
		return
	}

	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*answer += sum
	}

	sumOfNumbersDFS(root.Left, sum, answer)
	sumOfNumbersDFS(root.Right, sum, answer)
}
