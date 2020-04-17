package array

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/781/

Given an unsorted array return whether an increasing subsequence of length 3 exists or not in the array.

Formally the function should:

Return true if there exists i, j, k
such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.
Note: Your algorithm should run in O(n) time complexity and O(1) space complexity.

Example 1:
Input: [1,2,3,4,5]
Output: true

Example 2:
Input: [5,4,3,2,1]
Output: false

Example 3:
Input: [3,2,1,-1,5,6]
Output: true

*/
func IncreasingTriplet(nums []int) bool {
	return increasingTriplet(nums)
}
func increasingTriplet(nums []int) bool {
	if nums == nil || len(nums) <= 2 {
		return false
	}
	a, b := 1<<31, 1<<31
	for i := 0; i < len(nums); i++ {
		if nums[i] <= a { // nums[i]<=a
			a = nums[i]
		} else if nums[i] <= b { // a <= nums[i] <= b
			b = nums[i]
		} else { // a <= b <= nums[i]
			return true
		}
	}
	return false
}
