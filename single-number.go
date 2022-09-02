//https://leetcode.com/problems/single-number/

package leetcode_solutions_golang

func singleNumber(nums []int) int {
	a := 0
	for _, i := range nums {
		a ^= i
	}
	return a
}
