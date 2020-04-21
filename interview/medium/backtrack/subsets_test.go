package backtrack_test

import (
	"leetcode/interview/medium/backtrack"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	t.Log(backtrack.Subsets(nums))
	nums = []int{9, 0, 3, 5, 7}
	t.Log(backtrack.Subsets(nums))
}
