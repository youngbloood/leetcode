package array

import (
	"strconv"
	"strings"
)

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

/*
# Basic Calculator II
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/836/

Implement a basic calculator to evaluate a simple expression string.

The expression string contains only non-negative integers, +, -, *, / operators and empty spaces . The integer division should truncate toward zero.

Example 1:

Input: "3+2*2"
Output: 7
Example 2:

Input: " 3/2 "
Output: 1
Example 3:

Input: " 3+5 / 2 "
Output: 5
Note:

You may assume that the given expression is always valid.
Do not use the eval built-in library function.
*/
func Calculate(s string) int {
	return calculate(s)
}

// 12ms
func calculate(s string) int {
	s = strings.TrimSpace(s)
	if num, err := strconv.Atoi(s); err == nil {
		return num
	}

	adds := strings.Split(s, "+")
	var val int
	if len(adds) > 1 {
		for _, add := range adds {
			val += calculate(add)
		}
		return val
	}

	subs := strings.Split(s, "-")
	if len(subs) > 1 {
		if val == 0 {
			val = calculate(subs[0])
		}
		for i := 1; i < len(subs); i++ {
			val -= calculate(subs[i])
		}
		return val
	}

	// *,/ 依次运算

	var before int
	var isMul bool
	for i, b := range s {
		if b == '*' || b == '/' {
			if val == 0 {
				val = calculate(s[:i])
				before = i
				if b == '*' {
					isMul = true
				} else {
					isMul = false
				}
			} else {
				if isMul {
					val *= calculate(s[before+1 : i])
				} else {
					val /= calculate(s[before+1 : i])
				}
				if b == '*' {
					isMul = true
				} else {
					isMul = false
				}
				before = i
			}
		}
	}
	if isMul {
		val *= calculate(s[before+1:])
	} else {
		val /= calculate(s[before+1:])
	}

	return val
}

// leetcode优解:0ms
func calculate2(s string) int {
	stack := []int{}
	num := 0
	operator := '+'

	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			num = num*10 + int(s[i]-'0')
		}
		if i == len(s)-1 || (s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/') {
			if operator == '+' {
				stack = append(stack, num)
			} else if operator == '-' {
				stack = append(stack, num*(-1))
			} else if operator == '*' {
				stack[len(stack)-1] = stack[len(stack)-1] * num
			} else if operator == '/' {
				stack[len(stack)-1] = stack[len(stack)-1] / num
			}
			num = 0
			operator = rune(s[i])
		}
	}

	res := 0

	for _, v := range stack {
		res += v
	}

	return res

}

func isDigit(b byte) bool {
	if b-'0' < 10 {
		return true
	}
	return false
}

/*
# Container With Most Water
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/830/

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.





The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.



Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49



Hint #2
Start with the maximum width container and go to a shorter width container if there is a vertical line longer than the current containers shorter line. This way we are compromising on the width but we are looking forward to a longer length container.
*/

func MaxArea(height []int) int {
	return maxArea(height)
}
func maxArea(height []int) int {
	var area int
	allLenght := len(height) - 1
	for lenght := allLenght; lenght >= 1; lenght-- {
		for i := 0; i <= allLenght-lenght; i++ {
			high := height[i]
			if height[i+lenght] < high {
				high = height[i+lenght]
			}
			if val := high * lenght; val > area {
				area = val
			}
		}
	}
	return area
}

// Leetcode最优解
func maxArea2(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for left < right {
		leftHeight := height[left]
		rightHeight := height[right]
		width := right - left
		if leftHeight < rightHeight {
			if leftHeight*width > maxArea {
				maxArea = leftHeight * width
			}

			left++
		} else {
			if rightHeight*width > maxArea {
				maxArea = rightHeight * width
			}

			right--
		}
	}
	return maxArea
}
