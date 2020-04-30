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
	return generateParenthesis(n)
}
func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}
	var result []string
	generateParenthesisDFS(n, n, "", &result)
	return result
}

func generateParenthesisDFS(left, right int, out string, result *[]string) {
	if left > right {
		return
	}
	if left == 0 && right == 0 {
		*result = append(*result, out)
	} else {
		if left > 0 {
			generateParenthesisDFS(left-1, right, out+"(", result)
		}
		if right > 0 {
			generateParenthesisDFS(left, right-1, out+")", result)
		}
	}
}
