package linklist_test

import (
	"leetcode/interview/hard/linklist"
	"testing"
)

func initNodeList(vals ...int) *linklist.ListNode {

	if len(vals) == 0 {
		return nil
	}
	head := &linklist.ListNode{
		Val: vals[0],
	}
	next := head
	for i := 1; i < len(vals); i++ {
		next.Insert(vals[i])
		next = next.Next
	}
	return head
}
func TestMergeKLists(t *testing.T) {
	list := initNodeList(4, 2, 1, 3)
	t.Log(linklist.SortList(list))
	list = initNodeList(-1, 5, 3, 4, 0)
	t.Log(linklist.SortList(list))
	list = initNodeList()
	t.Log(linklist.SortList(list))
}

func BenchmarkMergeKLists(b *testing.B) {

	list := initNodeList(4, 2, 1, 3)
	for i := 0; i < b.N; i++ {
		linklist.SortList(list)
	}

	list = initNodeList(-1, 5, 3, 4, 0)
	for i := 0; i < b.N; i++ {
		linklist.SortList(list)
	}
	list = initNodeList()
	for i := 0; i < b.N; i++ {
		linklist.SortList(list)
	}
}

func initRandomNodeList(nodes ...[2]int) *linklist.Node {
	if nodes == nil || len(nodes) == 0 {
		return nil
	}
	root := &linklist.Node{Val: nodes[0][0]}
	list := []*linklist.Node{root}
	ptr := root
	for i := 1; i < len(nodes); i++ {
		ptr.Next = &linklist.Node{Val: nodes[i][0]}
		list = append(list, ptr.Next)
		ptr = ptr.Next
	}
	ptr = root
	var i int
	for ptr != nil {
		if nodes[i][1] == -1 {
			ptr = ptr.Next
			i++
			continue
		}
		ptr.Random = list[nodes[i][1]]
		ptr = ptr.Next
		i++
	}
	return root
}

func TestCopyRandomList(t *testing.T) {
	// [[7,null],[13,0],[11,4],[10,2],[1,0]]
	node := initRandomNodeList([2]int{7, -1},
		[2]int{13, 0},
		[2]int{11, 4},
		[2]int{10, 2},
		[2]int{1, 0})
	val := linklist.CopyRandomList(node)

	t.Log(val)
}
