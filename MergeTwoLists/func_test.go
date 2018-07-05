package mergeTwoLists

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
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
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
		{
			name: "second",
			args: args{
				l1: &ListNode{
					Val:  1,
					Next: nil,
				},
				l2: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
		{
			name: "third",
			args: args{
				l1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
				l2: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
