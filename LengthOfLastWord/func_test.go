package lengthOfLastWord

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
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
				s: "hello world",
			},
			want: 5,
		},
		{
			name: "second",
			args: args{
				s: "world",
			},
			want: 5,
		},
		{
			name: "third",
			args: args{
				s: "         world        ",
			},
			want: 5,
		},
		{
			name: "fourth",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "fifth",
			args: args{
				s: "               ",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
