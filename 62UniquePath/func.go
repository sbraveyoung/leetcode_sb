package uniquePaths

func uniquePaths(m int, n int) int {
	dp := make([][]int, m, m)
	for r := 0; r < m; r++ {
		dp[r] = make([]int, n, n)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if r == 0 || c == 0 {
				dp[r][c] = 1
			}
		}
	}

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			dp[r][c] = dp[r][c-1] + dp[r-1][c]
		}
	}

	return dp[m-1][n-1]
}
