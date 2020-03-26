package matrix

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}

/*
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
Example:

Consider the following matrix:

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
Given target = 5, return true.

Given target = 20, return false.
*/
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	if target < matrix[0][0] && target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return target == matrix[0][0]
	}

	// for _, row := range matrix {
	// 	for _, col := range row {
	// 		if col == target {
	// 			return true
	// 		}
	// 	}
	// }
	// return false
	return searchMatrix(matrix[:len(matrix)/2][:len(matrix[0])/2], target) ||
		searchMatrix(matrix[:len(matrix)/2][len(matrix[0])/2:], target) ||
		searchMatrix(matrix[len(matrix)/2:][:len(matrix[0])/2], target) ||
		searchMatrix(matrix[len(matrix)/2:][len(matrix[0])/2:], target)
}
