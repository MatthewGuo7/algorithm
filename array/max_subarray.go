package array

import "math"

// 最小子序和
// prefixSum[j+1] - prefixSum[i]
func maxSubArray(nums []int) int {
	if len(nums) == 0{
		return 0
	}

	minSum := 0
	maxSum := math.MinInt32
	sum := 0

	for _, num := range nums {
		sum += num

		curMaxSum := sum - minSum
		if curMaxSum > maxSum {
			maxSum = curMaxSum
		}

		if sum < minSum {
			minSum = sum
		}
	}

	return maxSum
}

// prefixSum[j+1] - prefixSum[i]
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	prefixSum := getPrefixSum(nums)
	preMin := getPreMin(prefixSum)

	max := math.MinInt32

	for index := 1; index < len(prefixSum);index++ {
		sum := prefixSum[index]	- preMin[index-1]
		if sum > max {
			max = sum
		}
	}

	return max
}

func getPreMin(prefixSum []int) []int{
	preMin := make([]int, len(prefixSum))
	preMin[0] = prefixSum[0]

	for index := 1; index < len(prefixSum); index++ {
		min := preMin[index-1]
		if prefixSum[index] < min {
			preMin[index] = prefixSum[index]
		} else {
			preMin[index] = min
		}
	}

	return preMin
}

func getPrefixSum(nums []int) []int {
	prefixSum := make([]int, len(nums)+1)
	prefixSum[0] = 0

	for index, num := range nums{
		prefixSum[index+1] = prefixSum[index] + num
	}

	return prefixSum
}
