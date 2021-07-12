package binary2

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BalancedTreeRet struct {
	isBalanced bool
	depth      int
}

func isBalanced(root *TreeNode) bool {
	return isBalanced(root)
}

func isBalancedHelper(root *TreeNode) BalancedTreeRet {
	if root == nil {
		return BalancedTreeRet{
			isBalanced: true,
			depth:      0,
		}
	}

	leftRet := isBalancedHelper(root.Left)
	rightRet := isBalancedHelper(root.Right)

	if leftRet.isBalanced == false || rightRet.isBalanced == false {
		return BalancedTreeRet{
			isBalanced: false,
			depth:      0,
		}
	}

	if math.Abs(float64(leftRet.depth)-float64(rightRet.depth)) > 1 {
		return BalancedTreeRet{
			isBalanced: false,
			depth:      0,
		}
	}

	return BalancedTreeRet{
		isBalanced: true,
		depth:      int(math.Max(float64(leftRet.depth), float64(rightRet.depth))) + 1,
	}
}
