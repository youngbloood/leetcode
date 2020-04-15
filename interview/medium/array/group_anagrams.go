package array

import (
	"sort"
)

/*

Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.
*/

func GroupAnagrams(strs []string) [][]string {
	return groupAnagrams(strs)
}
func groupAnagrams(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return nil
	}
	var result [][]string
	valMap := make(map[string]int, 0)
	var index int
	for _, str := range strs {
		valB := String(str)
		sort.Sort(valB)

		val := string(valB)
		if iter, ok := valMap[val]; !ok {
			valMap[val] = index
			result = append(result, []string{str})
			index++
		} else {
			result[iter] = append(result[iter], str)
		}
	}
	return result
}

type String []byte

func (s String) Less(i, j int) bool { return s[i] < s[j] }

func (s String) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s String) Len() int { return len(s) }
