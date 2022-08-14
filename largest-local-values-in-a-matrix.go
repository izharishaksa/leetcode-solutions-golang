//https://leetcode.com/problems/largest-local-values-in-a-matrix/

package leetcode_solutions_golang

func largestLocal(grid [][]int) [][]int {
	size := len(grid)
	result := make([][]int, size-2)
	for i := 0; i < size-2; i++ {
		result[i] = make([]int, size-2)
	}
	for i := 0; i < size-2; i++ {
		for j := 0; j < size-2; j++ {
			max := 0
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					if grid[k][l] > max {
						max = grid[k][l]
					}
				}
			}
			result[i][j] = max
		}
	}
	return result
}
