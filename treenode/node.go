package treenode

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode Definition for a binary tree node.
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

// GenerateTrees .
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

// InsertTreeBatch . 批量插入元素
func InsertTreeBatch(tree *TreeNode, vals ...int) {
	for _, v := range vals {
		InsertTree(tree, v)
	}
}

// InsertTree . 单个插入元素
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

// CopyTree . treenode的复制
func CopyTree(src *TreeNode) *TreeNode {
	if src == nil {
		return nil
	}
	dst := &TreeNode{Val: src.Val}
	dst.Right = CopyTree(src.Right)
	dst.Left = CopyTree(src.Left)
	return dst
}

// Copy . treenode的复制
func (tn *TreeNode) Copy() *TreeNode {
	if tn == nil {
		return nil
	}
	dst := &TreeNode{Val: tn.Val}
	dst.Right = tn.Right.Copy()
	dst.Left = tn.Left.Copy()
	return dst
}

// IsValidBST .
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

// IsSameTree .
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

// InorderTraversal .
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

// LevelOrder .
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

// PreorderTraversal .
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

// PostorderTraversal .
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

// IsSymmetric .
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

// HasPathSum .
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

// BuildTreeInPost .
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
func BuildTreeInPost(inorder []int, postorder []int) *TreeNode {
	return buildTreeInPost(inorder, postorder)
}

func buildTreeInPost(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	var root *TreeNode
	postPos := len(postorder) - 1
	root = &TreeNode{Val: postorder[postPos]}

	for inPos := 0; inPos < len(inorder); inPos++ {
		if postorder[postPos] == inorder[inPos] {
			// inorder中inPos左边是左子树
			root.Left = buildTreeInPost(inorder[:inPos], postorder[:inPos])
			// inorder中inPos右边是右子树
			root.Right = buildTreeInPost(inorder[inPos+1:], postorder[inPos:postPos])
			// 找到了就退出此次循环
			break
		}
	}
	return root
}

// BuildTreePreIn .
/*
# https://leetcode.com/explore/learn/card/data-structure-tree/133/conclusion/943/

Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/
func BuildTreePreIn(preorder []int, inorder []int) *TreeNode {
	return buildTreePreIn(preorder, inorder)
}
func buildTreePreIn(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	var root *TreeNode
	prePos := 0
	root = &TreeNode{Val: preorder[prePos]}

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

// LowestCommonAncestor .
/*
# https://leetcode.com/explore/learn/card/data-structure-tree/133/conclusion/932/

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]

					  3
					 / \
				   /     \
				  5       1
				 / \     / \
				6   2   0   8
				   / \
				  7   4

Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.
Example 2:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.


Note:

All of the nodes' values will be unique.
p and q are different and both values will exist in the binary tree.
*/
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestor(root, p, q)
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 查看p是否包含q
	if isContain := contain(p, q); isContain {
		return p
	}
	// 查看q是否包含p
	if isContain := contain(q, p); isContain {
		return q
	}

	// 寻找p的父节点集合
	pHead := p
	var pHeads []*TreeNode
	for pHead != root {
		pHead = findHead(root, root, pHead)
		pHeads = append(pHeads, pHead)
	}
	// 寻找q的父节点集合
	qHead := q
	var qHeads []*TreeNode
	for qHead != root {
		qHead = findHead(root, root, qHead)
		qHeads = append(qHeads, qHead)
	}

	// 倒序找出最小公共父节点
	var minLen int = len(pHeads)
	if len(pHeads) > len(qHeads) {
		minLen = len(qHeads)
	}

	for i := 0; i < minLen; i++ {
		if pHeads[len(pHeads)-1-i] != qHeads[len(qHeads)-1-i] {
			return pHeads[len(pHeads)-1-i+1]
		}
	}
	if len(pHeads) > len(qHeads) {
		return qHeads[0]
	}
	return pHeads[0]
}

