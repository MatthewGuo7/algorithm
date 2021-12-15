package two_pointers

import "math"

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0
	min := math.MaxInt32
	sum := 0
	for i := 0 ; i < len(nums);i++{
		for j < len(nums) && sum < target {
			sum += nums[j]
			j++
		}

		if sum >= target{
			length := j - i
			if length < min {
				min = length
			}
		}

		sum -= nums[i]
	}

	if min == math.MinInt32 {
		return 0
	}

	return min
}


