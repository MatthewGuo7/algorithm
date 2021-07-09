package string

func LongestCommonString(str string, str2 string) int {

	lcs := 0
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str2); j++ {
			lcs_temp := 0
			for str[i+lcs_temp] == str2[j+lcs_temp] && i+lcs_temp < len(str)&& j +lcs_temp < len(str2) {
				lcs_temp++
			}

			if lcs_temp > lcs {
				lcs = lcs_temp
			}
		}
	}


	return 0
}
