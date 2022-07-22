package leetcode_solutions_golang

import "container/list"

//https://leetcode.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	type Pos struct {
		row int16
		col int16
	}
	m := len(grid)
	n := len(grid[0])
	q := list.New()

	numIsland := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				numIsland++
				q.PushFront(Pos{int16(i), int16(j)})
				grid[i][j] = '0'
				for q.Len() > 0 {
					el := q.Front()
					q.Remove(el)
					curPos := el.Value.(Pos)
					if curPos.row-1 >= 0 && grid[curPos.row-1][curPos.col] == '1' {
						q.PushBack(Pos{curPos.row - 1, curPos.col})
						grid[curPos.row-1][curPos.col] = '0'
					}
					if curPos.row+1 < int16(m) && grid[curPos.row+1][curPos.col] == '1' {
						q.PushBack(Pos{curPos.row + 1, curPos.col})
						grid[curPos.row+1][curPos.col] = '0'
					}
					if curPos.col-1 >= 0 && grid[curPos.row][curPos.col-1] == '1' {
						q.PushBack(Pos{curPos.row, curPos.col - 1})
						grid[curPos.row][curPos.col-1] = '0'
					}
					if curPos.col+1 < int16(n) && grid[curPos.row][curPos.col+1] == '1' {
						q.PushBack(Pos{curPos.row, curPos.col + 1})
						grid[curPos.row][curPos.col+1] = '0'
					}
				}
			}
		}
	}
	return numIsland
}
