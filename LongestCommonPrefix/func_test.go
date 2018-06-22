package longestCommonPrefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
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
				strs: []string{
					"flower",
					"flow",
					"flight",
				},
			},
			want: "fl",
		},
		{
			name: "second",
			args: args{
				strs: []string{
					"dog",
					"racecar",
					"car",
				},
			},
			want: "",
		},
		{
			name: "third",
			args: args{
				strs: []string{
					"hello",
					"hello",
				},
			},
			want: "hello",
		},
		{
			name: "fourth",
			args: args{
				strs: []string{
					"aa",
					"a",
				},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
