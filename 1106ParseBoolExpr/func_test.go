package parseBoolExpr

import "testing"

func Test_parseBoolExpr(t *testing.T) {
	type args struct {
		expression string
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
				expression: "!(f)",
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				expression: "|(f,t)",
			},
			want: true,
		},
		{
			name: "third",
			args: args{
				expression: "&(t,f)",
			},
			want: false,
		},
		{
			name: "fourth",
			args: args{
				expression: "|(&(t,f,t),!(t))",
			},
			want: false,
		},
		{
			name: "fifth",
			args: args{
				expression: "&(t,!(f),&(t,f))",
			},
			want: false,
		},
		{
			name: "sixth",
			args: args{
				expression: "&(t,!(f),|(t,f))",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBoolExpr(tt.args.expression); got != tt.want {
				t.Errorf("parseBoolExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
