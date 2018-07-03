package rob

func rob(nums []int) int {
	//state transfer equation:
	// room num:      _ 1 2 3 1
	// money:pr,cr    0 _ _ _ _
	//      :pnr,cr   0 1 2 4 3
	//      :pr,cnr   0 0 1 2 4
	//      :pnr,cnr  0 0 0 1 2

	// pr:  prev rob
	// pnr: prev not rob
	// cr:  current rob
	// cnr: current not rob

	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	}
	dp := make([][4]int, length+1)
	money := 0
	for j := 0; j < 4; j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= length; i++ {
		pnr_max := 0
		if dp[i-1][2] > dp[i-1][3] {
			pnr_max = dp[i-1][2]
		} else {
			pnr_max = dp[i-1][3]
		}

		dp[i][0] = 0
		dp[i][1] = pnr_max + nums[i-1]
		dp[i][2] = dp[i-1][1]
		dp[i][3] = pnr_max
	}
	for j := 1; j < 4; j++ {
		if dp[length][j] > money {
			money = dp[length][j]
		}
	}
	return money
}
