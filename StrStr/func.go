package strstr

//KMP
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := calcNext(needle)
	i, j := 0, 0
	for ; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			//j = next[j-1]
			return i - len(needle) + 1
		}
	}
	return -1
}

func calcNext(pattern string) []int {
	next := make([]int, len(pattern))
	if len(pattern) == 0 {
		return next
	}
	next[0] = -1
	i, j := 0, -1
	for i < len(pattern)-1 {
		for j >= 0 && pattern[j] != pattern[i] {
			j = next[j]
		}
		i++
		j++
		next[i] = j
	}
	return next
}

/*
//Boyer-Moore
func strStr(haystack string, needle string) int {
}
*/

/*
//Aho-Corasick
func strStr(haystack string, needle string) int {
}
*/
