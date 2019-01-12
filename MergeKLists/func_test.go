package mergeKLists

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			args: args{
				lists: []*ListNode{
					{
						Val: 1,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
					{
						Val: 1,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
					{
						Val: 2,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val: 6,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			args: args{
				lists: []*ListNode{},
			},
			want: nil,
		},
		{
			args: args{
				lists: []*ListNode{
					nil, nil,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
