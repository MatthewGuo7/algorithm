package sortdemo2

import (
	"fmt"
	"testing"
)

func mergeSort(nums []int)  {
	if len(nums) == 0 {
		return
	}

	temp := make([]int,len(nums))

	mergeSortHelper(nums, 0, len(nums)-1, temp)
}

func mergeSortHelper(nums []int, start, end int, temp []int)  {
	if start >= end{
		return
	}
	mid := (start+end)/2
	mergeSortHelper(nums, start, mid, temp)
	mergeSortHelper(nums, mid+1, end, temp)
	merge(nums, start, end, temp)
}

func merge(nums []int, start, end int, temp []int)  {
	if start>=end{
		return
	}

	mid := (start+end)/2
	leftIndex:= start
	rightIndex := mid+1
	tempIndex := start

	for leftIndex <= mid && rightIndex <= end {
		if nums[leftIndex] < nums[rightIndex] {
			temp[tempIndex] = nums[leftIndex]
			tempIndex++
			leftIndex++
		} else {
			temp[tempIndex] = nums[rightIndex]
			tempIndex++
			rightIndex++
		}
	}

	for leftIndex <= mid {
		temp[tempIndex] = nums[leftIndex]
		tempIndex++
		leftIndex++
	}

	for rightIndex <= end {
		temp[tempIndex] = nums[rightIndex]
		tempIndex++
		rightIndex++
	}

	for index := start; index <= end; index++{
		nums[index] = temp[index]
	}
}

func TestMergeSort(t *testing.T)  {
	v := []int{3,6,4,8,7,10,1}
	mergeSort(v)
	fmt.Println(v)
}
