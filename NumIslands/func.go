package numIslands

func numIslands(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	if col == 0 {
		return 0
	}
	uf := newUnionFind(grid, row, col)
	uf.union(grid, row, col)
	return uf.count
}

type UnionFind struct {
	count int
	arr   []int
}

func newUnionFind(grid [][]byte, row, col int) (uf UnionFind) {
	uf.arr = make([]int, row*col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				uf.count++
			}
			uf.arr[i*col+j] = i*col + j
		}
	}
	return
}

func (uf *UnionFind) union(grid [][]byte, row, col int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				if i > 0 && grid[i][j] == grid[i-1][j] {
					b := (i-1)*col + j
					a := i*col + j
					ar := uf.findRoot(a)
					br := uf.findRoot(b)
					if ar != br {
						uf.arr[br] = ar
						uf.count--
					}
				}
				if j > 0 && grid[i][j] == grid[i][j-1] {
					b := i*col + j - 1
					a := i*col + j
					ar := uf.findRoot(a)
					br := uf.findRoot(b)
					if ar != br {
						uf.arr[br] = ar
						uf.count--
					}
				}
				if i < row-1 && grid[i][j] == grid[i+1][j] {
					b := (i+1)*col + j
					a := i*col + j
					ar := uf.findRoot(a)
					br := uf.findRoot(b)
					if ar != br {
						uf.arr[br] = ar
						uf.count--
					}
				}
				if j < col-1 && grid[i][j] == grid[i][j+1] {
					b := i*col + j + 1
					a := i*col + j
					ar := uf.findRoot(a)
					br := uf.findRoot(b)
					if ar != br {
						uf.arr[br] = ar
						uf.count--
					}
				}
			}
		}
	}
}

func (uf *UnionFind) findRoot(pos int) int {
	for pos != uf.arr[pos] {
		pos = uf.arr[pos]
	}
	return pos
}
