package FloodFill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	return modify(image, sr, sc, image[sr][sc], newColor)
}

func modify(image [][]int, sr int, sc int, originColor int, newColor int) [][]int {
	if len(image) == 0 {
		return image
	}
	if originColor == newColor {
		return image
	}
	raw, column := len(image), len(image[0])
	if image[sr][sc] == originColor {
		image[sr][sc] = newColor
	}
	for c := sc + 1; c < column; c++ {
		if image[sr][c] == originColor {
			modify(image, sr, c, originColor, newColor)
		} else {
			break
		}
	}
	for c := sc - 1; c >= 0; c-- {
		if image[sr][c] == originColor {
			modify(image, sr, c, originColor, newColor)
		} else {
			break
		}
	}
	for r := sr + 1; r < raw; r++ {
		if image[r][sc] == originColor {
			modify(image, r, sc, originColor, newColor)
		} else {
			break
		}
	}
	for r := sr - 1; r >= 0; r-- {
		if image[r][sc] == originColor {
			modify(image, r, sc, originColor, newColor)
		} else {
			break
		}
	}
	return image
}
