package node

// Node Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 根据前序遍历和中序遍历得出Node
func BuildTreePreIn(preorder []int, inorder []int) *Node {
	return buildTreePreIn(preorder, inorder)
}
func buildTreePreIn(preorder []int, inorder []int) *Node {
	if len(inorder) == 0 {
		return nil
	}
	var root *Node
	prePos := 0
	root = &Node{Val: preorder[prePos]}

	for inPos := 0; inPos < len(inorder); inPos++ {
		if preorder[prePos] == inorder[inPos] {
			// inorder中inPos左边是左子树
			root.Left = buildTreePreIn(preorder[prePos+1:inPos+1], inorder[:inPos])
			// inorder中inPos右边是右子树
			root.Right = buildTreePreIn(preorder[inPos+1:], inorder[inPos+1:])
			// 找到了就退出此次循环
			break
		}
	}
	return root
}

/*
# https://leetcode.com/explore/learn/card/data-structure-tree/133/conclusion/994/

You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.



Follow up:

You may only use constant extra space.
Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.


Example 1:
    1                      1  -> null
   / \                    / \
  2   3                  2 ->3  -> null
 / \ / \                / \ /  \
4  5 6  7              4->5->6->7  -> null

 Figure A		         Figure B


Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.


Constraints:

The number of nodes in the given tree is less than 4096.
-1000 <= node.val <= 1000
*/

func Connect(root *Node) *Node {
	return connect(root)
}
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectPerfectTree(root.Left, root.Right)
	return root
}

func connectPerfectTree(left, right *Node) {
	if left == nil && right == nil {
		return
	}
	if left != nil && right != nil {
		left.Next = right
	}
	connectPerfectTree(left.Left, left.Right)
	connectPerfectTree(right.Left, right.Right)
	connectPerfectTree(left.Right, right.Left)
}

/*
# https://leetcode.com/explore/learn/card/data-structure-tree/133/conclusion/1016/

Given a binary tree

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.



Follow up:

You may only use constant extra space.
Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.


Example 1:

    1                      1  -> null
   / \                    / \
  2   3                  2 ->3  -> null
 / \   \                / \   \
4   5   7              4-> 5 ->7  -> null

 Figure A		         Figure B


Input: root = [1,2,3,4,5,null,7]
Output: [1,#,2,3,#,4,5,7,#]
Explanation: Given the above binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.


Constraints:

The number of nodes in the given tree is less than 6000.
-100 <= node.val <= 100
*/
func Connect2(root *Node) *Node {
	return connect2(root)
}
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	src := connectNotPerfectTree(root, nil, 0)
	for _, level := range src {
		if len(level) > 1 {
			for i := 0; i < len(level)-1; i++ {
				level[i].Next = level[i+1]
			}
		}
	}
	return root
}

func connectNotPerfectTree(root *Node, src [][]*Node, level int) [][]*Node {
	if root == nil {
		return src
	}
	if len(src) > level {
		src[level] = append(src[level], root)
	} else {
		src = append(src, []*Node{root})
	}
	src = connectNotPerfectTree(root.Left, src, level+1)
	src = connectNotPerfectTree(root.Right, src, level+1)
	return src
}
