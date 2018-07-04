package hammingDistance

func hammingDistance(x int, y int) int {
	x ^= y
	count := 0
	if x < 0 {
		count++
		x &= 0x7fffffff
	}
	for x != 0 {
		if (x & 0x01) != 0 {
			count++
		}
		x >>= 1
	}
	return count
}
