//https://leetcode.com/problems/find-original-array-from-doubled-array/

package leetcode_solutions_golang

import "sort"

func findOriginalArray(changed []int) []int {
	if len(changed)%2 == 1 {
		return []int{}
	}

	sort.Ints(changed)
	result := make([]int, 0)
	var count [100001]int
	for _, i := range changed {
		count[i]++
	}
	for i := 100000; i >= 0; i -= 2 {
		half := i / 2
		for count[i] > 0 && count[half] > 0 {
			count[i]--
			count[half]--
			result = append(result, half)
		}
	}
	for i := 0; i <= 100000; i++ {
		if count[i] > 0 {
			return []int{}
		}
	}

	return result
}
