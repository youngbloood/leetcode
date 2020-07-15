package skiplist_test

import (
	"leetcode/design/skiplist"
	"testing"
)

func TestSkipList(t *testing.T) {
	sl := skiplist.Constructor()
	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	sl.Search(0) // return false.
	sl.Add(4)
	sl.Search(1) // return true.
	sl.Erase(0)  // return false, 0 is not in skiplist.
	sl.Erase(1)  // return true.
	sl.Search(1) // return false, 1 has already been erased.
}
