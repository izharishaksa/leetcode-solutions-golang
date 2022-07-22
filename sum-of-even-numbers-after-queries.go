package leetcode_solutions_golang

//https://leetcode.com/problems/sum-of-even-numbers-after-queries/
func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	ret := make([]int, len(queries))
	cur := 0
	for _, i := range nums {
		if i%2 == 0 {
			cur += i
		}
	}
	for i := 0; i < len(queries); i++ {
		val := queries[i][0]
		idx := queries[i][1]
		if nums[idx]%2 == 0 {
			cur -= nums[idx]
			if (nums[idx]+val)%2 == 0 {
				cur += val + nums[idx]
			}
		} else if (nums[idx]+val)%2 == 0 {
			cur += nums[idx] + val
		}
		nums[idx] += val
		ret[i] = cur
	}
	return ret
}
