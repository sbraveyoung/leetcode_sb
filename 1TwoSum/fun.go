package main

func twoSum(nums []int, target int) []int {
	var sub []int
	for index, _ := range nums {
		sub = append(sub, index)
	}
	QuickSort(nums, sub)
	var ret []int
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == target {
			ret = append(ret, sub[i])
			ret = append(ret, sub[j])
			if ret[0] > ret[1] {
				ret[0], ret[1] = ret[1], ret[0]
			}
			return ret
		}
		if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}
	}
	return ret
}

func QuickSort(nums []int, sub []int) {
	QSort(nums, 0, len(nums)-1, sub)
}

func QSort(nums []int, left, right int, sub []int) {
	if left >= right {
		return
	}
	mid := Partition(nums, left, right, sub)
	QSort(nums, left, mid-1, sub)
	QSort(nums, mid+1, right, sub)
}

func Partition(nums []int, left, right int, sub []int) int {
	if left >= right {
		return left
	}
	value := nums[left]
	vs := sub[left]
	i, j := left, right
	for i < j {
		for i < j && value <= nums[j] {
			j--
		}
		if i < j {
			nums[i] = nums[j]
			sub[i] = sub[j]
			i++
		}
		for i < j && nums[i] <= value {
			i++
		}
		if i < j {
			nums[j] = nums[i]
			sub[j] = sub[i]
			j--
		}
	}
	nums[i] = value
	sub[i] = vs
	return i
}
