package treenode

/*
# https://leetcode.com/explore/learn/card/introduction-to-data-structure-binary-search-tree/140/introduction-to-a-bst/1008/

Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

Calling next() will return the next smallest number in the BST.


Example:

		7
	   / \
	  3   15
	      / \
         9   20

BSTIterator iterator = new BSTIterator(root);
iterator.next();    // return 3
iterator.next();    // return 7
iterator.hasNext(); // return true
iterator.next();    // return 9
iterator.hasNext(); // return true
iterator.next();    // return 15
iterator.hasNext(); // return true
iterator.next();    // return 20
iterator.hasNext(); // return false


Note:

next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
You may assume that next() call will always be valid, that is, there will be at least a next smallest number in the BST when next() is called.
*/
type BSTIterator struct {
	iter int
	list []*TreeNode
}

func BSTConstructor(root *TreeNode) BSTIterator {
	var list []*TreeNode
	initBSTIterator(root, &list)
	return BSTIterator{
		list: list,
	}
}

func initBSTIterator(root *TreeNode, list *[]*TreeNode) {
	if root == nil {
		return
	}
	initBSTIterator(root.Left, list)
	*list = append(*list, root)
	initBSTIterator(root.Right, list)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	val := this.list[this.iter].Val
	this.iter++
	return val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.iter == len(this.list) {
		return false
	}
	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

/*
# https://leetcode.com/explore/learn/card/introduction-to-data-structure-binary-search-tree/141/basic-operations-in-a-bst/1003/

Given the root node of a binary search tree (BST) and a value to be inserted into the tree, insert the value into the BST. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Note that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

For example,

Given the tree:
        4
       / \
      2   7
     / \
    1   3
And the value to insert: 5
You can return this binary search tree:

         4
       /   \
      2     7
     / \   /
    1   3 5
This tree is also valid:

         5
       /   \
      2     7
     / \
    1   3
         \
          4
*/
func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	return insertIntoBST(root, val)
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	head := root
	for root != nil {
		switch {
		case root.Val > val:
			if root.Left == nil {
				root.Left = &TreeNode{
					Val: val,
				}
				return head
			}
			root = root.Left
		case root.Val < val:
			if root.Right == nil {
				root.Right = &TreeNode{
					Val: val,
				}
				return head
			}
			root = root.Right
		default:
			// 相等
			return head
		}
	}
	// 表示已经存在该值
	return head
}

/*
# https://leetcode.com/explore/learn/card/introduction-to-data-structure-binary-search-tree/141/basic-operations-in-a-bst/1006/

Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.
Note: Time complexity should be O(height of tree).

Example:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

Given key to delete is 3. So we find the node with value 3 and delete it.

One valid answer is [5,4,6,2,null,null,7], shown in the following BST.

    5
   / \
  4   6
 /     \
2       7

Another valid answer is [5,2,6,null,4,null,7].

    5
   / \
  2   6
   \   \
    4   7
*/
func DeleteNode(root *TreeNode, key int) *TreeNode {
	return deleteNode2(root, key)
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	head := root
	parent := root

	for root != nil {
		if root.Val > key {
			parent = root
			root = root.Left
		} else if root.Val < key {
			parent = root
			root = root.Right
		} else {
			// root==key
			if root.Left == nil && root.Right == nil {
				// root左右子节点均为空
				if parent.Left != nil && parent.Left.Val == key {
					parent.Left = nil
				} else if parent.Right != nil && parent.Right.Val == key {
					parent.Right = nil
				} else if parent == head {
					return nil
				}
			} else if root.Left != nil && root.Right != nil {
				// root左右子节点均不为空
				rightLeft := root.Right
				rightParent := root
				for !(rightLeft.Left == nil && rightLeft.Right == nil) {
					if rightLeft.Left != nil {
						rightParent = rightLeft
						rightLeft = rightLeft.Left
					} else if rightLeft.Right != nil && rightLeft != root.Right {
						rightParent = rightLeft
						rightLeft = rightLeft.Right
					} else {
						break
					}
				}
				root.Val = rightLeft.Val
				if rightParent == root {
					root.Right = nil

				} else {
					if rightParent.Left != nil && rightParent.Left.Val == root.Val {
						rightParent.Left = nil
					} else {
						rightParent.Right = nil
					}
				}
			} else {
				// root左右子节点一个为空
				if root == head {
					if head.Left != nil {
						head = head.Left
					} else {
						head = head.Right
					}
				} else if parent.Left.Val == key {
					// root位于parent的左节点
					if root.Left != nil {
						parent.Left = root.Left
					} else {
						parent.Left = root.Right
					}
				} else {
					// root位于parent的右节点
					if root.Left != nil {
						parent.Right = root.Left
					} else {
						parent.Right = root.Right
					}
				}

			}
			break
		}
	}
	return head
}

