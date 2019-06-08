package calculate

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// name: "first",
		// args: args{
		// s: "",
		// },
		// want: 0,
		// },
		// {
		// name: "second",
		// args: args{
		// s: "1+1",
		// },
		// want: 2,
		// },
		{
			name: "three",
			args: args{
				s: "1+2 -3+    4    -5      + 6",
			},
			want: 5,
		},
		{
			name: "four",
			args: args{
				s: "(1+2)-(3+4)  + (5  -6)+  7 - 8 +((2+3) - 4)",
			},
			want: -5,
		},
		{
			name: "five",
			args: args{
				s: "(1+(4+5+2)-3)+(6+8)",
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
