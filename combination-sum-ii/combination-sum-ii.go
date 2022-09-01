//https://leetcode.com/problems/combination-sum-ii/

package combination_sum_ii

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	set := make(map[string][]int, 0)
	backtrack(candidates, []int{}, target, 0, -1, set)

	result := make([][]int, 0)
	for _, v := range set {
		result = append(result, v)
	}

	return result
}

func backtrack(candidates, curNum []int, target, curTotal, curIndex int, result map[string][]int) {
	if curTotal == target {
		key := fmt.Sprintf("%v", curNum)
		temp := make([]int, len(curNum))
		copy(temp, curNum)
		result[key] = temp
		return
	}
	prev := -1
	for i := curIndex + 1; i < len(candidates); i++ {
		if curTotal+candidates[i] <= target && candidates[i] != prev {
			curNum = append(curNum, candidates[i])
			backtrack(candidates, curNum, target, curTotal+candidates[i], i, result)
			prev = candidates[i]
			curNum = curNum[:len(curNum)-1]
		}
	}

	return
}
