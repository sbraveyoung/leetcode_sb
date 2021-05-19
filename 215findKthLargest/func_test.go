package findKthLargest

import "testing"

func Test_findKthNum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				nums: []int{5, 4, 3, 2, 1, 6, 7, 8, 9, 0},
				k:    1,
			},
			wantRet: 9,
		},
		{
			name: "second",
			args: args{
				nums: []int{5, 4, 3, 2, 1, 6, 7, 8, 9, 0},
				k:    10,
			},
			wantRet: 0,
		},
		{
			name: "third",
			args: args{
				nums: []int{5, 4, 3, 2, 1, 6, 7, 8, 9, 0},
				k:    5,
			},
			wantRet: 5,
		},
		{
			name: "fourth",
			args: args{
				nums: []int{1},
				k:    1,
			},
			wantRet: 1,
		},
		{
			name: "fifth",
			args: args{
				nums: []int{99, 99},
				k:    1,
			},
			wantRet: 99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := findKthLargest(tt.args.nums, tt.args.k); gotRet != tt.wantRet {
				t.Errorf("findKthNum() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
