package two_pointers

func winSum(nums []int, k int) []int{
	if len(nums) == 0 {
		return nil
	}

	i := 0;
	j := 1

	ret := make([]int, 0)
	sum := 0
	for ; i < len(nums);i++{
		for j < len(nums) && j - i < k {
			sum += nums[j]
			j++
		}

		if j >= len(nums) {
			break
		}

		if j - i == k {
			ret = append(ret, sum)
		}

		sum -= nums[i]
	}

}