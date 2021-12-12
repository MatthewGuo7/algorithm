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

func twoSum2(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	i := 0;
	j := len(numbers) - 1

	for ; i < len(numbers); i++ {
		for j > i && numbers[i] + numbers[j] != target{
			j--
		}

		if j <= i {
			break
		}

		if numbers[i] + numbers[j] == target{
			return []int{i+1, j+2}
		}

	}

	return nil
}
