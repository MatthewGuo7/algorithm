package dfs

import (
	"fmt"
	"sort"
	"testing"
)

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) == 0 {
		return ret
	}

	visited := make(map[int]struct{}, 0)
	permute := make([]int, 0)

	permuteDFS(nums, visited, &permute, &ret)

	return ret
}

func permuteDFS(nums []int, visited map[int]struct{}, permute *[]int,
	permutes *[][]int) {

	if len(*permute) == len(nums) {
		tempPermute := make([]int, len(*permute))
		copy(tempPermute, *permute)
		*permutes = append(*permutes, tempPermute)
		return
	}

	for _, num := range nums {
		if _, ok := visited[num]; ok {
			continue
		}

		visited[num] = struct{}{}
		*permute = append(*permute, num)
		permuteDFS(nums, visited, permute, permutes)

		delete(visited, num)
		*permute = (*permute)[:len(*permute)-1]
	}

}

func permuteUnique(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) == 0 {
		return ret
	}

	//visited := make(map[int]struct{}, 0)
	visited := make([]int,len(nums))
	permute := make([]int, 0)

	sort.Ints(nums)

	permuteUniqueDFS(nums, visited, &permute, &ret)

	return ret
}



func permuteUniqueDFS(nums []int, visited []int, permute *[]int,
	permutes *[][]int) {

	if len(*permute) == len(nums) {
		tempPermute := make([]int, len(*permute))
		copy(tempPermute, *permute)
		*permutes = append(*permutes, tempPermute)
		return
	}

	for index, num := range nums {
		if visited[index] == 1 {
			continue
		}

		if index > 0 && (nums[index] == nums[index-1]) &&visited[index-1] == 0{
				continue
		}

		*permute = append(*permute, num)
		visited[index] = 1

		permuteUniqueDFS(nums, visited, permute, permutes)

		visited[index] = 0
		*permute = (*permute)[:len(*permute)-1]
	}
}

func permute2(nums []int) [][]int {
	ret := make([][]int, 0)

	if len(nums) == 0 {
		return  ret
	}
	visited := make([]int, len(nums))
	chosen := make([]int, 0)
	permuteRecursive(nums, 0, &visited, &chosen, &ret)

	return ret
}

func permuteRecursive(nums []int, index int,visited *[]int, chosen *[]int, answer *[][]int){

	if len(*chosen) == len(nums) {
		tmp := make([]int, len(*chosen))
		copy(tmp, *chosen)
		*answer = append(*answer, tmp)

		return
	}

	for start := 0; start < len(nums); start++{
		if (*visited)[start] == 1 {
			continue
		}

		*chosen = append(*chosen, nums[start])
		(*visited)[start] = 1
		permuteRecursive(nums, index+1, visited, chosen, answer)
		(*visited)[start] = 0
		*chosen = (*chosen)[:len(*chosen)-1]
	}
}

func TestPmutestUnique(t *testing.T) {
	input1 := [5]int{}
	fmt.Println(input1)
	//input := []int{3,3,0,3}
	//ret := permuteUnique(input)
	//fmt.Println(ret)
}
