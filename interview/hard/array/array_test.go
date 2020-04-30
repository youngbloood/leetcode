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
