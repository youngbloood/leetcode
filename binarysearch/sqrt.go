package binarysearch

/*
# https://leetcode.com/explore/learn/card/binary-search/125/template-i/950/

Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:
Input: 4
Output: 2

Example 2:
Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
*/

func MySqrt(x int) int {
	return mySqrt(x)
}

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}
	mid := x / 2
	left, right := 0, x
	for right >= left {
		square := mid * mid
		if square == x {
			return mid
		} else if square > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return mid
}
