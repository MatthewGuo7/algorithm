package binary2

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) < len(nums2) {
		return
	}

	index := m + n - 1

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[index] = nums1[m-1]
			index--
			m--
		} else {
			nums1[index] = nums2[n-1]
			index--
			n--
		}
	}

	for m > 0 {
		nums1[index] = nums1[m]
		index--
		m--
	}
}
