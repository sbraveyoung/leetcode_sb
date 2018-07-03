package isPowerOfThree

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	type args struct {
		n int
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
				n: 9,
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				n: 8,
			},
			want: false,
		},
		{
			name: "third",
			args: args{
				n: 2,
			},
			want: false,
		},
		{
			name: "fourth",
			args: args{
				n: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
