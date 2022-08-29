//https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package leetcode_solutions_golang

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	a := -1
	b := -1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			a = middle
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	left = 0
	right = len(nums) - 1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			b = middle
			left = middle + 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return []int{a, b}
}
