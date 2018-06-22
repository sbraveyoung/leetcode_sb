package climbStairs

func climbStairs(n int) int {
	left, right := 1, 1
	for i := 2; i <= n; i++ {
		left, right = right, left+right
	}
	return right
}
