//https://leetcode.com/problems/sort-colors/

package leetcode_solutions_golang

func sortColors(nums []int) {
	count := [3]int{0, 0, 0}
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	index := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < count[i]; j++ {
			nums[index] = i
			index++
		}
	}
}
