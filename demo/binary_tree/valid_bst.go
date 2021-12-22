package binary_tree

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	pre := math.MinInt32

	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		for root.Left != nil {
			stack = append(stack, root.Left)
			root = root.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.Val <= pre {
			return false
		}

		pre = top.Val
		root = root.Right
	}

	return true
}

func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return isValidBSTHelper(root, math.MaxInt32)
}

func isValidBSTHelper(root *TreeNode, pre int) bool {
	if root == nil {
		return false
	}

	if !isValidBSTHelper(root.Left, pre) {
		return false
	}

	if root.Left.Val <= pre {
		return false
	}

	pre = root.Left.Val

	return isValidBSTHelper(root.Right, pre)
}

type ValidBSTRet struct {
	isValid bool
	max     int
	min     int
}

func isValidBST3(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return isValidBST3Helper(root).isValid
}

func isValidBST3Helper(root *TreeNode) ValidBSTRet {
	if root == nil {
		return ValidBSTRet{
			isValid: true,
			max:     math.MinInt32,
			min:     math.MaxInt32,
		}
	}

	leftRet := isValidBST3Helper(root.Left)
	rightRet := isValidBST3Helper(root.Right)

	if !leftRet.isValid || !rightRet.isValid {
		return ValidBSTRet{
			isValid: false,
			max:     0,
			min:     0,
		}
	}

	if root.Left != nil && leftRet.max >= root.Val {
		return ValidBSTRet{
			isValid: false,
			max:     0,
			min:     0,
		}
	}

	if root.Right != nil && rightRet.min < root.Val {
		return ValidBSTRet{
			isValid: false,
			max:     0,
			min:     0,
		}
	}

	return ValidBSTRet{
		isValid: true,
		max:     int(math.Max(float64(root.Val), float64(rightRet.max))),
		min:     int(math.Min(float64(root.Val), float64(leftRet.min))),
	}
}
