package leetcode_solutions_golang

import "sort"

//https://leetcode.com/problems/maximize-the-topmost-element-after-k-moves/
func maximumTop(nums []int, k int) int {
	size := len(nums)
	if k < size {
		max := -1
		for i := 0; i < k-1; i++ {
			if nums[i] > max {
				max = nums[i]
			}
		}
		if nums[k] > max {
			return nums[k]
		}
		return max
	}
	maxPos := 0
	for i := 0; i < size; i++ {
		if nums[i] > nums[maxPos] {
			maxPos = i
		}
	}
	sort.Ints(nums)
	if size == 1 {
		if k%2 == 1 {
			return -1
		}
		return nums[0]
	}
	if k == size {
		if maxPos < size-1 {
			return nums[size-1]
		}
		return nums[size-2]
	}
	return nums[size-1]
}
