//https://leetcode.com/problems/search-insert-position/

package leetcode_solutions_golang

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left <= right {
		middle := (right + left) / 2
		if middle == len(nums) {
			return middle
		}
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			right = middle
			right--
		}
		if nums[middle] < target {
			left = middle
			left++
		}
	}
	return left
}
