package rotateImage

func rotate(matrix [][]int) {
	//because the matrix is n x n
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
		}
	}
}
