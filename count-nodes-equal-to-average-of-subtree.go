package leetcode_solutions_golang

import "math"

//https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/
func dfs(root *TreeNode) (totalValues, totalNodes, answer int) {
	if root == nil {
		return 0, 0, answer
	}
	totalLeft, nodeLeft, cur1 := dfs(root.Left)
	totalRight, nodeRight, cur2 := dfs(root.Right)
	totalValues = totalLeft + totalRight + root.Val
	totalNodes = nodeLeft + nodeRight + 1

	answer += cur1
	answer += cur2
	avg := int(math.Floor(float64(totalValues) / float64(totalNodes)))
	if avg == root.Val {
		answer++
	}
	return totalValues, totalNodes, answer
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	_, _, answer := dfs(root)
	return answer
}
