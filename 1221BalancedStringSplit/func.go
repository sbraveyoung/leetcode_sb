package balancedStringSplit

func balancedStringSplit(s string) int {
	count := 0
	lc, rc := 0, 0
	for _, c := range s {
		switch c {
		case 'L':
			lc++
		case 'R':
			rc++
		}
		if lc == rc {
			count++
			lc, rc = 0, 0
		}
	}
	return count
}
