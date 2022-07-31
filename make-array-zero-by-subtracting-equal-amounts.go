package leetcode_solutions_golang

//https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/
func minimumOperations(nums []int) int {
	count := make([]int, 101)
	step := 0
	for _, num := range nums {
		count[num]++
	}
	for i := 1; i <= 100; i++ {
		if count[i] > 0 {
			step++
		}
	}
	return step
}
