package leetcode_solutions_golang

import "math"

//https://leetcode.com/problems/intersection-of-two-arrays-ii/
func intersect(nums1 []int, nums2 []int) []int {
	arr1 := [1001]int{}
	arr2 := [1001]int{}

	for _, i := range nums1 {
		arr1[i]++
	}
	for _, i := range nums2 {
		arr2[i]++
	}

	result := []int{}
	for i := 0; i < 1001; i++ {
		min := int(math.Min(float64(arr1[i]), float64(arr2[i])))
		for j := 0; j < min; j++ {
			result = append(result, i)
		}
	}
	return result
}
