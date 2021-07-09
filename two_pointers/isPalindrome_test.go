package two_pointers

import "unicode"

func isPalindrome(s string) bool {
	lenS := len(s)
	if lenS == 0 {
		return false
	}

	left := 0
	right := lenS - 1
	for left < right {
		for left < right && !isValidByte(s[left]) {
			left++
		}

		for left < right && !isValidByte(s[right]) {
			right--
		}

		if left < right && !isEqual(s[left], s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isValidByte(b byte) bool {
	return unicode.IsDigit(rune(b)) || unicode.IsLetter(rune(b))
}

func isEqual(left, right byte) bool {
	return unicode.ToLower(rune(right)) == unicode.ToLower(rune(left))
}

func isPalindrome2(s string) bool {
	lenS := len(s)
	if lenS == 0 {
		return true
	}

	left, right := findDifference(s, 0, lenS-1)
	if left >= right {
		return true
	}

	return IsValidPalindrome(s, left+1, right) || IsValidPalindrome(s, left, right-1)
}

func IsValidPalindrome(s string, left, right int) bool {
	left, right = findDifference(s, left, right)
	return left >= right
}

func findDifference(s string, left, right int) (int, int) {
	for left < right && s[left] == s[right] {
		left++
		right--
	}

	return left, right
}