func contain(p, q *TreeNode) bool {
	if p == nil {
		return false
	}
	if p.Val == q.Val {
		return true
	}
	return contain(p.Left, q) || contain(p.Right, q)
}

func findHead(parent, tree, target *TreeNode) *TreeNode {
	if !contain(tree, target) {
		return nil
	}
	if tree.Val == target.Val {
		return parent
	}
	if r := findHead(tree, tree.Left, target); r != nil {
		return r
	}
	return findHead(tree, tree.Right, target)
}

/*Leetcode上的一个优解*/
func findLCA(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	// 分别找到p,q在左右子树中的位置，为l,r
	l, r := findLCA(root.Left, p, q), findLCA(root.Right, p, q)
	// 分别位于两边子树
	if l != nil && r != nil {
		return root
	}
	// 均位于左子树
	if l != nil {
		return l
	}
	// 均位于右子树
	return r
}

/*
# https://leetcode.com/explore/learn/card/data-structure-tree/133/conclusion/995/

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Example:

You may serialize the following tree:

    1
   / \
  2   3
     / \
    4   5

as "[1,2,3,null,null,4,5]"
Clarification: The above format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

Note: Do not use class member/global/static variables to store states. Your serialize and deserialize algorithms should be stateless.
*/

// 思路：值（左子树（左子树，右子树），右子树（左子树，右子树））
/*
	1
   / \
  2   3
 / \ / \
4  5 6  7

// Serialize1: 1(2,3(4,5))
// Serialize2: 1(2(4,5),3(6,7))
*/

// Codec .
type Codec struct {
	nodeSplit  string // 节点分割符
	leftSplit  string // 与rightSplit配对
	rightSplit string // 与leftSplit配对
}

// Constructor .
func Constructor() Codec {
	return Codec{nodeSplit: ",", leftSplit: "(", rightSplit: ")"}
}

// Serialize . a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
	return this.serialize(root)
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	leftStr := this.serialize(root.Left)
	rightStr := this.serialize(root.Right)
	var nodeStr string
	switch {
	case leftStr == "" && rightStr == "":
		nodeStr = fmt.Sprintf("%d", root.Val)
	case leftStr != "" && rightStr == "":
		nodeStr = fmt.Sprintf("%d%s%s%s%s", root.Val, this.leftSplit, leftStr, this.nodeSplit, this.rightSplit)
	case leftStr == "" && rightStr != "":
		nodeStr = fmt.Sprintf("%d%s%s%s%s", root.Val, this.leftSplit, this.nodeSplit, rightStr, this.rightSplit)
	default:
		nodeStr = fmt.Sprintf("%d%s%s%s%s%s", root.Val, this.leftSplit, leftStr, this.nodeSplit, rightStr, this.rightSplit)
	}
	return nodeStr
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	return this.deserialize(data)
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	val, err := strconv.Atoi(data)
	if err == nil {
		return &TreeNode{Val: val}
	}
	// 寻找第一个leftSplit
	firstSplit := strings.Index(data, this.leftSplit)
	if firstSplit < 0 {
		return nil
	}
	subStr := data[firstSplit+1 : len(data)-1]

	// 寻找分割点povit
	var povit, splitCount int
	for i, v := range subStr {
		if string(v) == this.nodeSplit && splitCount == 0 {
			povit = i
			break
		}
		if string(v) == this.leftSplit {
			splitCount++
			continue
		}
		if string(v) == this.rightSplit {
			splitCount--
			if splitCount == 0 {
				if i == len(subStr)-1 {
					povit = i
				} else {
					povit = i + 1
				}
				break
			}
		}
	}
	// 提取出head值，并解析左子树和右子树
	val, err = strconv.Atoi(data[:firstSplit])
	if err != nil {
		panic("not standard string")
	}
	root := &TreeNode{Val: val}
	leftNode, rightNode := subStr[:povit], subStr[povit+1:]
	root.Left = this.deserialize(leftNode)
	root.Right = this.deserialize(rightNode)
	return root
}
