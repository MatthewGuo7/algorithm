package binary_tree

import (
	"container/list"
	"strconv"
)

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

func binaryTreePaths2(root *TreeNode) []string {
	paths := make([]string, 0)
	path := list.New()
	binaryTreePathsHelper(root, path, &paths)

	return paths
}

func binaryTreePathsHelper(root *TreeNode, path *list.List, paths *[]string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		retPath := ""
		for cur := path.Front(); cur != nil; cur = cur.Next(){
			curNode := cur.Value.(*TreeNode)
			retPath += strconv.Itoa(curNode.Val)
			if cur.Next() != nil {
				retPath += "->"
			}
		}
		*paths = append(*paths, retPath)
	}

	path.PushBack(root.Left)
	last := path.Back()
	binaryTreePathsHelper(root.Left, path, paths)
	path.Remove(last)

	path.PushBack(root.Right)
	last = path.Back()
	binaryTreePathsHelper(root.Right, path, paths)
	path.Remove(last)
}