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

func TestBuildTreeInPost(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	tree := treenode.BuildTreeInPost(inorder, postorder)
	inorderTree := treenode.InorderTraversal(tree)
	postorderTree := treenode.PostorderTraversal(tree)
	t.Log(inorderTree)
	t.Log(postorderTree)
	t.Log(reflect.DeepEqual(inorder, inorderTree))
	t.Log(reflect.DeepEqual(postorder, postorderTree))
}

func TestBuildTreePreIn(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := treenode.BuildTreePreIn(preorder, inorder)
	inorderTree := treenode.InorderTraversal(tree)
	preorderTree := treenode.PreorderTraversal(tree)
	t.Log(preorderTree)
	t.Log(inorderTree)
	t.Log(reflect.DeepEqual(inorder, inorderTree))
	t.Log(reflect.DeepEqual(preorder, preorderTree))
}

func TestLowestCommonAncestor(t *testing.T) {
	preorder := []int{3, 5, 6, 2, 7, 4, 1, 0, 8}
	inorder := []int{6, 5, 7, 2, 4, 3, 0, 1, 8}
	tree := treenode.BuildTreePreIn(preorder, inorder)
	lca := treenode.LowestCommonAncestor(tree, tree.Left.Left, tree.Right.Left)
	t.Log(lca)

	preorder = []int{-1, 0, -2, 8, 4, 3}
	inorder = []int{8, -2, 0, 4, -1, 3}
	tree = treenode.BuildTreePreIn(preorder, inorder)
	lca = treenode.LowestCommonAncestor(tree, tree.Left.Right, tree.Left.Left.Left)
	t.Log(lca)
}

func TestCodec(t *testing.T) {
	preorder := []int{3, 5, 6, 2, 7, 4, 1, 0, 8}
	inorder := []int{6, 5, 7, 2, 4, 3, 0, 1, 8}
	root := treenode.BuildTreePreIn(preorder, inorder)
	obj := treenode.Constructor()
	data := obj.Serialize(root)
	ans := obj.Deserialize(data)
	t.Log(data)
	t.Log(ans)
}
