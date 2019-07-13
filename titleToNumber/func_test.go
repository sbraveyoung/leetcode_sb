package titleToNumber

import "testing"

func Test_titleToNumber(t *testing.T) {
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
				s: "A",
			},
			want: 1,
		},
		{
			name: "second",
			args: args{
				s: "Z",
			},
			want: 26,
		},
		{
			name: "third",
			args: args{
				s: "AB",
			},
			want: 28,
		},
		{
			name: "four",
			args: args{
				s: "ZY",
			},
			want: 701,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.s); got != tt.want {
				t.Errorf("titleToNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
