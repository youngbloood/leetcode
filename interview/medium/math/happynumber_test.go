package math_test

import (
	"leetcode/interview/medium/math"
	"testing"
)

func TestIsHappy(t *testing.T) {
	t.Log(math.IsHappy(19))
	t.Log(math.IsHappy(2))
	t.Log(math.IsHappy(3))
}
