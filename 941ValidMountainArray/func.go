package validMountainArray

//one
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	up := true
	for index := 1; index < len(A); index++ {
		if A[index-1] == A[index] {
			return false
		}
		if up {
			if A[index-1] > A[index] {
				if index == 1 {
					return false
				}
				up = false
			} else {
				continue
			}
		} else {
			if A[index-1] > A[index] {
				continue
			} else {
				return false
			}
		}
	}
	return !up
}

//two
//func validMountainArray(A []int) bool {
//	if len(A) < 3 {
//		return false
//	}
//
//	mid := -1
//	for start := 1; start < len(A); start++ {
//		if A[start-1] >= A[start] {
//			mid = start - 1
//			break
//		}
//	}
//	if mid == -1 {
//		return false
//	}
//	for end := len(A) - 1; end > 0; end-- {
//		if A[end] >= A[end-1] {
//			if mid == end {
//				return true
//			} else {
//				return false
//			}
//		}
//	}
//	return false
//}
