package reverseStr

import "fmt"

func reverseStr(s string, k int) string {
	var str string
	if s == "" || k <= 1 {
		return s
	}
	for index := 0; index < len(s); {
		left, right := index, index+k-1
		if right >= len(s) {
			right = len(s) - 1
		}
		for i := right; i >= left; i-- {
			str = fmt.Sprintf("%s%c", str, s[i])
		}
		left, right = right+1, right+k
		if left >= len(s) {
			break
		}
		if right >= len(s) {
			right = len(s) - 1
		}
		for i := left; i <= right; i++ {
			str = fmt.Sprintf("%s%c", str, s[i])
		}
		index = right + 1
	}
	return str
}
