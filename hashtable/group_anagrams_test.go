package hashtable

import "sort"

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	groups := make(map[string][]string)

	for _, str := range strs {
		tmpStr := []byte(str)
		sort.Slice(tmpStr, func(i, j int) bool {
			return tmpStr[i] < tmpStr[j]
		})

		groups[string(tmpStr)] = append(groups[string(tmpStr)], str)
	}

	ret := make([][]string,0, len(groups))
	for _, group := range groups {
		ret = append(ret, group)
	}

	return ret
}

func groupAnagrams2(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	groups := make(map[[26]int][]string)
	for _, str := range strs{
		k := [26]int{}
		for _, s := range str{
			k[s-'a']++
		}

		if _, ok := groups[k]; ok {
			groups[k] = append(groups[k], str)
		}
	}

	ret := make([][]string,0, len(groups))
	for _, group := range groups {
		ret = append(ret, group)
	}

	return ret
}
