//https://leetcode.com/problems/validate-binary-search-tree/

package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	order := make([]int, 0)
	inorder(root, &order)
	for i := 1; i < len(order); i++ {
		if order[i] <= order[i-1] {
			return false
		}
	}
	return true
}

func inorder(root *TreeNode, order *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, order)
	*order = append(*order, root.Val)
	inorder(root.Right, order)
}
