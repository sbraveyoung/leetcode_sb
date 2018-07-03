package countPrimes

func countPrimes(n int) int {
	buf := make([]int, n)
label:
	for cursor := 2; cursor < n; cursor++ {
		if buf[cursor] == -1 {
			continue label
		}
		for p := cursor * 2; p < n; p += cursor {
			buf[p] = -1
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if buf[i] != -1 {
			count++
		}
	}
	return count
}
