package binary_tree

func PathSum(root *TreeNode, sum int)[][]int{
	var path []int
	var result [][]int
	PathSumHelper(path, result, root, sum)
	return result
}

func PathSumHelper(path []int, result [][]int, root *TreeNode, sum int)  {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	if root.Val == sum {
		result = append(result, path)
	}

	PathSumHelper(path, result, root.Left, sum - root.Val)
	PathSumHelper(path, result, root.Right, sum - root.Val)
}

