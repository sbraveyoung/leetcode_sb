package myAtoi

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
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
				str: "42",
			},
			want: 42,
		},
		{
			name: "second",
			args: args{
				str: "    -42",
			},
			want: -42,
		},
		{
			name: "three",
			args: args{
				str: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "fourth",
			args: args{
				str: "words and 987",
			},
			want: 0,
		},
		{
			name: "fifth",
			args: args{
				str: "12345678901234567890",
			},
			want: 2147483647,
		},
		{
			name: "sixth",
			args: args{
				str: "-2147483648",
			},
			want: -2147483648,
		},
		{
			name: "seventh",
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "eighth",
			args: args{
				str: "-6147483648",
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
