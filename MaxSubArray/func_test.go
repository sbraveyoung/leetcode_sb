package maxSubArray

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "second",
			args: args{
				nums: []int{-1, -2, -3},
			},
			want: -1,
		},
		{
			name: "third",
			args: args{
				nums: []int{1, -2, 3, 4, 0, -5, 6, -7, -8, 9, 10},
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
