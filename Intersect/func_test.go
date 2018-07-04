package intersect

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
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
				nums1: []int{1, 2, 3, 4},
				nums2: []int{0, 2, 4, 6},
			},
			want: []int{2, 4},
		},
		{
			name: "second",
			args: args{
				nums1: []int{1, 2, 2, 4},
				nums2: []int{0, 2, 2, 4},
			},
			want: []int{2, 2, 4},
		},
		{
			name: "third",
			args: args{
				nums1: []int{1, 2, 3, 4},
				nums2: []int{5, 6, 7, 8},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.args.nums1, tt.args.nums2)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
