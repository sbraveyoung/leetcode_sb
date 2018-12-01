package maxArea

func maxArea(height []int) int {
	maxCapacity := 0
	for i, j := 0, len(height)-1; i < j; {
		minHeight := 0
		if height[i] < height[j] {
			minHeight = height[i]
		} else {
			minHeight = height[j]
		}
		capacity := (j - i) * minHeight
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		if capacity > maxCapacity {
			maxCapacity = capacity
		}
	}
	return maxCapacity
}
