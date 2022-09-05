//https://leetcode.com/problems/n-ary-tree-level-order-traversal/

package n_ary_tree_level_order_traversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		level := make([]int, 0)
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
		result = append(result, level)
	}
	return result
}
