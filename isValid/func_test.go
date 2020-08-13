package isValid

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
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
				s: "()",
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "third",
			args: args{
				s: "[{()}]",
			},
			want: true,
		},
		{
			name: "fourth",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "fifth",
			args: args{
				s: "((",
			},
			want: false,
		},
		{
			name: "sixth",
			args: args{
				s: "))",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
