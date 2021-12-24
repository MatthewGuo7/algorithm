package binary_tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricHelper(root, root)
}

func isSymmetricHelper(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSymmetricHelper(p.Left, q.Right) && isSymmetricHelper(p.Right, q.Left)
}






func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := make([]*TreeNode,0)
	queue = append(queue, root, root)

	for len(queue) > 0 {
		p, q := queue[0], queue[1]
		queue = queue[2:]

		if p == nil && q == nil {
			continue
		}

		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		queue = append(queue, p.Left, q.Right)
		queue = append(queue, p.Right, q.Left)
	}

	return true
}


