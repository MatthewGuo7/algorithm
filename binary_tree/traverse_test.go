package binary_tree

import (
	"container/list"
	"testing"
)

type PreOrder struct {
}

func (r *PreOrder) Traverse(root *TreeNode) []int {
	ret := make([]int, 0)
	r.TraverseHelper(root, &ret)

	return ret
}

//递归的定义
func (r *PreOrder) TraverseHelper(root *TreeNode, result *[]int) {

	//递归的出口
	if root == nil {
		return
	}

	//递归的拆解
	*result = append(*result, root.Val)
	r.TraverseHelper(root.Left, result)
	r.TraverseHelper(root.Right, result)
}

func (r *PreOrder) Travse2(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return nil
	}

	leftResult := r.Travse2(root.Left)
	rightResult := r.Travse2(root.Right)

	ret = append(ret, root.Val)
	ret = append(ret, leftResult...)
	ret = append(ret, rightResult...)

	return ret
}

func Inorder(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	stack := list.New()
	cur := root
	for stack.Len() > 0 || cur != nil {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			top := stack.Back()
			stack.Remove(top)

			ret = append(ret, top.Value.(*TreeNode).Val)

			right := top.Value.(*TreeNode).Right
			cur = right
		}
	}
	return ret
}

func PreOrder2(root *TreeNode) []int  {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		top := stack.Back()
		stack.Remove(top)

		node := top.Value.(*TreeNode)
		ret = append(ret, top.Value.(*TreeNode).Val)

		if node.Right != nil{
			stack.PushBack(node.Right)
		}

		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}

	return ret

}

// N叉树前序遍历
 type Node struct {
     Val int
     Children []*Node
 }

func preorder(root *Node) []int {
	ret := make([]int, 0)
	if root == nil {
		return nil
	}

	stack := make([]*Node,0)
	stack = append(stack, root)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		ret = append(ret, top.Val)
		stack = stack[:len(stack)-1]

		for index := len(top.Children)-1; index >= 0; index--{
			stack = append(stack, top.Children[index])
		}
	}

	return ret
}

func preorderHelper2(root *Node, answer *[]int) {
	if root == nil {
		return
	}

	*answer = append(*answer, root.Val)

	for _, child := range root.Children {
		preorderHelper2(child, answer)
	}
}

func TestTraveser(t *testing.T) {

}
