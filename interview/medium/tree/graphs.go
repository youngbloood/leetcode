package tree

/*
# Number of Islands
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/108/trees-and-graphs/792/

Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input:
11110
11010
11000
00000
Output: 1

Example 2:
Input:
11000
11000
00100
00011
Output: 3

// 解法：DFS和BFS解法
// # https://zhuanlan.zhihu.com/p/24986203
*/
func NumIslands(grid [][]byte) int {
	return numIslandsBFS(grid)
}
func numIslandsDFS(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	var count int
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' || visited[i][j] {
				continue
			}
			numIslandsWithDFS(grid, visited, i, j)
			count++
		}
	}
	return count
}

func numIslandsWithDFS(grid [][]byte, visited [][]bool, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' || visited[x][y] {
		return
	}
	visited[x][y] = true
	numIslandsWithDFS(grid, visited, x-1, y)
	numIslandsWithDFS(grid, visited, x+1, y)
	numIslandsWithDFS(grid, visited, x, y-1)
	numIslandsWithDFS(grid, visited, x, y+1)
}

func numIslandsBFS(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	var count int
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' || visited[i][j] {
				continue
			}
			count++
			queue := []int{i*n + j}

			for len(queue) != 0 {
				t := queue[0]
				queue = queue[1:]
				numIslandsWithBFS(grid, visited, &queue, t/n-1, t%n)
				numIslandsWithBFS(grid, visited, &queue, t/n, t%n-1)
				numIslandsWithBFS(grid, visited, &queue, t/n+1, t%n)
				numIslandsWithBFS(grid, visited, &queue, t/n, t%n+1)
			}
		}
	}
	return count
}

func numIslandsWithBFS(grid [][]byte, visited [][]bool, queue *[]int, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' || visited[x][y] {
		return
	}
	visited[x][y] = true
	*queue = append(*queue, x*len(grid[0])+y)
}
