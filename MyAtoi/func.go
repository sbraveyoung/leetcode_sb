package myAtoi

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	m := map[rune]int64{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	var ret int64
	var positive int64 = 1
	switch str[0] {
	case '-':
		positive = -1
		str = str[1:]
	case '+':
		positive = 1
		str = str[1:]
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		break
	default:
		return 0
	}
	for _, c := range str {
		if c < '0' || c > '9' {
			return int(ret * positive)
		}
		if ret*10+m[c] > math.MaxInt32 {
			if positive == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		ret = ret*10 + m[c]
	}
	return int(ret * positive)
}
