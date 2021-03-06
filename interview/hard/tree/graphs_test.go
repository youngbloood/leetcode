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

func TestLongestIncreasingPath(t *testing.T) {
	matrix := [][]int{
		[]int{9, 9, 4},
		[]int{6, 6, 8},
		[]int{2, 1, 1},
	}
	t.Log(tree.LongestIncreasingPath(matrix))

	matrix = [][]int{
		[]int{3, 4, 5},
		[]int{3, 2, 6},
		[]int{2, 2, 1},
	}
	t.Log(tree.LongestIncreasingPath(matrix))

	matrix = [][]int{}
	t.Log(tree.LongestIncreasingPath(matrix))

	matrix = [][]int{
		[]int{2, 1},
	}
	t.Log(tree.LongestIncreasingPath(matrix))

	// Expected：6
	matrix = [][]int{
		[]int{7, 8, 9},
		[]int{9, 7, 6},
		[]int{7, 2, 3},
	}
	t.Log(tree.LongestIncreasingPath(matrix))
	// Expected:2
	matrix = [][]int{
		[]int{2147483647, 1},
	}
	t.Log(tree.LongestIncreasingPath(matrix))

	matrix = [][]int{
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{19, 18, 17, 16, 15, 14, 13, 12, 11, 10},
		[]int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
		[]int{39, 38, 37, 36, 35, 34, 33, 32, 31, 30},
		[]int{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
		[]int{59, 58, 57, 56, 55, 54, 53, 52, 51, 50},
		[]int{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
		[]int{79, 78, 77, 76, 75, 74, 73, 72, 71, 70},
		[]int{80, 81, 82, 83, 84, 85, 86, 87, 88, 89},
		[]int{99, 98, 97, 96, 95, 94, 93, 92, 91, 90},
		[]int{100, 101, 102, 103, 104, 105, 106, 107, 108, 109},
		[]int{119, 118, 117, 116, 115, 114, 113, 112, 111, 110},
		[]int{120, 121, 122, 123, 124, 125, 126, 127, 128, 129},
		[]int{139, 138, 137, 136, 135, 134, 133, 132, 131, 130},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	t.Log(tree.LongestIncreasingPath(matrix))
}

func TestCountSmaller(t *testing.T) {
	nums := []int{5, 2, 6, 1}
	t.Log(tree.CountSmaller(nums))
}
