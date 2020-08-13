package restoreipaddresses

import (
	"reflect"
	"sort"
	"testing"
)

func Test_restoreIPAddresses(t *testing.T) {
	type args struct {
		s string
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
				s: "1234",
			},
			want: []string{
				"1.2.3.4",
			},
		},
		{
			name: "second",
			args: args{
				s: "255255255255",
			},
			want: []string{
				"255.255.255.255",
			},
		},
		{
			name: "third",
			args: args{
				s: "12234",
			},
			want: []string{
				"1.2.2.34",
				"1.2.23.4",
				"1.22.3.4",
				"12.2.3.4",
			},
		},
		{
			name: "fourth",
			args: args{
				s: "25525511135",
			},
			want: []string{
				"255.255.11.135",
				"255.255.111.35",
			},
		},
		{
			name: "fifth",
			args: args{
				s: "010010",
			},
			want: []string{
				"0.100.1.0",
				"0.10.0.10",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := restoreIpAddresses(tt.args.s)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIPAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
