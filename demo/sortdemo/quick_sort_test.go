package sortdemo

func QuickSort(values []int) {
	if len(values) == 0 {
		return
	}

}

func quickSortHelper(values []int, start, end int) {
	if start >= end {
		return
	}

	mid := end - (end-start)/2

	left := start
	right := end

	for left <= mid && mid <= right {
		for left <= mid && values[left] < values[mid] {
			left++
		}

		for right >= mid && values[right] < values[mid] {
			right--
		}

		if left <= right {
			values[left], values[right] = values[right], values[left]
		}
	}

	quickSortHelper(values, start, right)
	quickSortHelper(values, left, end)
}
