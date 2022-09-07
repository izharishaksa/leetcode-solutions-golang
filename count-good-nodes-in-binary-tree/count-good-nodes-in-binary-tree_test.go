package count_good_nodes_in_binary_tree

import "testing"

func Test_goodNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodNodes(tt.args.root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
