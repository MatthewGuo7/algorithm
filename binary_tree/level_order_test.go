package binary_tree

import (
	"container/list"
)

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)

	if root == nil {
		return ret
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		size := queue.Len()

		currentLevel := make([]int, 0, size)
		for i := 0; i < size; i++ {
			front := queue.Front()
			queue.Remove(front)

			currentNode := front.Value.(*TreeNode)
			left := front.Value.(*TreeNode).Left
			if left != nil {
				queue.PushBack(left)
			}

			right := front.Value.(*TreeNode).Right
			if right != nil {
				queue.PushBack(right)
			}

			currentLevel = append(currentLevel, currentNode.Val)
		}
		ret = append(ret, currentLevel)
	}

	return ret
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)

	if root == nil {
		return ret
	}

	queue := list.New()
	queue.PushBack(root)

	level := 0
	for queue.Len() > 0 {
		currentLevel := make([]int, 0)

		size := queue.Len()

		for ; size > 0; size-- {
			top := queue.Front()
			queue.Remove(top)

			node := top.Value.(*TreeNode)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			currentLevel = append(currentLevel, node.Val)
		}

		if level % 2 == 1 {
			currentLevel = swapArr(currentLevel)
		}

		ret = append(ret, currentLevel)
		level++
	}

	return ret
}

func swapArr(values []int)[]int  {
	start := 0;
	end := len(values) - 1

	for start < end {
		values[start], values[end] = values[end], values[start]

		start++
		end--
	}

	return values
}
