package array

/*
#  Spiral Matrix
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/828/

Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:

Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/
func SpiralOrder(matrix [][]int) []int {
	return spiralOrder(matrix)
}
func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return nil
	}
	result := make([]int, len(matrix)*len(matrix[0]))
	xn := len(matrix[0])
	yn := len(matrix)
	spiralOrderMatrix(matrix, xn, yn, 0, 0, 0, result)
	return result
}

func spiralOrderMatrix(matrix [][]int, xLen, yLen, xStart, yStart, i int, result []int) {
	if i == len(result) {
		return
	}
	var pos int
	// 注意位置的分割
	for pos = 0; pos < len(result)-i; pos++ {
		if pos <= xLen-1 {
			result[i+pos] = matrix[yStart][pos+xStart]
		} else if pos > xLen-1 && pos < xLen-1+yLen-1 {
			result[i+pos] = matrix[pos-(xLen-1)+yStart][xLen-1+xStart]
		} else if pos >= xLen-1+yLen-1 && pos < 2*(xLen-1)+yLen-1 {
			result[i+pos] = matrix[yLen-1+yStart][xLen+xStart-(pos-(xLen-1+yLen-1))-1]
		} else if pos < 2*(xLen-1+yLen-1) {
			result[i+pos] = matrix[yStart+yLen-(pos-(2*(xLen-1)+yLen-1))-1][xStart]
		} else {
			break
		}
	}
	spiralOrderMatrix(matrix, xLen-2, yLen-2, xStart+1, yStart+1, i+pos, result)
}
