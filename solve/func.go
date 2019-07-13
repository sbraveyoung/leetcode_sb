package solve

//https://leetcode-cn.com/problems/surrounded-regions/

func solve(board [][]byte) {
	rows := len(board)
	if rows <= 2 {
		return
	}
	cols := len(board[0])
	if cols <= 2 {
		return
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if (row != 0 && row != rows-1) && (col != 0 && col != cols-1) {
				continue
			}
			i, j := row, col
			if board[i][j] == 'O' {
				do(board, i, j)
			}
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'O' {
				board[row][col] = 'X'
			}
		}
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'P' {
				board[row][col] = 'O'
			}
		}
	}
}

func do(board [][]byte, i, j int) {
	board[i][j] = 'P'
	if i-1 >= 0 && board[i-1][j] == 'O' {
		do(board, i-1, j)
	}
	if j+1 < len(board[i]) && board[i][j+1] == 'O' {
		do(board, i, j+1)
	}
	if i+1 < len(board) && board[i+1][j] == 'O' {
		do(board, i+1, j)
	}
	if j-1 >= 0 && board[i][j-1] == 'O' {
		do(board, i, j-1)
	}
}
