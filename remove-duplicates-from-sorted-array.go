//https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package leetcode_solutions_golang

func removeDuplicates(nums []int) int {
	i := 0
	for j := i + 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
