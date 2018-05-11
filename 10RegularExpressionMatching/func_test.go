package RegularExpressionMatching

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
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
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "second",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "third",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "fourth",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "fifth",
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want: false,
		},
		{
			name: "sixth",
			args: args{
				s: "ssippi",
				p: "s*p*.",
			},
			want: false,
		},
		{
			name: "seventh",
			args: args{
				s: "a",
				p: ".*..a*",
			},
			want: false,
		},
		{
			name: "eighth",
			args: args{
				s: "ab",
				p: ".*..",
			},
			want: true,
		},
		{
			name: "ninth",
			args: args{
				s: "aaca",
				p: "ab*a*c*a",
			},
			want: true,
		},
		{
			name: "tenth",
			args: args{
				s: "aaba",
				p: "ab*a*c*a",
			},
			want: false,
		},
		{
			name: "eleventh",
			args: args{
				s: "aaa",
				p: "ab*a*c*a",
			},
			want: true,
		},
		{
			name: "twelveth",
			args: args{
				s: "aaa",
				p: ".*",
			},
			want: true,
		},
		{
			name: "thirteen",
			args: args{
				s: "a",
				p: "ab*",
			},
			want: true,
		},
		{
			name: "fourteen",
			args: args{
				s: "ab",
				p: ".*c",
			},
			want: false,
		},
		{
			name: "temp",
			args: args{
				s: "a",
				p: "a*a",
			},
			want: true,
		},
		{
			name: "fifteen",
			args: args{
				s: "bbbba",
				p: ".*a*a",
			},
			want: true,
		},
		{
			name: "sixteen",
			args: args{
				s: "",
				p: ".*",
			},
			want: true,
		},
		{
			name: "seventeen",
			args: args{
				s: "aaaaaaaaaaaaab",
				p: "a*a*a*a*a*a*a*a*a*a*c",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
