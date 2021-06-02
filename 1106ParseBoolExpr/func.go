package parseBoolExpr

func parseBoolExpr(expression string) bool {
	if expression == "" {
		return false
	}

	switch expression[0] {
	case 't':
		return true
	case 'f':
		return false
	case '!':
		expressions, _ := getContentInBracket(expression, 1)
		if len(expressions) != 1 {
			return false
		}
		return !parseBoolExpr(expressions[0])
	case '&':
		flag := true
		expressions, _ := getContentInBracket(expression, 1)
		for _, s := range expressions {
			flag = parseBoolExpr(s) && flag
			if !flag {
				return flag
			}
		}
		return flag
	case '|':
		flag := false
		expressions, _ := getContentInBracket(expression, 1)
		for _, s := range expressions {
			flag = parseBoolExpr(s) || flag
			if flag {
				return flag
			}
		}
		return flag
	default:
		return true
	}
}

func getContentInBracket(s string, start int) (expressions []string, end int) {
	if s == "" || start < 0 || start >= len(s) || s[start] != '(' {
		return
	}

	stack := 1
	for i := start + 1; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack++
		case ')':
			stack--
			if stack == 0 {
				expressions = append(expressions, s[start+1:i])
				end = i + 1
				break
			}
		case ',':
			if stack == 1 {
				expressions = append(expressions, s[start+1:i])
				start = i
			}
		}
	}
	return
}
