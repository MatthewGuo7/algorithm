package strings2

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 {
		return -1
	}

	for indexHayStack := 0; indexHayStack < len(haystack)-len(needle)+1; indexHayStack++ {
		indexNeedle := 0
		for ; indexNeedle < len(needle); indexNeedle++ {
			if haystack[indexHayStack+indexNeedle] == needle[indexNeedle] {
				continue
			}

			break
		}

		if indexNeedle == len(needle) {
			return indexHayStack
		}
	}

	return -1
}

//abcdef
//bcd
func strStr2(src string, target string) int {
	lenSrc, lenTarget := len(src), len(target)

	if lenTarget == 0 {
		return 0
	}

	if lenSrc == 0 {
		return -1
	}

	hashCodeTarget := 0
	for index := 0; index < lenTarget; index++ {
		hashCodeTarget += int(target[index])
	}

	hashCodeSrc := 0
	for index := 0; index < lenSrc; index++ {
		hashCodeSrc += int(src[index])
		//abc+d
		if index < lenTarget-1 {
			continue
		}

		//abcd-a
		if index >= lenTarget {
			hashCodeSrc -= int(src[index-lenTarget])
		}

		if hashCodeSrc == hashCodeTarget {
			if src[index-lenTarget+1:index+1] == target {
				return index - lenTarget + 1
			}
		}

	}

	return -1
}
