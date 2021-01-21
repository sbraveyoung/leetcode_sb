package balancedStringSplit

import "testing"

func Test_balancedStringSplit(t *testing.T) {
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
				s: "RLRRLLRLRL",
			},
			want: 4,
		},
		{
			name: "second",
			args: args{
				s: "RLLLLRRRLR",
			},
			want: 3,
		},
		{
			name: "third",
			args: args{
				s: "LLLLRRRR",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedStringSplit(tt.args.s); got != tt.want {
				t.Errorf("balancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
