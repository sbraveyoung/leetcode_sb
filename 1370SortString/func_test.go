package sortString

import "testing"

func Test_sortString(t *testing.T) {
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
				s: "aaaabbbbcccc",
			},
			want: "abccbaabccba",
		},
		{
			name: "second",
			args: args{
				s: "rat",
			},
			want: "art",
		},
		{
			name: "third",
			args: args{
				s: "leetcode",
			},
			want: "cdelotee",
		},
		{
			name: "fourth",
			args: args{
				s: "ggggggg",
			},
			want: "ggggggg",
		},
		{
			name: "fifth",
			args: args{
				s: "spo",
			},
			want: "ops",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.args.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
