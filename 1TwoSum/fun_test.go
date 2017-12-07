package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 3}
	ret := twoSum(nums, 6)
	if ret[0] != 0 && ret[1] != 2 {
		fmt.Println(ret[0], ret[1])
		t.Errorf("TestTwoSum fail.")
	}
}

func TestQuickSort(t *testing.T) {
	//nums := []int{8, 2, 4, 5, 6, 7, 3, 0, 1, 9}
	nums := []int{3, 2, 3}
	var sub []int
	for index, _ := range nums {
		sub = append(sub, index)
	}
	QuickSort(nums, sub)
	for i, _ := range nums {
		if i == len(nums)-1 {
			break
		}
		if nums[i] > nums[i+1] {
			fmt.Println(nums)
			fmt.Println(sub)
			t.Errorf("TestQuickSort fail.")
			break
		}
	}
}
