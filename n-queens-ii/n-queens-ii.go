//https://leetcode.com/problems/n-queens-ii/

package n_queens_ii

func totalNQueens(n int) int {
	count := 0
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	totalNQueensHelper(board, 0, &count)
	return count
}

func totalNQueensHelper(board [][]bool, row int, count *int) {
	if row == len(board) {
		*count++
		return
	}
	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			board[row][col] = true
			totalNQueensHelper(board, row+1, count)
			board[row][col] = false
		}
	}
}

func isValid(board [][]bool, row int, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] {
			return false
		}
	}
	return true
}
