package array_test

import (
	"leetcode/interview/medium/array"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	t.Log(array.ThreeSum(nums))

	nums = []int{-2, 0, 0, 2, 2}
	t.Log(array.ThreeSum(nums))

	nums = []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	t.Log(array.ThreeSum(nums))
}
