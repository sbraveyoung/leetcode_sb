package reverseString

import "testing"

func Test_reverseString(t *testing.T) {
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
				s: "hello",
			},
			want: "olleh",
		},
		{
			name: "second",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "third",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "fourth",
			args: args{
				s: "hello world",
			},
			want: "dlrow olleh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.s); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
