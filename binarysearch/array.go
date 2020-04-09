package binarysearch

/*
# https://leetcode.com/explore/learn/card/binary-search/138/background/1038/

Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.


Example 1:
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1


Note:

You may assume that all elements in nums are unique.
n will be in the range [1, 10000].
The value of each element in nums will be in the range [-9999, 9999].
*/

func Search(nums []int, target int) int {
	return search(nums, target)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

/* leetcode上的解法 */

func binarySearch(nums []int, l int, r int, tar int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if nums[mid] == tar {
		return mid
	}
	if nums[l] <= nums[mid] { // 此区间是升序区间
		if tar >= nums[l] && tar < nums[mid] { // tar在升序区间序列中
			return binarySearch(nums, l, mid-1, tar)
		}
		return binarySearch(nums, mid+1, r, tar)
	}

	if tar > nums[mid] && tar <= nums[r] { //tar在mid和r区间中
		return binarySearch(nums, mid+1, r, tar)
	}
	return binarySearch(nums, l, mid-1, tar)
}
func searchInLeetCode(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

/*
# https://leetcode.com/explore/learn/card/binary-search/125/template-i/952/

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
*/
func SearchWithRorated(nums []int, target int) int {
	return searchWithRorated(nums, target)
}

func searchWithRorated(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 && nums[0] == target {
		return 0
	} else if len(nums) == 1 && nums[0] != target {
		return -1
	}
	povit := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			povit = i
			break
		}
	}

	if povit >= 0 && (nums[povit] < target || nums[povit+1] > target) {
		return -1
	}
	var wait []int
	var inRight bool
	if povit < 0 {
		wait = nums
	} else if nums[len(nums)-1] >= target {
		wait = nums[povit+1:]
		inRight = true
	} else {
		wait = nums[:povit+1]
	}

	left, right := 0, len(wait)-1
	for right >= left {
		mid := (left + right) / 2
		if wait[mid] == target {
			if inRight {
				return mid + povit + 1
			}
			return mid
		} else if wait[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
