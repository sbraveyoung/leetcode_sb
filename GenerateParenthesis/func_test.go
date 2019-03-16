package generateParenthesis

import (
	"reflect"
	"sort"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				n: 0,
			},
			want: []string{},
		},
		{
			name: "second",
			args: args{
				n: 1,
			},
			want: []string{
				"()",
			},
		},
		{
			name: "third",
			args: args{
				n: 2,
			},
			want: []string{
				"()()",
				"(())",
			},
		},
		{
			name: "fourth",
			args: args{
				n: 3,
			},
			want: []string{
				"()()()",
				"(())()",
				"((()))",
				"()(())",
				"(()())",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !equal(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(got, want []string) bool {
	sort.Strings(got)
	sort.Strings(want)
	return reflect.DeepEqual(got, want)
}
