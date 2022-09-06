//https://leetcode.com/problems/non-decreasing-array/

package leetcode_solutions_golang

func checkPossibility(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	changed := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] && !changed {
			if i == 1 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
			changed = true
		} else if nums[i] < nums[i-1] && changed {
			return false
		}
	}
	return true
}
