package findKthLargest

func partition(nums []int, left, right int) (mid int) {
	if len(nums) == 0 || left >= right {
		return 0
	}

	tmp := left
	for left <= right {
		for left <= right && nums[left] >= nums[tmp] {
			left++
		}
		for left <= right && nums[right] <= nums[tmp] {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	left--
	nums[left], nums[tmp] = nums[tmp], nums[left]
	return left

}

func findKthNum(nums []int, k int) (ret int) {
	if len(nums) == 0 || k < 0 || k > len(nums) {
		return -1
	}

	mid := partition(nums, 0, len(nums)-1)
	if mid == k {
		return nums[mid]
	} else if mid > k {
		return findKthNum(nums[:mid], k)
	} else {
		return findKthNum(nums[mid+1:], k-mid-1)
	}
}

func findKthLargest(nums []int, k int) int {
	return findKthNum(nums, k-1)
}
