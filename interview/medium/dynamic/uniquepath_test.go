package dynamic_test

import (
	"leetcode/interview/medium/dynamic"
	"testing"
)

func TestUniquePaths(t *testing.T){
	t.Log(dynamic.UniquePaths(3,2))
	t.Log(dynamic.UniquePaths(7,3))
	t.Log(dynamic.UniquePaths(51,9))
}