package firstUniqChar

func firstUniqChar(s string) int {
	table := [256]int64{}
	length := len(s)
	for i := 0; i < length; i++ {
		table[s[i]]++
	}
	for i := 0; i < length; i++ {
		if table[s[i]] == 1 {
			return i
		}
	}
	return -1
}
