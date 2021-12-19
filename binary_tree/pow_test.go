package binary_tree

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {

	}

	temp := myPow(x, n/2)
	answer := temp * temp

	if n %2 == 1 {
		answer *= x
	}

	return answer
}
