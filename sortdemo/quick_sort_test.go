package sortdemo

func QuickSort(values []int) {
	if len(values) == 0 {
		return
	}

	quickSortHelper(values, 0, len(values) - 1)
}

func quickSortHelper(values []int, start, end int)  {

	left := start
	right := end

	midValue := values[(start+end)/2]

	for left <= right {
		for left <= right && values[left] < midValue {
			left++
		}

		for left <= right && values[right] < midValue {
			right--
		}

		if left <= right {
			values[left], values[right] = values[right], values[left]
			left++
			right--
		}
	}

	quickSortHelper(values, start, right)
	quickSortHelper(values, left, end)
}
