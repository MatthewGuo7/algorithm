package dfs

func combine(n int, k int) [][]int {
	answer := make([][]int, 0)
	chosen := make([]int,0)

	combineRecursive(n, k, 1, &chosen, &answer)

	return answer
}

func combineRecursive(n int, k int, index int, chosen *[]int, answer *[][]int) {
	if len(*chosen) > k || len(*chosen) + (n-index+1) < k {
		return
	}

	if index == n + 1 {
		if len(*chosen) == k {
			tmp := make([]int, len(*chosen))
			copy(tmp, *chosen)
			*answer = append(*answer, tmp)
		}

		return
	}

	combineRecursive(n, k, index+1, chosen, answer)

	*chosen = append(*chosen, index)
	combineRecursive(n, k, index+1, chosen, answer)
	*chosen = (*chosen)[:len(*chosen)-1]
}
