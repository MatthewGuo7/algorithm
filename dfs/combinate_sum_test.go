package dfs

import (
	"fmt"
	"sort"
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	combinatios := make([][]int, 0)
	if len(candidates) == 0 {
		return combinatios
	}

	candidates = removeDulplicatesAndSort(candidates)

	subSet := make([]int, 0)
	combinationSumDFS(candidates, 0, 0, target,
		&subSet, &combinatios)

	return combinatios
}

func removeDulplicatesAndSort(candiadates []int)  []int{
	m := make(map[int]struct{}, len(candiadates))
	for _, candiadate := range candiadates {
		m[candiadate] = struct{}{}
	}

	ret := make([]int, 0, len(m))
	for k,_ := range m {
		ret = append(ret, k)
	}

	sort.Ints(ret)

	return ret
}

func combinationSumDFS(candidates[]int, index int, sum int, target int,
	subSet *[]int,
	subSets *[][]int)  {

	if sum > target {
		return
	}

	if sum == target {
		temp := make([]int, len(*subSet))
		copy(temp, *subSet)
		fmt.Println(subSet)
		fmt.Println(temp)
		*subSets = append(*subSets, temp)
		return
	}

	//for _, candidate := range candidates {
	for tempIndex := index; tempIndex < len(candidates); tempIndex++{
		candidate := candidates[tempIndex]
		if sum + candidate > target {
			break
		}


		*subSet = append(*subSet, candidate)
		combinationSumDFS(candidates, tempIndex, sum + candidate, target, subSet, subSets)
		*subSet = (*subSet)[:len(*subSet)-1]
	}
}

func TestCombination(t *testing.T)  {
	values := []int{2,3,6,7}
	target := 7

	ret := combinationSum(values, target)
	fmt.Println(ret)
}


