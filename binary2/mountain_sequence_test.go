package binary2

func mountainSequence(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1

	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[mid+1] {
			start = mid
		} else if nums[mid] > nums[mid+1] {
			end = mid
		} else {
			start = mid
		}
	}

	if nums[start] > nums[end] {
		return start
	} else {
		return end
	}
}
