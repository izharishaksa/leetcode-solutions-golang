//https://leetcode.com/problems/majority-element/

package leetcode_solutions_golang

func majorityElement(nums []int) int {
	count := make(map[int]int)
	for _, i := range nums {
		count[i]++
	}
	res := 0
	max := 0
	for k, v := range count {
		if v > max {
			max = v
			res = k
		}
	}
	return res
}
