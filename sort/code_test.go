package sort_test

import (
	"leetcode/sort"
	"testing"
)

func initInts(vals ...int) []int {
	var result = make([]int, len(vals))
	for i, v := range vals {
		result[i] = v
	}
	return result
}
func TestTopDownSort(t *testing.T) {
	t.Log(sort.TopDownSort(initInts(3, 2, 4, 4, 5, 3, 8, 54, 2, 4, 9)))
}

func TestQuickSort(t *testing.T) {
	sort.QuickSort(initInts(3, 11, 4, 10, 5, 3, 8, 54, 2, 4, 9))
}

func BenchmarkQuickSort(b *testing.B) {
	sort.QuickSort(initInts(3, 2, 4, 4, 5, 3, 8, 54, 2, 4, 9))
}
