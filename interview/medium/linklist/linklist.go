package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/783/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1ptr, l2ptr := l1, l2
	var nextAdd bool
	var quotient int
	for l1ptr != nil {
		// 将l2ptr的值往l2ptr中相加
		if l2ptr != nil {
			l1ptr.Val += l2ptr.Val
		}
		// 如果上一位有进位，则加商
		if nextAdd {
			l1ptr.Val += quotient
		}
		quotient = l1ptr.Val / 10
		if quotient >= 1 {
			// 下一位需要进位
			nextAdd = true
		} else {
			// 下一位不需要进位
			nextAdd = false
		}

		// l1ptr取余
		l1ptr.Val %= 10

		// l2比l1长，将l1ptr.Next指向l2ptr.Next，将l2ptr置空
		if l1ptr.Next == nil && l2ptr != nil && l2ptr.Next != nil {
			l1ptr.Next = l2ptr.Next
			l2ptr = nil
		}
		// 如果l1ptr到达末尾，但还是有进位，则增加一个节点，break
		if l1ptr.Next == nil && quotient >= 1 {
			l1ptr.Next = &ListNode{Val: quotient}
			break
		}
		// l1ptr移动
		l1ptr = l1ptr.Next
		if l2ptr != nil {
			// l2ptr移动
			l2ptr = l2ptr.Next
		}
	}
	return l1
}

func reverseList(list *ListNode) *ListNode {

	pre := list
	host := list.Next
	var index int
	for host != nil {
		temp := host.Next
		if index == 0 {
			pre.Next = nil
		}
		host.Next = pre
		pre = host
		host = temp
		index++
	}
	return pre
}

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/784/

Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

Example 1:
Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL

Example 2:
Input: 2->1->3->5->6->4->7->NULL
Output: 2->3->6->7->1->5->4->NULL
Note:

The relative order inside both the even and odd groups should remain as it was in the input.
The first node is considered odd, the second node even and so on ...
*/
func OddEvenList(head *ListNode) *ListNode {
	return oddEvenList(head)
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	even := head
	var odd *ListNode
	oddEvenListCycle(head, 1, odd, even)
}

func oddEvenListCycle(head *ListNode, index int, odd, even *ListNode) {
	if head == nil {
		return
	}
	if index/2 == 0 {
		// odd
		if odd == nil {
			odd == head
		} else {
			odd.Next = head
			odd = odd.Next
		}
	} else {
		// even
		if even == nil {
			even == head
		} else {
			even.Next = head
			even = even.Next
		}
	}
	oddEvenListCycle(head.Next, index+1, odd, even)
}
