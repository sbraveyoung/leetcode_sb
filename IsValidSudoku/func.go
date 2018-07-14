package isValidSudoku

func isValidSudoku(board [][]byte) bool {
	dict := map[byte]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, '.': 0}
	row := [10][10]int{}
	col := [10][10]int{}
	block := [3][3][10]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if dict[board[i][j]] == 0 {
				continue
			}
			if row[i][dict[board[i][j]]] == dict[board[i][j]] {
				return false
			}
			if col[j][dict[board[i][j]]] == dict[board[i][j]] {
				return false
			}
			if block[i/3][j/3][dict[board[i][j]]] == dict[board[i][j]] {
				return false
			}
			row[i][dict[board[i][j]]] = dict[board[i][j]]
			col[j][dict[board[i][j]]] = dict[board[i][j]]
			block[i/3][j/3][dict[board[i][j]]] = dict[board[i][j]]
		}
	}
	return true
}
