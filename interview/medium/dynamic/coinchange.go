package dynamic

import (
	"sort"
)

/*
# Coin Change
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/111/dynamic-programming/809/

You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.
*/

// a*coins[0]+b*coins[1]+c*coins[2]+...+z*coins[25]=amount

func CoinChange(coins []int, amount int) int {
	return coinChange(coins, amount)
}
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	return peakCoin(coins, amount, 0)
}

func peakCoin(coins []int, amount, nums int) int {
	if amount == 0 {
		return 0
	}
	for i := len(coins) - 1; i >= 0; i-- {
		max := amount / coins[i]
		if max == 0 {
			continue
		}
		for j := max; j > 0; j-- {
			left := amount - j*coins[i]
			if left == 0 {
				return nums + j
			}
			elseNum := peakCoin(coins[:len(coins)-1], left, nums+j)
			if elseNum == -1 {
				continue
			} else {
				return elseNum
			}
		}
	}
	return -1
}
