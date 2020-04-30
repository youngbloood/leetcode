package dynamic_test

import (
	"leetcode/interview/medium/dynamic"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	t.Log(dynamic.CoinChange(coins, amount))

	coins = []int{2}
	amount = 3
	t.Log(dynamic.CoinChange(coins, amount))

	coins = []int{1}
	amount = 0
	t.Log(dynamic.CoinChange(coins, amount))

	coins = []int{2, 5, 10, 1}
	amount = 27
	t.Log(dynamic.CoinChange(coins, amount))

	// 此例测试不通过
	// OUTPUT： 26
	// EXPECTED: 20
	coins = []int{186, 419, 83, 408}
	amount = 6249
	t.Log(dynamic.CoinChange(coins, amount))
}
