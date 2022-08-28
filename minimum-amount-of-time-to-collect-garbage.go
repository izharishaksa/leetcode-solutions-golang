//https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/

package leetcode_solutions_golang

func garbageCollection(garbage []string, travel []int) int {
	total := 0
	garbageType := []string{"M", "P", "G"}
	for t := 0; t < 3; t++ {
		cost := 0
		for i := 0; i < len(garbage); i++ {
			if i > 0 {
				cost += travel[i-1]
			}
			count := 0
			for j := 0; j < len(garbage[i]); j++ {
				if garbage[i][j] == garbageType[t][0] {
					count++
				}
			}
			if count > 0 {
				total += count + cost
				cost = 0
			}
		}
	}
	return total
}
