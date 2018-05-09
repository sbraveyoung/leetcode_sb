package StringToInteger

func myAtoi(str string) int {
	if len(str) <= 0 {
		return 0
	}
	var ret int
	var i int
	var start int
	m := map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
	s := false
	for i = 0; i < len(str); i++ {
		if str[i] != ' ' {
			break
		}
	}
	if i >= len(str) {
		return 0
	}
	if str[i] == '-' {
		i++
		s = true
	} else if str[i] == '+' {
		i++
	}
	for ; i < len(str); i++ {
		if str[i] != '0' {
			break
		}
	}
	if i >= len(str) {
		return 0
	}
	start = i
	for ; i < len(str); i++ {
		switch str[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			ret = ret*10 + m[str[i]]
		default:
			goto back
		}
		if isOverFlow, num := intOverFlow(str[start:i+1], s); isOverFlow {
			return num
		}
	}
back:
	if s {
		return -ret
	}
	return ret
}

func intOverFlow(str string, s bool) (bool, int) {
	if len(str) > 10 {
		if s {
			return true, -2147483648
		} else {
			return true, 2147483647
		}
	} else if len(str) < 10 {
		return false, 0
	}
	//len(str) equal to 10
	max := "2147483647"
	min := "2147483648"
	var num string
	if s {
		num = min
	} else {
		num = max
	}
	for i := 0; i < 10; i++ {
		if str[i] == num[i] {
			continue
		}
		if str[i] > num[i] {
			if s {
				return true, -2147483648
			} else {
				return true, 2147483647
			}
		} else {
			return false, 0
		}
	}
	return false, 0
}
