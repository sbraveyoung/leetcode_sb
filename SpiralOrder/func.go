package spiralOrder

func spiralOrder(matrix [][]int) (ret []int) {
	if len(matrix) == 0 {
		return
	}
	left, top, right, bottom := 0, 0, len(matrix[0])-1, len(matrix)-1
	for left <= right && top <= bottom {
		if left <= right {
			ret = append(ret, matrix[top][left:right+1]...)
			top++
		} else {
			break
		}

		if top <= bottom {
			for row := top; row <= bottom; row++ {
				ret = append(ret, matrix[row][right])
			}
			right--
		} else {
			break
		}

		if left <= right {
			for col := right; col >= left; col-- {
				ret = append(ret, matrix[bottom][col])
			}
			bottom--
		} else {
			break
		}

		if top <= bottom {
			for row := bottom; row >= top; row-- {
				ret = append(ret, matrix[row][left])
			}
			left++
		} else {
			break
		}
	}
	return
}
