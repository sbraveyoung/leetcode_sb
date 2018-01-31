package main

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}
	chars := [256]byte{}
	chars[s[0]] = 1
	length := []int{1}
	sub := map[byte]int{s[0]: 0}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			length = append(length, 1)
		} else {
			if chars[s[i]] == 0 {
				chars[s[i]]++
				length = append(length, length[i-1]+1)
			} else {
				if i-sub[s[i]] > length[i-1]+1 {
					length = append(length, length[i-1]+1)
				} else {
					length = append(length, i-sub[s[i]])
				}
			}
		}
		sub[s[i]] = i
	}
	max := 0
	for i := 0; i < len(length); i++ {
		if max < length[i] {
			max = length[i]
		}
	}
	return max
}
