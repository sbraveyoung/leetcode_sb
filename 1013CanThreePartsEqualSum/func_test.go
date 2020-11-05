package canThreePartsEqualSum

import "testing"

func Test_canThreePartsEqualSum(t *testing.T) {
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
				A: []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				A: []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
			},
			want: false,
		},
		{
			name: "third",
			args: args{
				A: []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4},
			},
			want: true,
		},
		{
			name: "fourth",
			args: args{
				A: []int{0},
			},
			want: false,
		},
		{
			name: "fifth",
			args: args{
				A: []int{0, 0, 0},
			},
			want: true,
		},
		{
			name: "sixth",
			args: args{
				A: []int{1, 2, 3, 2, 1, 3, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "seventh",
			args: args{
				A: []int{18, 12, -18, 18, -19, -1, 10, 10},
			},
			want: true,
		},
		{
			name: "eighth",
			args: args{
				A: []int{1, -1, 1, -1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canThreePartsEqualSum(tt.args.A); got != tt.want {
				t.Errorf("canThreePartsEqualSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
