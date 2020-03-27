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
	combine(n, k)
}

// 列出C(n,k)的所有值
func combine(n int, k int) [][]int {
	if k > n {
		return nil
	}
	src := make([]int, n)
	for i := 0; i < n; i++ {
		src[i] = i + 1
	}
	return next(nil, src, nil, k)
}

func next(result [][]int, src []int, left []int, extra int) [][]int {
	if extra == 0 {
		result = append(result, src)
		return result
	}
	var s []int
	for i, v := range src {
		// 向s中添加最后一个元素
		s = append(s, v)
		left = src[i+1:]
		result = next(result, s, left, extra-1)
		// 移除最有一个元素
		s = s[:i]
	}
	return result
}
