package binary_tree

import (
	"container/list"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	stack := list.New()

	pre := math.MinInt64
	for stack.Len() > 0 || root != nil {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}

		curNode := stack.Back()
		stack.Remove(curNode)

		curVal := curNode.Value.(*TreeNode)

		if curVal.Val <= pre {
			return false
		}

		pre = curVal.Val

		root = curVal.Right

	}

	return true
}



func isValidBST3(root *TreeNode) bool {
	return NewValidBST().isValidBST2Helper(root)
}

type ValidBST struct {
	pre int
}

func NewValidBST() *ValidBST {
	return &ValidBST{pre: math.MinInt64}
}

func (r *ValidBST) isValidBST2Helper(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !r.isValidBST2Helper(root.Left) {
		return false
	}

	if root.Val <= r.pre {
		return false
	}

	r.pre = root.Val

	return r.isValidBST2Helper(root.Right)
}

func isValidBST2(root *TreeNode) bool {
	ret := isValidBSTHelper(root)
	return ret.isValidBst
}

func isValidBSTHelper(root *TreeNode) *ValidBSTRet {
	if root == nil {
		return &ValidBSTRet{
			isValidBst: true,
			max:        math.MinInt32,
			min:        math.MaxInt32,
		}
	}

	leftBST := isValidBSTHelper(root.Left)
	rightBST := isValidBSTHelper(root.Right)

	if !leftBST.isValidBst || !rightBST.isValidBst {
		return &ValidBSTRet{
			isValidBst: false,
			max:        0,
			min:        0,
		}
	}

	if root.Left != nil && leftBST.max >= root.Val {
		return &ValidBSTRet{
			isValidBst: false,
			max:        0,
			min:        0,
		}
	}

	if root.Right != nil && rightBST.min <= root.Val {
		return &ValidBSTRet{
			isValidBst: false,
			max:        0,
			min:        0,
		}
	}

	return &ValidBSTRet{
		isValidBst: true,
		max:        int(math.Max(float64(rightBST.max), float64(root.Val))),
		min:        int(math.Min(float64(leftBST.min), float64(root.Val))),
	}
}

type ValidBSTRet struct {
	isValidBst bool
	max        int
	min        int
}

func NewValidBSTRet() *ValidBSTRet {
	return &ValidBSTRet{
		isValidBst: false,
		max:        math.MinInt32,
		min:        math.MaxInt32,
	}
}