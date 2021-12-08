package array

import (
	"fmt"
	"testing"
)

// sum[r] - sum[l-1] = k
// sum[r] - k = sum[l-1]
/*
 i...j
 sum[j+1] - sum[i] = k => sum[j+1]-k = sum[i]
 */
func numberOfSubarrays(nums []int, k int) int {
	prefixSum := getPrefixSum(nums)

	sum2Index := make(map[int]int,len(prefixSum))
	for index := range prefixSum {
		sum2Index[index]=1
	}

	answer := 0
	for start := 0; start < len(nums); start++ {
		sum := prefixSum[start+1] - k

		if _, ok := sum2Index[sum]; ok {
			answer+=sum2Index[sum]
		}

		//for end := start+1; end< len(prefixSum); end++ {
		//	if prefixSum[end] -k  == prefixSum[start]  {
		//		answer++
		//	}
		//}
	}

	//sum2Index := make(map[int]int)

	//answer := 1
	//for index := 0; index < len(nums);index++ {
	//	sum := prefixSum[index+1]-k
	//
	//	if _, ok := sum2Index[sum]; ok {
	//		answer++
	//	}
	//
	//	sum2Index[prefixSum[index+1]] = index+1
	//}
	return answer
}


/*
 [1,1,2,1,1]
[0,1,2,4,5,6]
[0,1,0,0,1,0]
3
 */
func getPrefixSum(nums []int)[]int{
	prefixSum := make([]int, len(nums)+1)
	prefixSum[0] = 0

	for index, num := range nums {
		prefixSum[index + 1] = prefixSum[index] + num%2
	}

	return prefixSum
}

func TestNumber(t *testing.T) {
	input := []int{1,1,2,1,1}
	k := 3

	fmt.Println(numberOfSubarrays(input, k))
}
