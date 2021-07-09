package sortdemo

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0{
		return -1
	}

	return findKthLargetHelper(nums, 0, len(nums)-1,k)
}

func findKthLargetHelper(nums []int, start, end, k int) int  {
	if start == end {
		return nums[start]
	}
	left := start
	right := end

	midValue := nums[(left+right)/2]

	for left <= right {
		for left <= right && nums[left] > midValue {
			left++
		}
		for left <= right && nums[right] < midValue {
			right--
		}

		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	if start + k - 1 <= right{
		return findKthLargetHelper(nums, start, right, k)
	}

	if start + k - 1 >= left {
		return findKthLargetHelper(nums, left, end, k - (left-start))
	}

	return nums[right+1]
}