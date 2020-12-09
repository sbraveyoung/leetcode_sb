package splitIntoFibonacci

import (
	"strconv"
)

func splitIntoFibonacci(S string) []int {
	ret := []int{}
	if len(S) < 3 {
		return ret
	}

loopa:
	for a0, a1 := 0, 1; a1 <= len(S)-2; a1++ {
		a, erra := strconv.ParseInt(S[a0:a1], 10, 32)
		if erra != nil || (a != 0 && S[a0] == '0') {
			break loopa
		}
	loopb:
		for b0, b1 := a1, a1+1; b1 <= len(S)-1; b1++ {
			b, errb := strconv.ParseInt(S[b0:b1], 10, 32)
			if errb != nil || (b != 0 && S[b0] == '0') {
				break loopb
			}
			if split(S, a, b, b1, &ret) {
				break loopa
			}
		}
	}
	return ret
}

func split(S string, a, b int64, c0 int, ret *[]int) (ok bool) {
	if c0 == len(S) {
		return true
	}
loopc:
	for c1 := c0 + 1; c1 <= len(S); c1++ {
		c, errc := strconv.ParseInt(S[c0:c1], 10, 32)
		if errc != nil || (c != 0 && S[c0] == '0') {
			break loopc
		}
		if a+b < c {
			break loopc
		}
		if a+b == c {
			if c1 == len(S) {
				*ret = append(*ret, int(a), int(b), int(c))
				return true
			} else {
				*ret = append(*ret, int(a))
				ok = split(S, b, c, c1, ret)
				if !ok {
					*ret = (*ret)[:len(*ret)-1]
				} else {
					return true
				}
			}
		}
	}
	return false
}
