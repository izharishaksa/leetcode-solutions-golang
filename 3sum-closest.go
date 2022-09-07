//https://leetcode.com/problems/3sum-closest/

package leetcode_solutions_golang

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if sum < target {
				j++
			} else {
				k--
			}
			if math.Abs(float64(sum)-float64(target)) < math.Abs(float64(min)-float64(target)) {
				min = sum
			}
		}
	}
	return min
}
