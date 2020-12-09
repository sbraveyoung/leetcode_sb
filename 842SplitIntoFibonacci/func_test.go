package splitIntoFibonacci

import (
	"strconv"
	"testing"
)

func Test_splitIntoFibonacci(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				S: "123456579",
			},
			want: []int{123, 456, 579},
		},
		{
			name: "second",
			args: args{
				S: "11235813",
			},
			want: []int{1, 1, 2, 3, 5, 8, 13},
		},
		{
			name: "third",
			args: args{
				S: "112358130",
			},
			want: []int{},
		},
		{
			name: "fourth",
			args: args{
				S: "0123",
			},
			want: []int{},
		},
		{
			name: "fifth",
			args: args{
				S: "1101111",
			},
			want: []int{110, 1, 111},
		},
		{
			name: "sixth",
			args: args{
				S: "112357",
			},
			want: []int{},
		},
		{
			name: "seventh",
			args: args{
				S: "123",
			},
			want: []int{1, 2, 3},
		},
		{
			name: "eighth",
			args: args{
				S: "11",
			},
			want: []int{},
		},
		{
			name: "nineth",
			args: args{
				S: "1011",
			},
			want: []int{1, 0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitIntoFibonacci(tt.args.S); !right(tt.want, got) {
				t.Errorf("splitIntoFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func right(want, ret []int) bool {
	wantStr := ""
	for i := 0; i < len(want); i++ {
		wantStr += strconv.Itoa(want[i])
	}
	retStr := ""
	for i := 0; i < len(ret); i++ {
		retStr += strconv.Itoa(ret[i])
	}
	return wantStr == retStr
}
