package leetcode_solutions_golang

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]
		if index, isExist := set[comp]; isExist {
			return []int{index, i}
		}
		set[nums[i]] = i
	}
	return []int{-1, -1}
}
