//https://leetcode.com/problems/4sum/

package leetcode_solutions_golang

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	set := make(map[string][]int)
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			for k := j + 1; k < size; k++ {
				comp := target - (nums[i] + nums[j] + nums[k])
				if binarySearch(nums, comp, k+1, size-1) {
					res := []int{nums[i], nums[j], nums[k], comp}
					set[arrToString(res)] = res
				}
			}
		}
	}

	result := make([][]int, 0)
	for _, v := range set {
		result = append(result, v)
	}
	return result
}

func arrToString(arr []int) string {
	return fmt.Sprintf("%v", arr)
}

func binarySearch(arr []int, target, left, right int) bool {
	for left <= right {
		mid := (right + left) / 2
		if arr[mid] == target {
			return true
		}
		if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
