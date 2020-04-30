package array

/*
# Product of Array Except Self
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/827/


Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].

Example:

Input:  [1,2,3,4]
Output: [24,12,8,6]

Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array (including the whole array) fits in a 32 bit integer.

Note: Please solve it without division and in O(n).

Follow up:
Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)
*/
func ProductExceptSelf(nums []int) []int {
	return productExceptSelf(nums)
}

func productExceptSelf(nums []int) []int {
	var val int = 1
	var zeroNum int
	for i := range nums {
		if nums[i] != 0 {
			val *= nums[i]
		} else {
			zeroNum++
		}
	}

	if zeroNum == len(nums) {
		return nums
	}

	for i := range nums {
		switch zeroNum {
		case 0:
			nums[i] = val / nums[i]
		case 1:
			if nums[i] == 0 {
				nums[i] = val
			} else {
				nums[i] = 0
			}
		default:
			nums[i] = 0
		}
	}
	return nums
}
