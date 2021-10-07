package dfs

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	prefixSet := make(map[string]struct{})
	for index := range word {
		prefixSet[word[:index+1]] = struct{}{}
	}

	visited := make([][]int, len(board))
	for index := 0; index < len(board); index++ {
		visited[index] = make([]int, len(board[0]))
	}

	for row := 0; row < len(board); row++{
		for col := 0; col < len(board[0]); col++{
			visited[row][col] = 1
			ret := existDFS(board, row, col, string(board[row][col]), visited, prefixSet, word)

			if ret {
				return ret
			}
			visited[row][col] = 0
		}
	}

	return false
}

func existDFS(board [][]byte, x, y int, curWord string, visited [][]int,
	prefixSet map[string]struct{}, word string) bool {

	if _, ok := prefixSet[curWord]; !ok {
		return false
	}

	if curWord == word {
		return true
	}

	for index := 0; index < len(dx);index++{
		adjX := x + dx[index]
		adjY := y + dy[index]

		if !isInBoard(board,adjX, adjY) || visited[adjX][adjY] == 1 {
			continue
		}

		visited[adjX][adjY] = 1
		ret := existDFS(board, adjX, adjY, curWord + string(board[adjX][adjY]), visited,
			prefixSet, word)

		if ret {
			return ret
		}

		visited[adjX][adjY] = 0
	}


	return false
}
