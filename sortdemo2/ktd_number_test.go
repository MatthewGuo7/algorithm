package sortdemo2

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}

	return findKthLargestHelper(nums, 0, len(nums)-1, k)
}

func findKthLargestHelper(nums []int, start, end, k int) int {
	if start == end {
		return nums[start]
	}

	left := start
	right := end
	mid := (left + right) / 2

	for left <= right {
		for left <= right && nums[left] < nums[mid] {
			left++
		}

		for left <= right && nums[right] > nums[mid] {
			right--
		}

		if left <= right {
			nums[left], nums[right] = nums[right], nums[end]
			left++
			right--
		}
	}

	if start+(k-1) <= right {
		return findKthLargestHelper(nums, start, right, k)
	}

	if start+(k-1) >= left {
		return findKthLargestHelper(nums, left, end, k-(left-start))
	}

	return nums[right+1]
}
