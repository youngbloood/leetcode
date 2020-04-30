package other_test

import (
	"leetcode/interview/medium/other"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	t.Log(other.EvalRPN(tokens))
	tokens = []string{"4", "13", "5", "/", "+"}
	t.Log(other.EvalRPN(tokens))
	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	t.Log(other.EvalRPN(tokens))
	tokens = []string{"18"}
	t.Log(other.EvalRPN(tokens))
}
