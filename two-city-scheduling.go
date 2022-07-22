package leetcode_solutions_golang

import "sort"

//https://leetcode.com/problems/two-city-scheduling/submissions/
func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		costA := costs[i][0] - costs[i][1]
		costB := costs[j][0] - costs[j][1]
		if costA >= costB {
			return false
		}
		return true
	})

	total := 0
	counter := len(costs) / 2
	for _, v := range costs {
		if counter > 0 {
			total += v[0]
		} else {
			total += v[1]
		}
		counter--
	}

	return total
}
