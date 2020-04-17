package tree

import (
	"leetcode/treenode"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode treenode.TreeNode

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/787/

Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]

### goto $GOPATH/src/leetcode/treenode/node.go:923
*/

// golink:
func zigzagLevelOrder(root *TreeNode)

// golink:
func inorderTraversal(root *TreeNode) []int
