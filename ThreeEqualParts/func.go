package threeEqualParts

//https://leetcode-cn.com/problems/three-equal-parts/

func threeEqualParts(A []int) []int {
	none := []int{-1, -1}
	if len(A) < 3 {
		return none
	}

	prefixZero := 0
	var index int
	for index = 0; index < len(A); index++ {
		if A[index] != 0 {
			prefixZero = index
			break
		}
	}
	//all is 0
	if index == len(A) {
		return []int{0, 2}
	}

	var fs, fe, ss, se, ts, te int //firstStart,firstEnd ...
	fs = prefixZero
out:
	for fe = fs; fe < len(A)-2; fe++ {
		length := fe - fs + 1
		for ss = fe + 1; ss < len(A)-1; ss++ {
			if A[ss] != 0 {
				break
			}
		}
		for se = ss; se-ss < length; se++ {
			if se >= len(A) || A[se] != A[fs+se-ss] {
				continue out
			}
		}
		for ts = se; ts < len(A); ts++ {
			if A[ts] != 0 {
				break
			}
		}
		if ts+length != len(A) {
			continue out
		}
		for te = ts; te-ts < length; te++ {
			if te >= len(A) || A[te] != A[fs+te-ts] {
				continue out
			}
		}
		return []int{fe, se}
	}
	return none
}
