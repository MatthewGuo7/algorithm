package string

import (
	"fmt"
	"testing"
)

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 {
		return -1
	}

	n, m := len(haystack), len(needle)

	for index := 0; index < n-m+1; index++ {
		if haystack[index] != needle[0] {
			continue
		}

		i := index
		j := 0
		for j = 0; j < m; j++ {
			if haystack[i] == needle[j] {
				continue
			}
		}

		if j == m {
			return index
		}
	}

	return -1
}

const primeRK = 16777619

func strStr2(haystack string, needle string) int {
	n, m := len(haystack), len(needle)

	if m == 0 {
		return 0
	}

	if n == 0 {
		return -1
	}

	//needle's hash
	h := uint32(0)
	for i := 0; i < m; i++ {
		h += uint32(needle[i])
	}

	hashM := h

	hashN := uint32(0)
	h = uint32(0)
	for index := 0; index < n ; index++ {
		hashN += uint32(haystack[index])
		if index < m - 1 {
			continue
		}

		if index >= m {
			hashN -= uint32(haystack[index-m])
		}
		if hashN == hashM {
			if haystack[index - m +1:index+1] == needle {
				return index - m +1
			}
		}


	}

	return -1
}

func TestStr(t *testing.T) {
	s1 := "hello"
	s2 := "ll"
	fmt.Println(strStr2(s1, s2))
}
