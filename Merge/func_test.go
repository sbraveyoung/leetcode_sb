package merge

import "testing"

func TestMerge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !check(tt.args.nums1, tt.args.m+tt.args.n) {
				t.Errorf("Merge() = %v", tt.args.nums1)
			}
		})
	}
}

func check(nums []int, n int) bool {
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}
