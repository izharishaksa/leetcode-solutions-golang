//https://leetcode.com/problems/longest-increasing-path-in-a-matrix/

package leetcode_solutions_golang

var dp [][]int

func longestIncreasingPath(matrix [][]int) int {
	dp = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			max = maxInt(max, traverseMatrix(matrix, i, j))
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func traverseMatrix(matrix [][]int, i, j int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	max := 1
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) {
			continue
		}
		if matrix[x][y] > matrix[i][j] {
			max = maxInt(max, 1+traverseMatrix(matrix, x, y))
		}
	}
	dp[i][j] = max
	return max
}
