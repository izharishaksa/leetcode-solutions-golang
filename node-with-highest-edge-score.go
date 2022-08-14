//https://leetcode.com/problems/node-with-highest-edge-score/

package leetcode_solutions_golang

func edgeScore(edges []int) int {
	scores := make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		scores[edges[i]] += i
	}

	max := -1
	result := -1
	for i := 0; i < len(edges); i++ {
		if scores[i] > max {
			max = scores[i]
			result = i
		}
	}
	return result
}
