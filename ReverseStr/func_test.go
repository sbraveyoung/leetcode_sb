package reverseStr

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
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
				s: "abcdefg",
				k: 2,
			},
			want: "bacdfeg",
		},
		{
			name: "second",
			args: args{
				s: "abcdefg",
				k: 1,
			},
			want: "abcdefg",
		},
		{
			name: "third",
			args: args{
				s: "abcdefg",
				k: 5,
			},
			want: "edcbafg",
		},
		{
			name: "fourth",
			args: args{
				s: "abcdefg",
				k: 7,
			},
			want: "gfedcba",
		},
		{
			name: "fifth",
			args: args{
				s: "abcdefg",
				k: 8,
			},
			want: "gfedcba",
		},
		{
			name: "sixth",
			args: args{
				s: "",
				k: 8,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
