package array

import "sort"

/*
# 4Sum II
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/829/

Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.

To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500. All integers are in the range of -228 to 228 - 1 and the result is guaranteed to be at most 231 - 1.

Example:
Input:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

Output:
2

Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/

func FourSumCount(A []int, B []int, C []int, D []int) int {
	return fourSumCount(A, B, C, D)
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	var count int
	ab := make(map[int]int, len(A)*len(B)) // val:count
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			ab[A[i]+B[j]]++
		}
	}
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			val := -1 * (C[i] + D[j])
			count += ab[val]
		}
	}
	return count
}

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
	return findDuplicate(nums)
}

// 抽屉原理
func findDuplicate(nums []int) int {
	left, right := 1, len(nums)
	for left < right {
		mid := left + (right-left)/2
		cnt := 0
		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

/*
# Sliding Window Maximum
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/837/

Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

Follow up:
Could you solve it in linear time?

Example:

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
Explanation:

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

Constraints:

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
*/
func MaxSlidingWindow(nums []int, k int) []int {
	return maxSlidingWindow(nums, k)
}
func maxSlidingWindow(nums []int, k int) []int {
	// start, end := 0, k-1
	// var maxPtr int
	// result := make([]int, 0, len(nums)-k+1)
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] > max {
	// 		max = nums[i]
	// 	}
	// 	if i == end {

	// 	}
	// }
	return nil
}
