//https://leetcode.com/problems/binary-search/

package leetcode_solutions_golang

func search2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	pivot := (left + right) / 2
	for right >= left {
		if nums[pivot] == target {
			return pivot
		}
		if nums[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
		pivot = (left + right) / 2
	}
	return -1
}
