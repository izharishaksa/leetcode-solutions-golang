//https://leetcode.com/problems/reachable-nodes-with-restrictions/

package leetcode_solutions_golang

var routes map[int][]int
var total int
var abandoned map[int]bool

func reachableNodes(n int, edges [][]int, restricted []int) int {
	abandoned = make(map[int]bool, 0)
	for i := 0; i < len(restricted); i++ {
		abandoned[restricted[i]] = true
	}
	routes = make(map[int][]int, 0)
	for i := 0; i < len(edges); i++ {
		cur := edges[i]
		routes[cur[0]] = append(routes[cur[0]], cur[1])
		routes[cur[1]] = append(routes[cur[1]], cur[0])
	}

	total = 0
	trace(0)

	return total
}

func trace(cur int) {
	abandoned[cur] = true
	total++
	next := routes[cur]
	for i := 0; i < len(next); i++ {
		if !abandoned[next[i]] {
			trace(next[i])
		}
	}
}
