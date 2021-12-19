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

func maxDepth2(root *TreeNode) int {
	answer := 0
	maxDepth2Helper(root, 1, &answer)

	return answer
}

func maxDepth2Helper(root *TreeNode, curDepth int, answer *int) {
	if root == nil {
		return
	}

	if curDepth > *answer {
		*answer = curDepth
	}

	curDepth++
	maxDepth2Helper(root.Left, curDepth, answer)
	curDepth--

	curDepth++
	maxDepth2Helper(root.Right, curDepth, answer)
	curDepth--
}
