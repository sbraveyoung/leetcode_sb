package isPowerOfThree

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	for {
		if n == 1 {
			return true
		}
		if n%3 == 0 {
			n /= 3
		} else {
			return false
		}
	}
}
