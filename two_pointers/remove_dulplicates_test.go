package two_pointers

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 1
	i := 0
	for ; i < len(nums);i++{
		for j < len(nums) && (nums[j] == nums[i]) {
			j++
		}

		if j >= len(nums) {
			break
		}

		nums[i+1] = nums[j]
	}

	return i+1
}
