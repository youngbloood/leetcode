package treenode_test

import (
	"fmt"
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
func TestGenerateTrees(t *testing.T) {
	// head := initTreeNode(1, 2, 3)

	fmt.Println(treenode.GenerateTrees(3))
}
