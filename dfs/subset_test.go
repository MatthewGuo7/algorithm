package dfs

import "sort"

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) == 0 {
		return ret
	}

	subset := make([]int, 0)
	subsetsDFS(nums, 0, &subset, &ret)

	return ret
}

func subsetsDFS(nums []int, index int, subset *[]int, subsets *[][]int) {
	temp := make([]int, len(*subset))
	copy(temp, *subset)
	*subsets = append(*subsets, temp)

	for tempIndex := index; tempIndex < len(nums); tempIndex++ {
		*subset = append(*subset, nums[tempIndex])
		subsetsDFS(nums, index+1, subset, subsets)
		*subset = (*subset)[:len(*subset)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) == 0 {
		return ret
	}

	sort.Ints(nums)

	subset := make([]int, 0)
	subsetsWithDupDFS(nums, 0, &subset, &ret)

	return ret
}

func subsetsWithDupDFS(nums []int, offset int, subset *[]int, subsets *[][]int) {
	temp := make([]int, len(*subset))
	copy(temp,*subset)
	*subsets = append(*subsets, temp)

	for index := offset; index < len(nums); index++ {

		if index!= 0 && nums[index] == nums[index-1] && index > offset {
			continue
		}

		*subset = append(*subset, nums[index])
		subsetsWithDupDFS(nums, index+1, subset, subsets)
		*subset = (*subset)[:len(*subset)-1]
	}
}

func subsets2(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	ret := make([][]int, 0)
	curAnswer := make([]int,0)
	subsetsRecur(nums, 0, &curAnswer, &ret)

	return ret
}

func subsetsRecur(nums []int, index int, curAnswer *[]int, answer *[][]int) {
	if index == len(nums) {
		tmp := make([]int, len(*curAnswer))
		copy(tmp, *curAnswer)
		*answer = append(*answer, tmp)

		return
	}

	subsetsRecur(nums, index+1, curAnswer, answer)

	*curAnswer = append(*curAnswer, nums[index])
	subsetsRecur(nums, index+1, curAnswer, answer)
	*curAnswer = (*curAnswer)[:len(*curAnswer)-1]
}
