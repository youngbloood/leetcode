package dynamic

/*
# Unique Paths
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/111/dynamic-programming/808/

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?


start	_	_	_	_	_	  _
  _		_	_	_	_	_	  _
  _		_	_	_	_	_	finish

Above is a 7 x 3 grid. How many possible unique paths are there?



Example 1:

start	_	 _
  _		_	end

Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right
Example 2:

Input: m = 7, n = 3
Output: 28


Constraints:

1 <= m, n <= 100
It's guaranteed that the answer will be less than or equal to 2 * 10 ^ 9.
*/

// 右移数量 right=m-1
// 下移数量 down=n-1
// 排列结果=A((right+down),(right+down))/A(right,right)/A(down,down)
// P[i][j] = P[i - 1][j] + P[i][j - 1]

func UniquePaths(m int, n int) int {
	return uniquePaths3(m, n)
}

// NOTE: 超时
func uniquePaths(m int, n int) int {
	right := m - 1
	down := n - 1

	var count int
	uniquePathRobot(right, down, &count)
	return count
}

func uniquePathRobot(right, down int, count *int) {
	if right == 0 && down == 0 {
		*count += 1
		return
	}
	if right >= 1 {
		uniquePathRobot(right-1, down, count)
	}
	if down >= 1 {
		uniquePathRobot(right, down-1, count)
	}
}

// Fibonacci方法 // NOTE: 超时
func uniquePathsFibonacci(m int, n int) int {
	if m <= 1 || n <= 1 {
		return 1
	}
	return uniquePathsFibonacci(m-1, n) + uniquePathsFibonacci(m, n-1)
}

// 用path存储到达每个格子的路径数量
func uniquePaths3(m int, n int) int {
	path := make([][]int, n)
	for i := 0; i < n; i++ {
		path[i] = make([]int, m)
		path[i][0] = 1
	}
	for i := 0; i < m; i++ {
		path[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	return path[n-1][m-1]
}
