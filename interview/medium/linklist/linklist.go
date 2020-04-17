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

// 反转list
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
	odd := head       // 奇
	even := head.Next // 偶
	oddEnd, _ := oddEvenListSplit(head, 1, odd, even)
	oddEnd.Next = even
	return odd
}

// 根据index分割奇偶
func oddEvenListSplit(head *ListNode, index int, odd, even *ListNode) (oddEnd, evenEnd *ListNode) {
	if head == nil {
		return odd, even
	}
	next := head.Next
	if !(index == 1 || index == 2) {
		if index%2 != 0 {
			// odd
			odd.Next = head
			odd = odd.Next
			if next == nil {
				even.Next = nil
			}
		} else {
			// even
			even.Next = head
			even = even.Next
			if next == nil {
				odd.Next = nil
			}
		}
	}
	return oddEvenListSplit(next, index+1, odd, even)
}

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/785/

Write a program to find the node at which the intersection of two singly linked lists begins.

For example, the following two linked lists:

	A:	   	  a1 -> a2
					   \
					   	c1 -> c2 -> c3
					   /
	B:  b1 -> b2 -> b3


begin to intersect at node c1.

Example 1:

	A:	   	    4 -> 1
					   \
					   	8 -> 4 -> 5
					   /
	B:     5 -> 0 -> 1


Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
Output: Reference of the node with value = 8
Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,0,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.


Example 2:

	A:	   	0 -> 9 -> 1
					   \
					   	2 -> 4
					   /
	B:                3

Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Reference of the node with value = 2
Input Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [0,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.


Example 3:

	A:	   	2 -> 6 -> 4

	B:           1 -> 5


Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: null
Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.


Notes:

If the two linked lists have no intersection at all, return null.
The linked lists must retain their original structure after the function returns.
You may assume there are no cycles anywhere in the entire linked structure.
Your code should preferably run in O(n) time and use only O(1) memory.
*/
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	return getIntersectionNode(headA, headB)
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	var arrA, arrB []int
	ptr := headA
	for ptr != nil {
		arrA = append(arrA, ptr.Val)
		ptr = ptr.Next
	}
	ptr = headB
	for ptr != nil {
		arrB = append(arrB, ptr.Val)
		ptr = ptr.Next
	}

	ptrA, ptrB := len(arrA)-1, len(arrB)-1
	if arrA[ptrA] != arrB[ptrB] {
		return nil
	}
	for ptrA >= 0 && ptrB >= 0 {
		if arrA[ptrA] == arrB[ptrB] {
			ptrA--
			ptrB--
		} else {
			break
		}
	}

	times := 0
	if len(arrA) > len(arrB) {
		ptr = headB
		times = ptrB + 1
	} else {
		ptr = headA
		times = ptrA + 1
	}

	for ptr != nil && times > 0 {
		times--
		ptr = ptr.Next
	}

	return ptr
}
