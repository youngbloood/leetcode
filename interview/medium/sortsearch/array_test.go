package sortsearch_test

import (
	"leetcode/interview/medium/sortsearch"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	t.Log(sortsearch.TopKFrequent(nums, k))

	k = 1
	t.Log(sortsearch.TopKFrequent(nums, k))

	nums = []int{1}
	k = 1
	t.Log(sortsearch.TopKFrequent(nums, k))
}

func TestFindPeakElement(t *testing.T) {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	t.Log(sortsearch.FindPeakElement(nums))
	nums = []int{1}
	t.Log(sortsearch.FindPeakElement(nums))
}

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	t.Log(sortsearch.SearchRange(nums, target))
	target = 6
	t.Log(sortsearch.SearchRange(nums, target))

	nums = []int{1}
	target = 1
	t.Log(sortsearch.SearchRange(nums, target))
}

func TestMerge(t *testing.T) {
	intervals := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	t.Log(sortsearch.Merge(intervals))
	intervals = [][]int{
		[]int{1, 4},
		[]int{4, 5},
		[]int{5, 7},
	}

	t.Log(sortsearch.Merge(intervals))

	intervals = [][]int{
		[]int{1, 4},
		[]int{0, 4},
	}
	t.Log(sortsearch.Merge(intervals))

	intervals = [][]int{
		[]int{1, 4},
		[]int{0, 2},
		[]int{3, 5},
	}
	t.Log(sortsearch.Merge(intervals))

	// Expected: [[2,4],[5,5]]
	intervals = [][]int{
		[]int{2, 3},
		[]int{5, 5},
		[]int{2, 2},
		[]int{3, 4},
		[]int{3, 4},
	}
	t.Log(sortsearch.Merge(intervals))

	// Expected:[[1,3],[4,7]]
	intervals = [][]int{
		[]int{2, 3},
		[]int{2, 2},
		[]int{3, 3},
		[]int{1, 3},
		[]int{5, 7},
		[]int{2, 2},
		[]int{4, 6},
	}
	t.Log(sortsearch.Merge(intervals))
}
