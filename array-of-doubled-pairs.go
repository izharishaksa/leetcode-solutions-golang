//https://leetcode.com/problems/array-of-doubled-pairs/

package leetcode_solutions_golang

import "sort"

func canReorderDoubled(arr []int) bool {
	count := make(map[int]int)
	for _, i := range arr {
		count[i]++
	}

	keys := make([]int, 0)
	for key, _ := range count {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	for i := 0; i < len(keys); i++ {
		full := keys[i] * 2
		half := keys[i]
		countFull, isOk1 := count[full]
		countHalf, isOk2 := count[half]
		if !isOk1 || !isOk2 {
			continue
		}
		if countFull >= countHalf {
			count[full] -= countHalf
			count[half] -= countFull
		} else {
			count[half] -= countFull
			count[full] -= countHalf
		}
	}

	for i := 0; i < len(keys); i++ {
		if count[keys[i]] > 0 {
			return false
		}
	}

	return true
}
