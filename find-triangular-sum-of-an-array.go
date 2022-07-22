package leetcode_solutions_golang

//https://leetcode.com/problems/find-triangular-sum-of-an-array/
func triangularSum(nums []int) int {
	for {
		if len(nums) == 1 {
			return nums[0]
		}
		var cur []int
		for i := 1; i < len(nums); i++ {
			cur = append(cur, (nums[i-1]+nums[i])%10)
		}
		nums = cur
	}
}
