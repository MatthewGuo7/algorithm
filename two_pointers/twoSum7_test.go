package two_pointers

import (
	"math"
	"sort"
)

func twoSum7(nums []int, target int) []int {
	if len(nums) == 0 || len(nums) < 2 {
		return nil
	}

	target = int(math.Abs(float64(target)))

	sort.Ints(nums)
	for index := 0; index < len(nums); index++ {
		j := binarySearch(nums, 0, len(nums)-1, target+nums[index])
		if j != -1 {
			return []int{index, j}
		}
	}

	return nil
}

func binarySearch(nums []int, start, end int, target int) int {
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}

	if nums[start] == target {
		return start
	}

	if nums[end] == target {
		return end
	}

	return -1
}

func twoSum77(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	target = int(math.Abs(float64(target)))

	j := 1
	for i := 0; i < len(numbers)-1; i++ {
		j = int(math.Max(float64(j), float64(i+1)))
		for j < len(numbers) && numbers[j]-numbers[i] < target {
			j++
		}

		if j >= len(numbers) {
			break
		}

		if numbers[j]-numbers[i] == target {
			return []int{i, j}
		}
	}

	return nil
}

func twoSum72(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	i, j := 0, 0
	target = int(math.Abs(float64(target)))

	for ; i < len(numbers);i++{
		j = int(math.Max(float64(j), float64(i+1)))
		for j < len(numbers) && numbers[j]-numbers[i] < target {
			j++
		}

		if j >= len(numbers) {
			break
		}

		if numbers[j]-numbers[i] == target {
			return []int{i,j}
		}
	}

	return nil
}