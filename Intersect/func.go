package intersect

func intersect(nums1 []int, nums2 []int) []int {
	store := map[int]int{}
	ret := []int{}
	for _, v := range nums1 {
		store[v]++
	}
	for _, v := range nums2 {
		if store[v] != 0 {
			ret = append(ret, v)
			store[v]--
		}
	}
	return ret
}
