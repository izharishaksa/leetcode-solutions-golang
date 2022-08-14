//https://leetcode.com/problems/merge-similar-items/

package leetcode_solutions_golang

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	sum := make(map[int]int, 0)
	for i := 0; i < len(items1); i++ {
		cur, isExist := sum[items1[i][0]]
		if !isExist {
			sum[items1[i][0]] = items1[i][1]
		} else {
			sum[items1[i][0]] = cur + items1[i][1]
		}
	}
	for i := 0; i < len(items2); i++ {
		cur, isExist := sum[items2[i][0]]
		if !isExist {
			sum[items2[i][0]] = items2[i][1]
		} else {
			sum[items2[i][0]] = cur + items2[i][1]
		}
	}
	key := make([]int, 0)
	for k := range sum {
		key = append(key, k)
	}
	sort.Ints(key)
	merged := make([][]int, 0)
	for i := 0; i < len(key); i++ {
		merged = append(merged, []int{key[i], sum[key[i]]})
	}
	return merged
}
