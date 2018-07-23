package Shuffle

import (
	"fmt"
	"sort"
	"testing"
)

func TestSolution_Shuffle(t *testing.T) {
	type fields struct {
		origin []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "first",
			fields: fields{
				origin: []int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Solution{
				origin: tt.fields.origin,
			}
			for i := 0; i < 12; i++ {
				got := this.Shuffle()
				fmt.Println(got)
				if !check(this.origin, got) {
					t.Errorf("Solution.Shuffle() = %v", got)
				}
			}
		})
	}
}

func check(origin, got []int) bool {
	// I'm not consider the probability of call Shuffle() much more times.
	length := len(origin)
	if length != len(got) {
		return false
	}
	sort.Ints(origin)
	sort.Ints(got)
	for i := 0; i < length; i++ {
		if origin[i] != got[i] {
			return false
		}
	}
	return true
}
