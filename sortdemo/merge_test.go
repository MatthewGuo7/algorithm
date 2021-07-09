package sortdemo

import (
	"fmt"
	"testing"
)

func mergeSort(nums []int)  {
	temp := make([]int, len(nums))
	mergeSortHelper(nums, 0, len(nums)-1, temp)
}

func mergeSortHelper(nums []int, start, end int, temp []int)  {
	if start >= end {
		return
	}

	mid := (start+end)/2
	mergeSortHelper(nums,start, mid, temp)
	mergeSortHelper(nums, mid+1, end, temp)
	merge(nums, start, mid, end, temp)
}

func merge(nums []int, start, mid, end int, temp []int)  {
	leftIndex := start
	rightIndex := mid + 1
	index := start

	for leftIndex <= mid &&rightIndex <= end {
		if nums[leftIndex] < nums[rightIndex] {
			temp[index] = nums[leftIndex]
			leftIndex++
		} else {
			temp[index] = nums[rightIndex]
			rightIndex++
		}
		index++
	}

	for leftIndex <= mid {
		temp[index] = nums[leftIndex]
		index++
		leftIndex++
	}

	for rightIndex <= end {
		temp[index] = nums[rightIndex]
		index++
		rightIndex++
	}

	for index = start; index <= end; index++ {
		nums[index] = temp[index]
	}
}

func TestMergeSort(t *testing.T)  {
	v := []int{3,6,4,8,7,10,1}
	mergeSort(v)
	fmt.Println(v)
}