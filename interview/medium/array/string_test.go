package array_test

import (
	"leetcode/interview/medium/array"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	str := "babad"
	t.Log(array.LongestPalindrome(str))
	str = "cbbd"
	t.Log(array.LongestPalindrome(str))

	str = "aaabaaaa"
	t.Log(array.LongestPalindrome(str))

	str = "a"
	t.Log(array.LongestPalindrome(str))
}
