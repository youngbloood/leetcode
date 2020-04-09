package treenode_test

import (
	"fmt"
	"leetcode/treenode"
	"reflect"
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

func TestPreorderTraversal(t *testing.T) {
	tree1 := initTreeNode(4, 2, 6, 3, 5, 1, 7)
	fmt.Println(treenode.PreorderTraversal(tree1))
}

func TestPostorderTraversal(t *testing.T) {
	tree1 := initTreeNode(4, 2, 6, 3, 5, 1, 7)
	fmt.Println(treenode.PostorderTraversal(tree1))
}

func TestBuildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	tree := treenode.BuildTree(inorder, postorder)
	inorderTree := treenode.InorderTraversal(tree)
	postorderTree := treenode.PostorderTraversal(tree)
	t.Log(inorderTree)
	t.Log(postorderTree)
	t.Log(reflect.DeepEqual(inorder, inorderTree))
	t.Log(reflect.DeepEqual(postorder, postorderTree))
}
