package string

import (
	"fmt"
	"strconv"
	"testing"
)

func ReverseWords(src []byte) string {
	for index := 0; index < len(src); index++ {
		offset, needReverse := needReverse(src, index)
		if needReverse {
			word := doReverse(src[index : index+offset])
			fmt.Println(string(word))
			for i, v := range word {
				src[index+i] = v
			}
		}
		index += offset + 1
	}

	return string(src)
}

func doReverse(src []byte) []byte {
	left := 0
	right := len(src) - 1

	for left < right {
		(src)[left], (src)[right] = (src)[right], (src)[left]
		left++
		right--
	}

	return src
}

func needReverse(src []byte, index int) (offset int, needReverse bool) {
	j := index
	for ; j < len(src) - index +1; j++ {
		if !IsPunctuationOrSpace(src[j]) {
			continue
		}

		if !IsNumber(src[index:j]) {
			needReverse = true
			offset = j
			return
		}
	}

	return j+1, false
}

func IsNumber(src []byte) bool {
	if _, err := strconv.Atoi(string(src)); err != nil {
		return false
	}

	return true
}

func IsPunctuationOrSpace(b byte) bool {
	return b == ' ' || b == ',' || b == '.'
}

func TestReverseWords(t *testing.T) {
	src := "I have 36 books, 40 pens2."
	//want := "I evah 36 skoob, 40 2snep."

	ret := ReverseWords([]byte(src))
	fmt.Printf("%+v", ret)
}
