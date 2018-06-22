package removeNthFromEnd

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				n: 2,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
		{
			name: "second",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
				n: 2,
			},
			want: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
		{
			name: "third",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
				n: 1,
			},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
