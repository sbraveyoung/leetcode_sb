package main

func threeSum(nums []int) [][]int {
	var ret [][]int
	if len(nums) < 3 {
		return ret
	}
	if has3zero(nums) {
		ret = append(ret, []int{0, 0, 0})
	}
	QuickSort(nums)
	index := -1
	for i, _ := range nums {
		if nums[i] >= 0 {
			index = i
			break
		}
	}
	if index == -1 {
		return ret
	}
	for i := 0; i < index; i++ {
		for j := index; j < len(nums); j++ {
			sub := -1
			if nums[i]+nums[j] == 0 {
				sub = FindZero(nums, i, j)
			} else if nums[i]+nums[j] > 0 {
				sub = Find3thNum(nums, 0, index, nums[i]+nums[j], false, i)
			} else {
				sub = Find3thNum(nums, index, len(nums), nums[i]+nums[j], true, j)
			}
			if sub != -1 {
				var arr []int
				if sub < i {
					arr = []int{nums[sub], nums[i], nums[j]}
				} else if sub > j {
					arr = []int{nums[i], nums[j], nums[sub]}
				} else {
					arr = []int{nums[i], nums[sub], nums[j]}
				}
				if isUnique(ret, arr) {
					ret = append(ret, arr)
				}
			}
		}
	}
	return ret
}

func has3zero(nums []int) bool {
	count := 0
	for _, value := range nums {
		if value == 0 {
			count++
		}
	}
	return count >= 3
}

func FindZero(nums []int, NotEquali, NotEqualj int) int {
	for index, _ := range nums {
		if index == NotEquali || index == NotEqualj {
			continue
		}
		if nums[index] == 0 {
			return index
		}
	}
	return -1
}

func isUnique(arrs [][]int, arr []int) bool {
	for _, item := range arrs {
		if arr[0] == item[0] && arr[1] == item[1] && arr[2] == item[2] {
			return false
		}
	}
	return true
}

func Find3thNum(nums []int, left, right int, SumofOtherNum int, PositiveofFindNum bool, NotEqual int) int { //sub
	for left < right {
		mid := left + (right-left)/2
		if nums[mid]+SumofOtherNum == 0 {
			if mid != NotEqual {
				return mid
			} else {
				if mid-1 >= left && nums[mid-1] == nums[mid] {
					return mid - 1
				}
				if mid+1 < right && nums[mid+1] == nums[mid] {
					return mid + 1
				}
				return -1
			}
		}
		if nums[mid]+SumofOtherNum > 0 {
			if mid == NotEqual {
				right = mid - 1
			} else {
				right = mid
			}
		} else {
			if mid == NotEqual {
				left = mid + 2
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

func QuickSort(nums []int) {
	QSort(nums, 0, len(nums)-1)
}

func QSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := Partition(nums, left, right)
	QSort(nums, left, mid-1)
	QSort(nums, mid+1, right)
}

func Partition(nums []int, left, right int) int {
	if left >= right {
		return left
	}
	value := nums[left]
	i, j := left, right
	for i < j {
		for i < j && value <= nums[j] {
			j--
		}
		if i < j {
			nums[i] = nums[j]
			i++
		}
		for i < j && nums[i] <= value {
			i++
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}
	}
	nums[i] = value
	return i
}
