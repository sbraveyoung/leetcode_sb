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
			want: "olleh",
		},
		{
			name: "third",
			args: args{
				s: "hello world",
			},
			want: "olleh dlrow",
		},
		{
			name: "fourth",
			args: args{
				s: "hello world, leetcode!",
			},
			want: "olleh ,dlrow !edocteel",
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
