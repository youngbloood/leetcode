package array_test

import (
	"leetcode/interview/hard/array"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	t.Log(array.SpiralOrder(matrix))

	matrix = [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}
	t.Log(array.SpiralOrder(matrix))

	matrix = [][]int{
		[]int{1},
	}
	t.Log(array.SpiralOrder(matrix))
	matrix = [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}
	t.Log(array.SpiralOrder(matrix))
}
