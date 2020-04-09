package binarysearch_test

import (
	"fmt"
	"leetcode/binarysearch"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	fmt.Println(binarysearch.FindPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(binarysearch.FindPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(binarysearch.FindPeakElement([]int{1, 2}))
	fmt.Println(binarysearch.FindPeakElement([]int{2, 1}))
	fmt.Println(binarysearch.FindPeakElement([]int{1, 2, 3, 4}))
	fmt.Println(binarysearch.FindPeakElement([]int{1, 3, 2, 1}))
}
