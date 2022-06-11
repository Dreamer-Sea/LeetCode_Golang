package main

func main() {

}

func solveSudoku(board [][]byte) {
	var backtracking func(board [][]byte) bool
	backtracking = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				for k := '1'; k <= '9'; k++ {
					if !isValid4(board, byte(k), i, j) {
						continue
					}
					board[i][j] = byte(k)
					if backtracking(board) {
						return true
					}
					board[i][j] = '.'
				}
				return false
			}
		}
		return true
	}
	backtracking(board)
}

func isValid4(board [][]byte, val byte, x, y int) bool {
	for i := 0; i < 9; i++ {
		if board[i][y] == val {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[x][i] == val {
			return false
		}
	}
	startRow := (x / 3) * 3
	startCol := (y / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}
