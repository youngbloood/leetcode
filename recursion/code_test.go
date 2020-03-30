package recursion_test

import (
	"fmt"
	"leetcode/recursion"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	fmt.Println(recursion.LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

func TestPermute(t *testing.T) {
	fmt.Println(recursion.Permute([]int{1, 2, 3}))
}

func TestLetterCombinations(t *testing.T) {
	fmt.Println(recursion.LetterCombinations("23"))
}
