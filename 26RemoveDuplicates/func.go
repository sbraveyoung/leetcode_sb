package removeDuplicates

func removeDuplicates(nums []int) int {
	//[1,1,2]
	if len(nums) <= 0 {
		return 0
	}
	count := 1
	i, j := 0, 0
	for j = 1; j < len(nums); j++ {
		if nums[j] == nums[j-1] {
			continue
		}
		count++
		i++
		nums[i] = nums[j]
	}
	return count
}
