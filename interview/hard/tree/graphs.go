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

	circleFriend := make([][2]int, 0)
	for i := 0; i < height; i++ {
		for j := i; j < width; j++ {
			if M[i][j] == 1 || M[j][i] == 1 {
				circleFriend = append(circleFriend, [2]int{i, j})
			}
		}
	}
	// 处理
	if len(circleFriend) == 0 {
		return 0
	}
	relate := make([][]int, 0)

	if circleFriend[len(circleFriend)-1][1] != width-1 {
		return len(circleFriend) + 1
	}
	return len(circleFriend)
}
