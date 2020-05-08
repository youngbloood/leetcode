package tree_test

import (
	"leetcode/interview/hard/tree"
	"testing"
)

func TestSolve(t *testing.T) {
	board := [][]byte{
		[]byte("XXXX"),
		[]byte("XOOX"),
		[]byte("XXOX"),
		[]byte("XOXX"),
	}
	tree.Solve(board)
	t.Log(board)

	board = [][]byte{
		[]byte("XOXOXO"),
		[]byte("OXOXOX"),
		[]byte("XOXOXO"),
		[]byte("OXOXOX"),
	}
	tree.Solve(board)
	t.Log(board)

}

func TestFindCircleNum(t *testing.T) {
	M := [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 0},
		[]int{0, 0, 1},
	}
	t.Log(tree.FindCircleNum(M))

	M = [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{0, 1, 1},
	}
	t.Log(tree.FindCircleNum(M))

	// [[1,0,0],[0,1,0],[0,0,1]]
	M = [][]int{
		[]int{1, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 1},
	}
	t.Log(tree.FindCircleNum(M))

	// [[1,0,0,1],[0,1,1,0],[0,1,1,1],[1,0,1,1]]
	M = [][]int{
		[]int{1, 0, 0, 1},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 1, 1},
		[]int{1, 0, 1, 1},
	}
	t.Log(tree.FindCircleNum(M))
}
