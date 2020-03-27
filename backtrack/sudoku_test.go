package backtrack_test

import (
	"fmt"
	"leetcode/backtrack"
	"testing"
)

var zero byte = '.'

func TestSolveSudoku(t *testing.T) {
	board := [][]byte{
		[]byte{5, 3, zero, zero, 7, zero, zero, zero, zero},
		[]byte{6, zero, zero, 1, 9, 5, zero, zero, zero},
		[]byte{zero, 9, 8, zero, zero, zero, zero, 6, zero},

		[]byte{8, zero, zero, zero, 6, zero, zero, zero, 3},
		[]byte{4, zero, zero, 8, zero, 3, zero, zero, 1},
		[]byte{7, zero, zero, zero, 2, zero, zero, zero, 6},

		[]byte{zero, 6, zero, zero, zero, zero, 2, 8, zero},
		[]byte{zero, zero, zero, 4, 1, 9, zero, zero, 5},
		[]byte{zero, zero, zero, zero, 8, zero, zero, 7, 9},
	}

	backtrack.SolveSudoku(board)
	for _, v := range board {
		for _, val := range v {
			fmt.Printf("%v , ", val)
		}
		fmt.Println()
	}

}
