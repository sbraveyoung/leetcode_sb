package calculate

import "strconv"

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	return calc(s, 0, '+')
}

func calc(s string, now int, label byte) int {
	for index := 0; index < len(s); {
		switch s[index] {
		case '(':
			jndex := findMyRight(s[index:])
			next := calculate(s[index+1 : index+jndex])
			now = do(now, next, label)
			index += (jndex + 1)
		case '+', '-':
			label = s[index]
			index++
		case ' ':
			index++
		default: //num
			var i int
			for i = index; i < len(s); i++ {
				if s[i] < '0' || s[i] > '9' {
					break
				}
			}
			next, _ := strconv.Atoi(s[index:i])
			now = do(now, next, label)
			index = i
		}
	}
	return now
}

func do(now, next int, label byte) int {
	switch label {
	case '+':
		now += next
	case '-':
		now -= next
	default:
	}
	return now
}

func findMyRight(s string) (index int) {
	count := 0
	for i, v := range s {
		if v == '(' {
			count++
		} else if v == ')' {
			count--
		}
		if count == 0 {
			return i
		}
	}
	return -1
}
