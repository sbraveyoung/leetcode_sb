package validMountainArray

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		A []int
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
				A: []int{1, 2},
			},
			want: false,
		},
		{
			name: "second",
			args: args{
				A: []int{1, 2, 3},
			},
			want: false,
		},
		{
			name: "third",
			args: args{
				A: []int{1, 2, 1},
			},
			want: true,
		},
		{
			name: "fourth",
			args: args{
				A: []int{1, 2, 3, 3, 2, 1},
			},
			want: false,
		},
		{
			name: "fifth",
			args: args{
				A: []int{1, 2, 3, 4, 5, 4, 3, 3, 2, 1},
			},
			want: false,
		},
		{
			name: "sixth",
			args: args{
				A: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "seventh",
			args: args{
				A: []int{3, 2, 1},
			},
			want: false,
		},
		{
			name: "eighth",
			args: args{
				A: []int{3, 2, 1, 2, 3},
			},
			want: false,
		},
		{
			name: "ninth",
			args: args{
				A: []int{0, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "tenth",
			args: args{
				A: []int{1, 2, 3, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.A); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
