package array_test

import (
	"leetcode/interview/hard/array"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{1, 2, 0}
	t.Log(array.FirstMissingPositive(nums))
	nums = []int{3, 4, -1, 1}
	t.Log(array.FirstMissingPositive(nums))
	nums = []int{7, 8, 9, 11, 12}
	t.Log(array.FirstMissingPositive(nums))
	nums = []int{0, 2, 2, 1, 1}
	t.Log(array.FirstMissingPositive(nums))
	nums = []int{0, 2, 2, 1, 1, 4, 4}
	t.Log(array.FirstMissingPositive(nums))
	nums = []int{}
	t.Log(array.FirstMissingPositive(nums))
	nums = []int{2}
	t.Log(array.FirstMissingPositive(nums))
	nums = []int{1}
	t.Log(array.FirstMissingPositive(nums))
	nums = []int{3, 4, 0, 2}
	t.Log(array.FirstMissingPositive(nums))
}

func TestLongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	t.Log(array.LongestConsecutive(nums))
	nums = []int{0}
	t.Log(array.LongestConsecutive(nums))
	nums = []int{0, -1}
	t.Log(array.LongestConsecutive(nums))
	nums = []int{1, 2, 0, 1}
	t.Log(array.LongestConsecutive(nums))

	nums = []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	t.Log(array.LongestConsecutive(nums))

}
