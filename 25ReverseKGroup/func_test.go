package reverseKGroup

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
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
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
				k: 2,
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			name: "second",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
				k: 3,
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			name: "third",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
				k: 1,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			name: "fourth",
			args: args{
				head: nil,
				k:    1,
			},
			want: nil,
		},
		{
			name: "fifth",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
				k: 3,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
		{
			name: "sixth",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
				k: 2,
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !equal(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(head1, head2 *ListNode) bool {
	if head1 == nil && head2 == nil {
		return true
	}
	if head1 == nil || head2 == nil {
		return false
	}

	for ; ; head1, head2 = head1.Next, head2.Next {
		if head1 == nil && head2 == nil {
			return true
		}
		if head1 == nil || head2 == nil {
			return false
		}
		if head1.Val != head2.Val {
			return false
		}
	}
}
