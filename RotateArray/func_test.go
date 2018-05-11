package RotateArray

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "second",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    0,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "third",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    7,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "fourth",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    10,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !checkTwoArray(tt.args.nums, tt.want) {
				t.Errorf("rotate() = %v,want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func checkTwoArray(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}
