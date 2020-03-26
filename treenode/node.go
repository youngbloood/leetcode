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
