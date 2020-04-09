package treenode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // 中序遍历
// func MiddleOrder(tn *TreeNode) string {
// 	if tn.Left != nil {
// 		MiddleOrder(tn.Left)
// 	}

// }

/*
Given the root node of a binary search tree (BST) and a value. You need to find the node in the BST that the node's value equals the given value. Return the subtree rooted with that node. If such node doesn't exist, you should return NULL.

For example：
Given the tree:
        4
       / \
      2   7
     / \
    1   3

And the value to search: 2
You should return this subtree:

      2
     / \
    1   3
In the example above, if we want to search the value 5, since there is no node with value 5, we should return NULL.

Note that an empty tree is represented by NULL, therefore you would see the expected output (serialized tree format) as [], not null.
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		if root.Left == nil {
			return nil
		}
		return searchBST(root.Left, val)
	}
	if root.Right == nil {
		return nil
	}
	return searchBST(root.Right, val)
}

/*
# https://leetcode.com/explore/learn/card/data-structure-tree/17/solve-problems-recursively/535/

Given a binary tree, find its maximum depth.
The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
Note: A leaf is a node with no children.

Example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	ldepth := maxDepth(root.Left)
	rdepth := maxDepth(root.Right)
	if ldepth > rdepth {
		return ldepth + 1
	}
	return rdepth + 1
}

/*
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

Example:
Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
// 插入顺序
// 1,3,2
// 3,2,1
// 3,1,2
// 2,1,3
// 2,3,1   和2,1,3结果一样
// 1,2,3

*/
func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{&TreeNode{Val: 1}}
	}

	var allTree []*TreeNode
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = i + 1
	}

	for i := 0; i < n; i++ {
		head := &TreeNode{Val: vals[i]}
		allTree = insertTreeArr(allTree, head, vals[:i], vals[i+1:])
	}
	return allTree
}

func insertTreeArr(allTree []*TreeNode, head *TreeNode, left, right []int) []*TreeNode {
	if head == nil {
		return allTree
	}
	if len(left) == 0 && len(right) == 0 {
		return allTree
	}
	if len(left) <= 1 && len(right) <= 1 {
		InsertTreeBatch(head, left...)
		InsertTreeBatch(head, right...)
		allTree = append(allTree, head)
		return allTree
	}

	headTemp := CopyTree(head)
	for i, v := range left {
		headLeft := CopyTree(headTemp)
		InsertTree(headLeft, v)
		allTree = insertTreeArr(allTree, headLeft, left[:i], left[i+1:])
	}
	head = headTemp
	for i, v := range right {
		headRight := CopyTree(head)
		InsertTree(headRight, v)
		allTree = insertTreeArr(allTree, headRight, right[:i], right[i+1:])
	}
	return allTree
}

// 批量插入元素
func InsertTreeBatch(tree *TreeNode, vals ...int) {
	for _, v := range vals {
		InsertTree(tree, v)
	}
}

// 单个插入元素
func InsertTree(tree *TreeNode, value int) *TreeNode {
	if tree == nil {
		return tree
	}
	for tree != nil {
		switch {
		case tree.Val > value:
			if tree.Left == nil {
				tree.Left = &TreeNode{
					Val: value,
				}
				return tree.Left
			}
			tree = tree.Left

		case tree.Val < value:
			if tree.Right == nil {
				tree.Right = &TreeNode{
					Val: value,
				}
				return tree.Right
			}
			tree = tree.Right
		default:
			// 相等
			return tree
		}
	}
	// 表示已经存在该值
	return nil
}

// treenode的复制
func CopyTree(src *TreeNode) *TreeNode {
	if src == nil {
		return nil
	}
	dst := &TreeNode{Val: src.Val}
	dst.Right = CopyTree(src.Right)
	dst.Left = CopyTree(src.Left)
	return dst
}

// treenode的复制
func (tn *TreeNode) Copy() *TreeNode {
	if tn == nil {
		return nil
	}
	dst := &TreeNode{Val: tn.Val}
	dst.Right = tn.Right.Copy()
	dst.Left = tn.Left.Copy()
	return dst
}

func IsValidBST(root *TreeNode) bool {
	return isValidBST(root)
}

func isValidBST(root *TreeNode) bool {
	list := inOrderTree(root, nil)
	if len(list) <= 1 {
		return true
	}
	for i := 1; i < len(list); i++ {
		if list[i-1].Val >= list[i].Val {
			return false
		}
	}
	return true
}

// 中序遍历root
func inOrderTree(root *TreeNode, list []*TreeNode) []*TreeNode {
	if root == nil {
		return list
	}
	list = inOrderTree(root.Left, list)
	list = append(list, root)
	list = inOrderTree(root.Right, list)
	return list
}

// leetcode最优解
func isValidBST2(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root, l, r *TreeNode) bool {
	if root == nil {
		return true
	}
	left, right := true, true
	if l != nil {
		left = l.Val < root.Val
	}
	if r != nil {
		right = root.Val < r.Val
	}
	return left && right && isValid(root.Left, l, root) && isValid(root.Right, root, r)
}

