package leetcode_solutions_golang

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	result := 0
	mid := int(len(nums) / 2)
	for _, i := range nums {
		if nums[mid] >= i {
			result += nums[mid] - i
		} else {
			result += i - nums[mid]
		}
	}
	return result
}
