/*
# https://leetcode.com/explore/learn/card/n-ary-tree/
n-ary
*/
package nary

/**
 * Definition for a Node.
 */
type Node struct {
	Val      int
	Children []*Node
}

/*
# https://leetcode.com/explore/learn/card/n-ary-tree/130/traversal/925/

Given an n-ary tree, return the preorder traversal of its nodes' values.
Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

Follow up:
Recursive solution is trivial, could you do it iteratively?

Example 1:

			1
		  / | \
		 3  2  4
		/ \
       5   6

Input: root = [1,null,3,2,4,null,5,6]
Output: [1,3,5,6,2,4]

Example 2:

				1
		   /  /   \  \
		  2  3     4  5
			/ \    |  | \
		   6   7   8  9  10
			   |   |  |
			   11  12 13
			   |
			   14

Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]


Constraints:

The height of the n-ary tree is less than or equal to 1000
The total number of nodes is between [0, 10^4]
*/
func Preorder(root *Node) []int {
	return preorder(root)
}
func preorder(root *Node) []int {
	var src []int
	preorderList(root, &src)
	return src
}

func preorderList(root *Node, src *[]int) {
	if root == nil {
		return
	}
	*src = append(*src, root.Val)
	for _, leaf := range root.Children {
		preorderList(leaf, src)
	}
}

/*
# https://leetcode.com/explore/learn/card/n-ary-tree/130/traversal/926/

Given an n-ary tree, return the postorder traversal of its nodes' values.
Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

Follow up:
Recursive solution is trivial, could you do it iteratively?

Example 1:

			1
		  / | \
		 3  2  4
		/ \
       5   6

Input: root = [1,null,3,2,4,null,5,6]
Output: [5,6,3,2,4,1]

Example 2:

				1
		   /  /   \  \
		  2  3     4  5
			/ \    |  | \
		   6   7   8  9  10
			   |   |  |
			   11  12 13
			   |
			   14

Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [2,6,14,11,7,3,12,8,4,13,9,10,5,1]


Constraints:

The height of the n-ary tree is less than or equal to 1000
The total number of nodes is between [0, 10^4]
*/
func Postorder(root *Node) []int {
	return postorder(root)
}

func postorder(root *Node) []int {
	var src []int
	postorderList(root, &src)
	return src
}

func postorderList(root *Node, src *[]int) {
	if root == nil {
		return
	}
	for _, leaf := range root.Children {
		postorderList(leaf, src)
	}
	*src = append(*src, root.Val)
}

/*
# https://leetcode.com/explore/learn/card/n-ary-tree/130/traversal/915/

Given an n-ary tree, return the level order traversal of its nodes' values.
Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

Example 1:

			1
		  / | \
		 3  2  4
		/ \
       5   6

Input: root = [1,null,3,2,4,null,5,6]
Output: [[1],[3,2,4],[5,6]]

Example 2:

				1
		   /  /   \  \
		  2  3     4  5
			/ \    |  | \
		   6   7   8  9  10
			   |   |  |
			   11  12 13
			   |
			   14

Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]

Constraints:

The height of the n-ary tree is less than or equal to 1000
The total number of nodes is between [0, 10^4]
*/
func LevelOrder(root *Node) [][]int {
	return levelOrder(root)
}

func levelOrder(root *Node) [][]int {
	return levelOrderList(root, nil, 0)
}

func levelOrderList(root *Node, list [][]int, level int) [][]int {
	if root == nil {
		return nil
	}
	if len(list) <= level {
		list = append(list, []int{root.Val})
	} else {
		list[level] = append(list[level], root.Val)
	}
	for _, leaf := range root.Children {
		list = levelOrderList(leaf, list, level+1)
	}
	return list
}

/*
# https://leetcode.com/explore/learn/card/n-ary-tree/131/recursion/919/

Given a n-ary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

Example 1:

			1
		  / | \
		 3  2  4
		/ \
       5   6

Input: root = [1,null,3,2,4,null,5,6]
Output: 3

Example 2:

				1
		   /  /   \  \
		  2  3     4  5
			/ \    |  | \
		   6   7   8  9  10
			   |   |  |
			   11  12 13
			   |
			   14

Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: 5


Constraints:
The depth of the n-ary tree is less than or equal to 1000.
The total number of nodes is between [0, 10^4].
*/

func MaxDepth(root *Node) int {
	return maxDepth(root)
}
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var max int
	for _, leaf := range root.Children {
		lDepth := maxDepth(leaf)
		if lDepth > max {
			max = lDepth
		}
	}
	return max + 1
}
