package MaxProfit

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
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
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "second",
			args: args{
				prices: []int{1, 2, 3, 4, 5, 6},
			},
			want: 5,
		},
		{
			name: "third",
			args: args{
				prices: []int{6, 5, 4, 3, 2, 1},
			},
			want: 0,
		},
		{
			name: "fourth",
			args: args{
				prices: []int{1, 6, 5, 4, 3, 2},
			},
			want: 5,
		},
		{
			name: "fifth",
			args: args{
				prices: []int{5, 4, 3, 2, 1, 7},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
