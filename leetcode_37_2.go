package main

func main() {

}

func solveSudoku1(board [][]byte) {
	var backtracking func() bool
	backtracking = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				for k := '1'; k <= '9'; k++ {
					if !isValid5(board, byte(k), i, j) {
						continue
					}
					board[i][j] = byte(k)
					if backtracking() {
						return true
					}
					board[i][j] = '.'
				}
				return false
			}
		}
		return true
	}
	backtracking()
}

func isValid5(board [][]byte, val byte, row, col int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][col] == val {
			return false
		}
	}
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}
