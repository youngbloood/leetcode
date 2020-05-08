package array_test

import (
	"leetcode/interview/hard/array"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	t.Log(array.ProductExceptSelf(nums))
	nums = []int{0, 0}
	t.Log(array.ProductExceptSelf(nums))
	nums = []int{1, 0}
	t.Log(array.ProductExceptSelf(nums))

	// Expected: [0,0,0,0]
	nums = []int{1, 0, 0, 1}
	t.Log(array.ProductExceptSelf(nums))
	// Expected: [0,0,0]
	nums = []int{0, 4, 0}
	t.Log(array.ProductExceptSelf(nums))
}

func TestCalculate(t *testing.T) {
	s := "3+2*2"
	t.Log(array.Calculate(s))
	s = " 3/2 "
	t.Log(array.Calculate(s))
	s = " 3+5 / 2 "
	t.Log(array.Calculate(s))
	s = "1-1"
	t.Log(array.Calculate(s))
	// Expected: -2147483647
	s = "0-2147483647"
	t.Log(array.Calculate(s))
	s = "1+2*5/3+6/4*2"
	t.Log(array.Calculate(s))
}

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Log(array.MaxArea(height))
}
