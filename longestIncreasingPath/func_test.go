package longestIncreasingPath

import "testing"

func Test_longestIncreasingPath(t *testing.T) {
	type args struct {
		matrix [][]int
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
				matrix: [][]int{
					{9, 9, 4},
					{6, 6, 8},
					{2, 1, 1},
				},
			},
			want: 4,
		},
		{
			name: "second",
			args: args{
				matrix: [][]int{
					{3, 4, 5},
					{3, 2, 6},
					{2, 2, 1},
				},
			},
			want: 4,
		},
		{
			name: "third",
			args: args{
				matrix: [][]int{
					{9, 8, 7},
					{2, 1, 6},
					{3, 4, 5},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath(tt.args.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
