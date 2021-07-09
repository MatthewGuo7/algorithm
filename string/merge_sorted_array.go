package string

func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[index] = nums1[m-1]
			m--
		} else {
			nums1[index] = nums2[n-1]
			n--
		}

		index--
	}

	for n > 0 {
		nums1[index] = nums1[n]
		n--
		index--
	}

}

func merge2(nums1 []int, nums2 []int) []int {
	ret := make([]int, len(nums1)+len(nums2))

	i,j:= 0,0

	k := 0
	for i < len(nums1) && j < len(nums2){
		if nums1[i] > nums2[j] {
			ret[k] = nums1[i]
			i++
		} else {
			ret[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len(nums1) {
		ret[k] = nums1[i]
		k++
		i++
	}

	for j < len(nums2) {
		ret[k] = nums2[j]
		k++
		j++
	}

	return ret
}
