package leetcode_solutions_golang

//https://leetcode.com/problems/count-hills-and-valleys-in-an-array/
func countHillValley(nums []int) int {
	arr := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		arr = append(arr, nums[i])
	}
	result := 0
	for i := 1; i < len(arr)-1; i++ {
		if (arr[i] > arr[i-1] && arr[i] > arr[i+1]) || (arr[i] < arr[i-1] && arr[i] < arr[i+1]) {
			result++
		}
	}
	return result
}
