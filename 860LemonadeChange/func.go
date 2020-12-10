package lemonadeChange

func lemonadeChange(bills []int) bool {
	m := [3]int{}
	for _, bill := range bills {
		switch bill {
		case 5:
			m[0]++
		case 10:
			m[1]++
			if m[0] > 0 {
				m[0]--
			} else {
				return false
			}
		case 20:
			m[2]++
			if m[0] > 0 && m[1] > 0 {
				m[0]--
				m[1]--
			} else {
				if m[0] > 2 {
					m[0] -= 3
				} else {
					return false
				}
			}
		}
	}
	return true
}