/**
Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Example 1:
Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true

Example 2:
Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false

Example 3:
Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false
*/

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTree(p, q)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil && p.Val != q.Val {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

/*
# https://leetcode.com/explore/learn/card/recursion-ii/503/recursion-to-iteration/2774/
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

func InorderTraversal(root *TreeNode) []int {
	return inorderTraversal(root)
}

func inorderTraversal(root *TreeNode) []int {
	return inorderTree(root, nil)
}

func inorderTree(root *TreeNode, vals []int) []int {
	if root == nil {
		return vals
	}
	vals = inorderTree(root.Left, vals)
	vals = append(vals, root.Val)
	vals = inorderTree(root.Right, vals)
	return vals
}

/*
# https://leetcode.com/explore/learn/card/recursion-ii/503/recursion-to-iteration/2784/

Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/
func LevelOrder(root *TreeNode) [][]int {
	return levelOrder(root)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	vals := [][]int{[]int{root.Val}}
	return levelOrderTree(root.Left, root.Right, vals, 1)
}

func levelOrderTree(left, right *TreeNode, vals [][]int, level int) [][]int {
	if left == nil && right == nil {
		return vals
	}
	if level >= len(vals) {
		vals = append(vals, []int{})
	}
	if left != nil {
		vals[level] = append(vals[level], left.Val)
		vals = levelOrderTree(left.Left, left.Right, vals, level+1)
	}
	if right != nil {
		vals[level] = append(vals[level], right.Val)
		vals = levelOrderTree(right.Left, right.Right, vals, level+1)
	}
	return vals
}

/*
# https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/928/

Given a binary tree, return the preorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/
func PreorderTraversal(root *TreeNode) []int {
	return preorderTraversal(root)
}

func preorderTraversal(root *TreeNode) []int {
	return preorderTree(root, nil)
}

func preorderTree(root *TreeNode, src []int) []int {
	if root == nil {
		return src
	}
	src = append(src, root.Val)
	src = preorderTree(root.Left, src)
	src = preorderTree(root.Right, src)
	return src
}

/*
# https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/930/

Given a binary tree, return the postorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/
func PostorderTraversal(root *TreeNode) []int {
	return postorderTraversal(root)
}
func postorderTraversal(root *TreeNode) []int {
	return postorderTree(root, nil)
}

func postorderTree(root *TreeNode, src []int) []int {
	if root == nil {
		return src
	}
	src = postorderTree(root.Left, src)
	src = postorderTree(root.Right, src)
	src = append(src, root.Val)
	return src
}

/*
# https://leetcode.com/explore/learn/card/data-structure-tree/17/solve-problems-recursively/536/

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3


Follow up: Solve it both recursively and iteratively.
*/
func IsSymmetric(root *TreeNode) bool {
	return IsSymmetric(root)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil && p.Val != q.Val {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	return isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
}

/*
# https://leetcode.com/explore/learn/card/data-structure-tree/17/solve-problems-recursively/537/

Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Example:
Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
*/
func HasPathSum(root *TreeNode, sum int) bool {
	return hasPathSum(root, sum)
}

func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSumTree(root, sum, 0)
}

func hasPathSumTree(root *TreeNode, sumTarget, sumNow int) bool {
	if root == nil {
		return false
	}
	sumNow += root.Val
	// 必须是叶子节点
	if sumNow == sumTarget && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSumTree(root.Left, sumTarget, sumNow) || hasPathSumTree(root.Right, sumTarget, sumNow)
}

/*
# https://leetcode.com/explore/learn/card/data-structure-tree/133/conclusion/942/

Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/
func BuildTree(inorder []int, postorder []int) *TreeNode {
	return buildTree(inorder, postorder)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return buildTreeWithPos(inorder, postorder, 0, len(inorder)-1)
}

func buildTreeWithPos(inorder, postorder []int, startInOrder, endInOrder int) *TreeNode {
	var root *TreeNode
	for postPos := len(postorder) - 1; postPos >= 0; postPos-- {
		if root == nil {
			root = &TreeNode{Val: postorder[postPos]}
		}
		for inPos := startInOrder; inPos < endInOrder; inPos++ {
			if postorder[postPos] == inorder[inPos] {
				// inorder中inPos左边是左子树
				root.Left = buildTreeWithPos(inorder, postorder[:postPos], startInOrder, inPos)
				// inorder中inPos右边是右子树
				root.Right = buildTreeWithPos(inorder, postorder[:postPos], inPos, endInOrder)
			}
		}
	}
	return root
}

func insertLeft(head *TreeNode, val int) {
	if head.Left == nil {
		head.Left = &TreeNode{
			Val: val,
		}
	}
}

func insertRight(head *TreeNode, val int) {
	if head.Right == nil {
		head.Right = &TreeNode{
			Val: val,
		}
	}
}
