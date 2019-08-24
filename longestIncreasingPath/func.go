package longestIncreasingPath

//https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix
func longestIncreasingPath(matrix [][]int) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	if cols == 0 {
		return 0
	}
	info := make([][]int, rows)
	for row := 0; row < rows; row++ {
		info[row] = make([]int, cols)
		for col := 0; col < cols; col++ {
			info[row][col] = -1
		}
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			setLengthInfoMatrix(matrix, info, row, col)
		}
	}
	max := 0
	for row := range info {
		for col := range info[row] {
			if info[row][col] > max {
				max = info[row][col]
			}
		}
	}
	return max + 1 //add itself
}

func setLengthInfoMatrix(matrix, info [][]int, row, col int) {
	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) {
		return
	}
	if info[row][col] != -1 {
		return
	}
	info[row][col] = 0
	up, down, left, right := 0, 0, 0, 0
	if row > 0 && matrix[row-1][col] > matrix[row][col] {
		setLengthInfoMatrix(matrix, info, row-1, col)
		up = info[row-1][col] + 1
	}
	if col < len(matrix[0])-1 && matrix[row][col+1] > matrix[row][col] {
		setLengthInfoMatrix(matrix, info, row, col+1)
		right = info[row][col+1] + 1
	}
	if row < len(matrix)-1 && matrix[row+1][col] > matrix[row][col] {
		setLengthInfoMatrix(matrix, info, row+1, col)
		down = info[row+1][col] + 1
	}
	if col > 0 && matrix[row][col-1] > matrix[row][col] {
		setLengthInfoMatrix(matrix, info, row, col-1)
		left = info[row][col-1] + 1
	}
	max := up
	if down > max {
		max = down
	}
	if right > max {
		max = right
	}
	if left > max {
		max = left
	}
	info[row][col] = max
}
