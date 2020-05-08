package linklist

import (
	"sort"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Insert(val int) {
	head.Next = &ListNode{
		Val: val,
	}
}

/*
# Merge k Sorted Lists
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/117/linked-list/839/

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/
func MergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists(lists)
}
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	var result []int
	for _, list := range lists {
		for list != nil {
			result = append(result, list.Val)
			list = list.Next
		}
	}

	if len(result) == 0 {
		return nil
	}
	sort.Ints(result)
	head := &ListNode{Val: result[0]}
	root := head
	for i := 1; i < len(result); i++ {
		root.Next = &ListNode{Val: result[i]}
		root = root.Next
	}
	return head
}

/*
# Sort List
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/117/linked-list/840/

Sort a linked list in O(n log n) time using constant space complexity.

Example 1:
Input: 4->2->1->3
Output: 1->2->3->4

Example 2:
Input: -1->5->3->4->0
Output: -1->0->3->4->5
*/
func SortList(head *ListNode) *ListNode {
	return sortList2(head)
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	root := head
	rootPtr := head
	leftPtr := head.Next
	root.Next = nil

	for leftPtr != nil {
		// 循环往root中插入
		pre := root
		for {
			if rootPtr != nil && leftPtr.Val < rootPtr.Val {
				// 插入头部
				if rootPtr == root {
					move := leftPtr
					leftPtr = leftPtr.Next
					move.Next = root
					root = move
					rootPtr = root
				} else {
					// 插入pre和rootPtr之间
					move := leftPtr
					leftPtr = leftPtr.Next
					pre.Next = move
					move.Next = rootPtr
					rootPtr = root
				}
				break
			} else if rootPtr == nil {
				// 插入root的尾部
				move := leftPtr
				leftPtr = leftPtr.Next
				move.Next = nil
				pre.Next = move
				rootPtr = root
				break
			} else {
				pre = rootPtr
				rootPtr = rootPtr.Next
			}
		}
	}
	return root
}

// leetcode优解：采用分治法
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	l := sortList2(head)
	r := sortList2(mid)
	dummy := &ListNode{}
	n := dummy

	// 合并两个有序链表
	for r != nil && l != nil {
		if r.Val <= l.Val {
			n.Next = r
			r = r.Next
		} else {
			n.Next = l
			l = l.Next
		}
		n = n.Next
	}
	if r != nil {
		n.Next = r
	}
	if l != nil {
		n.Next = l
	}
	return dummy.Next
}

/*
# Copy List with Random Pointer
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/117/linked-list/841/

A linked list is given such that each node contains an additional random pointer which could point to any node in the list or null.

Return a deep copy of the list.

The Linked List is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

val: an integer representing Node.val
random_index: the index of the node (range from 0 to n-1) where random pointer points to, or null if it does not point to any node.


Example 1:


Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
Example 2:


Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]
Example 3:



Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]
Example 4:

Input: head = []
Output: []
Explanation: Given linked list is empty (null pointer), so return null.


Constraints:

-10000 <= Node.val <= 10000
Node.random is null or pointing to a node in the linked list.
Number of Nodes will not exceed 1000.

Hint #1
Just iterate the linked list and create copies of the nodes on the go. Since a node can be referenced from multiple nodes due to the random pointers, make sure you are not making multiple copies of the same node.

Hint #2
You may want to use extra space to keep old node ---> new node mapping to prevent creating multiples copies of same node.

Hint #3
We can avoid using extra space for old node ---> new node mapping, by tweaking the original linked list. Simply interweave the nodes of the old and copied list. For e.g.
Old List: A --> B --> C --> D
InterWeaved List: A --> A' --> B --> B' --> C --> C' --> D --> D'

Hint #4
The interweaving is done using next pointers and we can make use of interweaved structure to get the correct reference nodes for random pointers.
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	return copyRandomList(head)
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	headPtr := head
	// InterWeaved List
	for headPtr != nil {
		newNode := &Node{Val: headPtr.Val}
		newNode.Next = headPtr.Next
		headPtr.Next = newNode

		headPtr = headPtr.Next.Next
	}

	// Init InterWeaved List's Random
	pre, next := head, head.Next
	for next != nil {
		if pre.Random != nil {
			next.Random = pre.Random.Next
		}

		pre = pre.Next.Next
		if next.Next == nil {
			break
		}
		next = next.Next.Next
	}

	var result, resultPtr *Node
	// Extract InterWeaved List &  Remove InterWeaved List
	pre = head
	next = head.Next
	for next != nil {
		if result == nil {
			result = next
			resultPtr = result
		} else {
			resultPtr.Next = next
			resultPtr = resultPtr.Next
		}
		if next.Next == nil {
			break
		}
		next = next.Next.Next
		pre.Next = pre.Next.Next
		pre = pre.Next
	}
	// remove lasted node
	pre.Next = nil
	return result
}
