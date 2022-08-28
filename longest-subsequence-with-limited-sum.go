//https://leetcode.com/problems/longest-subsequence-with-limited-sum/

package leetcode_solutions_golang

import "sort"

func answerQueries(nums []int, queries []int) []int {
	res := make([]int, len(queries))
	sort.Ints(nums)
	for i := 0; i < len(queries); i++ {
		total := 0
		for j := 0; j < len(nums); j++ {
			if total+nums[j] <= queries[i] {
				total += nums[j]
				res[i]++
			}
		}
	}
	return res
}
