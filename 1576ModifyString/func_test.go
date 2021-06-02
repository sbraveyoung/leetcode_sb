package modifyString

import "testing"

func Test_modifyString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				s: "?zs",
			},
		},
		{
			name: "second",
			args: args{
				s: "ubv?w",
			},
		},
		{
			name: "third",
			args: args{
				s: "j?qg??b",
			},
		},
		{
			name: "fourth",
			args: args{
				s: "??yw?ipkj?",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifyString(tt.args.s); !check(got) {
				t.Errorf("modifyString() = %v", got)
			}
		})
	}
}

func check(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return false
		}
	}
	return true
}
