package generateMatrix

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	if n <= 0 {
		return matrix
	}

	ct := counter{c: 1}
	left, top, right, bottom := 0, 0, len(matrix[0])-1, len(matrix)-1
	for left <= right && top <= bottom {
		if left <= right {
			for col := left; col <= right; col++ {
				matrix[top][col] = ct.get()
			}
			top++
		} else {
			break
		}

		if top <= bottom {
			for row := top; row <= bottom; row++ {
				matrix[row][right] = ct.get()
			}
			right--
		} else {
			break
		}

		if left <= right {
			for col := right; col >= left; col-- {
				matrix[bottom][col] = ct.get()
			}
			bottom--
		} else {
			break
		}

		if top <= bottom {
			for row := bottom; row >= top; row-- {
				matrix[row][left] = ct.get()
			}
			left++
		} else {
			break
		}
	}
	return matrix
}

type counter struct {
	c int
}

func (ct *counter) get() int {
	ct.c++
	return ct.c - 1
}
