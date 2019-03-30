package isIsomporphic

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dict1 := map[byte]byte{}
	dict2 := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if val, ok := dict1[s[i]]; ok {
			if val != t[i] {
				return false
			}
		}
		if val, ok := dict2[t[i]]; ok {
			if val != s[i] {
				return false
			}
		}
		dict1[s[i]] = t[i]
		dict2[t[i]] = s[i]
	}
	return true
}
