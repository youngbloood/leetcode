package array_test

import (
	"leetcode/interview/medium/array"
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	t.Log(array.IncreasingTriplet(nums))
	nums = []int{5, 4, 3, 2, 1}
	t.Log(array.IncreasingTriplet(nums))
	nums = []int{3, 2, 1, -1, 5, 6}
	t.Log(array.IncreasingTriplet(nums))

	nums = []int{2, 1, 5, 0, 3}
	t.Log(array.IncreasingTriplet(nums))

	nums = []int{5, 1, 5, 5, 2, 5, 4}
	t.Log(array.IncreasingTriplet(nums))

	nums = []int{1, 0, 0, 0, 0, 10, 0, 0, 0, 0, 0, 1000}
	t.Log(array.IncreasingTriplet(nums))

	nums = []int{1, 0, 0, 2, 0, 0, 0, 0, 0, 0, -1, -1, 3}
	t.Log(array.IncreasingTriplet(nums))
}
