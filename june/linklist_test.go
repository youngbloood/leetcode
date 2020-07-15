package june_test

import (
	"leetcode/june"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	june.DeleteNode(nil, &june.ListNode{Val: 5})
}
