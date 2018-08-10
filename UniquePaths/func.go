package uniquePaths

func uniquePaths(m int, n int) int {
	matrix := [][]int{}
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, m))
		matrix[i][0] = 1
	}
	for i := 0; i < m; i++ {
		matrix[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			matrix[i][j] = matrix[i][j-1] + matrix[i-1][j]
		}
	}
	return matrix[n-1][m-1]
}
