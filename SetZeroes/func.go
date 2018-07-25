package setZeroes

const (
	ROW = iota
	COL
)

func setZeroes(matrix [][]int) {
	row := map[int]bool{}
	col := map[int]bool{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i], col[j] = true, true
			}
		}
	}
	for r, _ := range row {
		setRowOrColToZeroes(matrix, r, ROW)
	}
	for c, _ := range col {
		setRowOrColToZeroes(matrix, c, COL)
	}
}

func setRowOrColToZeroes(matrix [][]int, index int, direction int) {
	if direction == ROW {
		for i := 0; i < len(matrix[index]); i++ {
			matrix[index][i] = 0
		}
	} else {
		for i := 0; i < len(matrix); i++ {
			matrix[i][index] = 0
		}
	}
}
