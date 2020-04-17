package tree_test

import (
	"leetcode/interview/medium/tree"
	"leetcode/treenode"
	"testing"
)

func initTreeNode(vals ...int) *treenode.TreeNode {
	if len(vals) == 0 {
		return nil
	}
	var head = &treenode.TreeNode{Val: vals[0]}
	for i := 1; i < len(vals); i++ {
		// head.Insert(v)
		treenode.InsertTree(head, vals[i])
	}
	return head
}
func TestZigzagLevelOrder(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree1 := treenode.BuildTreePreIn(preorder, inorder)
	t.Log(tree.ZigzagLevelOrder((*tree.TreeNode)(tree1)))
}
