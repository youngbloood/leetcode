package array

/*
# Minimum Window Substring
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/838/


Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

Hint #1
Use two pointers to create a window of letters in S, which would have all the characters from T.

Hint #2
Since you have to find the minimum window in S which has all the characters from T, you need to expand and contract the window using the two pointers and keep checking the window for all the characters. This approach is also called Sliding Window Approach.

L ------------------------ R , Suppose this is the window that contains all characters of T

        L----------------- R , this is the contracted window. We found a smaller window that still contains all the characters in T

When the window is no longer valid, start expanding again using the right pointer.
*/
func MinWindow(s string, t string) string {
	return minWindow(s, t)
}

// sliding window
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	// Dictionary which keeps a count of all the unique characters in t.
	dictT := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		dictT[t[i]]++
	}

	// Number of unique characters in t, which need to be present in the desired window.
	required := len(dictT)

	// Left and Right pointer
	l, r := 0, 0

	// formed is used to keep track of how many unique characters in t
	// are present in the current window in its desired frequency.
	// e.g. if t is "AABC" then the window must have two A's, one B and one C.
	// Thus formed would be = 3 when all these conditions are met.
	formed := 0

	// Dictionary which keeps a count of all the unique characters in the current window.
	windowCounts := make(map[byte]int, 0)

	// ans list of the form (window length, left, right)
	ans := []int{-1, 0, 0}
	for r < len(s) {
		// Add one character from the right to the window
		c := s[r]
		windowCounts[c]++

		// If the frequency of the current character added equals to the
		// desired count in t then increment the formed count by 1.
		if windowCounts[c] == dictT[c] {
			formed++
		}
		for l <= r && formed == required {
			c = s[l]
			// Save the smallest window until now.
			if ans[0] == -1 || r-l+1 < ans[0] {
				ans[0] = r - l + 1
				ans[1] = l
				ans[2] = r
			}

			// The character at the position pointed by the
			// `Left` pointer is no longer a part of the window.
			windowCounts[c]--
			if windowCounts[c] < dictT[c] {
				formed--
			}

			// Move the left pointer ahead, this would help to look for a new window.
			l++
		}

		// Keep expanding the window once we are done contracting.
		r++
	}

	if ans[0] == -1 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}

// leetcode 优解
func minWindow2(s string, t string) string {
	result := ""
	tMap := make([]int, 128)
	sMap := make([]int, 128)
	required := 0
	for _, val := range t {
		if tMap[val] == 0 {
			required++
		}
		tMap[val]++
	}

	counter := 0
	for r, l := 0, 0; r < len(s); r++ {
		sMap[s[r]]++
		if sMap[s[r]] == tMap[s[r]] {
			counter++
		}
		if counter == required {
			for l < r {
				sMap[s[l]]--
				if sMap[s[l]] < tMap[s[l]] {
					counter--
					break
				}
				l++
			}
			if result == "" || len(s[l:r+1]) < len(result) {
				result = s[l : r+1]
			}
			l++
		}

	}

	return result
}
