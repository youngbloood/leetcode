package backtrack

import "sort"

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/796/

Solution
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

func Subsets(nums []int) [][]int {
	return subsets(nums)
}

func subsets(nums []int) [][]int {
	var result [][]int
	for i := 0; i <= len(nums); i++ {
		generateSets(nums[i:], i, nil, &result)
	}
	return result
}

func generateSets(nums []int, n int, subset []int, result *[][]int) {
	if n == 0 {
		sort.Ints(subset)
		*result = append(*result, subset)
	}
	for i, v := range nums {
		generateSets(nums[i+1:], n-1, append(subset, v), result)
	}
}

// [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 9, 0, 3, 5, 7
// output
// [[],[9],[0],[3],[5],[7],[9,0],[9,3],[9,5],[9,7],[0,3],[0,5],[0,7],[3,5],[3,7],[5,7],[9,0,3],[9,0,5],[9,0,7],[9,3,5],[9,3,7],[9,5,7],[0,3,5],[0,3,7],[0,5,7],[3,5,7],[9,0,3,7],[9,0,3,7],[9,0,5,7],[9,3,5,7],[0,3,5,7],[9,0,3,5,7]]

// Expected
// [[],[9],[0],[0,9],[3],[3,9],[0,3],[0,3,9],[5],[5,9],[0,5],[0,5,9],[3,5],[3,5,9],[0,3,5],[0,3,5,9],[7],[7,9],[0,7],[0,7,9],[3,7],[3,7,9],[0,3,7],[0,3,7,9],[5,7],[5,7,9],[0,5,7],[0,5,7,9],[3,5,7],[3,5,7,9],[0,3,5,7],[0,3,5,7,9]]
