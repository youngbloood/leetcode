package tree_test

import "testing"

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}
	t.Log(treenode.NumIslands(grid))
	grid = [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	t.Log(treenode.NumIslands(grid))
}
