package IsPalindrome

func isPalindrome(s string) bool {
	length := len(s)
	for i, j := 0, length-1; i < j; {
		for ; i < j && !isCharOrNum(s[i]); i++ {
		}
		for ; i < j && !isCharOrNum(s[j]); j-- {
		}
		if i < j && !isEqualCase(s[i], s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isCharOrNum(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	if b >= 'a' && b <= 'z' {
		return true
	}
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

func isEqualCase(x, y byte) bool {
	step := 0
	if 'A' > 'a' {
		step = 'A' - 'a'
	} else {
		step = 'a' - 'A'
	}
	a, b := int(x), int(y)
	if a > b {
		a, b = b, a
	}
	if a == b {
		return true
	} else { //a<b
		if a >= '0' && a <= '9' {
			return false
		} else {
			if a+step == b {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
