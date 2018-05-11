package SingleNumber

func singleNumber(nums []int) int {
	var xor int = 0
	for _, value := range nums {
		xor ^= value
	}
	return xor
}
