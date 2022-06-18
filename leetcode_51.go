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
	res := make([][]string, 0)
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	backtracking10(board, n, 0, &res)
	return res
}

func backtracking10(board [][]string, n, start int, res *[][]string) {
	if start == n {
		temp := make([]string, 0, len(board))
		for _, tmpStr := range board {
			temp = append(temp, strings.Join(tmpStr, ""))
		}
		*res = append(*res, temp)
		return
	}
	for i := 0; i < n; i++ {
		if board[start][i] != "." {
			continue
		}
		if !isValid1(board, start, i) {
			continue
		}
		board[start][i] = "Q"
		backtracking10(board, n, start+1, res)
		board[start][i] = "."
	}
}

func isValid1(board [][]string, row, col int) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] == "Q" {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
