// #https://leetcode.com/explore/learn/card/recursion-ii/503/recursion-to-iteration/2772/
package backtrack

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
func GenerateParenthesis(n int) []string {

}
func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}
	if n == 1 {
		return []string{"()"}
	}
}

func insert(src []string, n int) []string {

}
