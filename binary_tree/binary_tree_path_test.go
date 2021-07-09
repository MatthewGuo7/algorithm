package binary_tree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	retPaths := make([]string, 0)
	if root == nil {
		return retPaths
	}

	if (root.Left == nil) && (root.Right == nil) {
		retPaths = append(retPaths, strconv.Itoa(root.Val))
	}

	leftPaths := binaryTreePaths(root.Left)
	rightPaths := binaryTreePaths(root.Right)

	for _, path := range leftPaths {
		path = strconv.Itoa(root.Val) + "->" + path
		retPaths = append(retPaths, path)
	}

	for _, path := range rightPaths {
		path = strconv.Itoa(root.Val) + "->" + path
		retPaths = append(retPaths, path)
	}

	return retPaths
}
