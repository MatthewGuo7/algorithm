package two_pointers

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	ret := make([][]int,0)
	for i := 0; i < len(nums); i++{
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		answers := twoSum9(nums, i+1, 0-nums[i])
		for _,answer := range answers {
			ret = append(ret, []int{nums[i], answer[0], answer[1]})
		}
	}

	return ret
}

func twoSum9(nums []int, start int, target int) [][]int {
	i := start;
	j := len(nums) - 1

	ret := make([][]int,0)
	for ; i < len(nums); i++ {
		if i > start &&  nums[i] == nums[i-1] {
			continue
		}

		for j > i && nums[j] + nums[i] > target {
			j--
		}

		if j > i && nums[i] + nums[j] == target{
			ret = append(ret, []int{nums[i], nums[j]})
		}
	}

	return ret
}
