package isHappy

func isHappy(n int) bool {
	// 1 -> true
	// 2 -> 4,16,37,58,89,145,42,20,4 -> false
	// 3 -> 9,81,65,61,37 -> false
	// 4 -> false
	// 5 -> 25,29,85 -> false
	// 6 -> 36,45,41,17,50,25 -> false
	// 7 -> 49,97,130,10 -> true
	// 8 -> 64,52 -> false
	// 9 -> 81 -> false
	//10 -> true
	m := map[int]bool{1: true, 2: false, 3: false, 4: false, 5: false, 6: false, 7: true, 8: false, 9: false, 10: true}
	for n != 1 {
		next := getNext(n)
		for next%10 == 0 {
			next /= 10
		}
		if value, ok := m[next]; ok {
			return value
		}
		n = next
	}
	return true
}

func getNext(n int) int {
	ret := 0
	for 0 != n {
		ret += ((n % 10) * (n % 10))
		n /= 10
	}
	return ret
}
