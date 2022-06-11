package main

import (
	"fmt"
	"strings"
)

func main() {
	res := solveNQueens(4)
	fmt.Println(res)
}

func solveNQueens(n int) [][]string {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	res := make([][]string, 0)
	var backtracking func(int)
	backtracking = func(row int) {
		if row == n {
			tmp := make([]string, n)
			for i, rowStr := range board {
				tmp[i] = strings.Join(rowStr, "")
			}
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if !isValid5(board, n, row, i) {
				continue
			}
			board[row][i] = "Q"
			backtracking(row + 1)
			board[row][i] = "."
		}
	}
	backtracking(0)
	return res
}

func isValid5(board [][]string, n, row, col int) bool {
	for i := 0; i < n; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
