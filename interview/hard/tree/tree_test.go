package tree_test

import (
	"leetcode/interview/hard/tree"
	"leetcode/treenode"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	inorder := []int{7, 11, 2, 4, 13}
	postorder := []int{7, 2, 11, 13, 4}
	node := treenode.BuildTreeInPost(inorder, postorder)
	t.Log(tree.MaxPathSum((*tree.TreeNode)(node)))

	inorder = []int{2, 1, 3}
	postorder = []int{2, 3, 1}
	node = treenode.BuildTreeInPost(inorder, postorder)
	t.Log(tree.MaxPathSum((*tree.TreeNode)(node)))

	inorder = []int{9, -10, 15, 20, 7}
	postorder = []int{9, 15, 7, 20, -10}
	node = treenode.BuildTreeInPost(inorder, postorder)
	t.Log(tree.MaxPathSum((*tree.TreeNode)(node)))

	node = &treenode.TreeNode{
		Val: -3,
	}
	t.Log(tree.MaxPathSum((*tree.TreeNode)(node)))

	inorder = []int{-1, 2, -2}
	postorder = []int{-1, -2, 2}
	node = treenode.BuildTreeInPost(inorder, postorder)
	t.Log(tree.MaxPathSum((*tree.TreeNode)(node)))
}
