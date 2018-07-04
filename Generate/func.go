package generate

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		ret[i-1] = make([]int, i)
		ret[i-1][0], ret[i-1][i-1] = 1, 1
		for j := 1; j < i-1; j++ {
			ret[i-1][j] = ret[i-2][j-1] + ret[i-2][j]
		}
	}
	return ret
}
