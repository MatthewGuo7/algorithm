package string

func rotate(nums []int, k int)  {
	l := len(nums)
	if k > l {
		k = k % l
	}

	reverse(nums[0:k])
	reverse(nums[k:])
	reverse(nums)
}

func reverse(nums []int)  {
	for left, right := 0, len(nums) - 1; left<right; {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

