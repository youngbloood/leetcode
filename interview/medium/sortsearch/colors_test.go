package sortsearch_test

import (
	"leetcode/interview/medium/sortsearch"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortsearch.SortColors(nums)
	t.Log(nums)
	nums = []int{2, 0, 0, 1, 1, 0}
	sortsearch.SortColors(nums)
	t.Log(nums)
	nums = []int{1, 0, 2, 1, 1, 0}
	sortsearch.SortColors(nums)
	t.Log(nums)
	nums = []int{2, 0, 2, 2, 2, 0}
	sortsearch.SortColors(nums)
	t.Log(nums)
	nums = []int{2, 0, 1}
	sortsearch.SortColors(nums)
	t.Log(nums)
}
