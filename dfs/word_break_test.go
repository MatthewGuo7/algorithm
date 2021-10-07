package dfs

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}

	if len(wordDict) == 0 {
		return true
	}

	memo := make(map[int]bool)

	wordSet := wordSet(wordDict)

	return wordBreakDFS(s, wordSet, 0, memo)
}

func wordSet(wordDict []string) map[string]struct{} {
	ret := make(map[string]struct{}, len(wordDict))

	for _, word := range wordDict {
		ret[word] = struct{}{}
	}

	return ret
}

func wordBreakDFS(s string, wordDict map[string]struct{},
	index int, memo map[int]bool) bool {
	if exist, ok := memo[index]; ok {
		return exist
	}

	if index == len(s) {
		return true
	}

	for end := index + 1; end <= len(s); end++ {
		if _, ok := wordDict[s[index:end]]; !ok {
			continue
		}

		if wordBreakDFS(s, wordDict, end, memo) {
			return true
		}
	}
	
	memo[index] = false

	return false
}
