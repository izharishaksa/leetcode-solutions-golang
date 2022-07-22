package leetcode_solutions_golang

import "math"

//https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
func minCost(colors string, neededTime []int) int {
	max := 0
	total := 0
	result := 0
	for i := 0; i < len(colors); i++ {
		if i > 0 && colors[i] != colors[i-1] {
			result += total - max
			total = neededTime[i]
			max = neededTime[i]
		} else {
			total += neededTime[i]
			max = int(math.Max(float64(max), float64(neededTime[i])))
		}
	}
	result += total - max
	return result
}
