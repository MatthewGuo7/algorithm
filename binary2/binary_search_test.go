package binary2

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end int, target int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		return binarySearch(nums, mid+1, end, target)
	}

	return binarySearch(nums, start, mid-1, target)
}

func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			start = mid
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if nums[start] == target {
		return start
	}

	if nums[end] == target {
		return end
	}

	return -1
}
