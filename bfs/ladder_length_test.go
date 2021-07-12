package bfs

import (
	"container/list"
	"fmt"
	"testing"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	allWords := getAllWords(beginWord, endWord, wordList)
	if len(allWords) == 0 {
		return 0
	}

	queue := list.New()
	queue.PushBack(beginWord)

	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}

	length := 1
	for queue.Len() > 0 {
		size := queue.Len()

		for index := 0; index < size; index++ {
			curNode := queue.Front()
			queue.Remove(curNode)

			curWord := curNode.Value.(string)

			for nextWord := range getNextWord(curWord, allWords) {
				if _, ok := visited[nextWord]; ok {
					continue
				}

				if nextWord == endWord {
					return length + 1
				}

				visited[nextWord] = struct{}{}
				queue.PushBack(nextWord)
			}
		}

		length++
	}

	return 0
}

func getNextWord(curWord string, allWords map[string]struct{}) map[string]struct{} {
	ret := make(map[string]struct{})

	for index := byte('a'); index <= byte('z'); index++ {
		for indexWord := 0; indexWord < len(curWord); indexWord++ {
			curWortToMatch := []byte(curWord)
			if curWortToMatch[indexWord] == index {
				continue
			}

			curWortToMatch[indexWord] = index
			//fmt.Println(string(curWortToMatch))
			if _, ok := allWords[string(curWortToMatch)]; ok {
				ret[string(curWortToMatch)] = struct{}{}
			}
		}
	}

	return ret
}

func getAllWords(beginWord string, endWord string, wordList []string) map[string]struct{} {
	words := make(map[string]struct{}, len(wordList))
	for _, word := range wordList {
		if word == beginWord {
			continue
		}
		words[word] = struct{}{}
	}

	words[endWord] = struct{}{}

	return words
}

func TestGetworkd(t *testing.T)  {
	beginWord := "hit"
	endWorkd := "cog"
	dict := []string{"hot","dot","dog","lot","log",/*"cog"*/}

	fmt.Println(ladderLength(beginWord, endWorkd, dict))
}
