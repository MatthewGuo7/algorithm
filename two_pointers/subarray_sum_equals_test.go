package two_pointers

import "math"

func subarraySumEqualsK2(nums []int, k int) int {
	prefixSum := getPrefixSum(nums)

	sum2Index := make(map[int]int)
	answer := math.MaxInt
	for end := 0; end < len(nums); end++{
		prefixSumStart := prefixSum[end+1] - k
		if start, ok := sum2Index[prefixSumStart]; ok {
			length := end + 1 - start

			if length < answer {
				answer = length
			}
		}

		sum2Index[prefixSum[end+1]] = end +1
	}

	return answer
}

func getPrefixSum(nums []int)[]int  {
	prefixSum:= make([]int, len(nums)+1)

	prefixSum[0] = 0

	for index, num := range nums {
		prefixSum[index+1] = prefixSum[index] + num
	}

	return prefixSum
}
