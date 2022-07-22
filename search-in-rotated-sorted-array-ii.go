package leetcode_solutions_golang

//https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
func searchInRotatedSortedArray(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return true
		}
	}
	return false
}
