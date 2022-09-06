//https://leetcode.com/problems/binary-tree-pruning/

package binary_tree_pruning

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	isPruned := prune(root)
	if isPruned {
		return nil
	}
	return root
}

func prune(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isLeftPruned := prune(root.Left)
	if isLeftPruned {
		root.Left = nil
	}
	isRightPruned := prune(root.Right)
	if isRightPruned {
		root.Right = nil
	}
	return isLeftPruned && isRightPruned && root.Val == 0
}
