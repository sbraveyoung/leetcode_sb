package main

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{1, 2, -2, -1}
	ret := threeSum(nums)
	if len(ret) != 0 {
		fmt.Println(ret)
		t.Errorf("TestTwoSum fail0.")
	}
	for _, i := range ret {
		if len(i) != 3 {
			fmt.Println(ret)
			t.Errorf("TestTwoSum fail1.")
		}
		if i[0]+i[1]+i[2] != 0 {
			fmt.Println(ret)
			t.Errorf("TestTwoSum fail2.")
		}
	}
	nums = []int{0, 0, 0}
	ret = threeSum(nums)
	if len(ret) != 1 {
		fmt.Println(ret)
		t.Errorf("TestTwoSum fail0.")
	}
	for _, i := range ret {
		if len(i) != 3 {
			fmt.Println(ret)
			t.Errorf("TestTwoSum fail1.")
		}
		if i[0]+i[1]+i[2] != 0 {
			fmt.Println(ret)
			t.Errorf("TestTwoSum fail2.")
		}
	}

	nums = []int{0, 1, 2}
	ret = threeSum(nums)
	if len(ret) != 0 {
		fmt.Println(ret)
		t.Errorf("TestTwoSum fail0.")
	}
	for _, i := range ret {
		if len(i) != 3 {
			fmt.Println(ret)
			t.Errorf("TestTwoSum fail1.")
		}
		if i[0]+i[1]+i[2] != 0 {
			fmt.Println(ret)
			t.Errorf("TestTwoSum fail2.")
		}
	}

	nums = []int{0, 1, -1, 0, 0}
	ret = threeSum(nums)
	if len(ret) != 2 {
		fmt.Println(ret)
		t.Errorf("TestTwoSum fail0.")
	}
	for _, i := range ret {
		if len(i) != 3 {
			fmt.Println(ret)
			t.Errorf("TestTwoSum fail1.")
		}
		if i[0]+i[1]+i[2] != 0 {
			fmt.Println(ret)
			t.Errorf("TestTwoSum fail2.")
		}
	}

	nums = []int{-1, 0, 1, 2, -1, -4}
	ret = threeSum(nums)
	if len(ret) != 2 {
		fmt.Println(ret)
		t.Errorf("TestTwoSum fail0.")
	}
	for _, i := range ret {
		if len(i) != 3 {
			fmt.Println(ret)
			t.Errorf("TestTwoSum fail1.")
		}
		if i[0]+i[1]+i[2] != 0 {
			fmt.Println(ret)
			t.Errorf("TestTwoSum fail2.")
		}
	}
}
