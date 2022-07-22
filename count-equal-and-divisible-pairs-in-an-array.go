package leetcode_solutions_golang

//https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/
func countPairs(nums []int, k int) int {
	size := len(nums)
	result := 0
	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				result++
			}
		}
	}
	return result
}
