package MaxProfit

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	min := make([]int, length)
	min[0] = 0
	for i := 1; i < length; i++ {
		if prices[i] < prices[min[i-1]] {
			min[i] = i
		} else {
			min[i] = min[i-1]
		}
	}
	max := 0
	for i := 1; i < length; i++ {
		if prices[i]-prices[min[i]] > max {
			max = prices[i] - prices[min[i]]
		}
	}
	return max
}
