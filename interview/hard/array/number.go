package array

import "sort"

/*
# First Missing Positive
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/832/

Given an unsorted integer array, find the smallest missing positive integer.

Example 1:

Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2
Example 3:

Input: [7,8,9,11,12]
Output: 1
Note:

Your algorithm should run in O(n) time and uses constant extra space.
*/
func FirstMissingPositive(nums []int) int {
	return firstMissingPositive(nums)
}
func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	max := 1
	hasCompare := false
	hasFirst := false
	for _, v := range nums {
		if v <= 0 {
			continue
		}
		if v == max {
			hasCompare = true
			hasFirst = true
			continue
		} else if v == max+1 && hasFirst {
			hasCompare = true
			max = v
		} else {
			break
		}
	}
	if hasCompare {
		return max + 1
	}
	return max
}

/*
# Longest Consecutive Sequence
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/833/

Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Your algorithm should run in O(n) complexity.

Example:

Input: [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
*/

func LongestConsecutive(nums []int) int {
	return longestConsecutive(nums)
}
func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	maxLenght, lenght := 0, 1
	for i := range nums {
		if i == 0 {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			lenght++
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			if lenght > maxLenght {
				maxLenght = lenght
			}
			lenght = 1
		}
	}
	if lenght > maxLenght {
		return lenght
	}
	return maxLenght
}

/*
#  Find the Duplicate Number
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/834/

Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

Example 1:

Input: [1,3,4,2,2]
Output: 2
Example 2:

Input: [3,1,3,4,2]
Output: 3
Note:

You must not modify the array (assume the array is read only).
You must use only constant, O(1) extra space.
Your runtime complexity should be less than O(n2).
There is only one duplicate number in the array, but it could be repeated more than once.
*/
func FindDuplicate(nums []int) int {

}
func findDuplicate(nums []int) int {

}
