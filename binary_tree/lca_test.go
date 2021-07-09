package binary_tree


func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil || root.Right == nil {
		return root
	}

	lcaLeft := lowestCommonAncestor2(root.Left, p, q)
	lcaRight := lowestCommonAncestor2(root.Right, p,q)

	if lcaLeft != nil && lcaRight != nil {
		return root
	}

	if lcaLeft != nil {
		return lcaLeft
	}

	if lcaRight != nil {
		return lcaRight
	}

	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parents := make(map[*TreeNode]*TreeNode)
	dfs(root, parents)

	visited := make(map[*TreeNode]struct{})

	for p != nil {
		visited[p] = struct{}{}
		p = parents[p]
	}

	for q != nil {
		if _, ok := visited[q] ; ok {
			return q
		}

		q = parents[q]
	}
	return nil
}

func dfs(root *TreeNode, ret map[*TreeNode]*TreeNode) *TreeNode{
	ret = make(map[*TreeNode]*TreeNode)

	if root == nil {
		return nil
	}

	left := dfs(root.Left, ret)
	right := dfs(root.Right, ret)

	if left != nil {
		ret[left] = root
	}

	if right != nil {
		ret[right] = root
	}

	return nil
}