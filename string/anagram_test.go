package string

import "testing"

func IsAnagram(s, t string) bool  {
	if len(s)== 0 || len(t) == 0 {
		return false
	}

	if len(s) != len(t) {
		return false
	}

	letterCount := [256]int{}
	for k, v := range s{
		letterCount[v]++
		letterCount[k]--
	}

	for i, _ := range letterCount {
		if letterCount[i] < 0 {
			return false
		}
	}

	return true
}

func TestAnagram(t *testing.T) {

}
