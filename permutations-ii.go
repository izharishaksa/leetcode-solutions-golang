//https://leetcode.com/problems/permutations-ii/

package leetcode_solutions_golang

import "fmt"

func permuteUnique(nums []int) [][]int {
	temp := make(map[string][]int, 0)
	used := make([]bool, len(nums))
	res := make([]int, len(nums))
	count := 0
	permuteUniqueHelper(nums, res, used, count, temp)
	result := make([][]int, 0)
	for _, i := range temp {
		result = append(result, i)
	}
	return result
}

func permuteUniqueHelper(nums, cur []int, used []bool, count int, result map[string][]int) {
	if count == len(nums) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		key := fmt.Sprintf("%v", temp)
		result[key] = temp
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		cur[count] = nums[i]
		count++
		permuteUniqueHelper(nums, cur, used, count, result)
		count--
		used[i] = false
	}
}
