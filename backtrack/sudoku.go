package backtrack

/*
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
Empty cells are indicated by the character '.'.

Note:

The given board contain only digits 1-9 and the character '.'.
You may assume that the given Sudoku puzzle will have a single unique solution.
The given board size is always 9x9.
*/

var zero byte = '.'

func SolveSudoku(board [][]byte) {
	solveSudoku(board)
}

func solveSudoku(board [][]byte) {
	if board == nil {
		return
	}
	if len(board) != 9 || len(board[0]) != 9 {
		return
	}
	sudoku(board, 0, 0)
}

func sudoku(board [][]byte, row int, col int) bool {
	// 终结条件
	if row == 9 {
		return true
	}
	// 一行赋值完，进入下一行
	if col >= 9 {
		return sudoku(board, row+1, 0)
	}

	// 该cell已经有数据，直接进入该行的下一列
	if board[row][col] != zero {
		return sudoku(board, row, col+1)
	}

	// 该cell是空白，可填充数据
	for val := 1; val < 9; val++ {
		// 检查填充的val是否符合sudoku条件
		if hasConflict(board, row, col, byte(val)) {
			continue
		}
		// 满足条件，写入值
		setVal(board, row, col, byte(val))
		// 并进入该行的下一例
		if sudoku(board, row, col+1) {
			return true
		}
		// 不满足最终条件，移除该行该例值
		rmVal(board, row, col)
	}

	return false
}

func setVal(board [][]byte, row, col int, val byte) {
	board[row][col] = byte(val)
}

func rmVal(board [][]byte, row, col int) {
	board[row][col] = zero
}

// 冲突检查
func hasConflict(board [][]byte, row, col int, val byte) bool {

	for i := 0; i < 9; i++ {
		// 例检查
		if i != col && board[row][i] == val {
			return true
		}
		// 行检查
		if i != row && board[i][col] == val {
			return true
		}
	}
	// 小九宫格检查
	nRow := row / 3
	nCol := col / 3

	for inRow := nRow * 3; inRow < (nRow+1)*3; inRow++ {
		for inCol := nCol * 3; inCol < (nCol)*3; inCol++ {
			if (inRow != row && inCol != col) && board[inRow][inCol] == val {
				return true
			}
		}
	}
	return false
}
