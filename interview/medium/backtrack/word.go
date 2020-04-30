package backtrack

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/797/


Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.


Constraints:

board and word consists only of lowercase and uppercase English letters.
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3
*/
func Exist(board [][]byte, word string) bool {
	return exist(board, word)
}
func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[i] {
			if wordDFS(board, visited, word, 0, i, j) {
				return true
			}
			// 重置访问列表
			for vi := range visited {
				for vj := range visited[vi] {
					visited[vi][vj] = false
				}
			}
		}
	}
	return false
}

func wordDFS(board [][]byte, visited [][]bool, word string, n, i, j int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || n < 0 || n >= len(word) || visited[i][j] || board[i][j] != word[n] {
		return false
	}
	visited[i][j] = true
	if n == len(word)-1 && board[i][j] == word[n] {
		return true
	}
	exit := wordDFS(board, visited, word, n+1, i-1, j) ||
		wordDFS(board, visited, word, n+1, i, j-1) ||
		wordDFS(board, visited, word, n+1, i+1, j) ||
		wordDFS(board, visited, word, n+1, i, j+1)
	if exit {
		return exit
	}
	visited[i][j] = false
	return false
}
