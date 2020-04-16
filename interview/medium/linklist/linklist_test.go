package linklist_test

import (
	"leetcode/interview/medium/linklist"
	"testing"
)

func initLinkList(vals ...int) *linklist.ListNode {
	if len(vals) == 0 {
		return nil
	}

	root := &linklist.ListNode{Val: vals[0]}
	ptr := root
	for i := 1; i < len(vals); i++ {
		ptr.Next = &linklist.ListNode{Val: vals[i]}
		ptr = ptr.Next
	}
	return root
}
func TestAddTwoNumbers(t *testing.T) {

	l1 := initLinkList(2, 4, 3)
	l2 := initLinkList(5, 6, 4)

	l1 = linklist.AddTwoNumbers(l1, l2)
	t.Log(l1)
	t.Log(l2)

	l1 = initLinkList(1, 8)
	l2 = initLinkList(0)

	l1 = linklist.AddTwoNumbers(l1, l2)
	t.Log(l1)
	t.Log(l2)

	/*
		Output:
		[9,9]
		Expected:
		[0,9]
	*/
	l1 = initLinkList(9, 8)
	l2 = initLinkList(1)

	l1 = linklist.AddTwoNumbers(l1, l2)
	t.Log(l1)
	t.Log(l2)

	l1 = initLinkList(8, 9, 9)
	l2 = initLinkList(2)

	l1 = linklist.AddTwoNumbers(l1, l2)
	t.Log(l1)
	t.Log(l2)

	/*
			[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
		[5,6,4]
	*/
	l1 = initLinkList(1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1)
	l2 = initLinkList(5, 6, 4)

	l1 = linklist.AddTwoNumbers(l1, l2)
	t.Log(l1)
	t.Log(l2)

	l1 = initLinkList(0)
	l2 = initLinkList(7, 3)

	l1 = linklist.AddTwoNumbers(l1, l2)
	t.Log(l1)
	t.Log(l2)
}
