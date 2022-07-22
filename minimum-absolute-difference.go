package leetcode_solutions_golang

import (
	"math"
	"sort"
)

//https://leetcode.com/problems/minimum-absolute-difference/
func minimumAbsDifference(arr []int) [][]int {
	minAbs := 3000000
	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		curAbs := math.Abs(float64(arr[i] - arr[i-1]))
		minAbs = int(math.Min(float64(minAbs), curAbs))
	}
	result := [][]int{}
	for i := 1; i < len(arr); i++ {
		curAbs := int(math.Abs(float64(arr[i] - arr[i-1])))
		if curAbs == minAbs {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}
