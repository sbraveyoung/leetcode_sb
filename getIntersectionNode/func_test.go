package getIntersectionNode

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	base := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
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
				headA: &ListNode{
					Val:  1,
					Next: nil,
				},
				headB: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			want: nil,
		},
		{
			name: "second",
			args: args{
				headA: &ListNode{
					Val:  1,
					Next: base,
				},
				headB: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: base,
					},
				},
			},
			want: base,
		},
		{
			name: "third",
			args: args{
				headA: base,
				headB: base,
			},
			want: base,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
