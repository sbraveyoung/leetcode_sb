package MaxProfit

func maxProfit(prices []int) int {
	//[7,1,5,3,6,4] 7
	//[-6,4,-2,3,-2]

	//[1,2,3,4,5] 4
	//[1,1,1,1]

	//[7,6,4,3,1] 0
	//[-1,-2,-1,-2] 0
	subPrices := []int{}
	for i := 1; i < len(prices); i++ {
		subPrices = append(subPrices, prices[i]-prices[i-1])
	}
	return MaxPositiveSum(subPrices)
}

func MaxPositiveSum(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	mid := length / 2
	leftSum := MaxPositiveSum(nums[:mid])
	rightSum := MaxPositiveSum(nums[mid+1:])
	curSum := 0
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			curSum += nums[i]
		}
	}
	max := curSum
	if leftSum > max {
		max = leftSum
	}
	if rightSum > max {
		max = rightSum
	}
	return max
}
