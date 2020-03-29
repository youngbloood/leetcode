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
	return combine(n, k)
}

// 列出C(n,k)的所有值
func combine(n int, k int) [][]int {
	if k > n || n <= 0 || k <= 0 {
		return nil
	}
	src := make([]int, n)
	for i := 0; i < n; i++ {
		src[i] = i + 1
	}
	return next(nil, nil, src, k)
}

func next(result [][]int, src []int, left []int, extra int) [][]int {
	if extra == 0 {
		srcTmp := make([]int, len(src))
		copy(srcTmp, src)
		result = append(result, srcTmp)
		return result
	}
	var s []int
	for i, v := range left {
		// 向s中添加最后一个元素
		s = append(src, v)
		leftTemp := left[i+1:]
		result = next(result, s, leftTemp, extra-1)
		// 移除最有一个元素
		s = s[:len(s)-1]
	}
	return result
}
