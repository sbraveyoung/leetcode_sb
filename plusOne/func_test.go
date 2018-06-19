package plusOne

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
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
				digits: []int{0},
			},
			want: []int{1},
		},
		{
			name: "second",
			args: args{
				digits: []int{1, 0},
			},
			want: []int{1, 1},
		},
		{
			name: "third",
			args: args{
				digits: []int{9, 9},
			},
			want: []int{1, 0, 0},
		},
		{
			name: "fourth",
			args: args{
				digits: []int{1, 9},
			},
			want: []int{2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
