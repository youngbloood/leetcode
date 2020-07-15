package june

/*
# Reverse String
# https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/539/week-1-june-1st-june-7th/3350/


Write a function that reverses a string. The input string is given as an array of characters char[].

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

You may assume all the characters consist of printable ascii characters.



Example 1:

Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:

Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
*/
func ReverseString(s []byte) {
	reverseString(s)
}
func reverseString(s []byte) {
	slen := len(s)
	for i := 0; i < slen/2; i++ {
		s[i], s[slen-i-1] = s[slen-i-1], s[i]
	}
}
