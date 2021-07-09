package string

func partition(nums []int, target int) int{
	left := 0
	right := len(nums) - 1

	for left <= right {
		for left <= right && nums[left] < target {
			left++
		}

		for left <= right && nums[right] >= target {
			right--
		}

		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	return left
}
