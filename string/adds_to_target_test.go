package string

func AddsToTarget(values []int, target int) (int,  int)  {
	first, second := 0,0

	numToIndex := make(map[int]int)
	for k, v := range values{
		left := target - v
		if _, ok := numToIndex[left]; ok {
			first = numToIndex[left]
			second = k
			return first, second
		}
		numToIndex[k] = v
	}
	return -1, -1
}
