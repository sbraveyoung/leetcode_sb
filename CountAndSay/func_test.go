package countAndSay

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
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
				n: 1,
			},
			want: "1",
		},
		{
			name: "second",
			args: args{
				n: 2,
			},
			want: "11",
		},
		{
			name: "third",
			args: args{
				n: 3,
			},
			want: "21",
		},
		{
			name: "fourth",
			args: args{
				n: 4,
			},
			want: "1211",
		},
		{
			name: "fifth",
			args: args{
				n: 5,
			},
			want: "111221",
		},
		{
			name: "sixth",
			args: args{
				n: 6,
			},
			want: "312211",
		},
		{
			name: "seventh",
			args: args{
				n: 7,
			},
			want: "13112221",
		},
		{
			name: "eighth",
			args: args{
				n: 8,
			},
			want: "1113213211",
		},
		{
			name: "ninth",
			args: args{
				n: 9,
			},
			want: "31131211131221",
		},
		{
			name: "tenth",
			args: args{
				n: 10,
			},
			want: "13211311123113112211",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
