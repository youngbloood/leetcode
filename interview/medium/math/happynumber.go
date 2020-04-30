package math

/*
#  Happy Number
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/815/

Write an algorithm to determine if a number n is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Return True if n is a happy number, and False if not.

Example:

Input: 19
Output: true
Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
*/
func IsHappy(n int) bool {
	return isHappy(n)
}
func isHappy(n int) bool {
	set := make(map[int]struct{}, 0)
	return isHappyNum(n, set)
}

func isHappyNum(n int, set map[int]struct{}) bool {
	var all int
	for n > 0 {
		num := n % 10
		n /= 10
		if num == 0 {
			continue
		} else if num == 1 {
			all++
		} else {
			all += num * num
		}
	}
	if all == 1 {
		return true
	}
	if isCycle := isHappyCycle(set, all); isCycle {
		return false
	}
	return isHappyNum(all, set)
}

func isHappyCycle(set map[int]struct{}, num int) bool {
	if _, exist := set[num]; exist {
		return true
	}
	set[num] = struct{}{}
	return false
}
