package other

import "sort"

/*
# Majority Element
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/114/others/824/

Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:
Input: [3,2,3]
Output: 3

Example 2:
Input: [2,2,1,1,1,2,2]
Output: 2
*/

func MajorityElement(nums []int) int {
	return majorityElement(nums)
}

// 最多元素超过n/2，排序一下，中间那个肯定是最多的
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var res int
	count := 0
	for _, num := range nums {
		if count == 0 {
			res, count = num, 1
		} else {
			if num == res {
				count++
			} else {
				count--
			}
		}
	}

	return res
}
