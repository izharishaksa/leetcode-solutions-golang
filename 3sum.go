//https://leetcode.com/problems/3sum/

package leetcode_solutions_golang

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	mappedNums := make(map[int]map[int]bool, 0)
	for pos, i := range nums {
		cur, isOk := mappedNums[i]
		if !isOk {
			mappedNums[i] = make(map[int]bool, 0)
			cur, _ = mappedNums[i]
		}
		cur[pos] = true
		mappedNums[i] = cur
	}
	res := make([][]int, 0)
	added := make(map[string]bool)
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			toFind := -(nums[i] + nums[j])
			val, isExist := mappedNums[toFind]
			if toFind >= nums[j] && isExist {
				cur := fmt.Sprintf("%d,%d,%d", nums[i], nums[j], toFind)
				count := len(val)
				_, ii := val[i]
				if ii {
					count--
				}
				_, jj := val[j]
				if jj {
					count--
				}
				if _, ok := added[cur]; !ok && count > 0 {
					res = append(res, []int{nums[i], nums[j], toFind})
					added[cur] = true
				}
			}
		}
	}
	return res
}
