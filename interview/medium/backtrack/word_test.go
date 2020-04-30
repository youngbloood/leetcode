package backtrack_test

import (
	"leetcode/interview/medium/backtrack"
	"testing"
)

func TestWordExist(t *testing.T) {
	board := [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}
	word := "ABCCED"
	t.Log(backtrack.Exist(board, word))
	word = "SEE"
	t.Log(backtrack.Exist(board, word))
	word = "ABCB"
	t.Log(backtrack.Exist(board, word))

	board = [][]byte{
		[]byte("ABCE"),
		[]byte("SFES"),
		[]byte("ADEE"),
	}
	word = "ABCESEEEFS"
	t.Log(backtrack.Exist(board, word))
}
