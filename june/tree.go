package june

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*

# Invert Binary Tree
# https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/539/week-1-june-1st-june-7th/3347/

Example:

Input:

     4
   /   \
  2     7
 / \   / \
1   3 6   9
Output:

     4
   /   \
  7     2
 / \   / \
9   6 3   1
Trivia:
This problem was inspired by this original tweet by Max Howell:

Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
