package array

import "math"

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/779/

Given a string, find the length of the longest substring without repeating characters.

Example 1:
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func LengthOfLongestSubstring(s string) int {
	return 0
}
func lengthOfLongestSubstring(s string) int {
	return 0
}

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/780/

Solution
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1)
		l := int(math.Max(float64(l1), float64(l2)))
		if l > end-start {
			start = i - (l-1)/2
			end = i + l/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

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

}
func increasingTriplet(nums []int) bool {

}
