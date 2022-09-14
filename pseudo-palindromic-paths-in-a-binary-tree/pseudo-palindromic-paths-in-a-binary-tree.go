//https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/

package pseudo_palindromic_paths_in_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	result := 0
	count := [10]int{}
	count[root.Val]++
	traverse(root, count, &result)
	return result
}

func traverse(root *TreeNode, count [10]int, result *int) {
	if root.Right == nil && root.Left == nil {
		oddCount := 0
		for _, i := range count {
			if i%2 == 1 {
				oddCount++
			}
		}
		if oddCount < 2 {
			*result++
		}
	}
	if root.Left != nil {
		count[root.Left.Val]++
		traverse(root.Left, count, result)
		count[root.Left.Val]--
	}
	if root.Right != nil {
		count[root.Right.Val]++
		traverse(root.Right, count, result)
		count[root.Right.Val]--
	}
}
