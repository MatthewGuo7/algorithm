package subarray

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := math.MinInt32

	minSum := 0
	sum := 0
	for _, num := range nums {
		sum += num

		curMaxSum := sum - minSum
		if curMaxSum > max {
			max = curMaxSum
		}

		if sum < minSum {
			minSum = sum
		}
	}

	return max
}
