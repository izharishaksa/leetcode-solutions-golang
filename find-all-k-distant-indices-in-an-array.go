package leetcode_solutions_golang

import (
	"math"
	"sort"
)

//https://leetcode.com/problems/find-all-k-distant-indices-in-an-array/
func findKDistantIndices(nums []int, key int, k int) []int {
	result := make(map[int]bool, 0)
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			distant := int(math.Abs(float64(i - j)))
			if distant <= k && nums[j] == key {
				result[i] = true
			}
		}
	}
	resultSlice := make([]int, 0)
	for key, _ := range result {
		resultSlice = append(resultSlice, key)
	}
	sort.Ints(resultSlice)
	return resultSlice
}
