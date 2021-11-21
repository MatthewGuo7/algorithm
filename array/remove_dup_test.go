package array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

func moveZeroes(nums []int) {
	n := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n] = nums[i]
			n++
		}
	}

	for i := n; i < len(nums); i++ {
		nums[i] = 0
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	index := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}

		index--
	}

	for index >= 0 {
		if i >= 0 {
			nums1[index] = nums1[i]
			index--
			i--
		}

		if j >= 0 {
			nums1[index] = nums2[j]
			index--
			i--
		}
	}
}
