//https://leetcode.com/problems/unique-paths-iii/

package unique_paths_iii

func uniquePathsIII(grid [][]int) int {
	isVisited := make([][]bool, len(grid))
	totalStep := len(grid) * len(grid[0])
	startRow, startCol := 0, 0
	endRow, endCol := 0, 0
	for i := 0; i < len(grid); i++ {
		isVisited[i] = make([]bool, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == -1 {
				totalStep--
			}
			if grid[i][j] == 1 {
				startRow, startCol = i, j
			}
			if grid[i][j] == 2 {
				endRow, endCol = i, j
			}
		}
	}
	totalPaths := 0
	uniquePathsIIIHelper(grid, isVisited, startRow, startCol, endRow, endCol, 0, totalStep, &totalPaths)
	return totalPaths
}

func uniquePathsIIIHelper(grid [][]int, isVisited [][]bool, row, col, targetRow, targetCol, curStep, totalStep int, totalPaths *int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) || isVisited[row][col] || grid[row][col] == -1 {
		return
	}
	if row == targetRow && col == targetCol {
		if curStep+1 == totalStep {
			*totalPaths++
		}
		return
	}
	isVisited[row][col] = true
	uniquePathsIIIHelper(grid, isVisited, row-1, col, targetRow, targetCol, curStep+1, totalStep, totalPaths)
	uniquePathsIIIHelper(grid, isVisited, row+1, col, targetRow, targetCol, curStep+1, totalStep, totalPaths)
	uniquePathsIIIHelper(grid, isVisited, row, col-1, targetRow, targetCol, curStep+1, totalStep, totalPaths)
	uniquePathsIIIHelper(grid, isVisited, row, col+1, targetRow, targetCol, curStep+1, totalStep, totalPaths)
	isVisited[row][col] = false
}
