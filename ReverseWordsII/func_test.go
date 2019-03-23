package reverseWords

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "second",
			args: args{
				s: "hello",
			},
			want: "hello",
		},
		{
			name: "third",
			args: args{
				s: "hello world",
			},
			want: "world hello",
		},
		{
			name: "fourth",
			args: args{
				s: "hello world, leetcode!",
			},
			want: "leetcode! world, hello",
		},
		{
			name: "fifth",
			args: args{
				s: "   the sky      is  blue",
			},
			want: "blue is sky the",
		},
		{
			name: "sixth",
			args: args{
				s: "      ",
			},
			want: "",
		},
		{
			name: "seventh",
			args: args{
				s: "a good  example",
			},
			want: "example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
