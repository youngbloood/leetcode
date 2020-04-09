package binarysearch_test

import (
	"leetcode/binarysearch"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Log(binarysearch.Search([]int{-1, 0, 3, 5, 9, 12}, 9))
	t.Log(binarysearch.Search([]int{-1, 0, 3, 5, 9, 12}, 2))
	t.Log(binarysearch.Search([]int{-1, 0, 3, 5, 9, 12}, 13))
	t.Log(binarysearch.Search([]int{5}, 5))
	t.Log(binarysearch.Search([]int{2, 5}, 2))
}

func TestSearchWithRorated(t *testing.T) {
	t.Log(binarysearch.SearchWithRorated([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	t.Log(binarysearch.SearchWithRorated([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	t.Log(binarysearch.SearchWithRorated([]int{}, 5))
	t.Log(binarysearch.SearchWithRorated([]int{5}, 5))
	t.Log(binarysearch.SearchWithRorated([]int{5}, 4))
	t.Log(binarysearch.SearchWithRorated([]int{1, 3}, 1))
	t.Log(binarysearch.SearchWithRorated([]int{3, 1}, 3))
}
