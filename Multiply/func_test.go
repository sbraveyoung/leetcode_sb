package multiply

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
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
				num1: "2",
				num2: "3",
			},
			want: "6",
		},
		{
			name: "second",
			args: args{
				num1: "0",
				num2: "123",
			},
			want: "0",
		},
		{
			name: "third",
			args: args{
				num1: "10",
				num2: "35",
			},
			want: "350",
		},
		{
			name: "fourth",
			args: args{
				num1: "12",
				num2: "34",
			},
			want: "408",
		},
		{
			name: "fifth",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "56088",
		},
		{
			name: "sixth",
			args: args{
				num1: "999",
				num2: "999",
			},
			want: "998001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
