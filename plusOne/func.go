package plusOne

func plusOne(digits []int) []int {
	for index := len(digits) - 1; index >= 0; index-- {
		if digits[index] < 9 {
			digits[index]++
			return digits
		}
		digits[index] = 0
	}
	return append([]int{1}, digits...)
}
