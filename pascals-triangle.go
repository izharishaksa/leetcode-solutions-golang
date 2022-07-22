package leetcode_solutions_golang

//https://leetcode.com/problems/pascals-triangle/
func generate(numRows int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{1})
	for i := 2; i <= numRows; i++ {
		cur := make([]int, 0)
		cur = append(cur, 1)
		prev := res[len(res)-1]
		for j := 1; j < len(prev); j++ {
			cur = append(cur, prev[j]+prev[j-1])
		}
		cur = append(cur, 1)
		res = append(res, cur)
	}
	return res
}
