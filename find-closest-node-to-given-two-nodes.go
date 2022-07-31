package leetcode_solutions_golang

import "math"

//https://leetcode.com/problems/find-closest-node-to-given-two-nodes/
func closestMeetingNode(edges []int, node1 int, node2 int) int {
	distFromNode1 := make([]int, len(edges))
	distFromNode2 := make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		distFromNode1[i] = math.MaxInt32
		distFromNode2[i] = math.MaxInt32
	}
	distFromNode1[node1] = 0
	distFromNode2[node2] = 0
	step := 0
	curNode := node1
	for {
		node := edges[curNode]
		if node == -1 {
			break
		}
		if distFromNode1[node] != math.MaxInt32 {
			break
		}
		distFromNode1[node] = step + 1
		step++
		curNode = node
	}
	step = 0
	curNode = node2
	for {
		node := edges[curNode]
		if node == -1 {
			break
		}
		if distFromNode2[node] != math.MaxInt32 {
			break
		}
		distFromNode2[node] = step + 1
		step++
		curNode = node
	}
	min := math.MaxInt64
	index := -1
	for i := 0; i < len(edges); i++ {
		if distFromNode1[i] != math.MaxInt32 && distFromNode2[i] != math.MaxInt32 {
			curMax := int(math.Max(float64(distFromNode1[i]), float64(distFromNode2[i])))
			if curMax < min {
				min = curMax
				index = i
			}
		}
	}
	return index
}
