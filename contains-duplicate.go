package leetcode_solutions_golang

func containsDuplicate(nums []int) bool {
	arr := make(map[int]bool, 10000)
	for _, i := range nums {
		_, isAdded := arr[i]
		if isAdded {
			return true
		}
		arr[i] = true
	}
	return false
}
