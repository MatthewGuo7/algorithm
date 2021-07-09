package binary_tree

import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)

	return int(math.Max(float64(leftMax), float64(rightMax)) + 1)
}

type MaxDepth struct {
	MaxDepthRet int
}

func (r *MaxDepth) GetMaxDepth(root *TreeNode) {
	r.GetMaxDepthHelper(root, 0)
}

func (r *MaxDepth) GetMaxDepthHelper(root *TreeNode, curDepth int) {
	if root == nil {
		return
	}

	if curDepth > r.MaxDepthRet {
		r.MaxDepthRet = curDepth
	}

	r.GetMaxDepthHelper(root.Left, curDepth+1)
	r.GetMaxDepthHelper(root.Right, curDepth+1)

}
