package tree

/*
# Word Ladder
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/842/

Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list.
Note:

Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output: 5

Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
return its length 5.

Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: 0

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
*/
func LadderLength(beginWord string, endWord string, wordList []string) int {
	return ladderLength(beginWord, endWord, wordList)
}

// Breadth First Search
func ladderLength(beginWord string, endWord string, wordList []string) int {

	L := len(beginWord)

	//  // Dictionary to hold combination of words that can be formed,
	//  // from any given word. By changing one letter at a time.
	allComboDict := make(map[string][]string)

	for _, word := range wordList {
		for i := 0; i < L; i++ {

			// Key is the generic word
			// Value is a list of words which have the same intermediate generic word.
			newWord := word[:i] + "*" + word[i+1:]
			allComboDict[newWord] = append(allComboDict[newWord], word)
		}
	}

	type pair struct {
		word  string
		level int
	}

	// Queue for BFS
	Q := make([]pair, 0)

	Q = append(Q, pair{word: beginWord, level: 1})

	// Visited to make sure we don't repeat processing same word.
	visited := make(map[string]bool, 0)
	visited[beginWord] = true

	for len(Q) != 0 {
		node := Q[len(Q)-1]
		Q = Q[:len(Q)-1]
		word := node.word
		level := node.level
		for i := 0; i < L; i++ {

			// Intermediate words for current word
			newWord := word[:i] + "*" + word[i+1:]

			// Next states are all the words which share the same intermediate state.
			for _, adjacentWord := range allComboDict[newWord] {

				// If at any point if we find what we are looking for
				// i.e. the end word - we can return with the answer.
				if adjacentWord == endWord {
					return level + 1
				}

				// Otherwise, add it to the BFS Queue. Also mark it visited
				if !visited[adjacentWord] {
					visited[adjacentWord] = true
					Q = append(Q, pair{word: adjacentWord, level: level + 1})
				}
			}
		}
	}

	return 0
}
