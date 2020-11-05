package canThreePartsEqualSum

func canThreePartsEqualSum(A []int) bool {
	if len(A) < 3 {
		return false
	}
	left, right := 0, len(A)-1
	sum1, sum2, sum3 := A[left], 0, A[right]
	for i := left + 1; i < right; i++ {
		sum2 += A[i]
	}

	for left < right-1 {
		if sum1 == sum2 && sum2 == sum3 {
			return true
		}

		//中间大，两头小：将中间往小的一边削
		if sum2 >= sum1 && sum2 >= sum3 {
			if sum1 > sum3 {
				right--
				sum2 -= A[right]
				sum3 += A[right]
				continue
			} else {
				left++
				sum2 -= A[left]
				sum1 += A[left]
				continue
			}
		}

		//单调递增：中间往左边削
		if sum1 <= sum2 && sum2 <= sum3 {
			left++
			sum2 -= A[left]
			sum1 += A[left]
			continue
		}

		//中间小，两边大：将大的一头往中间削
		if sum1 >= sum2 && sum3 >= sum2 {
			if sum1 > sum3 {
				left++
				sum2 -= A[left]
				sum1 += A[left]
				continue
			} else {
				right--
				sum2 -= A[right]
				sum3 += A[right]
				continue
			}
		}

		//单调递减：中间往右边削
		if sum1 >= sum2 && sum2 >= sum3 {
			right--
			sum2 -= A[right]
			sum3 += A[right]
			continue
		}
	}
	return false
}
