package maxSubArray

func maxSubArray(nums []int) int {
	//eg: [-2,1,-3,4,-1,2,1,-5,4]
	length := len(nums)
	max := 0
	dp := make([]int, length, length)
label:
	for i := 0; i < length; i++ {
		if i == 0 {
			dp[i] = nums[i]
			max = dp[i]
			continue label
		}
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
