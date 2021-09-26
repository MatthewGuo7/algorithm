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

func binaryTreePaths3(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	stack := list.New()
	stack.PushBack(root)
	ret := make([]string,0)
	binaryTreePathsHelper2(root, stack, &ret)

	return ret
}

func binaryTreePathsHelper2(root *TreeNode, stack *list.List, paths *[]string)  {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		path := ""
		for cur := stack.Front(); cur != nil; cur = cur.Next() {
			curValue := cur.Value.(*TreeNode)
			path += strconv.Itoa(curValue.Val)
			if cur.Next() != nil {
				path += "->"
			}
		}
		*paths = append(*paths, path)
	}

	stack.PushBack(root.Left)
	binaryTreePathsHelper2(root.Left, stack, paths)
	stack.Remove(stack.Back())


	stack.PushBack(root.Right)
	binaryTreePathsHelper2(root.Right, stack, paths)
	stack.Remove(stack.Back())
}



func binaryTreePaths4(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	ret := make([]string,0)
	binaryTreePathsHelper4(root, "", &ret)

	return ret
}

func binaryTreePathsHelper4(root *TreeNode, path string, paths *[]string)  {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path)
		return
	}

	if root.Left != nil {
		binaryTreePathsHelper4(root.Left, path + "->" + strconv.Itoa(root.Left.Val), paths)
	}

	if root.Right != nil {
		binaryTreePathsHelper4(root.Right, path + "->" + strconv.Itoa(root.Left.Val), paths)
	}
}
