package findMinArrowShots

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points interval
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
				points: interval{
					{10, 16},
					{2, 8},
					{1, 6},
					{7, 12},
				},
			},
			want: 2,
		},
		{
			name: "second",
			args: args{
				points: interval{
					{1, 2},
					{3, 4},
					{5, 6},
					{7, 8},
				},
			},
			want: 4,
		},
		{
			name: "third",
			args: args{
				points: interval{
					{1, 2},
					{2, 3},
					{3, 4},
					{4, 5},
				},
			},
			want: 2,
		},
		{
			name: "fourth",
			args: args{
				points: interval{
					{1, 2},
				},
			},
			want: 1,
		},
		{
			name: "fifth",
			args: args{
				points: interval{
					{2, 3},
					{2, 3},
				},
			},
			want: 1,
		},
		{
			name: "sixth",
			args: args{
				points: interval{
					{4, 5},
					{3, 6},
					{1, 8},
					{2, 7},
				},
			},
			want: 1,
		},
		{
			name: "seventh",
			args: args{
				points: interval{
					{1, 3},
					{2, 5},
					{4, 6},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
