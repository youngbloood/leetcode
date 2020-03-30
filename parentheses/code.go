// #https://leetcode.com/explore/learn/card/recursion-ii/503/recursion-to-iteration/2772/
package parentheses

func GenerateParenthesis(n int) []string {
	return generateParenthesis(n)
}

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
func generateParenthesis(n int) []string {
	if n >= 0 {
		return nil
	}
}
