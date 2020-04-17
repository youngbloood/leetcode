package treenode_test

import (
	"leetcode/treenode"
	"testing"
)

func TestConstructor(t *testing.T) {
	preorder := []int{7, 3, 15, 9, 20}
	inorder := []int{3, 7, 9, 15, 20}
	root := treenode.BuildTreePreIn(preorder, inorder)
	obj := treenode.BSTConstructor(root)
	for obj.HasNext() {
		t.Log(obj.Next())
	}
}

func TestInsertIntoBST(t *testing.T) {
	root := &treenode.TreeNode{Val: 5}
	wait := []int{4, 2, 7, 1, 3}
	for _, v := range wait {
		root = treenode.InsertIntoBST(root, v)
	}
	t.Log(1)
}

func TestDeleteNode(t *testing.T) {
	preorder := []int{5, 3, 2, 4, 6, 7}
	inorder := []int{2, 3, 4, 5, 6, 7}
	root := treenode.BuildTreePreIn(preorder, inorder)

	// root = treenode.DeleteNode(root, 4)
	root = treenode.DeleteNode(root, 5)
	// root = treenode.DeleteNode(root, 6)

	root = &treenode.TreeNode{Val: 0}
	root = treenode.DeleteNode(root, 0)

	preorder = []int{1, 2}
	inorder = []int{1, 2}
	root = treenode.BuildTreePreIn(preorder, inorder)
	root = treenode.DeleteNode(root, 2)

	preorder = []int{3, 1, 2, 4}
	inorder = []int{1, 2, 3, 4}
	root = treenode.BuildTreePreIn(preorder, inorder)
	root = treenode.DeleteNode(root, 1)
}

func TestConstructorKthLargest(t *testing.T) {
	obj := treenode.ConstructorKthLargest(3, []int{4, 5, 8, 2})
	param_1 := obj.Add(3)
	t.Log(param_1)
	param_1 = obj.Add(5)
	t.Log(param_1)
	param_1 = obj.Add(10)
	t.Log(param_1)
	param_1 = obj.Add(9)
	t.Log(param_1)
	param_1 = obj.Add(4)
	t.Log(param_1)

	// ["KthLargest","add","add","add","add","add"]
	// [[1,[]],[-3],[-2],[-4],[0],[4]]
	obj = treenode.ConstructorKthLargest(1, []int{})
	param_1 = obj.Add(-3)
	param_1 = obj.Add(-2)
	param_1 = obj.Add(-4)
	param_1 = obj.Add(0)
	param_1 = obj.Add(4)

}

func TestKthSmallest(t *testing.T) {
	/*
				3
			   / \
			  1   4
		       \
				2
	*/
	preorder := []int{3, 1, 2, 4}
	inorder := []int{1, 2, 3, 4}
	tree1 := treenode.BuildTreePreIn(preorder, inorder)
	t.Log(treenode.KthSmallest(tree1, 1))
	t.Log(treenode.KthSmallest(tree1, 2))
	t.Log(treenode.KthSmallest(tree1, 3))
	t.Log(treenode.KthSmallest(tree1, 4))

	/*
			   5
		      / \
		     3   6
		    / \
		   2   4
		  /
		 1
	*/
	preorder = []int{5, 3, 2, 1, 4, 6}
	inorder = []int{1, 2, 3, 4, 5, 6}
	tree1 = treenode.BuildTreePreIn(preorder, inorder)
	t.Log(treenode.KthSmallest(tree1, 1))
	t.Log(treenode.KthSmallest(tree1, 2))
	t.Log(treenode.KthSmallest(tree1, 3))
	t.Log(treenode.KthSmallest(tree1, 4))
	t.Log(treenode.KthSmallest(tree1, 5))
	t.Log(treenode.KthSmallest(tree1, 6))

}
