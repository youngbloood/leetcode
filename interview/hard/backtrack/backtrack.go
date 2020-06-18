package backtrack

import "fmt"

/*
# Palindrome Partitioning
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/119/backtracking/852/


Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Example:

Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
*/
func Partition(s string) [][]string {
	return partition(s)
}

func partition(s string) [][]string {

	fmt.Println(partitionMaxLen(s))

	return nil
}

// 找出最长的回文串
func partitionMaxLen(s string) string {

	var maxLen, maxIndex int
	for i := 0; i < len(s); i++ {
		for pos := i - 1; pos >= 0 && i+i-pos < len(s); pos-- {
			if isPartition(s[pos : i+i-pos]) {
				if i-pos > maxLen {
					maxLen = i - pos
					maxIndex = i
				}
			}
		}
	}
	return s[maxIndex-maxLen : maxIndex+maxLen]
}

func isPartition(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

/*
# Word Search II
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/119/backtracking/853/


Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.



Example:

Input:
board = [
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]
words = ["oath","pea","eat","rain"]

Output: ["eat","oath"]


Note:

All inputs are consist of lowercase letters a-z.
The values of words are distinct.

Hint #1
You would need to optimize your backtracking to pass the larger test. Could you stop backtracking earlier?
Hint #2
If the current candidate does not exist in all words' prefix, you could stop backtracking immediately. What kind of data structure could answer such query efficiently? Does a hash table work? Why or why not? How about a Trie? If you would like to learn how to implement a basic trie, please work on this problem: Implement Trie (Prefix Tree) first.
*/
func findWords(board [][]byte, words []string) []string {
	// 	将每个byte的位置记录在map中
	boardMap := make(map[byte][]int, 0)

	for i, row := range board {
		for j, v := range row {
			boardMap[v] = append(boardMap[v], i*len(board)+j)
		}
	}

	var result []string

	for _, word := range words {
		// 判断每个word的相邻2个byte在board中是否相邻（上下位置和左右位置）
		for i := 0; i < len(word)-1; i++ {
			b1 := word[i]
			b2 := word[i+1]

			pos1s := boardMap[b1]
			pos2s := boardMap[b2]
			// 不存在
			if pos1s == nil || pos2s == nil {
				continue
			}

		}
	}
}
