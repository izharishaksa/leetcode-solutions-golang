package leetcode_solutions_golang

//https://leetcode.com/problems/find-the-duplicate-number/
func findDuplicate(nums []int) int {
	flag := make(map[int]bool, 0)
	result := nums[0]
	for _, i := range nums {
		_, isExist := flag[i]
		if isExist {
			return i
		}
		flag[i] = true
	}
	return result
}
