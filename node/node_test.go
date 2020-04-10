package node_test

import (
	"leetcode/node"
	"testing"
)

func TestConnect(t *testing.T) {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	inorder := []int{4, 2, 5, 1, 6, 3, 7}
	tree := node.BuildTreePreIn(preorder, inorder)
	tree = node.Connect(tree)
	t.Log(1)
}

func TestConnectNoPerfect(t *testing.T) {
	preorder := []int{1, 2}
	inorder := []int{2, 1}
	tree := node.BuildTreePreIn(preorder, inorder)
	tree = node.Connect2(tree)
	t.Log(1)

	preorder = []int{1, 2, 4, 5, 3, 6, 7}
	inorder = []int{4, 2, 5, 1, 6, 3, 7}
	tree = node.BuildTreePreIn(preorder, inorder)
	tree = node.Connect2(tree)
	t.Log(1)

	preorder = []int{1, 2, 4, 7, 5, 3, 6, 8}
	inorder = []int{7, 4, 2, 5, 1, 3, 6, 8}
	tree = node.BuildTreePreIn(preorder, inorder)
	tree = node.Connect2(tree)
	t.Log(1)
}
