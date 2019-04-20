package complex

import "testing"

func Test_complexNumberMultiply(t *testing.T) {
	type args struct {
		a string
		b string
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
				a: "1+1i",
				b: "1+1i",
			},
			want: "0+2i",
		},
		{
			name: "second",
			args: args{
				a: "1+-1i",
				b: "1+-1i",
			},
			want: "0+-2i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complexNumberMultiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
