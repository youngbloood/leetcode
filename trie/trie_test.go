package trie_test

import (
	"leetcode/trie"
	"testing"
)

func TestTrie(t *testing.T) {

	// ["Trie","insert","search","search","startsWith","insert","search"]
	// [[],["apple"],["apple"],["app"],["app"],["app"],["app"]]
	// [null,null,true,false,true,null,true]
	obj := trie.NewTrie()
	word := "apple"
	obj.Insert(word)
	t.Log(obj.Search(word))
	word = "app"
	t.Log(obj.Search(word))

	t.Log(obj.StartsWith(word))

	obj.Insert(word)
	t.Log(obj.Search(word))
}
