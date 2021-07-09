package binary_tree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		root := &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   nums[0],
		}
		return root
	}

	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}

func sortedArrayToBSTHelper(nums []int, start, end int) *TreeNode  {
	if start > end {
		return nil
	}

	if start == end {
		return &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   nums[start],
		}
	}

	mid := end - (end-start)/2
	root := &TreeNode{
		Val:   nums[mid],
	}

	root.Left = sortedArrayToBSTHelper(nums, start, mid-1)
	root.Right = sortedArrayToBSTHelper(nums, mid+1, end)

	return root
}
