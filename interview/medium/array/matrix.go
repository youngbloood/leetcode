package array

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/777/

Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

Example 1:

Input:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
Output:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
Example 2:

Input:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
Output:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
Follow up:

A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?

Hint #1:
If any cell of the matrix has a zero we can record its row and column number using additional memory. But if you don't want to use extra memory then you can manipulate the array instead. i.e. simulating exactly what the question says.

Hint #2:
Setting cell values to zero on the fly while iterating might lead to discrepancies. What if you use some other integer value as your marker? There is still a better approach for this problem with 0(1) space.
*/

func SetZeroes(matrix [][]int) {
	setZeroes(matrix)
}

func setZeroes(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 {
		return
	}
	m := len(matrix)
	n := len(matrix[0])
	rowZero, colZero := false, false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colZero = true
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			rowZero = true
		}
	}
	// 第一行第一列存放除去第一行第一列的子矩阵哪些位置为0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	// 根据第一行第一列设置除去第一行第一列的子矩阵哪些位置为0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if rowZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if colZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
