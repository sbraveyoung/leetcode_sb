package isHappy

import "testing"

func Test_isHappy(t *testing.T) {
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
				n: 10,
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				n: 7,
			},
			want: true,
		},
		{
			name: "third",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "fourth",
			args: args{
				n: 34,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
