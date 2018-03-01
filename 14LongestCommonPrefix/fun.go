package main

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	ret := ""
	ch := 'a'
	for i := 0; ; i++ {
		if i >= len(strs[0]) {
			return ret
		}
		ch = rune(strs[0][i])
		for _, str := range strs {
			if i >= len(str) {
				return ret
			}
			if rune(str[i]) != ch {
				return ret
			}
		}
		ret = ret + string(ch)
	}
	return ret
}
