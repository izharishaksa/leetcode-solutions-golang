package leetcode_solutions_golang

import "sort"

//https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := []int{}
	merged = append(merged, nums1...)
	merged = append(merged, nums2...)
	sort.Ints(merged)

	if len(merged)%2 == 0 {
		total := float64(merged[len(merged)/2-1]+merged[len(merged)/2]) / 2
		return float64(total)
	}

	return float64(merged[len(merged)/2])
}
