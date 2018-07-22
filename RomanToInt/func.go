package romanToInt

func romanToInt(s string) int {
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	num := 0
	for index := 0; index < len(s); {
		if index != len(s)-1 {
			switch s[index] {
			case 'I':
				if s[index+1] == 'V' {
					num += 4
					index += 2
					continue
				} else if s[index+1] == 'X' {
					num += 9
					index += 2
					continue
				}
			case 'X':
				if s[index+1] == 'L' {
					num += 40
					index += 2
					continue
				} else if s[index+1] == 'C' {
					num += 90
					index += 2
					continue
				}
			case 'C':
				if s[index+1] == 'D' {
					num += 400
					index += 2
					continue
				} else if s[index+1] == 'M' {
					num += 900
					index += 2
					continue
				}
			}
		}
		num += m[s[index]]
		index++
	}
	return num
}
