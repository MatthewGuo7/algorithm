package strings2

import (
	"fmt"
	"testing"
)

func longestPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := make(map[int32]int, 128)

	for _, b := range s {
		count[b]++
	}

	ret := 0
	isOdd := false
	for _, num := range count {
		ret += num

		if num%2 == 1 {
			ret--
			isOdd = true
		}
	}

	if isOdd {
		ret++
	}

	return ret
}

func longestPalindrome2(s string) string {
	lenS := len(s)
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	ret := ""
	maxLen := 0
	for i := 0; i < lenS; i++ {
		for j := i + 1; j <= lenS; j++ {
			if isPalindrome(s[i:j]) {
				if j-i+1 > maxLen {
					ret = s[i:j]
					maxLen = j - i + 1
				}
			}
		}
	}

	return ret
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func longestPalindrome3(s string) string {
	lenS := len(s)
	if lenS == 0 {
		return ""
	}

	longestRet := ""
	for index := 0; index < lenS; index++ {
		oddLongest := getPalindromeFrom(s, index, index)
		if len(oddLongest) > len(longestRet){
			longestRet = oddLongest
		}

		evenLongest := getPalindromeFrom(s, index, index+1)
		if len(evenLongest) > len(longestRet){
			longestRet = evenLongest
		}
	}

	return longestRet
}

func getPalindromeFrom(s string, left, right int) string {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}

	return s[left+1 : right]
}

func TestLP(t *testing.T) {
	src := "cbbd"
	ret := longestPalindrome2(src)
	fmt.Println(ret)
}