func deleteNode2(root *TreeNode, key int) *TreeNode {

	head := root
	keyParent := root
	keyNode := root
	for keyNode != nil {
		if keyNode.Val > key {
			keyParent = keyNode
			keyNode = keyNode.Left
		} else if keyNode.Val < key {
			keyParent = keyNode
			keyNode = keyNode.Right
		} else {
			if keyNode.Left == nil && keyNode.Right == nil {
				if keyParent.Left != nil && keyParent.Left.Val == key {
					keyParent.Left = nil
				} else if keyParent.Right != nil && keyParent.Right.Val == key {
					keyParent.Right = nil
				} else if keyParent == head {
					return nil
				}
			} else if keyNode.Left != nil && keyNode.Right != nil {
				if keyParent == head {

				}

			} else {

			}

			break
		}
	}
	return head
}

/*
# https://leetcode.com/explore/learn/card/introduction-to-data-structure-binary-search-tree/142/conclusion/1018/

Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream. For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.

Example:

int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
Note:
You may assume that nums' length ≥ k-1 and k ≥ 1.
*/

type KthLargest struct {
	head *TreeNode
	k    int
}

func ConstructorKthLargest(k int, nums []int) KthLargest {
	var head *TreeNode
	if nums != nil && len(nums) != 0 {
		head = &TreeNode{Val: nums[0]}
	}
	InsertTreeBatch(head, nums[1:]...)
	return KthLargest{head: head, k: k}
}

func (this *KthLargest) Add(val int) int {
	InsertTree(this.head, val)
	// 寻找第k个最大值
	return 0
}

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/790/

Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note:
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Example 1:
Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1

Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
Output: 3
Follow up:
What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?

Hint #1:
Try to utilize the property of a BST.

Hint #2:
Try in-order traversal. (Credits to @chan13)

Hint #3:
What if you could modify the BST node's structure?

Hint #4:
The optimal runtime complexity is O(height of BST).
*/
func KthSmallest(root *TreeNode, k int) int {
	return kthSmallest2(root, k)
}

func kthSmallest(root *TreeNode, k int) int {
	list := inorderTraversal(root)
	if len(list) < k {
		return -1
	}
	return list[k-1]
}

// The optimal runtime complexity is O(height of BST).
func kthSmallest2(root *TreeNode, k int) int {
	leftCount, val := kthSmallest2Count(root.Left, k)
	if val != nil {
		return val.Val
	}
	rightCount, val := kthSmallest2Count(root.Right, k-leftCount-1)
	if val != nil {
		return val.Val
	}
	if k > rightCount+leftCount {
		return -1
	}
	return root.Val
}

func kthSmallest2Count(root *TreeNode, k int) (count int, val *TreeNode) {
	if root == nil {
		return 0, nil
	}
	leftCount, val := kthSmallest2Count(root.Left, k)
	if val != nil {
		return 0, val
	}
	if k < leftCount {
		// 找到了
		val = root
		return
	}
	rightCount, val := kthSmallest2Count(root.Right, k-leftCount-1)
	if val != nil {
		return 0, val
	}
	count = leftCount + rightCount + 1
	if k <= count && k != 0 {
		// 找到了
		val = root
	}
	return
}
