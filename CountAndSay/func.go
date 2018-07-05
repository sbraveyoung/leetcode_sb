package countAndSay

import "strconv"

func countAndSay(n int) string {
	prev, cur := "1", "1"
	for i := 1; i < n; i++ {
		cur = ""
		num, value, nextPos := 0, "", 0
		for nextPos != -1 {
			num, value, nextPos = calc(prev, nextPos)
			cur += strconv.Itoa(num) + value
		}
		prev = cur
	}
	return cur
}

func calc(s string, pos int) (num int, value string, nextPos int) {
	value = s[pos : pos+1]
	i := pos + 1
	for ; i < len(s); i++ {
		if s[i] != value[0] {
			break
		}
	}
	if i == len(s) { //normal
		nextPos = -1
	} else { //break
		nextPos = i
	}
	return i - pos, value, nextPos
}
