package leetcode_solutions_golang

//https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
func repeatedNTimes(nums []int) int {
	target := len(nums) / 2
	repetitions := make([]int, 10001)
	for _, v := range nums {
		repetitions[v]++
		if repetitions[v] == target {
			return v
		}
	}
	return 0
}
