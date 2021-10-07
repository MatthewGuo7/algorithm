package dfs

import (
	"fmt"
	"testing"
)

var (
	dx = []int{0, 0, -1, 1}
	dy = []int{1, -1, 0, 0}
)

func findWords(board [][]byte, words []string) []string {
	ret := make([]string, 0)
	if len(board) == 0 || len(board[0]) == 0 {
		return ret
	}

	wordSet := make(map[string]struct{})
	prefixSet := make(map[string]struct{})

	for _, word := range words {
		for index := range word {
			prefixSet[word[:index+1]] = struct{}{}
		}

		wordSet[word] = struct{}{}
	}

	visited := make([][]int, len(board))
	for index := 0; index < len(board); index++ {
		visited[index] = make([]int, len(board[0]))
	}

	retWords := make(map[string]struct{})

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			visited[i][j] = 1
			findWordsDFS(board, i, j, string(board[i][j]), visited, prefixSet, wordSet, retWords)
			visited[i][j] = 0
		}
	}

	ret = make([]string, 0, len(retWords))
	for word := range retWords {
		ret = append(ret, word)
	}

	return ret
}

func findWordsDFS(board [][]byte, x, y int, curWord string, visited [][]int,
	prefixSet map[string]struct{}, wordSet map[string]struct{},
	retSet map[string]struct{}) {

	if _, ok := prefixSet[curWord]; !ok {
		return
	}

	if _, ok := wordSet[curWord]; ok {
		retSet[curWord] = struct{}{}
	}

	for index := 0; index < len(dx); index++ {
		adjX := x + dx[index]
		adjY := y + dy[index]

		if !isInBoard(board, adjX, adjY) || visited[adjX][adjY] == 1 {
			continue
		}

		visited[adjX][adjY] = 1
		findWordsDFS(board, adjX, adjY, curWord+string(board[adjX][adjY]),
			visited, prefixSet, wordSet, retSet)
		visited[adjX][adjY] = 0
	}
}

func isInBoard(board [][]byte, x, y int) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	if x < 0 || x >= len(board) {
		return false
	}

	if y < 0 || y >= len(board[0]) {
		return false
	}

	return true
}

func TestWordSearch2(t *testing.T) {
	// [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]
	//["oath","pea","eat","rain"]

	board := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}


	words := []string{"oath","pea","eat","rain"}
	ret := findWords(board, words)

	fmt.Println(ret)
}
