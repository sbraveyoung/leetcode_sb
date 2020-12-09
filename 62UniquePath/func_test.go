package uniquePaths

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
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
				m: 3,
				n: 7,
			},
			want: 28,
		},
		{
			name: "second",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "third",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
		{
			name: "fourth",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
		{
			name: "fifth",
			args: args{
				m: 1,
				n: 1,
			},
			want: 1,
		},
		{
			name: "sixth",
			args: args{
				m: 1,
				n: 2,
			},
			want: 1,
		},
		{
			name: "seventh",
			args: args{
				m: 2,
				n: 1,
			},
			want: 1,
		},
		{
			name: "eighth",
			args: args{
				m: 2,
				n: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
