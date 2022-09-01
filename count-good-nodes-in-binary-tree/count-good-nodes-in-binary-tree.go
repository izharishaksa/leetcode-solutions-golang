//https://leetcode.com/problems/count-good-nodes-in-binary-tree/

package count_good_nodes_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return count(root, -10e4)
}

func count(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}
	result := 0
	if root.Val >= max {
		max = root.Val
		result++
	}

	return result + count(root.Left, max) + count(root.Right, max)
}
