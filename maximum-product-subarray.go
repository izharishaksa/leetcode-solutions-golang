package leetcode_solutions_golang

import "math"

func maxProduct(nums []int) int {
	max := float64(nums[0])
	min := float64(nums[0])
	result := max
	for i := 1; i < len(nums); i++ {
		cur := float64(nums[i])
		temp := math.Max(cur, math.Max(cur*max, cur*min))
		min = math.Min(cur, math.Min(cur*max, cur*min))
		max = temp
		result = math.Max(result, max)
	}
	return int(result)
}
