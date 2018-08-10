package canJump

func canJump(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return true
	}
	dp := make([]bool, length)
	dp[0] = true
	arrive := nums[0]
	for i := 1; i < length; i++ {
		if arrive >= i {
			dp[i] = true
		}
		if dp[i] {
			if i+nums[i] > arrive {
				arrive = i + nums[i]
			}
		}
	}
	return dp[length-1]
}
