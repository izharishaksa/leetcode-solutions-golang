package leetcode_solutions_golang

import "math"

//https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1
	for left < right {
		curMax := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		if curMax > max {
			max = curMax
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
