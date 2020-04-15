package array_test

import (
	"leetcode/interview/medium/array"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	array.SetZeroes(matrix)
	t.Log(matrix)
}
