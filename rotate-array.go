package leetcode_solutions_golang

//https://leetcode.com/problems/rotate-array/
func rotate(nums []int, k int) {
	size := len(nums)
	temp := make([]int, size)
	for i := 0; i < size; i++ {
		temp[(i+k)%size] = nums[i]
	}
	for i := 0; i < size; i++ {
		nums[i] = temp[i]
	}
}
