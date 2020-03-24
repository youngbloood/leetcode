package listnode

import (
	"fmt"
	"strconv"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Insert(val int) {
	head.Next = &ListNode{
		Val: val,
	}
}

func (head *ListNode) Println() {
	var str string
	for head != nil {
		str += " " + strconv.Itoa(head.Val)
		head = head.Next
	}
	fmt.Println(str)
}

/*
Given a linked list, swap every two adjacent nodes and return its head.
You may not modify the values in the list's nodes, only nodes itself may be changed.

Example:
Given 1->2->3->4, you should return the list as 2->1->4->3.
*/
func SwapPairs(head *ListNode) *ListNode {
	var tail, pre, next *ListNode
	pre = head
	next = head
	var i int
	for next != nil {
		if i%2 != 0 {
			pre.Next = next.Next
			next.Next = pre
			if i == 1 {
				head = next
			}
			if tail != nil {
				tail.Next = next
			}
			next = pre
			pre = pre.Next
			tail = next
		}
		next = next.Next
		i++
	}
	return head
}

/*
Reverse a singly linked list.

Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
*/
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var first, pre, next *ListNode
	first = head
	pre = head
	next = head.Next
	for next != nil {
		if next.Next == nil {
			head = next
			next.Next = pre
			break
		}
		temp := pre
		// 向链尾移动
		pre = next
		next = next.Next
		// 反指
		pre.Next = temp
	}
	first.Next = nil
	return head
}

/**
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// l1存head小的那个listnode
	if l1.Val >= l2.Val {
		l1, l2 = l2, l1
	}
	head := l1
	var pre, next, l2ptr *ListNode
	pre = l1
	next = l1.Next
	l2ptr = l2

	for l2ptr != nil {
		if next == nil {
			pre.Next = l2ptr
			break
		}
		// 比较pre和next的值，在此区间内插入
		if l2ptr.Val >= pre.Val && l2ptr.Val <= next.Val {
			pre.Next = l2ptr
			l2ptrTemp := l2ptr.Next
			l2ptr.Next = next
			// pre向后移动
			pre = l2ptr
			l2ptr = l2ptrTemp
		} else {
			next = next.Next
			pre = pre.Next
		}
	}

	return head
}
