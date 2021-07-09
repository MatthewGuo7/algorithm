package string

func CanCompose(newsPaper, message string) bool {
	if len(newsPaper) < len(message) {
		return false
	}

	numNewsPaper := make(map[int32]int32)
	for _, v := range newsPaper {
		numNewsPaper[v]++
	}

	for _, v := range message {
		if numNewsPaper[v] == 0 {
			return false
		}

		numNewsPaper[v]--

		if numNewsPaper[v] < 0 {
			return false
		}
	}


	return false
}
