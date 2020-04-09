package binarysearch

/*
# https://leetcode.com/explore/learn/card/binary-search/125/template-i/951/

We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I'll tell you whether the number is higher or lower.

You call a pre-defined API guess(int num) which returns 3 possible results (-1, 1, or 0):

-1 : My number is lower
 1 : My number is higher
 0 : Congrats! You got it!

Example :
Input: n = 10, pick = 6
Output: 6
*/
func GuessNumber(n int) int {
	return guessNumber(n)
}

func guessNumber(n int) int {
	var left, right int
	for {
		val := guess(n)
		if val == 0 {
			return n
		} else if val < 0 {
			left = n
			right = 2 * n
		} else {
			right = n
		}
		n = (left + right) / 2
	}
}

func guess(num int) int {
	if num > 6 {
		return 1
	} else if num < 6 {
		return -1
	}
	return 0
}
