package two_pointers

func twoSum(nums []int, target int) []int {
	ret := []int{
		-1,
		-1,
	}

	if len(nums) == 0 {

		return ret
	}

	exists := make(map[int]int)
	for index, num := range nums {
		left := target-num
		if i, ok := exists[left]; ok {
			ret[0] = i
			ret[1] = index

			return ret
		}

		exists[num] = index
	}

	return ret
}
