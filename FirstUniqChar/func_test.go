package firstUniqChar

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
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
				s: "leetcode",
			},
			want: 0,
		},
		{
			name: "second",
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			name: "third",
			args: args{
				s: "abcdefg",
			},
			want: 0,
		},
		{
			name: "fourth",
			args: args{
				s: "abcdeabcde",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
