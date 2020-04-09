package binarysearch

/*
# https://leetcode.com/explore/learn/card/binary-search/126/template-ii/948/

A peak element is an element that is greater than its neighbors.

Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that nums[-1] = nums[n] = -∞.

Example 1:
Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.

Example 2:
Input: nums = [1,2,1,3,5,6,4]
Output: 1 or 5
Explanation: Your function can return either index number 1 where the peak element is 2,
             or index number 5 where the peak element is 6.
Note:

Your solution should be in logarithmic complexity.
*/

func FindPeakElement(nums []int) int {
	return findPeakElement(nums)
}

func findPeakElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		} else {
			return 1
		}
	}

	left, right := 0, len(nums)-1
	if has, index := findPeak(nums, left, right); has {
		return index
	}
	return -1
}

func findPeak(nums []int, left, right int) (bool, int) {
	mid := (left + right) / 2

	if mid == 0 && nums[mid] > nums[mid+1] {
		return true, mid
	} else if mid == 0 {
		return false, -1
	}
	if mid == len(nums)-1 && nums[mid] > nums[mid-1] {
		return true, mid
	} else if mid == len(nums)-1 {
		return false, -1
	}

	if nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1] {
		return true, mid
	}

	if has, id := findPeak(nums, left, right/2); has {
		return has, id
	}
	return findPeak(nums, right/2+1, right)
}
