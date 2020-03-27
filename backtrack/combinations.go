package backtrack

/*
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

Example:

Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/
func Combine(n int, k int) [][]int {

}

// 列出C(n,k)的所有值
func combine(n int, k int) [][]int {
	src := make([]int, n)
	for i := 0; i < n; i++ {
		src[i] = i + 1
	}

}

func next(result [][]int, src []int, extra int) []int {
	if extra == 0 {
		result = append(result, src)
		return result
	}
}
