package dfs

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	combinations := make([]string,0)

	if len(digits) == 0 {
		return combinations
	}

	letterCombinationsHelper(digits, 0, "", &combinations)

	return combinations
}

func letterCombinationsHelper(digits string, index int, combination string, combinations *[]string)  {

	if index == len(digits) {
		*combinations = append(*combinations, combination)
		return
	}

	curLetters := phoneMap[string(digits[index])]
	for _, letter := range curLetters {
		letterCombinationsHelper(digits, index+1, combination+string(letter), combinations)
	}
}

//func ksum(a []int, k int, target int) [][]int {
//	ret := make([][]int, 0)
//	if len(a) == 0 {
//		return ret
//	}
//
//
//}
//
//func ksumdfs(a []int, index int, curtarget int,
//	k int, target int,
//	subset *[]int,
//	subsets *[][]int)  {
//
//	if index >= len(a) {
//		return
//	}
//
//	if index == k && curtarget == target {
//		*subsets = append(*subsets, *subset)
//		return
//	}
//
//	for _, value := range a{
//		*subset = append(*subset, value)
//		ksumdfs(a, index+1, curtarget + value, k, target, subset, subsets)
//		*subset = (*subset)[:len(*subset)-1]
//	}
//
//1}
