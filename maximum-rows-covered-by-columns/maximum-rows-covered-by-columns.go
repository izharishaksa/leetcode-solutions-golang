//https://leetcode.com/problems/maximum-rows-covered-by-columns/

package maximum_rows_covered_by_columns

func maximumRows(mat [][]int, cols int) int {
	max := 0
	isUsed := make([]bool, len(mat[0]))
	oneInRow := make([]int, len(mat))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				oneInRow[i]++
			}
		}
	}
	dfs(mat, oneInRow, []int{}, isUsed, -1, cols, &max)
	return max
}

func dfs(mat [][]int, oneInRow, curCol []int, isUsed []bool, curPos, totalCols int, max *int) {
	if len(curCol) == totalCols {
		totalRowCovered := 0
		totalOne := make([]int, len(mat))
		for i := 0; i < len(curCol); i++ {
			for j := 0; j < len(mat); j++ {
				if mat[j][curCol[i]] == 1 {
					totalOne[j]++
				}
			}
		}
		for i := 0; i < len(mat); i++ {
			if oneInRow[i] <= totalOne[i] {
				totalRowCovered++
			}
		}
		if totalRowCovered > *max {
			*max = totalRowCovered
		}
		return
	}
	for i := curPos + 1; i < len(mat[0]); i++ {
		if isUsed[i] {
			continue
		}
		isUsed[i] = true
		curCol = append(curCol, i)
		dfs(mat, oneInRow, curCol, isUsed, i, totalCols, max)
		curCol = curCol[:len(curCol)-1]
		isUsed[i] = false
	}
}
