package kthgrammar

/*
On the first row, we write a 0. Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.
Given row N and index K, return the K-th indexed symbol in row N. (The values of K are 1-indexed.) (1 indexed).

Examples:
Input: N = 1, K = 1
Output: 0

Input: N = 2, K = 1
Output: 0

Input: N = 2, K = 2
Output: 1

Input: N = 4, K = 5
Output: 1

Explanation:
row 1: 0
row 2: 01
row 3: 0110
row 4: 01101001
row 5: 0110100110010110

Note:

N will be an integer in the range [1, 30].
K will be an integer in the range [1, 2^(N-1)].
*/

func KthGrammar(N int, K int) int {
	if N < 1 || K < 1 {
		return -1
	}
	switch {
	case N == 1 && K == 1:
		return 0
	case N == 2 && K == 1:
		return 0
	case N == 2 && K == 2:
		return 1
	}

	KTemp := K
	//
	N--
	K = (K + 1) / 2
	val := KthGrammar(N, K)

	switch val {
	case 0:
		if KTemp%2 == 0 {
			return 1
		}
	case 1:
		if KTemp%2 != 0 {
			return 1
		}
	}
	return 0
}
