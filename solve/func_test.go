package solve

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	type want args
	tests := []struct {
		name string
		args args
		want want
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
			want: want{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'X', 'X', 'X'},
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			if !check(tt.args.board, tt.want.board) {
				t.Errorf("solve() = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}

func check(input, output [][]byte) bool {
	if len(input) != len(output) {
		return false
	}
	for i := 0; i < len(input); i++ {
		if len(input[i]) != len(output[i]) {
			return false
		}
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] != output[i][j] {
				return false
			}
		}
	}
	return true
}
