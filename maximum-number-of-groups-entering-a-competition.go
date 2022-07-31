package leetcode_solutions_golang

import "sort"

//https://leetcode.com/problems/maximum-number-of-groups-entering-a-competition/
func maximumGroups(grades []int) int {
	sort.Ints(grades)
	totalGroup := 0
	prevCount := 0
	prevTotal := 0
	curCount := 0
	curTotal := 0
	for i := 0; i < len(grades); i++ {
		curCount++
		curTotal += grades[i]
		if curCount > prevCount && curTotal > prevTotal {
			totalGroup++
			prevCount = curCount
			prevTotal = curTotal
			curCount = 0
			curTotal = 0
		}
	}
	if curCount > prevCount && curTotal > prevTotal {
		totalGroup++
	}
	return totalGroup
}
