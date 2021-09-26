package binary_tree


func lowestCommonAncestor4(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	leftLCA := lowestCommonAncestor4(root.Left, p, q)
	righLCA := lowestCommonAncestor4(root.Right, p, q)

	if leftLCA != nil && righLCA != nil {
		return root
	}

	if leftLCA != nil {
		return leftLCA
	}

	if righLCA != nil {
		return righLCA
	}

	return nil
}

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

func lowestCommonAncestor5(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	parents := make(map[*TreeNode]*TreeNode)
	dfs2(root, parents)

	visited := make(map[*TreeNode]struct{})

	for p != nil {
		visited[p] = struct{}{}
		p = parents[p]
	}

	for q != nil {
		if _, ok := visited[q]; ok {
			return q
		}

		q = parents[q]
	}

	return nil
}

func dfs2(root *TreeNode, parents map[*TreeNode]*TreeNode){
	if root == nil {
		return
	}

	if root.Left != nil {
		parents[root.Left] = root
	}

	if root.Right != nil {
		parents[root.Right] = root
	}

	dfs2(root.Left, parents)
	dfs2(root.Right, parents)
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

type ParentTreeNode struct {
	parent, left, right *ParentTreeNode
}


func lowestCommonAncestor3(root, p, q *ParentTreeNode) *ParentTreeNode{
	if root == nil {
		return nil
	}

	parentSet := make(map[*ParentTreeNode]struct{})
	for cur := p; cur != nil; cur = cur.parent {
		parentSet[cur] = struct{}{}
	}

	for cur := q; cur != nil; cur = cur.parent{
		if _, ok := parentSet[cur] ; ok{
			return cur
		}
	}

	return nil
}
