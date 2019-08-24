package intervalIntersection

func intervalIntersection(A, B [][]int) (C [][]int) {
	iterA, iterB := 0, 0
	for iterA < len(A) && iterB < len(B) {
		if A[iterA][0] < B[iterB][0] {
			if A[iterA][1] < B[iterB][0] {
				iterA++
				continue
			} else {
				if A[iterA][1] < B[iterB][1] {
					C = append(C, []int{B[iterB][0], A[iterA][1]})
					iterA++
				} else {
					C = append(C, B[iterB])
					iterB++
				}
			}
		} else {
			if A[iterA][0] <= B[iterB][1] {
				if A[iterA][1] < B[iterB][1] {
					C = append(C, A[iterA])
					iterA++
				} else {
					C = append(C, []int{A[iterA][0], B[iterB][1]})
					iterB++
				}
			} else {
				iterB++
				continue
			}
		}
	}
	return
}
