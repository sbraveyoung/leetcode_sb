package matrixScore

import "testing"

func Test_matrixScore(t *testing.T) {
	type args struct {
		A [][]int
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
				A: [][]int{
					{0, 0, 1, 1},
					{1, 0, 1, 0},
					{1, 1, 0, 0},
				},
			},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixScore(tt.args.A); got != tt.want {
				t.Errorf("matrixScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
