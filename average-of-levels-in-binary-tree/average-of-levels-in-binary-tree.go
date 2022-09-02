//https://leetcode.com/problems/average-of-levels-in-binary-tree/

package average_of_levels_in_binary_tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []float64
var count []int

func averageOfLevels(root *TreeNode) []float64 {
	result = make([]float64, 0)
	count = make([]int, 0)
	traverse(root, 0)
	return result
}

func traverse(root *TreeNode, level int) {
	if root != nil {
		fmt.Println(result, count)
		if level == len(count) {
			count = append(count, 1)
			result = append(result, float64(root.Val))
		} else {
			temp := float64(count[level]) * result[level]
			count[level]++
			temp += float64(root.Val)
			temp /= float64(count[level])
			result[level] = temp
		}
		traverse(root.Left, level+1)
		traverse(root.Right, level+1)
	}
}
