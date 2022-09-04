//https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/

package vertical_order_traversal_of_a_binary_tree

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Data struct {
	Val   int
	Depth int
}

func verticalTraversal(root *TreeNode) [][]int {
	m := make(map[int]map[int][]Data)
	traverse(root, 0, 0, m)
	result := make([][]int, 0)
	for i := -1000; i <= 1000; i++ {
		if _, ok := m[i]; !ok {
			continue
		}
		tmp := make([]Data, 0)
		for j := 0; j <= 1000; j++ {
			if _, ok := m[i][j]; !ok {
				continue
			}
			tmp = append(tmp, m[i][j]...)
		}
		sort.Slice(tmp, func(i, j int) bool {
			if tmp[i].Depth == tmp[j].Depth {
				return tmp[i].Val < tmp[j].Val
			}
			return tmp[i].Depth < tmp[j].Depth
		})
		tmp2 := make([]int, 0)
		for _, v := range tmp {
			tmp2 = append(tmp2, v.Val)
		}
		result = append(result, tmp2)
	}
	return result
}

func traverse(root *TreeNode, x, y int, m map[int]map[int][]Data) {
	if root == nil {
		return
	}
	if _, ok := m[x]; !ok {
		m[x] = make(map[int][]Data)
	}
	if _, ok := m[x][y]; !ok {
		m[x][y] = make([]Data, 0)
	}
	m[x][y] = append(m[x][y], Data{Val: root.Val, Depth: y})
	traverse(root.Left, x-1, y+1, m)
	traverse(root.Right, x+1, y+1, m)
}
