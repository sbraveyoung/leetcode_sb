package moveZeroes

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums   []int
		origin []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
		},
		{
			name: "second",
			args: args{
				nums: []int{0, 0, 3, 4, 5},
			},
		},
		{
			name: "third",
			args: args{
				nums: []int{1, 2, 3, 0, 0},
			},
		},
		{
			name: "fourth",
			args: args{
				nums: []int{0, 2, 0, 4, 0},
			},
		},
		{
			name: "fifth",
			args: args{
				nums: []int{1, 0, 3, 0, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.origin = tt.args.nums
			moveZeroes(tt.args.nums)
			if !check(tt.args.origin, tt.args.nums) {
				t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.args.origin)
			}
		})
	}
}

func check(origin, nums []int) bool {
	i, j := 0, 0
	length := len(origin)
	if length != len(nums) {
		return false
	}
	for ; i < length; i++ {
		if origin[i] == 0 {
			continue
		}
		if origin[i] != nums[j] {
			return false
		}
		j++
	}
	for ; j < length; j++ {
		if nums[j] != 0 {
			return false
		}
	}
	return true
}
