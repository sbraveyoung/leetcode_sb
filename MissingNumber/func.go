package missingNumber

func missingNumber(nums []int) int {
	var sum uint64
	max := 0
	exist_zero := false
	for _, v := range nums {
		sum += uint64(v)
		if v > max {
			max = v
		}
		if v == 0 {
			exist_zero = true
		}
	}
	value := max
	should := 0
	if max%2 != 0 {
		should = max
		max--
	}
	should += (max + 1) * (max / 2)
	if should-int(sum) == 0 {
		if exist_zero {
			return value + 1
		} else {
			return 0
		}
	}
	return should - int(sum)
}
