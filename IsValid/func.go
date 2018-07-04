package isValid

func isValid(s string) bool {
	stack := []rune{}
	pair := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, c := range s {
		if len(stack) == 0 || c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if pair[stack[len(stack)-1]] == c {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
