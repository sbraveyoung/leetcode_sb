package isAnagram

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	table := [256]int{}
	for _, b := range s {
		table[b]++
	}
	for _, b := range t {
		table[b]--
	}
	for _, v := range table {
		if v != 0 {
			return false
		}
	}
	return true
}
