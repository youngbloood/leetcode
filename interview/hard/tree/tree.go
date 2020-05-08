package tree

import (
	"leetcode/treenode"
	"math"
)

// Definition for a binary tree node.
type TreeNode treenode.TreeNode

/*
# Binary Tree Maximum Path Sum
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/845/

Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

Example 1:

Input: [1,2,3]

       1
      / \
     2   3

Output: 6

Example 2:
Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42
*/
func MaxPathSum(root *TreeNode) int {
	return maxPathSum(root)
}

func maxPathSum(root *TreeNode) int {
	var all = math.MinInt32
	depthWaySum := maxPathSumDFS(root, &all)
	if all > depthWaySum {
		return all
	}
	return depthWaySum
}

func maxPathSumDFS(root *TreeNode, all *int) int {
	if root == nil {
		return 0
	}

	left := maxPathSumDFS((*TreeNode)(root.Left), all)
	if left < 0 {
		left = 0
	}
	right := maxPathSumDFS((*TreeNode)(root.Right), all)
	if right < 0 {
		right = 0
	}
	allNode := left + right + root.Val
	if allNode > *all {
		*all = allNode
	}
	if left > right {
		return left + root.Val
	}
	return right + root.Val
}
