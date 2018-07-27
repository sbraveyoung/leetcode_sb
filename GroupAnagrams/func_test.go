package groupAnagrams

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				strs: []string{
					"abc", "de", "fghij",
					"bca", "cba", "ed", "acb",
				},
			},
			want: [][]string{
				{"abc", "bca", "cba", "acb"},
				{"de", "ed"},
				{"fghij"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
