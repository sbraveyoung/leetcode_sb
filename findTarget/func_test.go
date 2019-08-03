package findTarget

import "testing"

func Test_findTarget(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
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
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 9,
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 28,
			},
			want: false,
		},
		{
			name: "third",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 7,
			},
			want: true,
		},
		{
			name: "fourth",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 5,
			},
			want: true,
		},
		{
			name: "fifth",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 10,
			},
			want: true,
		},
		{
			name: "sixth",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 11,
			},
			want: true,
		},
		{
			name: "seventh",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 12,
			},
			want: true,
		},
		{
			name: "eighth",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				k: 4,
			},
			want: true,
		},
		{
			name: "ninth",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: -4,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				k: -1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
