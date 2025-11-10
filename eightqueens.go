package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	board := make([]int, 8)
	solve(board, 0)
}

func solve(board []int, col int) {
	if col == 8 {
		printSolution(board)
		return
	}

	for row := 1; row <= 8; row++ {
		if isSafe(board, col, row) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

func isSafe(board []int, col int, row int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row {
			return false
		}
		if abs(board[i]-row) == abs(i-col) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printSolution(board []int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune('0' + board[i]))
	}
	z01.PrintRune('\n')
}
