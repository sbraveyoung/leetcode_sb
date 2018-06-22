package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	ret := strs[0]
	last := len(ret)
label:
	for _, str := range strs[1:] {
		i := 0
		for ; i < last && i < len(str); i++ {
			if str[i] != ret[i] {
				last = i
				continue label
			}
		}
		if i != last {
			last = i
		}
	}
	return ret[:last]
}
