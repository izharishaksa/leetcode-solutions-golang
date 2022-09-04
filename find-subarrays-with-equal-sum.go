//https://leetcode.com/problems/find-subarrays-with-equal-sum/

package leetcode_solutions_golang

func findSubarrays(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		a := nums[i] + nums[i+1]
		for j := i + 1; j < len(nums)-1; j++ {
			b := nums[j] + nums[j+1]
			if a == b {
				return true
			}
		}
	}
	return false
}
