package countNegatives

func countNegatives(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	count := 0
	rows, cols := len(grid), len(grid[0])
	for row, col := 0, cols-1; row < rows && col >= 0; {
		if grid[row][col] < 0 {
			count += (rows - row)
			col--
		} else {
			row++
		}
	}
	return count
}
