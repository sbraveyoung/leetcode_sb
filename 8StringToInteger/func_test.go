package StringToInteger

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
		//TODO: Add test cases.
		{
			name: "first",
			args: args{str: "01234"},
			want: 1234,
		},
		{
			name: "second",
			args: args{str: "  -42"},
			want: -42,
		},
		{
			name: "third",
			args: args{str: "4293 with words"},
			want: 4293,
		},
		{
			name: "fourth",
			args: args{str: "words and 987"},
			want: 0,
		},
		{
			name: "fifth",
			args: args{str: "-2199999999"},
			want: -2147483648,
		},
		{
			name: "sixth",
			args: args{str: " "},
			want: 0,
		},
		{
			name: "seventh",
			args: args{str: "  -00000000000012345678"},
			want: -12345678,
		},
		{
			name: "eighth",
			args: args{str: "1095502006p8"},
			want: 1095502006,
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
