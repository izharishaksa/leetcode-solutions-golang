//https://leetcode.com/problems/pacific-atlantic-water-flow/

package leetcode_solutions_golang

func pacificAtlantic(heights [][]int) [][]int {
	isVisited := make([][]bool, len(heights))
	for i := 0; i < len(heights); i++ {
		isVisited[i] = make([]bool, len(heights[0]))
	}
	result := make([][]int, 0)
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if !traverseTheStream(heights, i, j, 0, 0, isVisited) {
				continue
			}
			if traverseTheStream(heights, i, j, len(heights)-1, len(heights[0])-1, isVisited) {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func traverseTheStream(heights [][]int, row, col, targetRow, targetCol int, isVisited [][]bool) bool {
	if row == targetRow || col == targetCol {
		return true
	}

	dirRow := []int{0, 1, 0, -1}
	dirCol := []int{1, 0, -1, 0}
	for i := 0; i < 4; i++ {
		newRow := row + dirRow[i]
		newCol := col + dirCol[i]
		if newRow < 0 || newRow >= len(heights) || newCol < 0 || newCol >= len(heights[0]) {
			continue
		}
		if heights[newRow][newCol] <= heights[row][col] && !isVisited[newRow][newCol] {
			isVisited[newRow][newCol] = true
			res := traverseTheStream(heights, newRow, newCol, targetRow, targetCol, isVisited)
			isVisited[newRow][newCol] = false
			if res {
				return true
			}
		}
	}
	return false
}
