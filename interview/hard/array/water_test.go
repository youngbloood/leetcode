package array_test

import (
	"leetcode/interview/hard/array"
	"testing"
)

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Log(array.MaxArea(height))
}
