package binary_tree

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		size := len(queue)
		curLevel := make([]int, 0, size)

		for index := 0; index < size; index++ {
			first := queue[index]

			curLevel = append(curLevel, first.Val)

			if first.Left != nil {
				queue = append(queue)
			}

			if first.Right != nil {
				queue = append(queue)
			}
		}

		queue = queue[size:]
		ret = append(ret, curLevel)
	}

	return ret
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	answer := make([][]int, 0)
	if root == nil {
		return answer
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	level := 1

	for len(queue) > 0 {
		size := len(queue)
		level++

		curLevel := make([]int, 0, size)

		for index := 0; index < size; index++ {
			node := queue[index]

			if level%2 == 1 {
				curLevel = append(curLevel, node.Val)
			} else {
				curLevel = append([]int{node.Val}, curLevel...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]

		answer = append(answer, curLevel)
	}

	return answer
}
