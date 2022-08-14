//https://leetcode.com/problems/check-if-there-is-a-valid-partition-for-the-array/

package leetcode_solutions_golang

func validPartition(nums []int) bool {
	memo := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = -1
	}
	return partition(0, nums, memo)
}

func partition(curPos int, nums, memo []int) bool {
	if curPos == len(nums) {
		return true
	}
	if memo[curPos] != -1 {
		return memo[curPos] == 1
	}
	if curPos+1 < len(nums) && nums[curPos] == nums[curPos+1] {
		result := partition(curPos+2, nums, memo)
		if result {
			memo[curPos] = 1
			return true
		}
	}
	if curPos+2 < len(nums) && nums[curPos] == nums[curPos+1] && nums[curPos] == nums[curPos+2] {
		result := partition(curPos+3, nums, memo)
		if result {
			memo[curPos] = 1
			return true
		}
	}
	if curPos+2 < len(nums) && nums[curPos]+1 == nums[curPos+1] && nums[curPos]+2 == nums[curPos+2] {
		result := partition(curPos+3, nums, memo)
		if result {
			memo[curPos] = 1
			return true
		}
	}
	memo[curPos] = 0
	return false
}
