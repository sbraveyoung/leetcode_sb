package moveZeroes

func moveZeroes(nums []int) {
	i, j := 0, 0
	length := len(nums)
	for ; j < length; j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}
	for ; i < length; i++ {
		nums[i] = 0
	}
}
