package canJump

func canJump(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return true
	}
	arrive := nums[0]
	for i := 1; i < length; i++ {
		if arrive >= i {
			if i+nums[i] > arrive {
				arrive = i + nums[i]
			}
		}
	}
	return arrive >= length-1
}
