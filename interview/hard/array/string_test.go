package array_test

import (
	"leetcode/interview/hard/array"
	"testing"
)

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	tar := "ABC"
	t.Log(array.MinWindow(s, tar))

	s = "ADOBECODEBANC"
	tar = "ABCC"
	t.Log(array.MinWindow(s, tar))
}
