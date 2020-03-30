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

	fmt.Println(treenode.GenerateTrees(4))
}

func TestCopyTree(t *testing.T) {
	tree1 := initTreeNode(1, 2, 3)
	fmt.Println("tree1 = ", tree1)
	// tree2 := tree1.Copy()
	tree2 := treenode.CopyTree(tree1)
	fmt.Println("tree2 = ", tree2)
}

func TestIsValidBST(t *testing.T) {
	tree1 := initTreeNode(1, 2, 3)
	fmt.Println(treenode.IsValidBST(tree1))
}

func TestIsSameTree(t *testing.T) {
	tree1 := initTreeNode(1, 2, 3)
	tree2 := initTreeNode(1, 2, 3)
	fmt.Println(treenode.IsSameTree(tree1, tree2))
}

func TestLevelOrder(t *testing.T) {
	tree1 := initTreeNode(4, 2, 6, 3, 5, 1, 7)
	fmt.Println(treenode.LevelOrder(tree1))
}
