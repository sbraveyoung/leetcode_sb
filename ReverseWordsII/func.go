package reverseWords

import "fmt"

func reverseWords(s string) string {
	ret := ""
	for i := len(s) - 1; i >= 0; {
		i = findNoneSpace(s, i)
		start := findSpace(s, i)
		if start > i {
			break
		}
		ret = fmt.Sprintf("%s%s ", ret, s[start:i+1])
		i = start - 1
	}
	if len(ret) == 0 {
		return ""
	}
	return ret[:len(ret)-1]
}

func findSpace(bs string, index int) int {
	if index < 0 || index >= len(bs) {
		return -1
	}
	for index >= 0 {
		if bs[index] == ' ' {
			return index + 1
		}
		index--
	}
	return index + 1
}
func findNoneSpace(bs string, index int) int {
	if index < 0 || index >= len(bs) {
		return -1
	}
	for index >= 0 {
		if bs[index] != ' ' {
			return index
		}
		index--
	}
	return 0
}
