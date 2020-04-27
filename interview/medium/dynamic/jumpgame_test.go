package dynamic_test

import (
	"leetcode/interview/medium/dynamic"
	"testing"
)

func TestCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	t.Log(dynamic.CanJump(nums))
	nums = []int{3, 2, 1, 0, 4}
	t.Log(dynamic.CanJump(nums))
}
