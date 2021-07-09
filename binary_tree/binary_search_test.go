package binary_tree

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	start := 0
	end := len(nums) - 1

	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			end = mid
		} else if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
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

func findMin(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1

	for start+1 < end {
		mid := end - (end-start)/2
		if nums[mid] < nums[end] {
			end = mid
		} else {
			start = mid
		}
	}

	if nums[start] < nums[end] {
		return nums[start]
	} else {
		return nums[end]
	}

}

func peakIndexInMountainArray(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {
		return arr[0]
	}

	if len(arr) == 2 {
		if arr[0] < arr[1] {
			return 1
		} else {
			return 0
		}
	}

	start := 0
	end := len(arr) - 1

	for start+1 < end {
		mid := end - (end-start)/2
		if arr[start] < arr[mid] && arr[mid] <= arr[mid+1] {
			start = mid
		} else {
			end = mid
		}
	}

	if arr[start] < arr[end] {
		return start
	} else {
		return end
	}
}

func searchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1

	for start + 1 < end {
		mid := end-(end-start)/2
		if nums[mid] > nums[start] {
			if nums[start] < target && target < nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else {
			if nums[mid] < target && target < nums[end] {
				start = mid
			} else  {
				end = mid
			}
		}
	}

	if nums[start] == target{
		return start
	}

	if nums[end] == target{
		return end
	}

	return -1
}
