package isIsomporphic

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
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
				s: "",
				t: "",
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				s: "egg",
				t: "add",
			},
			want: true,
		},
		{
			name: "third",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		},
		{
			name: "fourth",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
		{
			name: "fifth",
			args: args{
				s: "ab",
				t: "aa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
