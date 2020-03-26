/*
pseudocode template:

Python:

def backtrack(candidate):
    if find_solution(candidate):
        output(candidate)
        return

    # iterate all possible candidates.
    for next_candidate in list_of_candidates:
        if is_valid(next_candidate):
            # try this partial candidate solution
            place(next_candidate)
            # given the candidate, explore further.
            backtrack(next_candidate)
            # backtrack
            remove(next_candidate)
*/
package backtrack

/*
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

Given an integer n, return the number of distinct solutions to the n-queens puzzle.

Example:

Input: 4
Output: 2
Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
[
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
*/
func TotalNQueens(n int) int {
	return totalNQueens(n)
}

func totalNQueens(n int) int {
	var board [][]byte
	for row := 0; row < n; row++ {
		board = append(board, make([]byte, n))
	}
	return backtrackNQueen(board, n, 0, 0)
}

func backtrackNQueen(broad [][]byte, n, row, count int) int {
	for col := 0; col < n; col++ {
		if !isUnderAttack(broad, n, row, col) {
			placeQueen(broad, row, col)
			if row+1 == n {
				count++
			} else {
				count = backtrackNQueen(broad, n, row+1, count)
			}
			removeQueen(broad, row, col)
		}
	}
	return count
}

func placeQueen(board [][]byte, row, col int) {
	board[row][col] = 1
}
func removeQueen(board [][]byte, row, col int) {
	board[row][col] = 0
}

func isUnderAttack(board [][]byte, n, row, col int) bool {

	for i := 0; i < n; i++ {
		if board[row][i] == 1 {
			return true
		}
	}

	for i := 0; i < n; i++ {
		if board[i][col] == 1 {
			return true
		}
	}

	rowTemp, colTemp := row, col
	for rowTemp > 0 && colTemp < n-1 {
		rowTemp--
		colTemp++
		if board[rowTemp][colTemp] == 1 {
			return true
		}
	}
	rowTemp, colTemp = row, col
	for rowTemp < n-1 && colTemp > 0 {
		rowTemp++
		colTemp--
		if board[rowTemp][colTemp] == 1 {
			return true
		}
	}
	rowTemp, colTemp = row, col
	for rowTemp > 0 && colTemp > 0 {
		rowTemp--
		colTemp--
		if board[rowTemp][colTemp] == 1 {
			return true
		}
	}
	rowTemp, colTemp = row, col
	for rowTemp < n-1 && colTemp < n-1 {
		rowTemp++
		colTemp++
		if board[rowTemp][colTemp] == 1 {
			return true
		}
	}

	return false
}
