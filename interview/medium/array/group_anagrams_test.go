package array_test

import (
	"leetcode/interview/medium/array"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}
	t.Log(array.GroupAnagrams(strs))
	strs = []string{
		"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc",
	}
	t.Log(array.GroupAnagrams(strs))
}
