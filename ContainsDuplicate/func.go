package ContainsDuplicate

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	store := map[int]bool{}
	for _, v := range nums {
		if store[v] {
			return true
		}
		store[v] = true
	}
	return false
}
