//https://leetcode.com/problems/number-of-arithmetic-triplets/

package leetcode_solutions_golang

func arithmeticTriplets(nums []int, diff int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					total++
				}
			}
		}
	}
	return total
}
