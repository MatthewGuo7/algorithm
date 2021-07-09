package binary_tree

import "math"


func isBalanced(root *TreeNode) bool {
	return isBalancedHelper(root).isBalanced
}

type BalancedTreeRet struct {
	depth      int
	isBalanced bool
}

func isBalancedHelper(root *TreeNode) BalancedTreeRet {
	if root == nil {
		return BalancedTreeRet{0, true}
	}

	leftRet := isBalancedHelper(root.Left)
	rightRet := isBalancedHelper(root.Right)

	if leftRet.isBalanced == false || rightRet.isBalanced == false {
		return BalancedTreeRet{
			depth:      -1,
			isBalanced: false,
		}
	}

	if math.Abs(float64(leftRet.depth-rightRet.depth)) > 1 {
		return BalancedTreeRet{
			depth:      -1,
			isBalanced: false,
		}
	}

	depth := int(math.Max(float64(leftRet.depth), float64(leftRet.depth)) + 1)
	return BalancedTreeRet{
		depth:      depth,
		isBalanced: true,
	}
}
