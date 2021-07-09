package binary_tree

type MinimumSubTree struct {
	retSubTree *TreeNode
	retSum     int
}

func (r *MinimumSubTree) FindMinimumSubTree(root *TreeNode) *TreeNode {
	r.findMinimumSubTreeHelper(root)
	return r.retSubTree
}

func (r *MinimumSubTree) findMinimumSubTreeHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := r.findMinimumSubTreeHelper(root.Left) + r.findMinimumSubTreeHelper(root.Right) + root.Val

	if sum < r.retSum {
		r.retSum = sum
		r.retSubTree = root
	}

	return sum
}
