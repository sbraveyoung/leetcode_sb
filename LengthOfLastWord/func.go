package lengthOfLastWord

func lengthOfLastWord(s string) int {
	cnt := 0
	if len(s) <= 0 {
		return cnt
	}
	index := len(s) - 1
	for ; index >= 0; index-- {
		if s[index] != ' ' {
			break
		}
	}
	for ; index >= 0; index-- {
		if s[index] == ' ' {
			break
		}
		cnt++
	}
	return cnt
}
