package intToRoman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
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
				num: 1,
			},
			want: "I",
		},
		{
			name: "second",
			args: args{
				num: 2,
			},
			want: "II",
		},
		{
			name: "three",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "four",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "five",
			args: args{
				num: 5,
			},
			want: "V",
		},
		{
			name: "six",
			args: args{
				num: 6,
			},
			want: "VI",
		},
		{
			name: "seven",
			args: args{
				num: 9,
			},
			want: "IX",
		},
		{
			name: "eight",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "nine",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
