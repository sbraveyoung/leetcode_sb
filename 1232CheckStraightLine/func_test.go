package checkStraightLine

import "testing"

func Test_checkStraightLine(t *testing.T) {
	type args struct {
		coordinates [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				coordinates: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{4, 5},
					{5, 6},
					{6, 7},
				},
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				coordinates: [][]int{
					{1, 1},
					{2, 2},
					{3, 4},
					{4, 5},
					{5, 6},
					{7, 7},
				},
			},
			want: false,
		},
		{
			name: "third",
			args: args{
				coordinates: [][]int{
					{0, 0},
					{0, 1},
					{0, 2},
					{0, 3},
					{0, 4},
				},
			},
			want: true,
		},
		{
			name: "fourth",
			args: args{
				coordinates: [][]int{
					{0, 0},
					{1, 0},
					{2, 0},
					{3, 0},
					{4, 0},
				},
			},
			want: true,
		},
		{
			name: "fifth",
			args: args{
				coordinates: [][]int{
					{0, 0},
					{1, 1},
					{2, 2},
					{3, 3},
					{4, 4},
				},
			},
			want: true,
		},
		{
			name: "sixth",
			args: args{
				coordinates: [][]int{
					{0, 0},
					{0, 1},
					{0, -1},
				},
			},
			want: true,
		},
		{
			name: "seventh",
			args: args{
				coordinates: [][]int{
					{1, 1},
					{2, 2},
					{2, 0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStraightLine(tt.args.coordinates); got != tt.want {
				t.Errorf("checkStraightLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
