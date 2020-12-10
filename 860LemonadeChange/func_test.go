package lemonadeChange

import "testing"

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				bills: []int{5, 5, 5, 10, 20},
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				bills: []int{5, 5, 10},
			},
			want: true,
		},
		{
			name: "third",
			args: args{
				bills: []int{10, 10},
			},
			want: false,
		},
		{
			name: "fourth",
			args: args{
				bills: []int{5, 5, 10, 10, 20},
			},
			want: false,
		},
		{
			name: "fifth",
			args: args{
				bills: []int{5, 5, 5},
			},
			want: true,
		},
		{
			name: "sixth",
			args: args{
				bills: []int{10, 5, 5, 5},
			},
			want: false,
		},
		{
			name: "seventh",
			args: args{
				bills: []int{5, 10, 5, 10, 5},
			},
			want: true,
		},
		{
			name: "eighth",
			args: args{
				bills: []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5},
			},
			want: true,
		},
		{
			name: "ninth",
			args: args{
				bills: []int{5, 5, 5, 5, 20, 20, 5, 5, 20, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lemonadeChange(tt.args.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
