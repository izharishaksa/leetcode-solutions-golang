package leetcode_solutions_golang

//https://leetcode.com/problems/word-search/

import "container/list"

type BoardState struct {
	isUsed    [][]bool
	curRow    int
	curCol    int
	nextIndex int
}

func newBoardState(nRow, nCol, row, col, index int) *BoardState {
	state := &BoardState{
		isUsed: func() [][]bool {
			res := make([][]bool, nRow)
			for i := 0; i < nRow; i++ {
				res[i] = make([]bool, nCol)
			}
			return res
		}(),
		curRow:    row,
		curCol:    col,
		nextIndex: index,
	}
	state.isUsed[row][col] = true
	return state
}

func exist(board [][]byte, word string) bool {
	rowSize := len(board)
	colSize := len(board[0])
	q := list.New()
	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			if word[0] == board[row][col] {
				q.PushBack(newBoardState(rowSize, colSize, row, col, 1))
			}
		}
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for q.Len() > 0 {
		el := q.Front()
		q.Remove(el)
		curState := el.Value.(*BoardState)
		if curState.nextIndex == len(word) {
			return true
		}
		for _, dir := range directions {
			nextRow := curState.curRow + dir[0]
			nextCol := curState.curCol + dir[1]
			if nextRow >= 0 && nextRow < rowSize &&
				nextCol >= 0 && nextCol < colSize &&
				!curState.isUsed[nextRow][nextCol] &&
				board[nextRow][nextCol] == word[curState.nextIndex] {
				newState := newBoardState(rowSize, colSize, nextRow, nextCol, curState.nextIndex+1)
				for i := 0; i < rowSize; i++ {
					for j := 0; j < colSize; j++ {
						newState.isUsed[i][j] = curState.isUsed[i][j]
					}
				}
				newState.isUsed[nextRow][nextCol] = true
				q.PushFront(newState)
			}
		}
	}

	return false
}
