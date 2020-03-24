package listnode_test

import (
	node "leetcode/listnode"
	"testing"
)

func initNodeList(vals ...int) *node.ListNode {

	if len(vals) == 0 {
		return nil
	}
	head := &node.ListNode{
		Val: vals[0],
	}
	next := head
	for i := 1; i < len(vals); i++ {
		next.Insert(vals[i])
		next = next.Next
	}
	return head
}

func TestNodeSwap(t *testing.T) {
	list := initNodeList(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// list = &node.ListNode{
	// 	Val: 1,
	// 	Next: &node.ListNode{
	// 		Val: 2,
	// 		Next: &node.ListNode{
	// 			Val: 3,
	// 			Next: &node.ListNode{
	// 				Val: 4,
	// 			},
	// 		},
	// 	},
	// }
	list.Println()
	result := node.SwapPairs(list)
	result.Println()
}

func TestListNodeReverse(t *testing.T) {
	list := initNodeList(1, 2, 3, 4, 5, 6)
	list.Println()
	result := node.ReverseList(list)
	result.Println()
}

func TestMergeTwoLists(t *testing.T) {
	l1 := initNodeList(2)
	l2 := initNodeList(1)
	r := node.MergeTwoLists(l1, l2)
	r.Println()
}
