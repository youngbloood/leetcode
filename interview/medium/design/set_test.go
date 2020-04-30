package design_test

import (
	"leetcode/interview/medium/design"
	"testing"
)

// NOTE: 随机顺序于leetcode不一致通不过测试
func TestConstructor(t *testing.T) {
	obj := design.Constructor()
	// val := 1
	// t.Log(obj.Insert(val))
	// t.Log(obj.Remove(val))
	// t.Log(obj.Insert(2))
	// t.Log(obj.GetRandom())

	/*
			Input:
		["RandomizedSet","insert","insert","remove","insert","remove","getRandom"]
		[[],[0],[1],[0],[2],[1],[]]
		Output:
		[null,true,true,true,true,true,1]
		Expected:
		[null,true,true,true,true,true,2]
	*/

	t.Log(obj.GetRandom())
	t.Log(obj.Insert(1))
	t.Log(obj.Insert(10))
	// t.Log(obj.Remove(0))
	t.Log(obj.Insert(20))
	t.Log(obj.Insert(30))

	for i := 0; i <= 30; i++ {
		t.Log(obj.GetRandom())
	}

}
