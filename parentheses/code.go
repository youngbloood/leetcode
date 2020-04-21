// #https://leetcode.com/explore/learn/card/recursion-ii/503/recursion-to-iteration/2772/
package parentheses

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
	generateParenthesisBFS(n, n, "", &result)
	return result
}

func generateParenthesisBFS(left, right int, out string, result *[]string) {
	if left > right {
		return
	}
	if left == 0 && right == 0 {
		*result = append(*result, out)
	}
	if left > 0 {
		generateParenthesisBFS(left-1, right, out+"(", result)
	}
	if right > 0 {
		generateParenthesisBFS(left, right-1, out+")", result)
	}
}
