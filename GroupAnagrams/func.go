package groupAnagrams

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		base := convertToBase(strs[i])
		if v, ok := m[base]; ok {
			m[base] = append(v, strs[i])
		} else {
			m[base] = []string{strs[i]}
		}
	}
	ret := [][]string{}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

func convertToBase(str string) string {
	m := make([]int, 26)
	for _, v := range str {
		m[v-'a']++
	}
	ret := ""
	for k, v := range m {
		for i := 0; i < v; i++ {
			ret += string('a' + k)
		}
	}
	return ret
}

/*
func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	for i := 0; i < len(strs); i++ {
		j := 0
		for ; j < len(ret); j++ {
			if len(strs[i]) != len(ret[j][0]) {
				continue
			}
			if isSame(strs[i], ret[j][0]) {
				ret[j] = append(ret[j], strs[i])
				break
			}
		}
		if j >= len(ret) {
			ret = append(ret, []string{strs[i]})
		}
	}
	return ret
}

func isSame(one, two string) bool {
	dict := map[rune]int{}
	for _, b := range one {
		if _, ok := dict[b]; ok {
			dict[b]++
		} else {
			dict[b] = 1
		}
	}
	for _, b := range two {
		if v, ok := dict[b]; ok && v > 0 {
			dict[b]--
		} else {
			return false
		}
	}
	for _, v := range dict {
		if v != 0 {
			return false
		}
	}
	return true
}
*/
