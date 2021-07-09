package string

type Bound struct {
	Low, High int
}

func longestConsecutive(nums []int) int {
	maxLen := 0

	m := make(map[int]*Bound)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			continue
		}

		low, high := v, v

		if _, ok := m[low-1]; ok {
			low = m[low-1].Low
		}

		if _, ok := m[high+1]; ok {
			high = m[high+1].High
		}

		m[low].High = high
		m[v].High = high

		m[high].Low = low
		m[v].Low = low

		if high-low+1 > maxLen {
			maxLen = high-low+1
		}
	}

	return maxLen
}
