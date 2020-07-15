// # https://leetcode.com/problems/design-skiplist/
package skiplist

import (
	"math/rand"
	"time"
)

type Skiplist struct {
	val  int
	next []*Skiplist
	r    *rand.Rand
}

func Constructor() Skiplist {
	var s Skiplist
	s1 := rand.NewSource(time.Now().UnixNano())
	s.r = rand.New(s1)
	return s
}

func (this *Skiplist) Locate(target int) []*Skiplist {
	l := len(this.next) - 1
	prev := make([]*Skiplist, l+1)
	current := this
	for current != nil {
		next := current.next[l]
		if next == nil || next.val >= target {
			prev[l] = current
			if l == 0 {
				break
			} else {
				l--
			}
		} else {
			prev[l] = current
			current = next
		}
	}
	return prev
}

func (this *Skiplist) Search(target int) bool {
	l := len(this.next) - 1
	current := this
	for current != nil {
		next := current.next[l]
		if next == nil || next.val > target {
			if l == 0 {
				return false
			} else {
				l--
			}
		} else if next.val == target {
			return true
		} else {
			current = next
		}
	}

	return false
}

func (this *Skiplist) Add(num int) {
	l := len(this.next)
	var prev []*Skiplist
	if l > 0 {
		prev = this.Locate(num)
	}

	n := 1
	for this.r.Intn(2) != 0 {
		n++
	}

	node := &Skiplist{val: num}
	node.next = make([]*Skiplist, n)
	for ; n > l; n-- {
		this.next = append(this.next, node)
	}

	// a->c 之中插入b
	// b.next=a.next
	// a.next=b
	for i := 0; i < n; i++ {
		node.next[i] = prev[i].next[i]
		prev[i].next[i] = node
	}
}

func (this *Skiplist) Erase(num int) bool {
	if len(this.next) == 0 {
		return false
	}
	prev := this.Locate(num)

	node := prev[0].next[0]
	if node == nil || node.val != num {
		return false
	}
	for i := 0; i < len(node.next); i++ {
		prev[i].next[i] = node.next[i]
	}
	return true
}
