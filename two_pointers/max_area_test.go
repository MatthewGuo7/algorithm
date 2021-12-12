package two_pointers

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	i, j := 0, len(height)-1
	answer := 0

	for i < j {
		length := j - i
		heigh := height[i]
		if heigh > height[j] {
			heigh = height[j]
			j--
		} else {
			i++
		}

		area :=  heigh * length
		if area > answer {
			answer = area
		}
	}

	return answer
}
