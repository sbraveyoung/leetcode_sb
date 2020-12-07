package matrixScore

import (
	"math"
)

func matrixScore(A [][]int) int {
	if len(A) == 0 || len(A[0]) == 0 {
		return 0
	}

	for i, a := range A {
		if a[0] == 0 {
			for j := 0; j < len(a); j++ {
				reverse(A, i, j)
			}
		}
	}

	for j := 1; j < len(A[0]); j++ {
		sub10 := 0
		for i := 0; i < len(A); i++ {
			if A[i][j] == 1 {
				sub10++
			} else {
				sub10--
			}
		}
		if sub10 < 0 { //count0 is more than count1
			for i := 0; i < len(A); i++ {
				reverse(A, i, j)
			}
		}
	}
	return calc(A)
}

func reverse(A [][]int, i, j int) {
	if A[i][j] == 0 {
		A[i][j] = 1
	} else {
		A[i][j] = 0
	}
}

func calc(A [][]int) int {
	res := 0
	for i := 0; i < len(A); i++ {
		rRes := 0
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 1 {
				rRes += int(math.Pow(float64(2), float64(len(A[i])-1-j)))
			}
		}
		res += rRes
	}
	return res
}
