package tree

/*
# Surrounded Regions
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/843/

Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example:

X X X X
X O O X
X X O X
X O X X

After running your function, the board should be:

X X X X
X X X X
X X X X
X O X X
Explanation:

Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.
*/
func Solve(board [][]byte) {
	solve(board)
}

// dfs
func solve(board [][]byte) {
	if board == nil || len(board) <= 2 {
		return
	}

	height := len(board)
	width := len(board[0])

	visited := make([][]bool, height)
	notSet := make([][]bool, height)
	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
		notSet[i] = make([]bool, width)
	}

	for i := 0; i < width; i++ {
		if board[0][i] == 'O' {
			solveDFS(board, visited, notSet, height, width, 0, i)
		}
		if board[height-1][i] == 'O' {
			solveDFS(board, visited, notSet, height, width, height-1, i)
		}
	}

	for i := 1; i < height-1; i++ {
		if board[i][0] == 'O' {
			solveDFS(board, visited, notSet, height, width, i, 0)
		}
		if board[i][width-1] == 'O' {
			solveDFS(board, visited, notSet, height, width, i, width-1)
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if !notSet[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func solveDFS(board [][]byte, visited [][]bool, notSet [][]bool, height, width, x, y int) {
	if x < 0 || x >= height || y < 0 || y >= width || board[x][y] != 'O' || visited[x][y] {
		return
	}
	visited[x][y] = true
	notSet[x][y] = true
	solveDFS(board, visited, notSet, height, width, x-1, y)
	solveDFS(board, visited, notSet, height, width, x, y-1)
	solveDFS(board, visited, notSet, height, width, x+1, y)
	solveDFS(board, visited, notSet, height, width, x, y+1)
}

/*
# Friend Circles
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/846/

There are N students in a class. Some of them are friends, while some are not. Their friendship is transitive in nature. For example, if A is a direct friend of B, and B is a direct friend of C, then A is an indirect friend of C. And we defined a friend circle is a group of students who are direct or indirect friends.

Given a N*N matrix M representing the friend relationship between students in the class. If M[i][j] = 1, then the ith and jth students are direct friends with each other, otherwise not. And you have to output the total number of friend circles among all the students.

Example 1:
Input:
[[1,1,0],
 [1,1,0],
 [0,0,1]]
Output: 2
Explanation:The 0th and 1st students are direct friends, so they are in a friend circle.
The 2nd student himself is in a friend circle. So return 2.

Example 2:
Input:
[[1,1,0],
 [1,1,1],
 [0,1,1]]
Output: 1
Explanation:The 0th and 1st students are direct friends, the 1st and 2nd students are direct friends,
so the 0th and 2nd students are indirect friends. All of them are in the same friend circle, so return 1.
Note:
N is in range [1,200].
M[i][i] = 1 for all students.
If M[i][j] = 1, then M[j][i] = 1.
*/
func FindCircleNum(M [][]int) int {
	return findCircleNum(M)
}
func findCircleNum(M [][]int) int {
	height := len(M)
	width := len(M[0])

	visited := make([][]bool, height)
	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
	}

	var circleNum int
	for i := 0; i < height; i++ {
		for j := i; j < width; j++ {
			if visited[i][j] || M[i][j] != 1 {
				continue
			}
			circleNum++
			findCircleNumDFS(M, visited, height, width, i, j)
		}
	}

	return circleNum
}

func findCircleNumDFS(M [][]int, visited [][]bool, height, width int, x, y int) {
	if x < 0 || x >= height || y < 0 || y >= width || visited[x][y] || M[x][y] == 0 {
		return
	}
	visited[x][y] = true

	for j := 0; j < width; j++ {
		if j == y {
			continue
		}
		findCircleNumDFS(M, visited, height, width, x, j)
	}
	for i := 0; i < width; i++ {
		if i == x {
			continue
		}
		findCircleNumDFS(M, visited, height, width, i, y)
	}
}

// leetcode优解
func findCircleNum2(M [][]int) int {
	var numCircles int
	visited := make([]bool, len(M))
	for i := 0; i < len(M); i++ {
		if visited[i] {
			continue
		}
		numCircles++
		dfs(i, M, visited)
	}
	return numCircles
}

func dfs(node int, adjMatrix [][]int, visited []bool) {
	visited[node] = true

	for neighbor, friendship := range adjMatrix[node] {
		if friendship != 1 || visited[neighbor] {
			continue
		}
		dfs(neighbor, adjMatrix, visited)
	}
}

/*
# Course Schedule
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/847/

There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?



Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
			 To take course 1 you should have finished course 0. So it is possible.

Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0, and to take course 0 you should
             also have finished course 1. So it is impossible.


Constraints:

The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
You may assume that there are no duplicate edges in the input prerequisites.
1 <= numCourses <= 10^5
*/
func CanFinish(numCourses int, prerequisites [][]int) bool {
	return canFinish(numCourses, prerequisites)
}
func canFinish(numCourses int, prerequisites [][]int) bool {

	// table := make(map[int][]int, len(prerequisites))

	// for _, pre := range prerequisites {
	// 	table[pre[0]] = pre
	// }
	return false
}

/*
# Longest Increasing Path in a Matrix
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/118/trees-and-graphs/849/

Given an integer matrix, find the length of the longest increasing path.

From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).

Example 1:

Input: nums =
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
Output: 4
Explanation: The longest increasing path is [1, 2, 6, 9].

Example 2:
Input: nums =
[
  [3,4,5],
  [3,2,6],
  [2,2,1]
]
Output: 4
Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
*/
func LongestIncreasingPath(matrix [][]int) int {
	return longestIncreasingPath(matrix)
}
func longestIncreasingPath(matrix [][]int) int {

	height := len(matrix)
	width := len(matrix[0])

	visited := make([][]bool, height)
	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
	}

	var max int

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			queue := []int{i*width + j}
			var length int
			for len(queue) != 0 {
				t := queue[0]
				queue = queue[1:]
				before := matrix[i][j]
				longestIncreasingPathBFS(matrix, visited, t/width-1, t%width, before, &queue, &length)
				longestIncreasingPathBFS(matrix, visited, t/width, t%width-1, before, &queue, &length)
				longestIncreasingPathBFS(matrix, visited, t/width+1, t%width, before, &queue, &length)
				longestIncreasingPathBFS(matrix, visited, t/width, t%width+1, before, &queue, &length)
			}
			if length > max {
				max = length
			}
			// 重置访问路径
			for l := 0; l < height; l++ {
				visited[l] = make([]bool, width)
			}
		}
	}
	return max
}

func longestIncreasingPathBFS(matrix [][]int, visited [][]bool, x, y, before int, queue *[]int, length *int) {
	if x < 0 || x >= len(matrix[0]) || y < 0 || y >= len(matrix) || visited[x][y] || matrix[x][y] <= before {
		// visited[x][y] = false
		return
	}
	visited[x][y] = true
	*length = *length + 1
	*queue = append(*queue, y*len(matrix[0])+x)
	before = matrix[x][y]
}
