package other

import "strconv"

/*
# Evaluate Reverse Polish Notation
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/114/others/823/


Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Valid operators are +, -, *, /. Each operand may be an integer or another expression.

Note:

Division between two integers should truncate toward zero.
The given RPN expression is always valid. That means the expression would always evaluate to a result and there won't be any divide by zero operation.
Example 1:

Input: ["2", "1", "+", "3", "*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
Example 2:

Input: ["4", "13", "5", "/", "+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
Example 3:

Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
Output: 22
Explanation:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/
func EvalRPN(tokens []string) int {
	return evalRPN(tokens)
}
func evalRPN(tokens []string) int {
	stack := make(stack, 0)
	for i, token := range tokens {
		val, _ := strconv.Atoi(token)
		switch token {
		case "+", "-", "*", "/":
			val2 := (&stack).Pop()
			val1 := (&stack).Pop()
			switch token {
			case "+":
				val = val1 + val2
			case "-":
				val = val1 - val2
			case "*":
				val = val1 * val2
			case "/":
				val = val1 / val2
			}
			fallthrough
		default:
			(&stack).Push(val)
		}
		if i == len(tokens)-1 {
			return val
		}
	}
	return 0
}

type stack []int

func (s *stack) Push(i int) {
	*s = append(*s, i)
}
func (s *stack) Pop() int {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}
